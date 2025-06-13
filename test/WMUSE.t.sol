pragma solidity ^0.8.0;

import "../contracts/mevm/WMUSE.sol";
import "forge-std/Test.sol";

contract WMUSETest is Test {
    WETH9 wmuse;

    address alice = address(0xabc);
    address bob = address(0xdef);

    function setUp() public {
        wmuse = new WETH9();
        vm.deal(alice, 10 ether);
        vm.deal(bob, 10 ether);
    }

    function testDeposit() public {
        vm.startPrank(alice);
        wmuse.deposit{ value: 1 ether }();

        assertEq(wmuse.balanceOf(alice), 1 ether);
        assertEq(wmuse.totalSupply(), 1 ether);

        vm.stopPrank();
    }

    function testWithdraw() public {
        vm.startPrank(alice);
        wmuse.deposit{ value: 2 ether }();

        wmuse.withdraw(1 ether);

        assertEq(wmuse.balanceOf(alice), 1 ether);
        assertEq(alice.balance, 9 ether);
        assertEq(wmuse.totalSupply(), 1 ether);

        vm.stopPrank();
    }

    function testTransfer() public {
        vm.startPrank(alice);
        wmuse.deposit{ value: 2 ether }();

        wmuse.transfer(bob, 1 ether);

        assertEq(wmuse.balanceOf(alice), 1 ether);
        assertEq(wmuse.balanceOf(bob), 1 ether);

        vm.stopPrank();
    }

    function testApproveAndTransferFrom() public {
        vm.startPrank(alice);
        wmuse.deposit{ value: 2 ether }();

        wmuse.approve(bob, 1 ether);
        assertEq(wmuse.allowance(alice, bob), 1 ether);

        vm.stopPrank();

        vm.startPrank(bob);
        wmuse.transferFrom(alice, bob, 1 ether);

        assertEq(wmuse.balanceOf(alice), 1 ether);
        assertEq(wmuse.balanceOf(bob), 1 ether);
        assertEq(wmuse.allowance(alice, bob), 0);

        vm.stopPrank();
    }

    function testDepositAndReceiveFallback() public {
        vm.prank(alice);
        (bool success,) = address(wmuse).call{ value: 1 ether }("");
        assertTrue(success);

        assertEq(wmuse.balanceOf(alice), 1 ether);
        assertEq(wmuse.totalSupply(), 1 ether);
    }

    function testWithdrawRevertsIfInsufficientBalance() public {
        vm.startPrank(alice);
        wmuse.deposit{ value: 1 ether }();

        vm.expectRevert();
        wmuse.withdraw(2 ether);

        vm.stopPrank();
    }

    function testTransferRevertsIfInsufficientBalance() public {
        vm.startPrank(alice);
        wmuse.deposit{ value: 1 ether }();

        vm.expectRevert();
        wmuse.transfer(bob, 2 ether);

        vm.stopPrank();
    }

    function testTransferFromRevertsIfInsufficientAllowance() public {
        vm.startPrank(alice);
        wmuse.deposit{ value: 1 ether }();

        wmuse.approve(bob, 0.5 ether);

        vm.stopPrank();

        vm.startPrank(bob);
        vm.expectRevert();
        wmuse.transferFrom(alice, bob, 1 ether);

        vm.stopPrank();
    }

    function testTotalSupply() public {
        vm.startPrank(alice);
        wmuse.deposit{ value: 1 ether }();

        vm.stopPrank();

        vm.startPrank(bob);
        wmuse.deposit{ value: 2 ether }();

        vm.stopPrank();

        assertEq(wmuse.totalSupply(), 3 ether);
    }
}
