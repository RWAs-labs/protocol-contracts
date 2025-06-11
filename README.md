# MuseChain Protocol Contracts

Contracts of official protocol contracts deployed by the core MuseChain team.

## Codebase

The protocol contracts codebase is separated into two sections:

- `mevm`: contracts deployed on MuseChain
- `evm`: contracts deployed on EVM connected chains (Ethereum, Base, etc..)

**MEVM Contracts**

- `GatewayMEVM`: entrypoints for interaction. The users call the gateway contract to initiate cctx. The gateway contract is also the contract calling the target when a smart contract call is initiated on a connected chain
- `MRC20`: is a ERC20 compliant contract that represents fungible assets from connected chains
- `WMUSE`: wrapped MUSE (fork of WETH)

**EVM Contracts**

- `GatewayEVM`: similar to GatewayMEVM for connected chains. Entrypoint for users.
- `ERC20Custody`: hold ERC20 assets being sent to MuseChain
- `MuseConnector`: manage MUSE for connected chains. There are two version of the contract:
    - Native: when the MUSE tokens are native to the chain (MUSE where initially deployed as a ERC20 on Ethereum, not emitted fully on MuseChain). Use lock/unlock model.
    - Non-Native: when MUSE tokens where never native to the chains but withdrawn from MuseChain. Use mint/burn model.
- TSS (EOA): Threshold-signature-scheme wallet, this address holds the gas token sent to MuseChain (like Ethers)

## Usage

**Install dependencies**

```shell
$ yarn
$ forge soldeer update
```

**Build**

```shell
$ forge build
```

**Test**

```shell
$ forge test
```

**Format**

```shell
$ forge fmt
```

**Deploy**

```shell
$ forge script script/<DeployScript>.s.sol:<DeployScript> --rpc-url <your_rpc_url> --private-key <your_private_key>
```

To view detailed instructions on how to deploy the contracts, please refer to the [Deployment Guide](./scripts/deploy/readme.md).

This guide covers all steps required to deploy the contracts, including environment setup, commands, and best practices.
