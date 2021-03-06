/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
	"errors"
	"fmt"
	"strconv"
	"encoding/json"
	//"time"
	//"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type User struct{
	Name string `json:"Name"`
	Id int `json:"Id"`
}

type Marble struct{
	ID string `json:"ID"`
	Color string `json:"Color"`
	Size int `json:"Size"`
	cuser User `json:"cuser"`
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// ============================================================================================================================
// Init - reset all the things
// ============================================================================================================================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	var err error
	var MarbleA User

	MarbleA = User{"Parth", 125}
	jsonAsBytes, _ := json.Marshal(MarbleA)

	err = stub.PutState("A", jsonAsBytes)
	if err != nil{
		return nil, errors.New("asd")
	}

	return nil, nil
}

// ============================================================================================================================
// Invoke - Our entry point for Invocations
// ============================================================================================================================
func (t *SimpleChaincode) Invoke (stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error

	valAsbytes, err := stub.GetState("A")

	if err != nil{
		return nil, errors.New("asd")
	}
	var marbleIndex User
	json.Unmarshal(valAsbytes, &marbleIndex)	
	marbleIndexReturn := []byte(strconv.Itoa(marbleIndex.Id))
	//marbleIndexReturn := []byte(marbleIndex.Id)

	return marbleIndexReturn, nil
}

// ============================================================================================================================
// Query - Our entry point for Queries
// ============================================================================================================================
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error

	valAsbytes, err := stub.GetState("A")

	if err != nil{
		return nil, errors.New("asd")
	}
	var marbleIndex User
	json.Unmarshal(valAsbytes, &marbleIndex)
	marbleIndexReturn := []byte(strconv.Itoa(marbleIndex.Id))
	//marbleIndexReturn := []byte(marbleIndex.Id)

	return marbleIndexReturn, nil
}