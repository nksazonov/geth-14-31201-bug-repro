package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	ecrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
)

var SIMULATED_CHAIN_ID = big.NewInt(1337)
var DEFAULT_FUNDING_AMOUNT = big.NewInt(0).SetUint64(1000000000000000000) // 1 ETH

func main() {
	deployerPvk, err := ecrypto.GenerateKey()
	if err != nil {
		fmt.Printf("failed to generate deployer private key: %s\n", err)
		return
	}

	opts, err := bind.NewKeyedTransactorWithChainID(deployerPvk, SIMULATED_CHAIN_ID)
	if err != nil {
		fmt.Printf("failed to create transactor: %s\n", err)
	}

	alloc := map[common.Address]types.Account{}
	alloc[opts.From] = types.Account{
		Balance: DEFAULT_FUNDING_AMOUNT,
	}

	// NOTE: creating with default config
	backend := simulated.NewBackend(alloc)

	_, _, _, err = DeployDummyPreShanghai(opts, backend.Client())
	if err != nil {
		fmt.Printf("failed to create tx to deploy dummy (pre shanghai) contract: %s\n", err)
		return
	}

	fmt.Println("created tx to deploy dummy (pre shanghai) contract")

	// NOTE: if this is enabled, the post shanghai contract will be deployed successfully!
	// backend.Commit()

	_, _, _, err = DeployDummyPostShanghai(opts, backend.Client())
	if err != nil {
		fmt.Printf("failed to create tx to deploy dummy (post shanghai) contract: %s\n", err)
		return
	}

	fmt.Println("created tx to deploy dummy (post shanghai) contract")
}
