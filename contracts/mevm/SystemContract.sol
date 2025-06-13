// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "./interfaces/IMRC20.sol";
import "./interfaces/UniversalContract.sol";

/**
 * @dev Custom errors for SystemContract
 */
interface SystemContractErrors {
    error CallerIsNotFungibleModule();
    error InvalidTarget();
    error CantBeIdenticalAddresses();
    error CantBeZeroAddress();
    error ZeroAddress();
}

/**
 * @dev The system contract it's called by the protocol to interact with the blockchain.
 * Also includes a lot of tools to make easier to interact with MuseChain.
 */
contract SystemContract is SystemContractErrors {
    /// @notice Map to know the gas price of each chain given a chain id.
    mapping(uint256 => uint256) public gasPriceByChainId;
    /// @notice Map to know the MRC20.sol address of a token given a chain id, ex mETH, mBNB etc.
    mapping(uint256 => address) public gasCoinMRC20ByChainId;
    // @dev: Map to know uniswap V2 pool of MUSE/MRC20.sol given a chain id. This refer to the build in uniswap deployed at
    // genesis.
    mapping(uint256 => address) public gasMusePoolByChainId;

    /// @notice Fungible address is always the same, it's on protocol level.
    address public constant FUNGIBLE_MODULE_ADDRESS = 0x735b14BB79463307AAcBED86DAf3322B1e6226aB;
    /// @notice Uniswap V2 addresses.
    address public immutable uniswapv2FactoryAddress;
    address public immutable uniswapv2Router02Address;
    /// @notice Address of the wrapped MUSE to interact with Uniswap V2.
    address public wMuseContractAddress;
    /// @notice Address of MEVM Muse Connector.
    address public museConnectorMEVMAddress;

    /// @notice Custom SystemContract errors.
    event SystemContractDeployed();
    event SetGasPrice(uint256, uint256);
    event SetGasCoin(uint256, address);
    event SetGasMusePool(uint256, address);
    event SetWMuse(address);
    event SetConnectorMEVM(address);

    /**
     * @dev Only fungible module can deploy a system contract.
     */
    constructor(address wmuse_, address uniswapv2Factory_, address uniswapv2Router02_) {
        if (msg.sender != FUNGIBLE_MODULE_ADDRESS) revert CallerIsNotFungibleModule();
        wMuseContractAddress = wmuse_;
        uniswapv2FactoryAddress = uniswapv2Factory_;
        uniswapv2Router02Address = uniswapv2Router02_;
        emit SystemContractDeployed();
    }

    /**
     * @dev Deposit foreign coins into MRC20.sol and call user specified contract on mEVM.
     * @param context, context data for deposit.
     * @param mrc20, mrc20 address for deposit.
     * @param amount, amount to deposit.
     * @param target, contract address to make a call after deposit.
     * @param message, calldata for a call.
     */
    function depositAndCall(
        mContext calldata context,
        address mrc20,
        uint256 amount,
        address target,
        bytes calldata message
    )
        external
    {
        if (msg.sender != FUNGIBLE_MODULE_ADDRESS) revert CallerIsNotFungibleModule();
        if (target == FUNGIBLE_MODULE_ADDRESS || target == address(this)) revert InvalidTarget();

        IMRC20(mrc20).deposit(target, amount);
        mContract(target).onCrossChainCall(context, mrc20, amount, message);
    }

    /**
     * @dev Sort token addresses lexicographically. Used to handle return values from pairs sorted in the order.
     * @param tokenA, tokenA address.
     * @param tokenB, tokenB address.
     * @return token0 token1, returns sorted token addresses,.
     */
    function sortTokens(address tokenA, address tokenB) internal pure returns (address token0, address token1) {
        if (tokenA == tokenB) revert CantBeIdenticalAddresses();
        (token0, token1) = tokenA < tokenB ? (tokenA, tokenB) : (tokenB, tokenA);
        if (token0 == address(0)) revert CantBeZeroAddress();
    }

    /**
     * @dev Calculates the CREATE2 address for a pair without making any external calls.
     * @param factory, factory address.
     * @param tokenA, tokenA address.
     * @param tokenB, tokenB address.
     * @return pair tokens pair address.
     */
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

    /**
     * @dev Fungible module updates the gas price oracle periodically.
     * @param chainID, chain id.
     * @param price, new gas price.
     */
    function setGasPrice(uint256 chainID, uint256 price) external {
        if (msg.sender != FUNGIBLE_MODULE_ADDRESS) revert CallerIsNotFungibleModule();
        gasPriceByChainId[chainID] = price;
        emit SetGasPrice(chainID, price);
    }

    /**
     * @dev Setter for gasCoinMRC20ByChainId map.
     * @param chainID, chain id.
     * @param mrc20, MRC20.sol address.
     */
    function setGasCoinMRC20(uint256 chainID, address mrc20) external {
        if (msg.sender != FUNGIBLE_MODULE_ADDRESS) revert CallerIsNotFungibleModule();
        gasCoinMRC20ByChainId[chainID] = mrc20;
        emit SetGasCoin(chainID, mrc20);
    }

    /**
     * @dev Set the pool wmuse/erc20 address.
     * @param chainID, chain id.
     * @param erc20, pair for uniswap wmuse/erc20.
     */
    function setGasMusePool(uint256 chainID, address erc20) external {
        if (msg.sender != FUNGIBLE_MODULE_ADDRESS) revert CallerIsNotFungibleModule();
        address pool = uniswapv2PairFor(uniswapv2FactoryAddress, wMuseContractAddress, erc20);
        gasMusePoolByChainId[chainID] = pool;
        emit SetGasMusePool(chainID, pool);
    }

    /**
     * @dev Setter for wrapped MUSE address.
     * @param addr, wmuse new address.
     */
    function setWMUSEContractAddress(address addr) external {
        if (msg.sender != FUNGIBLE_MODULE_ADDRESS) revert CallerIsNotFungibleModule();
        if (addr == address(0)) revert ZeroAddress();
        wMuseContractAddress = addr;
        emit SetWMuse(wMuseContractAddress);
    }

    /**
     * @dev Setter for museConnector MEVM Address
     * @param addr, muse connector new address.
     */
    function setConnectorMEVMAddress(address addr) external {
        if (msg.sender != FUNGIBLE_MODULE_ADDRESS) revert CallerIsNotFungibleModule();
        if (addr == address(0)) revert ZeroAddress();
        museConnectorMEVMAddress = addr;
        emit SetConnectorMEVM(museConnectorMEVMAddress);
    }
}
