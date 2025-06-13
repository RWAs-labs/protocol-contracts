// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "./utils/ReceiverEVM.sol";

import "../contracts/evm/ERC20Custody.sol";
import { GatewayEVM } from "../contracts/evm/GatewayEVM.sol";
import "../contracts/evm/MuseConnectorNonNative.sol";
import "./utils/TestERC20.sol";

import "./utils/SenderMEVM.sol";

import { SystemContractMock } from "./utils/SystemContractMock.sol";

import { GatewayMEVM } from "../contracts/mevm/GatewayMEVM.sol";
import { IGatewayMEVM } from "../contracts/mevm/GatewayMEVM.sol";
import { CallOptions, IGatewayMEVMErrors, IGatewayMEVMEvents } from "../contracts/mevm/interfaces/IGatewayMEVM.sol";

import { IGatewayEVMErrors } from "../contracts/evm/interfaces/IGatewayEVM.sol";
import { IGatewayEVMEvents } from "../contracts/evm/interfaces/IGatewayEVM.sol";

import "../contracts/mevm/MRC20.sol";

import "./utils/IReceiverEVM.sol";

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import { Upgrades } from "openzeppelin-foundry-upgrades/Upgrades.sol";

contract GatewayEVMMEVMTest is
    Test,
    IGatewayEVMErrors,
    IGatewayEVMEvents,
    IGatewayMEVMEvents,
    IGatewayMEVMErrors,
    IReceiverEVMEvents
{
    // evm
    using SafeERC20 for IERC20;

    GatewayEVM gatewayEVM;
    ERC20Custody custody;
    MuseConnectorNonNative museConnector;
    TestERC20 token;
    TestERC20 muse;
    ReceiverEVM receiverEVM;
    address ownerEVM;
    address destination;
    address tssAddress;
    MessageContext arbitraryCallMessageContext = MessageContext({ sender: address(0) });

    // mevm
    address payable proxyMEVM;
    GatewayMEVM gatewayMEVM;
    SenderMEVM senderMEVM;
    SystemContractMock systemContract;
    MRC20 mrc20;
    address ownerMEVM;

    CallOptions callOptions;
    RevertOptions revertOptions;

    function setUp() public {
        // evm
        ownerEVM = address(this);
        destination = address(0x1234);
        tssAddress = address(0x5678);
        ownerMEVM = address(0x4321);

        token = new TestERC20("test", "TTK");
        muse = new TestERC20("muse", "MUSE");

        address proxy = Upgrades.deployUUPSProxy(
            "GatewayEVM.sol", abi.encodeCall(GatewayEVM.initialize, (tssAddress, address(muse), ownerEVM))
        );
        gatewayEVM = GatewayEVM(proxy);
        proxy = Upgrades.deployUUPSProxy(
            "ERC20Custody.sol", abi.encodeCall(ERC20Custody.initialize, (address(gatewayEVM), tssAddress, ownerEVM))
        );
        custody = ERC20Custody(proxy);
        proxy = Upgrades.deployUUPSProxy(
            "MuseConnectorNonNative.sol",
            abi.encodeCall(
                MuseConnectorNonNative.initialize, (address(gatewayEVM), address(muse), tssAddress, ownerEVM)
            )
        );
        museConnector = MuseConnectorNonNative(proxy);

        vm.deal(tssAddress, 1 ether);

        vm.startPrank(ownerEVM);
        gatewayEVM.setCustody(address(custody));
        gatewayEVM.setConnector(address(museConnector));
        vm.stopPrank();

        token.mint(ownerEVM, 1_000_000);
        token.transfer(address(custody), 500_000);

        receiverEVM = new ReceiverEVM();

        // mevm
        proxyMEVM = payable(
            Upgrades.deployUUPSProxy(
                "GatewayMEVM.sol", abi.encodeCall(GatewayMEVM.initialize, (address(muse), ownerMEVM))
            )
        );
        gatewayMEVM = GatewayMEVM(proxyMEVM);

        senderMEVM = new SenderMEVM(address(gatewayMEVM));
        address protocolAddress = address(0x735b14BB79463307AAcBED86DAf3322B1e6226aB);
        vm.startPrank(protocolAddress);
        systemContract = new SystemContractMock(address(0), address(0), address(0));
        mrc20 = new MRC20("TOKEN", "TKN", 18, 1, CoinType.Muse, 0, address(systemContract), address(gatewayMEVM));
        systemContract.setGasCoinMRC20(1, address(mrc20));
        systemContract.setGasPrice(1, 1);
        mrc20.deposit(ownerMEVM, 1_000_000);
        mrc20.deposit(address(senderMEVM), 1_000_000);
        vm.stopPrank();

        vm.prank(ownerMEVM);
        mrc20.approve(address(gatewayMEVM), 1_000_000);

        vm.deal(tssAddress, 1 ether);

        callOptions = CallOptions({ gasLimit: 100_000, isArbitraryCall: false });

        revertOptions = RevertOptions({
            revertAddress: address(0x321),
            callOnRevert: true,
            abortAddress: address(0x321),
            revertMessage: "",
            onRevertGasLimit: 0
        });
    }

    function testCallReceiverEVMFromMEVM() public {
        string memory str = "Hello!";
        uint256 num = 42;
        bool flag = true;
        uint256 value = 1 ether;

        // Encode the function call data and call on mevm
        bytes memory message = abi.encodeWithSelector(receiverEVM.receivePayable.selector, str, num, flag);
        vm.prank(ownerMEVM);
        vm.expectEmit(true, true, true, true, address(gatewayMEVM));
        emit Called(
            address(ownerMEVM),
            address(mrc20),
            abi.encodePacked(receiverEVM),
            message,
            CallOptions({ gasLimit: 100_000, isArbitraryCall: true }),
            revertOptions
        );
        gatewayMEVM.call(
            abi.encodePacked(receiverEVM),
            address(mrc20),
            message,
            CallOptions({ gasLimit: 100_000, isArbitraryCall: true }),
            revertOptions
        );

        // Call execute on evm
        vm.deal(address(gatewayEVM), value);
        vm.expectEmit(true, true, true, true, address(gatewayEVM));
        emit Executed(address(receiverEVM), value, message);
        vm.prank(tssAddress);
        gatewayEVM.execute{ value: value }(arbitraryCallMessageContext, address(receiverEVM), message);
    }

    function testCallReceiverEVMFromSenderMEVM() public {
        string memory str = "Hello!";
        uint256 num = 42;
        bool flag = true;
        uint256 value = 1 ether;

        // Encode the function call data and call on mevm
        bytes memory message = abi.encodeWithSelector(receiverEVM.receivePayable.selector, str, num, flag);
        bytes memory data = abi.encodeWithSignature(
            "call(bytes,address,bytes,(uint256,bool),(address,bool,address,bytes,uint256))",
            abi.encodePacked(receiverEVM),
            address(mrc20),
            message,
            callOptions,
            revertOptions
        );
        vm.expectCall(address(gatewayMEVM), 0, data);
        vm.prank(ownerMEVM);
        senderMEVM.callReceiver(abi.encodePacked(receiverEVM), address(mrc20), str, num, flag);

        // Call execute on evm
        vm.deal(address(gatewayEVM), value);
        vm.expectEmit(true, true, true, true, address(receiverEVM));
        emit ReceivedPayable(address(gatewayEVM), value, str, num, flag);
        vm.expectEmit(true, true, true, true, address(gatewayEVM));
        emit Executed(address(receiverEVM), value, message);
        vm.prank(tssAddress);
        gatewayEVM.execute{ value: value }(arbitraryCallMessageContext, address(receiverEVM), message);
    }

    function testWithdrawAndCallReceiverEVMFromMEVM() public {
        string memory str = "Hello!";
        uint256 num = 42;
        bool flag = true;
        uint256 value = 1 ether;

        // Encode the function call data and call on mevm
        bytes memory message = abi.encodeWithSelector(receiverEVM.receivePayable.selector, str, num, flag);
        uint256 expectedGasFee = 100_000;
        vm.expectEmit(true, true, true, true, address(gatewayMEVM));
        emit WithdrawnAndCalled(
            ownerMEVM,
            0,
            abi.encodePacked(receiverEVM),
            address(mrc20),
            500_000,
            expectedGasFee,
            mrc20.PROTOCOL_FLAT_FEE(),
            message,
            CallOptions({ gasLimit: 100_000, isArbitraryCall: true }),
            revertOptions
        );
        vm.prank(ownerMEVM);
        gatewayMEVM.withdrawAndCall(
            abi.encodePacked(receiverEVM),
            500_000,
            address(mrc20),
            message,
            CallOptions({ gasLimit: 100_000, isArbitraryCall: true }),
            revertOptions
        );

        // Check the balance after withdrawal
        uint256 balanceOfAfterWithdrawal = mrc20.balanceOf(ownerMEVM);
        assertEq(balanceOfAfterWithdrawal, 500_000 - expectedGasFee);

        // Call execute on evm
        vm.deal(address(gatewayEVM), value);
        vm.expectEmit(true, true, true, true, address(receiverEVM));
        emit ReceivedPayable(address(gatewayEVM), value, str, num, flag);
        vm.expectEmit(true, true, true, true, address(gatewayEVM));
        emit Executed(address(receiverEVM), value, message);
        vm.prank(tssAddress);
        gatewayEVM.execute{ value: value }(arbitraryCallMessageContext, address(receiverEVM), message);
    }

    function testWithdrawAndCallReceiverEVMFromSenderMEVM() public {
        string memory str = "Hello!";
        uint256 num = 42;
        bool flag = true;
        uint256 value = 1 ether;

        // Encode the function call data and call on mevm
        uint256 senderBalanceBeforeWithdrawal = IMRC20(mrc20).balanceOf(address(senderMEVM));
        bytes memory message = abi.encodeWithSelector(receiverEVM.receivePayable.selector, str, num, flag);
        bytes memory data = abi.encodeWithSignature(
            "withdrawAndCall(bytes,uint256,address,bytes,(uint256,bool),(address,bool,address,bytes,uint256))",
            abi.encodePacked(receiverEVM),
            500_000,
            address(mrc20),
            message,
            callOptions,
            revertOptions
        );
        vm.expectCall(address(gatewayMEVM), 0, data);
        vm.prank(ownerMEVM);
        senderMEVM.withdrawAndCallReceiver(abi.encodePacked(receiverEVM), 500_000, address(mrc20), str, num, flag);

        // Call execute on evm
        vm.deal(address(gatewayEVM), value);
        vm.expectEmit(true, true, true, true, address(receiverEVM));
        emit ReceivedPayable(address(gatewayEVM), value, str, num, flag);
        vm.expectEmit(true, true, true, true, address(gatewayEVM));
        emit Executed(address(receiverEVM), value, message);
        vm.prank(tssAddress);
        gatewayEVM.execute{ value: value }(arbitraryCallMessageContext, address(receiverEVM), message);

        // Check the balance after withdrawal
        uint256 senderBalanceAfterWithdrawal = IMRC20(mrc20).balanceOf(address(senderMEVM));
        // Expected gas fee 1
        assertEq(senderBalanceAfterWithdrawal, senderBalanceBeforeWithdrawal - 500_000 - 100_000);
    }
}
