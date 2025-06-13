// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "../../contracts/mevm/interfaces/IMRC20.sol";
import "../../contracts/mevm/interfaces/UniversalContract.sol";

interface SystemContractErrors {
    error CallerIsNotFungibleModule();

    error InvalidTarget();

    error CantBeIdenticalAddresses();

    error CantBeZeroAddress();
}

contract SystemContractMock is SystemContractErrors {
    mapping(uint256 => uint256) public gasPriceByChainId;
    mapping(uint256 => address) public gasCoinMRC20ByChainId;
    mapping(uint256 => address) public gasMusePoolByChainId;

    address public wMuseContractAddress;
    address public uniswapv2FactoryAddress;
    address public uniswapv2Router02Address;

    event SystemContractDeployed();
    event SetGasPrice(uint256, uint256);
    event SetGasCoin(uint256, address);
    event SetGasMusePool(uint256, address);
    event SetWMuse(address);

    constructor(address wmuse_, address uniswapv2Factory_, address uniswapv2Router02_) {
        wMuseContractAddress = wmuse_;
        uniswapv2FactoryAddress = uniswapv2Factory_;
        uniswapv2Router02Address = uniswapv2Router02_;
        emit SystemContractDeployed();
    }

    // fungible module updates the gas price oracle periodically
    function setGasPrice(uint256 chainID, uint256 price) external {
        gasPriceByChainId[chainID] = price;
        emit SetGasPrice(chainID, price);
    }

    function setGasCoinMRC20(uint256 chainID, address mrc20) external {
        gasCoinMRC20ByChainId[chainID] = mrc20;
        emit SetGasCoin(chainID, mrc20);
    }

    function setWMUSEContractAddress(address addr) external {
        wMuseContractAddress = addr;
        emit SetWMuse(wMuseContractAddress);
    }

    // returns sorted token addresses, used to handle return values from pairs sorted in this order
    function sortTokens(address tokenA, address tokenB) internal pure returns (address token0, address token1) {
        if (tokenA == tokenB) revert CantBeIdenticalAddresses();
        (token0, token1) = tokenA < tokenB ? (tokenA, tokenB) : (tokenB, tokenA);
        if (token0 == address(0)) revert CantBeZeroAddress();
    }

    function uniswapv2PairFor(address factory, address tokenA, address tokenB) public pure returns (address pair) {
        (address token0, address token1) = sortTokens(tokenA, tokenB);
        pair = address(
            uint160(
                uint256(
                    keccak256(
                        abi.encodePacked(
                            hex"ff",
                            factory,
                            keccak256(abi.encodePacked(token0, token1)),
                            hex"96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f" // init code hash
                        )
                    )
                )
            )
        );
    }

    function onCrossChainCall(address target, address mrc20, uint256 amount, bytes calldata message) external {
        mContext memory context = mContext({ sender: msg.sender, origin: "", chainID: block.chainid });
        IMRC20(mrc20).transfer(target, amount);
        mContract(target).onCrossChainCall(context, mrc20, amount, message);
    }
}
