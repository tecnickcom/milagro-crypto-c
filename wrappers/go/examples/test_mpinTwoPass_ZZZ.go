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

	"github.com/miracl/amcl-go-wrapper"
)

var HASH_TYPE_MPIN = amcl.SHA256

func main() {
	// Assign the End-User an ID
	IDstr := "testUser@miracl.com"
	ID := []byte(IDstr)
	fmt.Printf("ID: ")
	fmt.Printf("%x\n\n", ID)

	// Epoch time in days
	date := amcl.Today()

	// PIN variable to create token
	PIN1 := -1
	// PIN variable to authenticate
	PIN2 := -1

	// Seed value for Random Number Generator
	seedHex := "ac4509d6"
	seed, err := hex.DecodeString(seedHex)
	if err != nil {
		fmt.Println("Error decoding seed value")
		return
	}
	rng := amcl.CreateCSPRNG(seed)

	// Generate Master Secret Share 1
	rtn, MS1 := amcl.RandomGenerate_ZZZ(&rng)
	if rtn != 0 {
		fmt.Println("RandomGenerate_ZZZ Error:", rtn)
		return
	}
	fmt.Printf("MS1: 0x")
	fmt.Printf("%x\n", MS1[:])

	// Destroy MS1
	defer amcl.CleanMemory(MS1[:])

	// Generate Master Secret Share 2
	rtn, MS2 := amcl.RandomGenerate_ZZZ(&rng)
	if rtn != 0 {
		fmt.Println("RandomGenerate_ZZZ Error:", rtn)
		return
	}
	fmt.Printf("MS2: 0x")
	fmt.Printf("%x\n", MS2[:])

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
	fmt.Printf("SS1: 0x")
	fmt.Printf("%x\n", SS1[:])

	// Destroy SS1
	defer amcl.CleanMemory(SS1[:])

	// Generate server secret share 2
	rtn, SS2 := amcl.GetServerSecret_ZZZ(MS2[:])
	if rtn != 0 {
		fmt.Println("GetServerSecret_ZZZ Error:", rtn)
		return
	}
	fmt.Printf("SS2: 0x")
	fmt.Printf("%x\n", SS2[:])

	// Destroy SS2
	defer amcl.CleanMemory(SS2[:])

	// Combine server secret shares
	rtn, SS := amcl.RecombineG2_ZZZ(SS1[:], SS2[:])
	if rtn != 0 {
		fmt.Println("RecombineG2_ZZZ(SS1, SS2) Error:", rtn)
		return
	}
	fmt.Printf("SS: 0x")
	fmt.Printf("%x\n", SS[:])

	// Destroy SS
	defer amcl.CleanMemory(SS[:])

	// Generate client secret share 1
	rtn, CS1 := amcl.GetClientSecret_ZZZ(MS1[:], HCID)
	if rtn != 0 {
		fmt.Println("GetClientSecret_ZZZ Error:", rtn)
		return
	}
	fmt.Printf("Client Secret Share CS1: 0x")
	fmt.Printf("%x\n", CS1[:])

	// Destroy CS1
	defer amcl.CleanMemory(CS1[:])

	// Generate client secret share 2
	rtn, CS2 := amcl.GetClientSecret_ZZZ(MS2[:], HCID)
	if rtn != 0 {
		fmt.Println("GetClientSecret_ZZZ Error:", rtn)
		return
	}
	fmt.Printf("Client Secret Share CS2: 0x")
	fmt.Printf("%x\n", CS2[:])

	// Destroy CS2
	defer amcl.CleanMemory(CS2[:])

	// Combine client secret shares
	CS := make([]byte, amcl.G1S_ZZZ)
	rtn, CS = amcl.RecombineG1_ZZZ(CS1[:], CS2[:])
	if rtn != 0 {
		fmt.Println("RecombineG1_ZZZ Error:", rtn)
		return
	}
	fmt.Printf("Client Secret CS: 0x")
	fmt.Printf("%x\n", CS[:])

	// Destroy CS
	defer amcl.CleanMemory(CS[:])

	// Generate time permit share 1
	rtn, TP1 := amcl.GetClientPermit_ZZZ(HASH_TYPE_MPIN, date, MS1[:], HCID)
	if rtn != 0 {
		fmt.Println("GetClientPermit_ZZZ Error:", rtn)
		return
	}
	fmt.Printf("TP1: 0x")
	fmt.Printf("%x\n", TP1[:])

	// Destroy TP1
	defer amcl.CleanMemory(TP1[:])

	// Generate time permit share 2
	rtn, TP2 := amcl.GetClientPermit_ZZZ(HASH_TYPE_MPIN, date, MS2[:], HCID)
	if rtn != 0 {
		fmt.Println("GetClientPermit_ZZZ Error:", rtn)
		return
	}
	fmt.Printf("TP2: 0x")
	fmt.Printf("%x\n", TP2[:])

	// Destroy TP2
	defer amcl.CleanMemory(TP2[:])

	// Combine time permit shares
	rtn, TP := amcl.RecombineG1_ZZZ(TP1[:], TP2[:])
	if rtn != 0 {
		fmt.Println("RecombineG1_ZZZ(TP1, TP2) Error:", rtn)
		return
	}

	// Destroy TP
	defer amcl.CleanMemory(TP[:])

	// Client extracts PIN1 from secret to create Token
	for PIN1 < 0 {
		fmt.Printf("Please enter PIN to create token: ")
		fmt.Scan(&PIN1)
	}

	rtn, TOKEN := amcl.ExtractPIN_ZZZ(HASH_TYPE_MPIN, ID[:], PIN1, CS[:])
	if rtn != 0 {
		fmt.Printf("FAILURE: EXTRACT_PIN rtn: %d\n", rtn)
		return
	}
	fmt.Printf("Client Token TK: 0x")
	fmt.Printf("%x\n", TOKEN[:])

	// Destroy TOKEN
	defer amcl.CleanMemory(TOKEN[:])

	//////   Client   //////

	for PIN2 < 0 {
		fmt.Printf("Please enter PIN to authenticate: ")
		fmt.Scan(&PIN2)
	}

	////// Client Pass 1 //////
	// Send U and UT to server
	var X [amcl.PGS_ZZZ]byte
	fmt.Printf("X: 0x")
	fmt.Printf("%x\n", X[:])
	rtn, XOut, SEC, U, UT := amcl.Client1_ZZZ(HASH_TYPE_MPIN, date, ID, &rng, X[:], PIN2, TOKEN[:], TP[:])
	if rtn != 0 {
		fmt.Printf("FAILURE: CLIENT rtn: %d\n", rtn)
		return
	}
	fmt.Printf("XOut: 0x")
	fmt.Printf("%x\n", XOut[:])

	// Destroy X
	defer amcl.CleanMemory(X[:])
	// Destroy XOut
	defer amcl.CleanMemory(XOut[:])
	// Destroy SEC
	defer amcl.CleanMemory(SEC[:])

	//////   Server Pass 1  //////
	/* Calculate H(ID) and H(T|H(ID)) (if time permits enabled), and maps them to points on the curve HID and HTID resp. */
	HID, HTID := amcl.Server1_ZZZ(HASH_TYPE_MPIN, date, ID)

	/* Send Y to Client */
	rtn, Y := amcl.RandomGenerate_ZZZ(&rng)
	if rtn != 0 {
		fmt.Println("RandomGenerate_ZZZ Error:", rtn)
		return
	}
	fmt.Printf("Y: 0x")
	fmt.Printf("%x\n", Y[:])

	// Destroy Y
	defer amcl.CleanMemory(Y[:])

	/* Client Second Pass: Inputs Client secret SEC, x and y. Outputs -(x+y)*SEC */
	rtn, V := amcl.Client2_ZZZ(XOut[:], Y[:], SEC[:])
	if rtn != 0 {
		fmt.Printf("FAILURE: CLIENT_2 rtn: %d\n", rtn)
	}

	// Destroy V
	defer amcl.CleanMemory(V[:])

	/* Server Second Pass. Inputs hashed client id, random Y, -(x+y)*SEC, xID and xCID and Server secret SST. E and F help Kangaroo_ZZZs to find error. */
	/* If PIN error not required, set E and F = null */

	rtn, E, F := amcl.Server2_ZZZ(date, HID[:], HTID[:], nil, Y[:], SS[:], U[:], UT[:], V[:], true)
	if rtn != 0 {
		fmt.Printf("FAILURE: Server2 rtn: %d\n", rtn)
	}
	fmt.Printf("HID: 0x")
	fmt.Printf("%x\n", HID[:])
	fmt.Printf("HTID: 0x")
	fmt.Printf("%x\n", HTID[:])

	if rtn != 0 {
		fmt.Printf("Authentication failed Error Code %d\n", rtn)
		err := amcl.Kangaroo_ZZZ(E[:], F[:])
		if err != 0 {
			fmt.Printf("PIN Error %d\n", err)
		}
		return
	}

	if rtn != 0 {
		fmt.Printf("Authentication failed Error Code %d\n", rtn)
		return
	} else {
		fmt.Printf("Authenticated ID: %s \n", IDstr)
	}
}
