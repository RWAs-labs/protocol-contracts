// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "../contracts/mevm/SystemContract.sol";

import "./utils/TestUniversalContract.sol";

import "./utils/WMUSE.sol";
import "./utils/upgrades/GatewayMEVMUpgradeTest.sol";

import "../contracts/mevm/GatewayMEVM.sol";
import "../contracts/mevm/MRC20.sol";
import "../contracts/mevm/interfaces/IGatewayMEVM.sol";
import "../contracts/mevm/interfaces/IMRC20.sol";
import { Upgrades } from "openzeppelin-foundry-upgrades/Upgrades.sol";

contract GatewayMEVMInboundTest is Test, IGatewayMEVMEvents, IGatewayMEVMErrors {
    address payable proxy;
    GatewayMEVM gateway;
    MRC20 mrc20;
    WETH9 museToken;
    SystemContract systemContract;
    TestUniversalContract testUniversalContract;
    address owner;
    address addr1;
    address protocolAddress;
    RevertOptions revertOptions;
    CallOptions callOptions;

    error ZeroAddress();
    error LowBalance();
    error MUSENotSupported();

    event WithdrawnV2(
        address indexed sender,
        uint256 indexed chainId,
        bytes receiver,
        address mrc20,
        uint256 value,
        uint256 gasfee,
        uint256 protocolFlatFee,
        bytes message,
        CallOptions callOptions,
        RevertOptions revertOptions
    );

    uint256 constant MIN_GAS_LIMIT = 100_000;

    function setUp() public {
        owner = address(this);
        addr1 = address(0x1234);

        museToken = new WETH9();

        proxy = payable(
            Upgrades.deployUUPSProxy(
                "GatewayMEVM.sol", abi.encodeCall(GatewayMEVM.initialize, (address(museToken), owner))
            )
        );
        gateway = GatewayMEVM(proxy);

        protocolAddress = gateway.PROTOCOL_ADDRESS();
        testUniversalContract = new TestUniversalContract();

        vm.startPrank(protocolAddress);
        systemContract = new SystemContract(address(0), address(0), address(0));
        mrc20 = new MRC20("TOKEN", "TKN", 18, 1, CoinType.Gas, 0, address(systemContract), address(gateway));
        systemContract.setGasCoinMRC20(1, address(mrc20));
        systemContract.setGasPrice(1, 1);
        vm.deal(protocolAddress, 1_000_000_000);
        museToken.deposit{ value: 10 }();
        museToken.approve(address(gateway), 10);
        mrc20.deposit(owner, 100_000_000);
        vm.stopPrank();

        vm.startPrank(owner);
        mrc20.approve(address(gateway), 100_000_000);
        museToken.deposit{ value: 10 }();
        museToken.approve(address(gateway), 10);
        vm.stopPrank();

        revertOptions = RevertOptions({
            revertAddress: address(0x321),
            callOnRevert: true,
            abortAddress: address(0x321),
            revertMessage: "",
            onRevertGasLimit: 0
        });

        callOptions = CallOptions({ gasLimit: MIN_GAS_LIMIT, isArbitraryCall: true });
    }

    function testWithdrawMRC20() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = mrc20.balanceOf(owner);

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Withdrawn(
            owner,
            0,
            abi.encodePacked(addr1),
            address(mrc20),
            amount,
            0,
            mrc20.PROTOCOL_FLAT_FEE(),
            "",
            CallOptions({ gasLimit: 0, isArbitraryCall: true }),
            revertOptions
        );
        gateway.withdraw(abi.encodePacked(addr1), amount, address(mrc20), revertOptions);

        uint256 ownerBalanceAfter = mrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore - amount, ownerBalanceAfter);
    }

    function testWithdrawMRC20FailsIfNoBalanceForGasFee() public {
        uint256 amount = 1;
        uint256 ownerBalance = mrc20.balanceOf(owner);
        mrc20.transfer(address(0x123), ownerBalance);

        vm.prank(protocolAddress);
        mrc20.updateGasLimit(10);

        // Get the gas fee information from the contract
        (address gasMRC20, uint256 gasFee) = mrc20.withdrawGasFeeWithGasLimit(10);

        vm.expectRevert(
            abi.encodeWithSelector(
                GasFeeTransferFailed.selector,
                gasMRC20,
                protocolAddress, // The PROTOCOL_ADDRESS constant
                gasFee
            )
        );

        gateway.withdraw(abi.encodePacked(addr1), amount, address(mrc20), revertOptions);
    }

    function testWithdrawMRC20FailsIfNoBalanceForTransfer() public {
        uint256 amount = 2;
        uint256 ownerBalance = mrc20.balanceOf(owner);
        mrc20.transfer(address(0x123), ownerBalance - 1);

        // Assuming MRC20TransferFailed now takes parameters:
        address tokenAddress = address(mrc20);
        address from = owner;
        address to = address(gateway);

        vm.expectRevert(abi.encodeWithSelector(MRC20TransferFailed.selector, tokenAddress, from, to, amount));

        gateway.withdraw(abi.encodePacked(addr1), amount, address(mrc20), revertOptions);
    }

    function testWithdrawMRC20FailsIfMessageSizeExceeded() public {
        revertOptions.revertMessage = new bytes(gateway.MAX_MESSAGE_SIZE() + 1);

        uint256 messageSize = revertOptions.revertMessage.length;
        uint256 maxSize = gateway.MAX_MESSAGE_SIZE();

        vm.expectRevert(abi.encodeWithSelector(MessageSizeExceeded.selector, messageSize, maxSize));

        gateway.withdraw(abi.encodePacked(addr1), 2, address(mrc20), revertOptions);
    }

    function testWithdrawMRC20FailsIsAmountIs0() public {
        vm.expectRevert(InsufficientMRC20Amount.selector);
        gateway.withdraw(abi.encodePacked(addr1), 0, address(mrc20), revertOptions);
    }

    function testWithdrawMRC20FailsIfReceiverIsZeroAddress() public {
        vm.expectRevert(ZeroAddress.selector);
        gateway.withdraw(abi.encodePacked(""), 1, address(mrc20), revertOptions);
    }

    function testWithdrawMRC20FailsIfNoAllowance() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = mrc20.balanceOf(owner);

        // Remove allowance for gateway
        vm.prank(owner);
        mrc20.approve(address(gateway), 0);

        vm.expectRevert();
        gateway.withdraw(abi.encodePacked(addr1), amount, address(mrc20), revertOptions);

        // Check that balance didn't change
        uint256 ownerBalanceAfter = mrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore, ownerBalanceAfter);
    }

    function testWithdrawAndCallMRC20FailsIfReceiverIsZeroAddress() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert(ZeroAddress.selector);
        gateway.withdrawAndCall(
            abi.encodePacked(""),
            1,
            address(mrc20),
            message,
            CallOptions({ gasLimit: MIN_GAS_LIMIT, isArbitraryCall: false }),
            revertOptions
        );
    }

    function testWithdrawAndCallMRC20FailsIfMessageSizeExceeded() public {
        bytes memory message = new bytes(gateway.MAX_MESSAGE_SIZE() / 2);
        revertOptions.revertMessage = new bytes(gateway.MAX_MESSAGE_SIZE() / 2 + 1);

        uint256 messageSize = message.length + revertOptions.revertMessage.length;
        uint256 maxSize = gateway.MAX_MESSAGE_SIZE();

        vm.expectRevert(abi.encodeWithSelector(MessageSizeExceeded.selector, messageSize, maxSize));

        gateway.withdrawAndCall(abi.encodePacked(addr1), 1, address(mrc20), message, callOptions, revertOptions);
    }

    function testWithdrawAndCallMRC20FailsIfGasLimitIsZero() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert(InsufficientGasLimit.selector);
        gateway.withdrawAndCall(
            abi.encodePacked(addr1),
            1,
            address(mrc20),
            message,
            CallOptions({ gasLimit: 0, isArbitraryCall: false }),
            revertOptions
        );
    }

    function testWithdrawAndCallMRC20FailsIfGasLimitIsBelowMin() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert(InsufficientGasLimit.selector);
        gateway.withdrawAndCall(
            abi.encodePacked(addr1),
            1,
            address(mrc20),
            message,
            CallOptions({ gasLimit: MIN_GAS_LIMIT - 1, isArbitraryCall: false }),
            revertOptions
        );
    }

    function testWithdrawAndCallMRC20FailsIfAmountIsZero() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert(InsufficientMRC20Amount.selector);
        gateway.withdrawAndCall(
            abi.encodePacked(addr1),
            0,
            address(mrc20),
            message,
            CallOptions({ gasLimit: MIN_GAS_LIMIT, isArbitraryCall: false }),
            revertOptions
        );
    }

    function testWithdrawMRC20WithMessageFailsIfNoAllowance() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = mrc20.balanceOf(owner);

        // Remove allowance for gateway
        vm.prank(owner);
        mrc20.approve(address(gateway), 0);

        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert();
        gateway.withdrawAndCall(
            abi.encodePacked(addr1),
            amount,
            address(mrc20),
            message,
            CallOptions({ gasLimit: MIN_GAS_LIMIT, isArbitraryCall: false }),
            revertOptions
        );

        // Check that balance didn't change
        uint256 ownerBalanceAfter = mrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore, ownerBalanceAfter);
    }

    function testWithdrawMRC20WithMessage() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = mrc20.balanceOf(owner);

        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        uint256 expectedGasFee = MIN_GAS_LIMIT;
        uint256 gasLimit = MIN_GAS_LIMIT;
        vm.expectEmit(true, true, true, true, address(gateway));
        emit WithdrawnAndCalled(
            owner,
            0,
            abi.encodePacked(addr1),
            address(mrc20),
            amount,
            expectedGasFee,
            mrc20.PROTOCOL_FLAT_FEE(),
            message,
            CallOptions({ gasLimit: gasLimit, isArbitraryCall: true }),
            revertOptions
        );
        gateway.withdrawAndCall(
            abi.encodePacked(addr1),
            amount,
            address(mrc20),
            message,
            CallOptions({ gasLimit: gasLimit, isArbitraryCall: true }),
            revertOptions
        );

        uint256 ownerBalanceAfter = mrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore - amount - expectedGasFee, ownerBalanceAfter);
    }

    function testWithdrawAndCallMRC20WithCallOptsFailsIfReceiverIsZeroAddress() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert(ZeroAddress.selector);
        gateway.withdrawAndCall(abi.encodePacked(""), 1, address(mrc20), message, callOptions, revertOptions);
    }

    function testWithdrawAndCallMRC20WithCallOptsFailsIfMessageSizeExceeded() public {
        bytes memory message = new bytes(gateway.MAX_MESSAGE_SIZE() / 2);
        revertOptions.revertMessage = new bytes(gateway.MAX_MESSAGE_SIZE() / 2 + 1);

        uint256 messageSize = message.length + revertOptions.revertMessage.length;
        uint256 maxSize = gateway.MAX_MESSAGE_SIZE();

        vm.expectRevert(abi.encodeWithSelector(MessageSizeExceeded.selector, messageSize, maxSize));

        gateway.withdrawAndCall(abi.encodePacked(addr1), 1, address(mrc20), message, callOptions, revertOptions);
    }

    function testWithdrawAndCallMRC20WithCallOptsFailsIfGasLimitIsZero() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        callOptions.gasLimit = 0;
        vm.expectRevert(InsufficientGasLimit.selector);
        gateway.withdrawAndCall(abi.encodePacked(addr1), 1, address(mrc20), message, callOptions, revertOptions);
    }

    function testWithdrawAndCallMRC20WithCallOptsFailsIfAmountIsZero() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert(InsufficientMRC20Amount.selector);
        gateway.withdrawAndCall(abi.encodePacked(addr1), 0, address(mrc20), message, callOptions, revertOptions);
    }

    function testWithdrawMRC20WithMessageWithCallOptsFailsIfNoAllowance() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = mrc20.balanceOf(owner);

        // Remove allowance for gateway
        vm.prank(owner);
        mrc20.approve(address(gateway), 0);

        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert();
        gateway.withdrawAndCall(abi.encodePacked(addr1), amount, address(mrc20), message, callOptions, revertOptions);

        // Check that balance didn't change
        uint256 ownerBalanceAfter = mrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore, ownerBalanceAfter);
    }

    function testWithdrawMRC20WithCallOptsWithMessage() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = mrc20.balanceOf(owner);

        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        uint256 expectedGasFee = MIN_GAS_LIMIT;
        uint256 gasLimit = MIN_GAS_LIMIT;
        vm.expectEmit(true, true, true, true, address(gateway));
        emit WithdrawnAndCalled(
            owner,
            0,
            abi.encodePacked(addr1),
            address(mrc20),
            amount,
            expectedGasFee,
            mrc20.PROTOCOL_FLAT_FEE(),
            message,
            CallOptions({ gasLimit: gasLimit, isArbitraryCall: true }),
            revertOptions
        );

        gateway.withdrawAndCall(
            abi.encodePacked(addr1),
            amount,
            address(mrc20),
            message,
            CallOptions({ gasLimit: gasLimit, isArbitraryCall: true }),
            revertOptions
        );

        uint256 ownerBalanceAfter = mrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore - amount - expectedGasFee, ownerBalanceAfter);
    }

    function testWithdrawMUSEFailsIfAmountIsZero() public {
        // TODO: replace error to check once MUSE supported back
        // https://github.com/muse-chain/protocol-contracts/issues/394
        // vm.expectRevert(InsufficientMuseAmount.selector);
        vm.expectRevert(MUSENotSupported.selector);

        gateway.withdraw(abi.encodePacked(addr1), 0, 1, revertOptions);
    }

    function testWithdrawMUSEFailsIfMessageSizeExceeded() public {
        revertOptions.revertMessage = new bytes(gateway.MAX_MESSAGE_SIZE() + 1);

        // TODO: replace error to check once MUSE supported back
        // https://github.com/muse-chain/protocol-contracts/issues/394
        // vm.expectRevert(MessageSizeExceeded.selector);
        vm.expectRevert(MUSENotSupported.selector);

        gateway.withdraw(abi.encodePacked(addr1), 1, 1, revertOptions);
    }

    function testWithdrawMUSEFailsIfReceiverIsZeroAddress() public {
        // TODO: replace error to check once MUSE supported back
        // https://github.com/muse-chain/protocol-contracts/issues/394
        // vm.expectRevert(ZeroAddress.selector);
        vm.expectRevert(MUSENotSupported.selector);

        gateway.withdraw(abi.encodePacked(""), 0, 1, revertOptions);
    }

    function testWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsZero() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);

        // TODO: replace error to check once MUSE supported back
        // https://github.com/muse-chain/protocol-contracts/issues/394
        // vm.expectRevert(InsufficientMuseAmount.selector);
        vm.expectRevert(MUSENotSupported.selector);

        gateway.withdrawAndCall(abi.encodePacked(addr1), 0, 1, message, callOptions, revertOptions);
    }

    function testWithdrawAndCallMUSEWithCallOptsFailsIfMessageSizeExceeded() public {
        bytes memory message = new bytes(gateway.MAX_MESSAGE_SIZE() / 2);
        revertOptions.revertMessage = new bytes(gateway.MAX_MESSAGE_SIZE() / 2 + 1);

        // TODO: replace error to check once MUSE supported back
        // https://github.com/muse-chain/protocol-contracts/issues/394
        // vm.expectRevert(MessageSizeExceeded.selector);
        vm.expectRevert(MUSENotSupported.selector);

        gateway.withdrawAndCall(abi.encodePacked(addr1), 1, 1, message, callOptions, revertOptions);
    }

    function testWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsReceiverIsZeroAddress() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);

        // TODO: replace error to check once MUSE supported back
        // https://github.com/muse-chain/protocol-contracts/issues/394
        // vm.expectRevert(ZeroAddress.selector);
        vm.expectRevert(MUSENotSupported.selector);

        gateway.withdrawAndCall(abi.encodePacked(""), 1, 1, message, callOptions, revertOptions);
    }

    function testWithdrawMUSE() public {
        uint256 amount = 1;
        // uint256 gatewayBalanceBefore = museToken.balanceOf(address(gateway));
        // uint256 protocolBalanceBefore = protocolAddress.balance;
        uint256 chainId = 1;

        // vm.expectEmit(true, true, true, true, address(gateway));
        emit Withdrawn(
            owner,
            chainId,
            abi.encodePacked(addr1),
            address(museToken),
            amount,
            0,
            0,
            "",
            CallOptions({ gasLimit: 0, isArbitraryCall: true }),
            revertOptions
        );

        // TODO: remove error once MUSE supported back
        // https://github.com/muse-chain/protocol-contracts/issues/394
        vm.expectRevert(MUSENotSupported.selector);

        gateway.withdraw(abi.encodePacked(addr1), amount, chainId, revertOptions);

        // uint256 ownerBalanceAfter = museToken.balanceOf(owner);
        // assertEq(ownerBalanceBefore - 1, ownerBalanceAfter);

        // uint256 gatewayBalanceAfter = museToken.balanceOf(address(gateway));
        // assertEq(gatewayBalanceBefore, gatewayBalanceAfter);

        // // Verify amount is transfered to protocol address
        // assertEq(protocolBalanceBefore + 1, protocolAddress.balance);
    }

    function testWithdrawMUSEFailsIfNoAllowance() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = museToken.balanceOf(owner);
        uint256 gatewayBalanceBefore = museToken.balanceOf(address(gateway));
        uint256 protocolBalanceBefore = protocolAddress.balance;
        uint256 chainId = 1;

        // Remove allowance for gateway
        vm.prank(owner);
        museToken.approve(address(gateway), 0);

        vm.expectRevert();
        gateway.withdraw(abi.encodePacked(addr1), amount, chainId, revertOptions);

        // Verify balances not changed
        uint256 ownerBalanceAfter = museToken.balanceOf(owner);
        assertEq(ownerBalanceBefore, ownerBalanceAfter);

        uint256 gatewayBalanceAfter = museToken.balanceOf(address(gateway));
        assertEq(gatewayBalanceBefore, gatewayBalanceAfter);

        assertEq(protocolBalanceBefore, protocolAddress.balance);
    }

    function testWithdrawMUSEFailsIfNoBalance() public {
        uint256 amount = 1;
        uint256 ownerBalance = museToken.balanceOf(owner);
        museToken.transfer(address(0x123), ownerBalance);
        uint256 chainId = 1;

        vm.expectRevert();
        gateway.withdraw(abi.encodePacked(addr1), amount, chainId, revertOptions);
    }

    function testWithdrawMUSEWithCallOptsWithMessage() public {
        uint256 amount = 1;
        // uint256 ownerBalanceBefore = museToken.balanceOf(owner);
        // uint256 gatewayBalanceBefore = museToken.balanceOf(address(gateway));
        // uint256 protocolAddressBalanceBefore = protocolAddress.balance;
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        uint256 chainId = 1;

        // vm.expectEmit(true, true, true, true, address(gateway));
        emit Withdrawn(
            owner,
            chainId,
            abi.encodePacked(addr1),
            address(museToken),
            amount,
            0,
            0,
            message,
            callOptions,
            revertOptions
        );

        // TODO: remove error once MUSE supported back
        // vm.expectRevert(InsufficientGasLimit.selector);
        vm.expectRevert(MUSENotSupported.selector);

        gateway.withdrawAndCall(abi.encodePacked(addr1), amount, chainId, message, callOptions, revertOptions);

        // uint256 ownerBalanceAfter = museToken.balanceOf(owner);
        // assertEq(ownerBalanceBefore - 1, ownerBalanceAfter);

        // uint256 gatewayBalanceAfter = museToken.balanceOf(address(gateway));
        // assertEq(gatewayBalanceBefore, gatewayBalanceAfter);

        // // Verify amount is transfered to fungible module
        // assertEq(protocolAddressBalanceBefore + 1, protocolAddress.balance);
    }

    function testWithdrawMUSEWithCallOptsWithMessageFailsIfNoAllowance() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = museToken.balanceOf(owner);
        uint256 gatewayBalanceBefore = museToken.balanceOf(address(gateway));
        uint256 protocolAddressBalanceBefore = protocolAddress.balance;
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        uint256 chainId = 1;

        // Remove allowance for gateway
        vm.prank(owner);
        museToken.approve(address(gateway), 0);

        vm.expectRevert();
        gateway.withdrawAndCall(abi.encodePacked(addr1), amount, chainId, message, callOptions, revertOptions);

        // Verify balances not changed
        uint256 ownerBalanceAfter = museToken.balanceOf(owner);
        assertEq(ownerBalanceBefore, ownerBalanceAfter);

        uint256 gatewayBalanceAfter = museToken.balanceOf(address(gateway));
        assertEq(gatewayBalanceBefore, gatewayBalanceAfter);

        assertEq(protocolAddressBalanceBefore, protocolAddress.balance);
    }

    function testWithdrawMUSEWithCallOptsWithMessageFailsIfGasLimitIsZero() public {
        uint256 amount = 1;
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        uint256 chainId = 1;

        // Remove allowance for gateway
        vm.prank(owner);
        museToken.approve(address(gateway), 0);

        callOptions.gasLimit = 0;

        // TODO: replace error to check once MUSE supported back
        // https://github.com/muse-chain/protocol-contracts/issues/394
        // vm.expectRevert(InsufficientGasLimit.selector);
        vm.expectRevert(MUSENotSupported.selector);

        gateway.withdrawAndCall(abi.encodePacked(addr1), amount, chainId, message, callOptions, revertOptions);
    }

    function testCallWithCallOptsFailsIfReceiverIsZeroAddress() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert(ZeroAddress.selector);
        gateway.call(abi.encodePacked(""), address(mrc20), message, callOptions, revertOptions);
    }

    function testCallWithCallOptsFailsIfMessageSizeExceeded() public {
        bytes memory message = new bytes(gateway.MAX_MESSAGE_SIZE() / 2);
        revertOptions.revertMessage = new bytes(gateway.MAX_MESSAGE_SIZE() / 2 + 1);

        uint256 messageSize = message.length + revertOptions.revertMessage.length;
        uint256 maxSize = gateway.MAX_MESSAGE_SIZE();

        vm.expectRevert(abi.encodeWithSelector(MessageSizeExceeded.selector, messageSize, maxSize));

        gateway.withdrawAndCall(
            abi.encodePacked(addr1),
            1,
            address(mrc20),
            message,
            CallOptions({ gasLimit: MIN_GAS_LIMIT, isArbitraryCall: false }),
            revertOptions
        );
    }

    function testCallWithCallOptsFailsIfGasLimitIsZero() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        callOptions.gasLimit = 0;
        vm.expectRevert(InsufficientGasLimit.selector);
        gateway.call(abi.encodePacked(addr1), address(mrc20), message, callOptions, revertOptions);
    }

    function testCallWithCallOptsFailsIfGasLimitIsBelowMin() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        callOptions.gasLimit = MIN_GAS_LIMIT - 1;
        vm.expectRevert(InsufficientGasLimit.selector);
        gateway.call(abi.encodePacked(addr1), address(mrc20), message, callOptions, revertOptions);
    }

    function testCallWithCallOpts() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectEmit(true, true, true, true, address(gateway));

        emit Called(owner, address(mrc20), abi.encodePacked(addr1), message, callOptions, revertOptions);
        gateway.call(abi.encodePacked(addr1), address(mrc20), message, callOptions, revertOptions);
    }

    function testUpgradeAndWithdrawMRC20() public {
        // upgrade
        Upgrades.upgradeProxy(proxy, "GatewayMEVMUpgradeTest.sol", "", owner);
        GatewayMEVMUpgradeTest gatewayUpgradeTest = GatewayMEVMUpgradeTest(proxy);
        // withdraw
        uint256 amount = 1;
        uint256 ownerBalanceBefore = mrc20.balanceOf(owner);

        vm.expectEmit(true, true, true, true, address(gatewayUpgradeTest));
        emit WithdrawnV2(
            owner,
            0,
            abi.encodePacked(addr1),
            address(mrc20),
            amount,
            0,
            mrc20.PROTOCOL_FLAT_FEE(),
            "",
            CallOptions({ gasLimit: 0, isArbitraryCall: true }),
            revertOptions
        );
        gatewayUpgradeTest.withdraw(abi.encodePacked(addr1), amount, address(mrc20), revertOptions);

        uint256 ownerBalanceAfter = mrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore - amount, ownerBalanceAfter);
    }
}

contract GatewayMEVMOutboundTest is Test, IGatewayMEVMEvents, IGatewayMEVMErrors {
    address payable proxy;
    GatewayMEVM gateway;
    MRC20 mrc20;
    WETH9 museToken;
    SystemContract systemContract;
    TestUniversalContract testUniversalContract;
    address owner;
    address addr1;
    address protocolAddress;
    RevertOptions revertOptions;
    RevertContext revertContext;
    AbortContext abortContext;

    event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message);
    event ContextDataRevert(RevertContext revertContext);
    event ContextDataAbort(AbortContext abortContext);

    error ZeroAddress();
    error EnforcedPause();
    error AccessControlUnauthorizedAccount(address account, bytes32 neededRole);

    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    uint256 constant MIN_GAS_LIMIT = 100_000;

    function setUp() public {
        owner = address(this);
        addr1 = address(0x1234);

        museToken = new WETH9();

        proxy = payable(
            Upgrades.deployUUPSProxy(
                "GatewayMEVM.sol", abi.encodeCall(GatewayMEVM.initialize, (address(museToken), owner))
            )
        );
        gateway = GatewayMEVM(proxy);

        protocolAddress = gateway.PROTOCOL_ADDRESS();

        testUniversalContract = new TestUniversalContract();

        vm.startPrank(protocolAddress);
        systemContract = new SystemContract(address(0), address(0), address(0));
        mrc20 = new MRC20("TOKEN", "TKN", 18, 1, CoinType.Gas, 0, address(systemContract), address(gateway));
        systemContract.setGasCoinMRC20(1, address(mrc20));
        systemContract.setGasPrice(1, 1);
        vm.deal(protocolAddress, 1_000_000_000);
        museToken.deposit{ value: 10 }();
        museToken.approve(address(gateway), 10);
        mrc20.deposit(owner, MIN_GAS_LIMIT);
        vm.stopPrank();

        vm.startPrank(owner);
        mrc20.approve(address(gateway), MIN_GAS_LIMIT);
        museToken.deposit{ value: 10 }();
        museToken.approve(address(gateway), 10);
        vm.stopPrank();

        revertContext = RevertContext({ sender: owner, asset: address(0), amount: 1, revertMessage: "" });
        abortContext = AbortContext({
            sender: abi.encodePacked(owner),
            asset: address(0),
            amount: 1,
            outgoing: false,
            chainID: 1,
            revertMessage: ""
        });
    }

    function testDepositFailsIfMRC20IsZeroAddress() public {
        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.deposit(address(0), 1, addr1);
    }

    function testDepositFailsIfTargetIsZeroAddress() public {
        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.deposit(address(mrc20), 1, address(0));
    }

    function testDepositFailsIfAmountIs0() public {
        vm.prank(protocolAddress);
        vm.expectRevert(InsufficientMRC20Amount.selector);
        gateway.deposit(address(mrc20), 0, addr1);
    }

    function testDeposit() public {
        uint256 amount = 1;
        uint256 balanceBefore = mrc20.balanceOf(addr1);
        assertEq(0, balanceBefore);

        vm.prank(protocolAddress);
        gateway.deposit(address(mrc20), amount, addr1);

        uint256 balanceAfter = mrc20.balanceOf(addr1);
        assertEq(amount, balanceAfter);
    }

    function testDepositTogglePause() public {
        vm.prank(protocolAddress);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, protocolAddress, PAUSER_ROLE));
        gateway.pause();

        vm.prank(protocolAddress);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, protocolAddress, PAUSER_ROLE));
        gateway.unpause();

        vm.prank(owner);
        gateway.pause();

        uint256 amount = 1;

        vm.expectRevert(EnforcedPause.selector);
        vm.prank(protocolAddress);
        gateway.deposit(address(mrc20), amount, addr1);

        vm.prank(owner);
        gateway.unpause();

        uint256 balanceBefore = mrc20.balanceOf(addr1);
        assertEq(0, balanceBefore);

        vm.prank(protocolAddress);
        gateway.deposit(address(mrc20), amount, addr1);

        uint256 balanceAfter = mrc20.balanceOf(addr1);
        assertEq(amount, balanceAfter);
    }

    function testDepositFailsIfSenderNotProtocol() public {
        uint256 amount = 1;
        uint256 balanceBefore = mrc20.balanceOf(addr1);
        assertEq(0, balanceBefore);

        vm.prank(owner);
        vm.expectRevert(CallerIsNotProtocol.selector);
        gateway.deposit(address(mrc20), amount, addr1);

        uint256 balanceAfter = mrc20.balanceOf(addr1);
        assertEq(0, balanceAfter);
    }

    function testDepositFailsIfTargetIsGateway() public {
        uint256 amount = 1;

        vm.prank(protocolAddress);
        vm.expectRevert(InvalidTarget.selector);
        gateway.deposit(address(mrc20), amount, address(gateway));
    }

    function testDepositFailsIfTargetIsProtocol() public {
        uint256 amount = 1;
        vm.prank(protocolAddress);
        vm.expectRevert(InvalidTarget.selector);
        gateway.deposit(address(mrc20), amount, protocolAddress);
    }

    function testExecuteFailsIfMRC20IsZeroAddress() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.execute(context, address(0), 1, address(testUniversalContract), message);
    }

    function testExecuteFailsIfTargetIsZeroAddress() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.execute(context, address(mrc20), 1, address(0), message);
    }

    function testExecuteUniversalContractFailsIfZeroAddress() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.execute(context, address(0), 1, address(testUniversalContract), message);
    }

    function testExecuteUniversalContract() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.expectEmit(true, true, true, true, address(testUniversalContract));
        emit ContextData(abi.encodePacked(gateway), protocolAddress, 1, address(gateway), "hello");
        vm.prank(protocolAddress);
        gateway.execute(context, address(mrc20), 1, address(testUniversalContract), message);
    }

    function testExecuteUniversalContractFailsIfSenderIsNotProtocol() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.expectRevert(CallerIsNotProtocol.selector);
        vm.prank(owner);
        gateway.execute(context, address(mrc20), 1, address(testUniversalContract), message);
    }

    function testExecuteRevertUniversalContractFailsIfTargetIsZeroAddress() public {
        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.executeRevert(address(0), revertContext);
    }

    function testExecuteRevertUniversalContract() public {
        vm.expectEmit(true, true, true, true, address(testUniversalContract));
        emit ContextDataRevert(revertContext);
        vm.prank(protocolAddress);
        gateway.executeRevert(address(testUniversalContract), revertContext);
    }

    function testExecuteRevertUniversalContractIfSenderIsNotProtocol() public {
        vm.expectRevert(CallerIsNotProtocol.selector);
        vm.prank(owner);
        gateway.executeRevert(address(testUniversalContract), revertContext);
    }

    function testDepositMRC20AndCallUniversalContractFailsIfMRC20IsZeroAddress() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndCall(context, address(0), 1, address(testUniversalContract), message);
    }

    function testDepositMRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndCall(context, address(mrc20), 1, address(0), message);
    }

    function testDepositMRC20AndCallUniversalContractFailsIfAmountIsZero() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(InsufficientMRC20Amount.selector);
        gateway.depositAndCall(context, address(mrc20), 0, address(testUniversalContract), message);
    }

    function testDepositMRC20AndCallUniversalContract() public {
        uint256 balanceBefore = mrc20.balanceOf(address(testUniversalContract));
        assertEq(0, balanceBefore);

        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.expectEmit(true, true, true, true, address(testUniversalContract));
        emit ContextData(abi.encodePacked(gateway), protocolAddress, 1, address(gateway), "hello");
        vm.prank(protocolAddress);
        gateway.depositAndCall(context, address(mrc20), 1, address(testUniversalContract), message);

        uint256 balanceAfter = mrc20.balanceOf(address(testUniversalContract));
        assertEq(1, balanceAfter);
    }

    function testDepositMRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.expectRevert(CallerIsNotProtocol.selector);
        vm.prank(owner);
        gateway.depositAndCall(context, address(mrc20), 1, address(testUniversalContract), message);
    }

    function testDepositMRC20AndCallUniversalContractIfTargetIsProtocol() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.expectRevert(InvalidTarget.selector);
        vm.prank(protocolAddress);
        gateway.depositAndCall(context, address(mrc20), 1, protocolAddress, message);
    }

    function testDepositMRC20AndCallUniversalContractIfTargetIsGateway() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.expectRevert(InvalidTarget.selector);
        vm.prank(protocolAddress);
        gateway.depositAndCall(context, address(mrc20), 1, address(gateway), message);
    }

    function testDepositAndRevertMRC20AndCallUniversalContractFailsIfMRC20IsZeroAddress() public {
        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndRevert(address(0), 1, address(testUniversalContract), revertContext);
    }

    function testDepositAndRevertMRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() public {
        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndRevert(address(mrc20), 1, address(0), revertContext);
    }

    function testDepositAndRevertMRC20AndCallUniversalContractFailsIfAmountIsZero() public {
        vm.prank(protocolAddress);
        vm.expectRevert(InsufficientMRC20Amount.selector);
        gateway.depositAndRevert(address(mrc20), 0, address(testUniversalContract), revertContext);
    }

    function testDepositAndRevertMRC20AndCallUniversalContract() public {
        uint256 balanceBefore = mrc20.balanceOf(address(testUniversalContract));
        assertEq(0, balanceBefore);

        vm.expectEmit(true, true, true, true, address(testUniversalContract));
        emit ContextDataRevert(revertContext);
        vm.prank(protocolAddress);
        gateway.depositAndRevert(address(mrc20), 1, address(testUniversalContract), revertContext);

        uint256 balanceAfter = mrc20.balanceOf(address(testUniversalContract));
        assertEq(1, balanceAfter);
    }

    function testDepositAndRevertMRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() public {
        vm.expectRevert(CallerIsNotProtocol.selector);
        vm.prank(owner);
        gateway.depositAndRevert(address(mrc20), 1, address(testUniversalContract), revertContext);
    }

    function testDepositAndRevertMRC20AndCallUniversalContractFailsITargetIsProtocol() public {
        vm.expectRevert(InvalidTarget.selector);
        vm.prank(protocolAddress);
        gateway.depositAndRevert(address(mrc20), 1, protocolAddress, revertContext);
    }

    function testDepositAndRevertMRC20AndCallUniversalContractFailsITargetIsGateway() public {
        vm.expectRevert(InvalidTarget.selector);
        vm.prank(protocolAddress);
        gateway.depositAndRevert(address(mrc20), 1, address(gateway), revertContext);
    }

    function testDepositMUSEAndCallUniversalContractFailsIfTargetIsZeroAddress() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndCall(context, 1, address(0), message);
    }

    function testDepositMUSEAndCallUniversalContractFailsIfTargetIsAmountIsZero() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(InsufficientMuseAmount.selector);
        gateway.depositAndCall(context, 0, address(mrc20), message);
    }

    function testDepositMUSEAndCallUniversalContractFailsIfZeroAddress() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndCall(context, 1, address(0), message);
    }

    function testDepositMUSEAndCallUniversal() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndCall(context, 1, address(0), message);
    }

    function testDepositMUSEAndCallUniversalContract() public {
        uint256 amount = 1;
        uint256 protocolBalanceBefore = museToken.balanceOf(protocolAddress);
        uint256 gatewayBalanceBefore = museToken.balanceOf(address(gateway));
        uint256 destinationBalanceBefore = address(testUniversalContract).balance;
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.expectEmit(true, true, true, true, address(testUniversalContract));
        emit ContextData(abi.encodePacked(gateway), protocolAddress, amount, address(gateway), "hello");
        vm.prank(protocolAddress);
        gateway.depositAndCall(context, amount, address(testUniversalContract), message);

        uint256 protocolBalanceAfter = museToken.balanceOf(protocolAddress);
        assertEq(protocolBalanceBefore - amount, protocolBalanceAfter);

        uint256 gatewayBalanceAfter = museToken.balanceOf(address(gateway));
        assertEq(gatewayBalanceBefore, gatewayBalanceAfter);

        // Verify amount is transfered to destination
        assertEq(destinationBalanceBefore + amount, address(testUniversalContract).balance);
    }

    function testDepositMUSEAndCallUniversalContractFailsIfSenderIsNotProtocol() public {
        uint256 amount = 1;
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.expectRevert(CallerIsNotProtocol.selector);
        vm.prank(owner);
        gateway.depositAndCall(context, amount, address(testUniversalContract), message);
    }

    function testDepositMUSEAndCallUniversalContractFailsIfTargetIsProtocol() public {
        uint256 amount = 1;
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.expectRevert(InvalidTarget.selector);
        vm.prank(protocolAddress);
        gateway.depositAndCall(context, amount, protocolAddress, message);
    }

    function testDepositMUSEAndCallUniversalContractFailsIfTargetIsGateway() public {
        uint256 amount = 1;
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.expectRevert(InvalidTarget.selector);
        vm.prank(protocolAddress);
        gateway.depositAndCall(context, amount, address(gateway), message);
    }

    function testExecuteAbortUniversalContract() public {
        vm.expectEmit(true, true, true, true, address(testUniversalContract));
        emit ContextDataAbort(abortContext);
        vm.prank(protocolAddress);
        gateway.executeAbort(address(testUniversalContract), abortContext);
    }

    function testExecuteAbortUniversalContractFailsIfTargetIsZeroAddress() public {
        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.executeAbort(address(0), abortContext);
    }
}
