// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "../contracts/mevm/SystemContract.sol";
import "./utils/WMUSE.sol";

import "../contracts/mevm/GatewayMEVM.sol";
import "../contracts/mevm/MRC20.sol";
import "../contracts/mevm/interfaces/IGatewayMEVM.sol";
import "../contracts/mevm/interfaces/IMRC20.sol";
import { Upgrades } from "openzeppelin-foundry-upgrades/Upgrades.sol";

contract MRC20Test is Test, MRC20Errors {
    MRC20 mrc20;
    SystemContract systemContract;
    GatewayMEVM gateway;
    address payable proxy;
    WETH9 museToken;

    address owner;
    address addr1;
    address protocolAddress;

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

        vm.startPrank(protocolAddress);
        systemContract = new SystemContract(address(0), address(0), address(0));
        mrc20 = new MRC20("TOKEN", "TKN", 18, 1, CoinType.Gas, 0, address(systemContract), address(gateway));
        systemContract.setGasCoinMRC20(1, address(mrc20));
        systemContract.setGasPrice(1, 1);
        vm.deal(protocolAddress, 1_000_000_000);
        vm.deal(proxy, 1_000_000_000);
        mrc20.deposit(owner, 100_000);
        vm.stopPrank();
    }

    function testMRC20BasicInfo() public {
        // silence view function warning
        mrc20 = mrc20;

        string memory name = mrc20.name();
        assertEq("TOKEN", name);

        string memory symbol = mrc20.symbol();
        assertEq("TKN", symbol);

        uint8 decimals = mrc20.decimals();
        assertEq(18, decimals);

        uint256 totalSupply = mrc20.totalSupply();
        assertEq(100_000, totalSupply);
    }

    function testUpdateNameAndSymbol() public {
        assertEq("TOKEN", mrc20.name());
        assertEq("TKN", mrc20.symbol());

        vm.prank(protocolAddress);
        mrc20.setName("TOKEN2");
        vm.prank(protocolAddress);
        mrc20.setSymbol("TKN2");

        assertEq("TOKEN2", mrc20.name());
        assertEq("TKN2", mrc20.symbol());
    }

    function testUpdateNameAndSymbolFailsIfSenderIsNotProtocol() public {
        vm.expectRevert(CallerIsNotFungibleModule.selector);
        mrc20.setName("TOKEN2");

        vm.expectRevert(CallerIsNotFungibleModule.selector);
        mrc20.setSymbol("TKN2");
    }

    function testTransfer() public {
        uint256 balanceStart = mrc20.balanceOf(addr1);
        assertEq(0, balanceStart);

        uint256 amount = 50_000;
        mrc20.transfer(addr1, amount);

        uint256 balanceAfter = mrc20.balanceOf(addr1);
        assertEq(amount, balanceAfter);
    }

    function testTransferFailsIfNoBalance() public {
        uint256 balanceStart = mrc20.balanceOf(addr1);
        assertEq(0, balanceStart);

        uint256 amount = 500_000;
        mrc20.approve(owner, amount);
        vm.expectRevert(LowBalance.selector);
        mrc20.transfer(addr1, amount);
    }

    function testTransferFailsIfRecipientIsZeroAddress() public {
        uint256 amount = 500_000;
        mrc20.approve(owner, amount);
        vm.expectRevert(ZeroAddress.selector);
        mrc20.transfer(address(0), amount);
    }

    function testTransferFrom() public {
        uint256 balanceStart = mrc20.balanceOf(addr1);
        assertEq(0, balanceStart);

        uint256 amount = 50_000;
        mrc20.approve(owner, amount);
        mrc20.transferFrom(owner, addr1, amount);

        uint256 balanceAfter = mrc20.balanceOf(addr1);
        assertEq(amount, balanceAfter);
    }

    function testTransferFromFailsIfNoAllowance() public {
        uint256 balanceStart = mrc20.balanceOf(addr1);
        assertEq(0, balanceStart);

        uint256 allowance = mrc20.allowance(owner, owner);
        assertEq(0, allowance);

        uint256 amount = 50_000;
        vm.expectRevert(LowAllowance.selector);
        mrc20.transferFrom(owner, addr1, amount);
    }

    function testTransferFromFailsIfNoBalance() public {
        uint256 balanceStart = mrc20.balanceOf(addr1);
        assertEq(0, balanceStart);

        uint256 amount = 500_000;
        mrc20.approve(owner, amount);
        vm.expectRevert(LowBalance.selector);
        mrc20.transferFrom(owner, addr1, amount);
    }

    function testTransferFromFailsIfSenderIsZeroAddress() public {
        uint256 balanceStart = mrc20.balanceOf(addr1);
        assertEq(0, balanceStart);

        uint256 amount = 500_000;
        mrc20.approve(owner, amount);
        vm.expectRevert(ZeroAddress.selector);
        mrc20.transferFrom(address(0), addr1, amount);
    }

    function testTransferFromFailsIfRecipientIsZeroAddress() public {
        uint256 amount = 500_000;
        mrc20.approve(owner, amount);
        vm.expectRevert(ZeroAddress.selector);
        mrc20.transferFrom(owner, address(0), amount);
    }

    function testBurn() public {
        uint256 balanceStart = mrc20.balanceOf(owner);
        assertEq(100_000, balanceStart);
        uint256 totalSupplyStart = mrc20.totalSupply();
        assertEq(100_000, totalSupplyStart);

        mrc20.burn(50_000);
        uint256 balanceAfter = mrc20.balanceOf(owner);
        assertEq(50_000, balanceAfter);
        uint256 totalSupplyAfter = mrc20.totalSupply();
        assertEq(50_000, totalSupplyAfter);
    }

    function testApproveFailsIfRecipientIsZeroAddress() public {
        vm.expectRevert(ZeroAddress.selector);
        mrc20.approve(address(0), 10);
    }

    function testBurnFailsIfNoBalance() public {
        vm.expectRevert(LowBalance.selector);
        mrc20.burn(150_000);
    }

    function testDeposit() public {
        uint256 totalSupplyStart = mrc20.totalSupply();
        assertEq(100_000, totalSupplyStart);

        vm.prank(proxy);
        mrc20.deposit(owner, 100_000);

        uint256 totalSupplyEnd = mrc20.totalSupply();
        assertEq(200_000, totalSupplyEnd);
    }

    function testDepositFailsIfRecipientIsZeroAddress() public {
        vm.prank(proxy);
        vm.expectRevert(ZeroAddress.selector);
        mrc20.deposit(address(0), 100_000);
    }

    function testWithdrawGasFee() public {
        vm.prank(protocolAddress);
        uint256 gasLimit = 10;
        uint256 protocolFlatFee = 10;

        mrc20.updateGasLimit(10);

        vm.prank(protocolAddress);
        mrc20.updateProtocolFlatFee(10);

        (address gasMRC20, uint256 gasFee) = mrc20.withdrawGasFee();
        assertEq(address(mrc20), gasMRC20);
        assertEq(gasLimit + protocolFlatFee, gasFee);
    }

    function testWithdrawGasFeeFailsIfGasCoinNotSetForChainId() public {
        vm.prank(protocolAddress);
        systemContract.setGasCoinMRC20(1, address(0));

        vm.expectRevert(ZeroGasCoin.selector);
        mrc20.withdrawGasFee();
    }

    function testWithdrawGasFeeFailsIfGasPriceNotSetForChainId() public {
        vm.prank(protocolAddress);
        systemContract.setGasPrice(1, 0);

        vm.expectRevert(ZeroGasPrice.selector);
        mrc20.withdrawGasFee();
    }

    function testWithdraw() public {
        uint256 gasLimit = 10;
        uint256 protocolFlatFee = 10;

        vm.prank(protocolAddress);
        mrc20.updateGasLimit(gasLimit);

        vm.prank(protocolAddress);
        mrc20.updateProtocolFlatFee(protocolFlatFee);

        uint256 balanceStart = mrc20.balanceOf(owner);
        assertEq(100_000, balanceStart);
        uint256 totalSupplyStart = mrc20.totalSupply();
        assertEq(100_000, totalSupplyStart);

        uint256 protocolAddressBalanceStart = mrc20.balanceOf(protocolAddress);

        mrc20.approve(address(mrc20), 50_000);
        mrc20.withdraw(abi.encodePacked(addr1), 50_000);

        uint256 protocolAddressBalanceAfter = mrc20.balanceOf(protocolAddress);
        assertEq(protocolAddressBalanceStart + gasLimit + protocolFlatFee, protocolAddressBalanceAfter);

        uint256 balanceAfter = mrc20.balanceOf(owner);
        assertEq(50_000 - gasLimit - protocolFlatFee, balanceAfter);
        uint256 totalSupplyAfter = mrc20.totalSupply();
        assertEq(50_000, totalSupplyAfter);
    }

    function testWithdrawFailsIfNoBalance() public {
        uint256 gasLimit = 10;
        uint256 protocolFlatFee = 100_000_000;

        vm.prank(protocolAddress);
        mrc20.updateGasLimit(gasLimit);

        vm.prank(protocolAddress);
        mrc20.updateProtocolFlatFee(protocolFlatFee);

        mrc20.approve(address(mrc20), 200_000_000);

        vm.expectRevert(LowBalance.selector);
        mrc20.withdraw(abi.encodePacked(addr1), 100);
    }

    function testWithdrawFailsIfNoAllowance() public {
        uint256 gasLimit = 10;
        uint256 protocolFlatFee = 10;

        vm.prank(protocolAddress);
        mrc20.updateGasLimit(gasLimit);

        vm.prank(protocolAddress);
        mrc20.updateProtocolFlatFee(protocolFlatFee);

        mrc20.approve(address(mrc20), 0);

        vm.expectRevert();
        mrc20.withdraw(abi.encodePacked(addr1), 1);
    }

    function testDepositFailsIfSenderIsNotGateway() public {
        vm.expectRevert(InvalidSender.selector);
        mrc20.deposit(owner, 100_000);
    }

    function testUpdateSystemContractAddress() public {
        vm.prank(protocolAddress);
        mrc20.updateSystemContractAddress(address(0x3211));
        assertEq(mrc20.SYSTEM_CONTRACT_ADDRESS(), address(0x3211));
    }

    function testUpdateSystemContractAddressFailsIfZeroAddress() public {
        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        mrc20.updateSystemContractAddress(address(0));
    }

    function testUpdateSystemContractAddressFailsIfSenderIsNotProtocol() public {
        vm.expectRevert(CallerIsNotFungibleModule.selector);
        mrc20.updateSystemContractAddress(address(0x3211));
    }

    function testUpdateGatewayAddress() public {
        vm.prank(protocolAddress);
        mrc20.updateGatewayAddress(address(0x3211));
        assertEq(mrc20.gatewayAddress(), address(0x3211));
    }

    function testUpdateGatewayAddressFailsIfZeroAddress() public {
        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        mrc20.updateGatewayAddress(address(0));
    }

    function testUpdateGatewayAddressFailsIfSenderIsNotProtocol() public {
        vm.expectRevert(CallerIsNotFungibleModule.selector);
        mrc20.updateGatewayAddress(address(0x3211));
    }

    function testUpdateGasLimit() public {
        vm.prank(protocolAddress);
        mrc20.updateGasLimit(10);
        assertEq(10, mrc20.GAS_LIMIT());
    }

    function testUpdateGasLimitFailsIfSenderIsNotProtocol() public {
        vm.expectRevert(CallerIsNotFungibleModule.selector);
        mrc20.updateGasLimit(10);
    }

    function testUpdateProtocolFlatFee() public {
        vm.prank(protocolAddress);
        mrc20.updateProtocolFlatFee(10);
        assertEq(10, mrc20.PROTOCOL_FLAT_FEE());
    }

    function testUpdateProtocolFlatFeeFailsIfSenderIsNotProtocol() public {
        vm.expectRevert(CallerIsNotFungibleModule.selector);
        mrc20.updateProtocolFlatFee(10);
    }
}
