/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"feedFactory/chaincode"
)

func main() {
	feedChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating feedFactory chaincode: %v", err)
	}

	if err := feedChaincode.Start(); err != nil {
		log.Panicf("Error starting feedFactory chaincode: %v", err)
	}
}
