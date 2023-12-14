/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"cryptoFungi/chaincode"
)

func main() {
	fungiChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating cypytofungi chaincode: %v", err)
	}

	if err := fungiChaincode.Start(); err != nil {
		log.Panicf("Error starting cypytofungi chaincode: %v", err)
	}
}
