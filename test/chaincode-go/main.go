/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
	student "chaincode-go/contracts"
)

func main() {
	studentSmartContract, err := contractapi.NewChaincode(&student.SmartContract{})
	if err != nil {
		log.Panicf("Error creating student chaincode: %v", err)
	}

	if err := studentSmartContract.Start(); err != nil {
		log.Panicf("Error starting student chaincode: %v", err)
	}
}
