// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "./utils/ReceiverEVM.sol";

import "./utils/TestERC20.sol";
import "./utils/upgrades/MuseConnectorNativeUpgradeTest.sol";

import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import { Upgrades } from "openzeppelin-foundry-upgrades/Upgrades.sol";

import "./utils/IReceiverEVM.sol";

import "../contracts/evm/ERC20Custody.sol";
import "../contracts/evm/GatewayEVM.sol";
import "../contracts/evm/MuseConnectorNative.sol";
import "../contracts/evm/interfaces/IGatewayEVM.sol";
import "../contracts/evm/interfaces/IMuseConnector.sol";

contract MuseConnectorNativeTest is
    Test,
    IGatewayEVMErrors,
    IGatewayEVMEvents,
    IReceiverEVMEvents,
    IMuseConnectorEvents
{
    using SafeERC20 for IERC20;

    GatewayEVM gateway;
    ReceiverEVM receiver;
    ERC20Custody custody;
    MuseConnectorNative museConnector;
    TestERC20 museToken;
    address owner;
    address destination;
    address tssAddress;
    address foo;
    RevertContext revertContext;
    MessageContext arbitraryCallMessageContext = MessageContext({ sender: address(0) });

    error EnforcedPause();
    error AccessControlUnauthorizedAccount(address account, bytes32 neededRole);

    event WithdrawnV2(address indexed to, uint256 amount);

    bytes32 public constant WITHDRAWER_ROLE = keccak256("WITHDRAWER_ROLE");
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    bytes32 public constant TSS_ROLE = keccak256("TSS_ROLE");
    bytes32 public constant DEFAULT_ADMIN_ROLE = 0x00;

    function setUp() public {
        owner = address(this);
        destination = address(0x1234);
        tssAddress = address(0x5678);
        foo = address(0x9876);

        museToken = new TestERC20("muse", "MUSE");

        address proxy = Upgrades.deployUUPSProxy(
            "GatewayEVM.sol", abi.encodeCall(GatewayEVM.initialize, (tssAddress, address(museToken), owner))
        );
        gateway = GatewayEVM(proxy);
        proxy = Upgrades.deployUUPSProxy(
            "ERC20Custody.sol", abi.encodeCall(ERC20Custody.initialize, (address(gateway), tssAddress, owner))
        );
        custody = ERC20Custody(proxy);
        proxy = Upgrades.deployUUPSProxy(
            "MuseConnectorNative.sol",
            abi.encodeCall(MuseConnectorNative.initialize, (address(gateway), address(museToken), tssAddress, owner))
        );
        museConnector = MuseConnectorNative(proxy);

        receiver = new ReceiverEVM();

        vm.deal(tssAddress, 1 ether);

        vm.startPrank(owner);
        gateway.setCustody(address(custody));
        gateway.setConnector(address(museConnector));
        vm.stopPrank();

        museToken.mint(address(museConnector), 5_000_000);

        vm.deal(tssAddress, 1 ether);
        revertContext = RevertContext({ sender: owner, asset: address(museToken), amount: 1, revertMessage: "" });
    }

    function testTSSUpgrade() public {
        address newTSSAddress = address(0x4321);

        bool newTSSAddressWithdrawerRole = museConnector.hasRole(WITHDRAWER_ROLE, newTSSAddress);
        assertFalse(newTSSAddressWithdrawerRole);
        bool newTSSAddressWhitelisterRole = museConnector.hasRole(TSS_ROLE, newTSSAddress);
        assertFalse(newTSSAddressWhitelisterRole);

        bool oldTSSAddressHasWithdrawerRole = museConnector.hasRole(WITHDRAWER_ROLE, tssAddress);
        assertTrue(oldTSSAddressHasWithdrawerRole);
        bool oldTSSAddressHasWhitelisterRole = museConnector.hasRole(TSS_ROLE, tssAddress);
        assertTrue(oldTSSAddressHasWhitelisterRole);

        vm.startPrank(owner);
        vm.expectEmit(true, true, true, true, address(museConnector));
        emit UpdatedMuseConnectorTSSAddress(tssAddress, newTSSAddress);
        museConnector.updateTSSAddress(newTSSAddress);
        assertEq(newTSSAddress, museConnector.tssAddress());

        newTSSAddressWithdrawerRole = museConnector.hasRole(WITHDRAWER_ROLE, newTSSAddress);
        assertTrue(newTSSAddressWithdrawerRole);
        newTSSAddressWhitelisterRole = museConnector.hasRole(TSS_ROLE, newTSSAddress);
        assertTrue(newTSSAddressWhitelisterRole);

        oldTSSAddressHasWithdrawerRole = museConnector.hasRole(WITHDRAWER_ROLE, tssAddress);
        assertFalse(oldTSSAddressHasWithdrawerRole);
        oldTSSAddressHasWhitelisterRole = museConnector.hasRole(TSS_ROLE, tssAddress);
        assertFalse(oldTSSAddressHasWhitelisterRole);
    }

    function testTSSUpgradeFailsIfSenderIsNotTSSUpdater() public {
        vm.startPrank(tssAddress);
        vm.expectRevert(
            abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, tssAddress, DEFAULT_ADMIN_ROLE)
        );
        museConnector.updateTSSAddress(owner);
    }

    function testTSSUpgradeFailsIfZeroAddress() public {
        vm.startPrank(owner);
        vm.expectRevert(ZeroAddress.selector);
        museConnector.updateTSSAddress(address(0));
    }

    function testWithdraw() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";
        uint256 balanceBefore = museToken.balanceOf(destination);
        assertEq(balanceBefore, 0);

        bytes memory data = abi.encodeWithSignature("transfer(address,uint256)", destination, amount);
        vm.expectCall(address(museToken), 0, data);
        vm.expectEmit(true, true, true, true, address(museConnector));
        emit Withdrawn(destination, amount);
        vm.prank(tssAddress);
        museConnector.withdraw(destination, amount, internalSendHash);
        uint256 balanceAfter = museToken.balanceOf(destination);
        assertEq(balanceAfter, amount);
    }

    function testWithdrawTogglePause() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";

        vm.prank(foo);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, foo, PAUSER_ROLE));
        museConnector.pause();

        vm.prank(foo);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, foo, PAUSER_ROLE));
        museConnector.unpause();

        vm.prank(tssAddress);
        museConnector.pause();

        vm.expectRevert(EnforcedPause.selector);
        vm.prank(tssAddress);
        museConnector.withdraw(destination, amount, internalSendHash);

        vm.prank(owner);
        museConnector.unpause();

        uint256 balanceBefore = museToken.balanceOf(destination);
        assertEq(balanceBefore, 0);

        bytes memory data = abi.encodeWithSignature("transfer(address,uint256)", destination, amount);
        vm.expectCall(address(museToken), 0, data);
        vm.expectEmit(true, true, true, true, address(museConnector));
        emit Withdrawn(destination, amount);
        vm.prank(tssAddress);
        museConnector.withdraw(destination, amount, internalSendHash);
        uint256 balanceAfter = museToken.balanceOf(destination);
        assertEq(balanceAfter, amount);
    }

    function testWithdrawFailsIfSenderIsNotWithdrawer() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, owner, WITHDRAWER_ROLE));
        museConnector.withdraw(destination, amount, internalSendHash);
    }

    function testWithdrawAndCallReceiveERC20() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";
        bytes memory data =
            abi.encodeWithSignature("receiveERC20(uint256,address,address)", amount, address(museToken), destination);
        uint256 balanceBefore = museToken.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeMuseConnector = museToken.balanceOf(address(museConnector));

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(gateway), amount);
        vm.expectCall(address(museToken), 0, transferData);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedERC20(address(gateway), amount, address(museToken), destination);
        vm.expectEmit(true, true, true, true, address(museConnector));
        emit WithdrawnAndCalled(address(receiver), amount, data);
        vm.prank(tssAddress);
        museConnector.withdrawAndCall(arbitraryCallMessageContext, address(receiver), amount, data, internalSendHash);

        // Verify that the tokens were transferred to the destination address
        uint256 balanceAfter = museToken.balanceOf(destination);
        assertEq(balanceAfter, amount);

        // Verify that muse connector doesn't get more tokens
        uint256 balanceAfterMuseConnector = museToken.balanceOf(address(museConnector));
        assertEq(balanceAfterMuseConnector, balanceBeforeMuseConnector - amount);

        // Verify that the approval was reset
        uint256 allowance = museToken.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = museToken.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testWithdrawAndCallReceiveOnCall() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";
        address sender = address(0x123);
        bytes memory message = bytes("1");
        uint256 balanceBefore = museToken.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeMuseConnector = museToken.balanceOf(address(museConnector));

        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedOnCall(sender, message);
        vm.expectEmit(true, true, true, true, address(museConnector));
        emit WithdrawnAndCalled(address(receiver), amount, message);
        vm.prank(tssAddress);
        museConnector.withdrawAndCall(
            MessageContext({ sender: sender }), address(receiver), amount, message, internalSendHash
        );

        // Verify that the no tokens were transferred to the destination address
        uint256 balanceAfter = museToken.balanceOf(destination);
        assertEq(balanceAfter, 0);

        // Verify that muse connector doesn't get more tokens
        uint256 balanceAfterMuseConnector = museToken.balanceOf(address(museConnector));
        assertEq(balanceAfterMuseConnector, balanceBeforeMuseConnector);

        // Verify that the approval was reset
        uint256 allowance = museToken.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = museToken.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testWithdrawAndCallReceiveOnCallTNotAllowedWithArbitraryCall() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";
        address sender = address(0x123);
        bytes memory message = bytes("1");
        bytes memory data = abi.encodeWithSignature("onCall((address),bytes)", sender, message);

        vm.expectRevert(NotAllowedToCallOnCall.selector);
        vm.prank(tssAddress);
        museConnector.withdrawAndCall(arbitraryCallMessageContext, address(receiver), amount, data, internalSendHash);
    }

    function testWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";
        bytes memory data =
            abi.encodeWithSignature("receiveERC20(uint256,address,address)", amount, address(museToken), destination);

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, owner, WITHDRAWER_ROLE));
        museConnector.withdrawAndCall(arbitraryCallMessageContext, address(receiver), amount, data, internalSendHash);
    }

    function testWithdrawAndCallReceiveNoParams() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";
        bytes memory data = abi.encodeWithSignature("receiveNoParams()");
        uint256 balanceBefore = museToken.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeMuseConnector = museToken.balanceOf(address(museConnector));

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(gateway), amount);
        vm.expectCall(address(museToken), 0, transferData);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedNoParams(address(gateway));
        vm.expectEmit(true, true, true, true, address(museConnector));
        emit WithdrawnAndCalled(address(receiver), amount, data);
        vm.prank(tssAddress);
        museConnector.withdrawAndCall(arbitraryCallMessageContext, address(receiver), amount, data, internalSendHash);

        // Verify that the no tokens were transferred to the destination address
        uint256 balanceAfter = museToken.balanceOf(destination);
        assertEq(balanceAfter, 0);

        // Verify that muse connector doesn't get more tokens
        uint256 balanceAfterMuseConnector = museToken.balanceOf(address(museConnector));
        assertEq(balanceAfterMuseConnector, balanceBeforeMuseConnector);

        // Verify that the approval was reset
        uint256 allowance = museToken.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = museToken.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testWithdrawAndCallReceiveERC20Partial() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";
        bytes memory data = abi.encodeWithSignature(
            "receiveERC20Partial(uint256,address,address)", amount, address(museToken), destination
        );
        uint256 balanceBefore = museToken.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeMuseConnector = museToken.balanceOf(address(museConnector));

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(gateway), amount);
        vm.expectCall(address(museToken), 0, transferData);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedERC20(address(gateway), amount / 2, address(museToken), destination);
        vm.expectEmit(true, true, true, true, address(museConnector));
        emit WithdrawnAndCalled(address(receiver), amount, data);
        vm.prank(tssAddress);
        museConnector.withdrawAndCall(arbitraryCallMessageContext, address(receiver), amount, data, internalSendHash);

        // Verify that the tokens were transferred to the destination address
        uint256 balanceAfter = museToken.balanceOf(destination);
        assertEq(balanceAfter, amount / 2);

        // Verify that muse connector doesn't get more tokens
        uint256 balanceAfterMuseConnector = museToken.balanceOf(address(museConnector));
        assertEq(balanceAfterMuseConnector, balanceBeforeMuseConnector - amount / 2);

        // Verify that the approval was reset
        uint256 allowance = museToken.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = museToken.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testWithdrawAndRevert() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";
        bytes memory data = abi.encodePacked("hello");
        uint256 balanceBefore = museToken.balanceOf(address(receiver));
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeMuseConnector = museToken.balanceOf(address(museConnector));

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(gateway), amount);
        vm.expectCall(address(museToken), 0, transferData);
        // Verify that onRevert callback was called
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedRevert(address(gateway), revertContext);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Reverted(address(receiver), address(museToken), amount, data, revertContext);
        vm.expectEmit(true, true, true, true, address(museConnector));
        emit WithdrawnAndReverted(address(receiver), amount, data, revertContext);
        vm.prank(tssAddress);
        museConnector.withdrawAndRevert(address(receiver), amount, data, internalSendHash, revertContext);

        // Verify that the tokens were transferred to the receiver address
        uint256 balanceAfter = museToken.balanceOf(address(receiver));
        assertEq(balanceAfter, amount);

        // Verify that muse connector doesn't get more tokens
        uint256 balanceAfterMuseConnector = museToken.balanceOf(address(museConnector));
        assertEq(balanceAfterMuseConnector, balanceBeforeMuseConnector - amount);

        // Verify that the approval was reset
        uint256 allowance = museToken.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = museToken.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testWithdrawAndRevertFailsIfSenderIsNotWithdrawer() public {
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";
        bytes memory data = abi.encodePacked("hello");

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, owner, WITHDRAWER_ROLE));
        museConnector.withdrawAndRevert(address(receiver), amount, data, internalSendHash, revertContext);
    }

    function testUpgradeAndWithdraw() public {
        // upgrade
        Upgrades.upgradeProxy(address(museConnector), "MuseConnectorNativeUpgradeTest.sol", "", owner);
        MuseConnectorNativeUpgradeTest museConnectorV2 = MuseConnectorNativeUpgradeTest(address(museConnector));
        // withdraw
        uint256 amount = 100_000;
        bytes32 internalSendHash = "";
        uint256 balanceBefore = museToken.balanceOf(destination);
        assertEq(balanceBefore, 0);

        bytes memory data = abi.encodeWithSignature("transfer(address,uint256)", destination, amount);
        vm.expectCall(address(museToken), 0, data);
        vm.expectEmit(true, true, true, true, address(museConnectorV2));
        emit WithdrawnV2(destination, amount);
        vm.prank(tssAddress);
        museConnectorV2.withdraw(destination, amount, internalSendHash);
        uint256 balanceAfter = museToken.balanceOf(destination);
        assertEq(balanceAfter, amount);
    }
}
