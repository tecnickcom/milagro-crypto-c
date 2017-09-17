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
	"encoding/hex"
	"fmt"
	"sync"
	"time"

	"github.com/miracl/amcl-go-wrapper"
)

const numRoutines = 1000

var (
	HASH_TYPE_MPIN = amcl.SHA256
	wg             sync.WaitGroup
)

func run(rng *amcl.RandNG) {

	// Assign the End-User an ID
	IDstr := "testUser@miracl.com"
	ID := []byte(IDstr)

	// Epoch time in days
	date := amcl.Today()

	// Epoch time in seconds
	timeValue := amcl.GetTime()

	// PIN variable to create token
	PIN := 1111

	// Message to sign
	MESSAGE := []byte("test sign message")

	// Generate Master Secret Share 1
	rtn, MS1 := amcl.RandomGenerate_ZZZ(rng)
	if rtn != 0 {
		fmt.Println("RandomGenerate_ZZZ Error:", rtn)
		return
	}
	// Destroy MS1
	defer amcl.CleanMemory(MS1[:])

	// Generate Master Secret Share 2
	rtn, MS2 := amcl.RandomGenerate_ZZZ(rng)
	if rtn != 0 {
		fmt.Println("RandomGenerate_ZZZ Error:", rtn)
		return
	}
	// Destroy MS2
	defer amcl.CleanMemory(MS2[:])

	// Either Client or TA calculates Hash(ID)
	HCID := amcl.HashId(HASH_TYPE_MPIN, hashBytes_ZZZ, ID)

	// Generate server secret share 1
	rtn, SS1 := amcl.GetServerSecret_ZZZ(MS1[:])
	if rtn != 0 {
		fmt.Println("GetServerSecret_ZZZ Error:", rtn)
		return
	}
	// Destroy SS1
	defer amcl.CleanMemory(SS1[:])

	// Generate server secret share 2
	rtn, SS2 := amcl.GetServerSecret_ZZZ(MS2[:])
	if rtn != 0 {
		fmt.Println("GetServerSecret_ZZZ Error:", rtn)
		return
	}
	// Destroy SS2
	defer amcl.CleanMemory(SS2[:])

	// Combine server secret shares
	rtn, SS := amcl.RecombineG2_ZZZ(SS1[:], SS2[:])
	if rtn != 0 {
		fmt.Println("RecombineG2_ZZZ(SS1, SS2) Error:", rtn)
		return
	}
	// Destroy SS
	defer amcl.CleanMemory(SS[:])

	// Generate client secret share 1
	rtn, CS1 := amcl.GetClientSecret_ZZZ(MS1[:], HCID)
	if rtn != 0 {
		fmt.Println("GetClientSecret_ZZZ Error:", rtn)
		return
	}
	// Destroy CS1
	defer amcl.CleanMemory(CS1[:])

	// Generate client secret share 2
	rtn, CS2 := amcl.GetClientSecret_ZZZ(MS2[:], HCID)
	if rtn != 0 {
		fmt.Println("GetClientSecret_ZZZ Error:", rtn)
		return
	}
	// Destroy CS2
	defer amcl.CleanMemory(CS2[:])

	// Combine client secret shares
	CS := make([]byte, amcl.G1S_ZZZ)
	rtn, CS = amcl.RecombineG1_ZZZ(CS1[:], CS2[:])
	if rtn != 0 {
		fmt.Println("RecombineG1_ZZZ Error:", rtn, SS, CS)
		return
	}
	// Destroy CS
	defer amcl.CleanMemory(CS[:])

	// Generate time permit share 1
	rtn, TP1 := amcl.GetClientPermit_ZZZ(HASH_TYPE_MPIN, date, MS1[:], HCID)
	if rtn != 0 {
		fmt.Println("GetClientPermit_ZZZ Error:", rtn)
		return
	}
	// Destroy TP1
	defer amcl.CleanMemory(TP1[:])

	// Generate time permit share 2
	rtn, TP2 := amcl.GetClientPermit_ZZZ(HASH_TYPE_MPIN, date, MS2[:], HCID)
	if rtn != 0 {
		fmt.Println("GetClientPermit_ZZZ Error:", rtn)
		return
	}
	// Destroy TP2
	defer amcl.CleanMemory(TP2[:])

	// Combine time permit shares
	rtn, TP := amcl.RecombineG1_ZZZ(TP1[:], TP2[:])
	if rtn != 0 {
		fmt.Println("RecombineG1_ZZZ(TP1, TP2) Error:", rtn, TP)
		return
	}
	// Destroy TP
	defer amcl.CleanMemory(TP[:])

	rtn, TOKEN := amcl.ExtractPIN_ZZZ(HASH_TYPE_MPIN, ID[:], PIN, CS[:])
	if rtn != 0 {
		fmt.Printf("FAILURE: EXTRACT_PIN rtn: %d\n", rtn)
		return
	}
	// Destroy TOKEN
	defer amcl.CleanMemory(TOKEN[:])

	// --- Client ---
	// Send U, UT, V, timeValue and Message to server
	var X [amcl.PGS_ZZZ]byte
	rtn, _, _, SEC, U, UT := amcl.Client_ZZZ(HASH_TYPE_MPIN, date, ID[:], rng, X[:], PIN, TOKEN[:], TP[:], MESSAGE[:], timeValue)
	if rtn != 0 {
		fmt.Printf("FAILURE: CLIENT rtn: %d\n", rtn, SEC, U, UT)
		return
	}

	// --- Server ---
	rtn, _, _, _, E, F := amcl.Server_ZZZ(HASH_TYPE_MPIN, date, timeValue, SS[:], U[:], UT[:], SEC[:], ID[:], nil, MESSAGE[:], true)
	if rtn != 0 {
		fmt.Printf("FAILURE: SERVER rtn: %d\n", rtn)
	}
	if rtn != 0 {
		fmt.Printf("Authentication failed Error Code %d\n", rtn)
		err := amcl.Kangaroo_ZZZ(E[:], F[:])
		if err != 0 {
			fmt.Printf("PIN Error %d\n", err)
		}
		return
	}

	wg.Done()
}

func main() {
	// Seed value for Random Number Generator (RNG)
	seedHex := "ac4509d6"
	seed, err := hex.DecodeString(seedHex)
	if err != nil {
		fmt.Println("Error decoding seed value")
		return
	}
	rng := amcl.CreateCSPRNG(seed)

	wg = sync.WaitGroup{}

	fmt.Printf("Stating %v go routines...\n", numRoutines)
	wg.Add(numRoutines)
	t := time.Now()
	for x := 0; x < numRoutines; x++ {
		go run(&rng)
	}
	wg.Wait()

	fmt.Printf("Done: %v \n", time.Now().Sub(t).Seconds())
}
