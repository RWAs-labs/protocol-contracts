// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "../../contracts/mevm/interfaces/ISystem.sol";
import "../../contracts/mevm/interfaces/IMRC20.sol";

/**
 * @dev Custom errors for MRC20.sol
 */
interface MRC20Errors {
    // @dev: Error thrown when caller is not the fungible module
    error CallerIsNotFungibleModule();
    error InvalidSender();
    error GasFeeTransferFailed();
    error ZeroGasCoin();
    error ZeroGasPrice();
    error LowAllowance();
    error LowBalance();
    error ZeroAddress();
}

// NOTE: this is exactly the same as MRC20.sol, except gateway contract address is set at deployment
// and used to allow deposit. This is first version, it might change in the future.
contract MRC20 is IMRC20Metadata, MRC20Errors, MRC20Events {
    /// @notice Fungible address is always the same, maintained at the protocol level
    address public constant FUNGIBLE_MODULE_ADDRESS = 0x735b14BB79463307AAcBED86DAf3322B1e6226aB;
    /// @notice Chain id.abi
    uint256 public immutable CHAIN_ID;
    /// @notice Coin type, checkout Interfaces.sol.
    CoinType public immutable COIN_TYPE;
    /// @notice System contract address.
    /// @dev Name is in upper case to maintain compatibility with MRC20.sol v1
    address public SYSTEM_CONTRACT_ADDRESS;
    /// @notice Gas limit.
    /// @dev Name is in upper case to maintain compatibility with MRC20.sol v1
    uint256 public GAS_LIMIT;
    /// @notice Protocol flat fee.
    /// @dev Name is in upper case to maintain compatibility with MRC20.sol v1
    uint256 public override PROTOCOL_FLAT_FEE;

    mapping(address => uint256) private _balances;
    mapping(address => mapping(address => uint256)) private _allowances;
    uint256 private _totalSupply;
    string private _name;
    string private _symbol;
    uint8 private _decimals;

    /// @notice Gateway contract address.
    /// @dev This variable is added at last position to maintain storage layout with MRC20.sol.sol v1
    address public gatewayAddress;

    function _msgSender() internal view virtual returns (address) {
        return msg.sender;
    }

    /**
     * @dev Only fungible module modifier.
     */
    modifier onlyFungible() {
        if (msg.sender != FUNGIBLE_MODULE_ADDRESS) revert CallerIsNotFungibleModule();
        _;
    }

    /**
     * @dev The only one allowed to deploy new MRC20.sol is fungible address.
     */
    constructor(
        string memory name_,
        string memory symbol_,
        uint8 decimals_,
        uint256 chainid_,
        CoinType coinType_,
        uint256 gasLimit_,
        address systemContractAddress_,
        address gatewayAddress_
    ) {
        if (systemContractAddress_ == address(0) || gatewayAddress_ == address(0)) revert ZeroAddress();
        _name = name_;
        _symbol = symbol_;
        _decimals = decimals_;
        CHAIN_ID = chainid_;
        COIN_TYPE = coinType_;
        GAS_LIMIT = gasLimit_;
        SYSTEM_CONTRACT_ADDRESS = systemContractAddress_;
        gatewayAddress = gatewayAddress_;
    }

    /**
     * @dev MRC20.sol name
     * @return name as string
     */
    function name() public view virtual override returns (string memory) {
        return _name;
    }

    /**
     * @dev Name can be updated by fungible module account.
     */
    function setName(string memory newName) external override onlyFungible {
        _name = newName;
    }

    /**
     * @dev Symbol can be updated by fungible module account.
     */
    function setSymbol(string memory newSymbol) external override onlyFungible {
        _symbol = newSymbol;
    }

    /**
     * @dev MRC20.sol symbol.
     * @return symbol as string.
     */
    function symbol() public view virtual override returns (string memory) {
        return _symbol;
    }

    /**
     * @dev MRC20.sol decimals.
     * @return returns uint8 decimals.
     */
    function decimals() public view virtual override returns (uint8) {
        return _decimals;
    }

    /**
     * @dev MRC20.sol total supply.
     * @return returns uint256 total supply.
     */
    function totalSupply() public view virtual override returns (uint256) {
        return _totalSupply;
    }

    /**
     * @dev Returns MRC20.sol balance of an account.
     * @param account, account address for which balance is requested.
     * @return uint256 account balance.
     */
    function balanceOf(address account) public view virtual override returns (uint256) {
        return _balances[account];
    }

    /**
     * @dev Returns MRC20.sol balance of an account.
     * @param recipient, recipiuent address to which transfer is done.
     * @return true/false if transfer succeeded/failed.
     */
    function transfer(address recipient, uint256 amount) public virtual override returns (bool) {
        _transfer(_msgSender(), recipient, amount);
        return true;
    }

    /**
     * @dev Returns token allowance from owner to spender.
     * @param owner, owner address.
     * @return uint256 allowance.
     */
    function allowance(address owner, address spender) public view virtual override returns (uint256) {
        return _allowances[owner][spender];
    }

    /**
     * @dev Approves amount transferFrom for spender.
     * @param spender, spender address.
     * @param amount, amount to approve.
     * @return true/false if succeeded/failed.
     */
    function approve(address spender, uint256 amount) public virtual override returns (bool) {
        _approve(_msgSender(), spender, amount);
        return true;
    }

    /**
     * @dev Transfers tokens from sender to recipient.
     * @param sender, sender address.
     * @param recipient, recipient address.
     * @param amount, amount to transfer.
     * @return true/false if succeeded/failed.
     */
    function transferFrom(address sender, address recipient, uint256 amount) public virtual override returns (bool) {
        _transfer(sender, recipient, amount);

        uint256 currentAllowance = _allowances[sender][_msgSender()];
        if (currentAllowance < amount) revert LowAllowance();

        _approve(sender, _msgSender(), currentAllowance - amount);

        return true;
    }

    /**
     * @dev Burns an amount of tokens.
     * @param amount, amount to burn.
     * @return true/false if succeeded/failed.
     */
    function burn(uint256 amount) external override returns (bool) {
        _burn(msg.sender, amount);
        return true;
    }

    function _transfer(address sender, address recipient, uint256 amount) internal virtual {
        if (sender == address(0)) revert ZeroAddress();
        if (recipient == address(0)) revert ZeroAddress();

        uint256 senderBalance = _balances[sender];
        if (senderBalance < amount) revert LowBalance();

        _balances[sender] = senderBalance - amount;
        _balances[recipient] += amount;

        emit Transfer(sender, recipient, amount);
    }

    function _mint(address account, uint256 amount) internal virtual {
        if (account == address(0)) revert ZeroAddress();

        _totalSupply += amount;
        _balances[account] += amount;
        emit Transfer(address(0), account, amount);
    }

    function _burn(address account, uint256 amount) internal virtual {
        if (account == address(0)) revert ZeroAddress();

        uint256 accountBalance = _balances[account];
        if (accountBalance < amount) revert LowBalance();

        _balances[account] = accountBalance - amount;
        _totalSupply -= amount;

        emit Transfer(account, address(0), amount);
    }

    function _approve(address owner, address spender, uint256 amount) internal virtual {
        if (owner == address(0)) revert ZeroAddress();
        if (spender == address(0)) revert ZeroAddress();

        _allowances[owner][spender] = amount;
        emit Approval(owner, spender, amount);
    }

    /**
     * @dev Deposits corresponding tokens from external chain, only callable by Fungible module.
     * @param to, recipient address.
     * @param amount, amount to deposit.
     * @return true/false if succeeded/failed.
     */
    function deposit(address to, uint256 amount) external override returns (bool) {
        if (
            msg.sender != FUNGIBLE_MODULE_ADDRESS && msg.sender != SYSTEM_CONTRACT_ADDRESS
                && msg.sender != gatewayAddress
        ) revert InvalidSender();
        _mint(to, amount);
        emit Deposit(abi.encodePacked(FUNGIBLE_MODULE_ADDRESS), to, amount);
        return true;
    }

    /**
     * @dev Withdraws gas fees.
     * @return returns the MRC20.sol address for gas on the same chain of this MRC20.sol, and calculates the gas fee for
     * withdraw()
     */
    function withdrawGasFee() public view override returns (address, uint256) {
        address gasMRC20 = ISystem(SYSTEM_CONTRACT_ADDRESS).gasCoinMRC20ByChainId(CHAIN_ID);
        if (gasMRC20 == address(0)) revert ZeroGasCoin();

        uint256 gasPrice = ISystem(SYSTEM_CONTRACT_ADDRESS).gasPriceByChainId(CHAIN_ID);
        if (gasPrice == 0) {
            revert ZeroGasPrice();
        }
        uint256 gasFee = gasPrice * GAS_LIMIT + PROTOCOL_FLAT_FEE;
        return (gasMRC20, gasFee);
    }

    /**
     * @dev Withdraws gas fees with specified gasLimit
     * @return returns the MRC20.sol address for gas on the same chain of this MRC20.sol, and calculates the gas fee for
     * withdraw()
     */
    function withdrawGasFeeWithGasLimit(uint256 gasLimit) public view override returns (address, uint256) {
        address gasMRC20 = ISystem(SYSTEM_CONTRACT_ADDRESS).gasCoinMRC20ByChainId(CHAIN_ID);
        if (gasMRC20 == address(0)) revert ZeroGasCoin();

        uint256 gasPrice = ISystem(SYSTEM_CONTRACT_ADDRESS).gasPriceByChainId(CHAIN_ID);
        if (gasPrice == 0) {
            revert ZeroGasPrice();
        }
        uint256 gasFee = gasPrice * gasLimit + PROTOCOL_FLAT_FEE;
        return (gasMRC20, gasFee);
    }

    /**
     * @dev Withraws MRC20.sol tokens to external chains, this function causes cctx module to send out outbound tx to the
     * outbound chain
     * this contract should be given enough allowance of the gas MRC20.sol to pay for outbound tx gas fee.
     * @param to, recipient address.
     * @param amount, amount to deposit.
     * @return true/false if succeeded/failed.
     */
    function withdraw(bytes memory to, uint256 amount) external override returns (bool) {
        (address gasMRC20, uint256 gasFee) = withdrawGasFee();
        if (!IMRC20(gasMRC20).transferFrom(msg.sender, FUNGIBLE_MODULE_ADDRESS, gasFee)) {
            revert GasFeeTransferFailed();
        }
        _burn(msg.sender, amount);
        emit Withdrawal(msg.sender, to, amount, gasFee, PROTOCOL_FLAT_FEE);
        return true;
    }

    /**
     * @dev Updates system contract address. Can only be updated by the fungible module.
     * @param addr, new system contract address.
     */
    function updateSystemContractAddress(address addr) external onlyFungible {
        if (addr == address(0)) revert ZeroAddress();
        SYSTEM_CONTRACT_ADDRESS = addr;
        emit UpdatedSystemContract(addr);
    }

    /**
     * @dev Updates gateway contract address. Can only be updated by the fungible module.
     * @param addr, new gateway contract address.
     */
    function updateGatewayAddress(address addr) external onlyFungible {
        if (addr == address(0)) revert ZeroAddress();
        gatewayAddress = addr;
        emit UpdatedGateway(addr);
    }

    /**
     * @dev Updates gas limit. Can only be updated by the fungible module.
     * @param gasLimit_, new gas limit.
     */
    function updateGasLimit(uint256 gasLimit_) external onlyFungible {
        GAS_LIMIT = gasLimit_;
        emit UpdatedGasLimit(GAS_LIMIT);
    }

    /**
     * @dev Updates protocol flat fee. Can only be updated by the fungible module.
     * @param protocolFlatFee_, new protocol flat fee.
     */
    function updateProtocolFlatFee(uint256 protocolFlatFee_) external onlyFungible {
        PROTOCOL_FLAT_FEE = protocolFlatFee_;
        emit UpdatedProtocolFlatFee(PROTOCOL_FLAT_FEE);
    }
}
