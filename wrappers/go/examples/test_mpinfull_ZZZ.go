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

	amcl "github.com/miracl/amcl-go-wrapper"
)

var HASH_TYPE_MPIN = amcl.SHA256

func main() {
	// Assign the End-User an ID
	IDstr := "testUser@miracl.com"
	ID := []byte(IDstr)
	fmt.Printf("ID: ")
	fmt.Printf("%x\n\n", ID)

	// Epoch time in days
	date := 0
	fmt.Println("date: ", date)

	// Epoch time in seconds
	timeValue := amcl.GetTime()
	fmt.Println("timeValue: ", timeValue)

	// PIN variable to create token
	PIN1 := 0
	// PIN variable to authenticate
	PIN2 := 0

	// Seed value for Random Number Generator (RNG)
	seedHex := "ac4509d6"
	seed, err := hex.DecodeString(seedHex)
	if err != nil {
		fmt.Println("Error decoding seed value")
		return
	}
	rng := amcl.CreateCSPRNG(seed)

	// Message to sign
	var MESSAGE []byte
	// MESSAGE := []byte("test sign message")

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

	// Precomputation
	rtn, G1, G2 := amcl.Precompute_ZZZ(TOKEN[:], HCID)
	if rtn != 0 {
		fmt.Println("Precompute_ZZZ(TOKEN[:], HCID) Error:", rtn)
		return
	}

	// Destroy G2
	defer amcl.CleanMemory(G2[:])
	// Destroy G1
	defer amcl.CleanMemory(G1[:])

	for PIN2 < 0 {
		fmt.Printf("Please enter PIN to authenticate: ")
		fmt.Scan(&PIN2)
	}

	// Send U, V, timeValue and Message to server
	var X [amcl.PGS_ZZZ]byte
	fmt.Printf("X: 0x")
	fmt.Printf("%x\n", X[:])
	rtn, XOut, Y1, V, U, _ := amcl.Client_ZZZ(HASH_TYPE_MPIN, date, ID[:], &rng, X[:], PIN2, TOKEN[:], nil, MESSAGE[:], timeValue)
	if rtn != 0 {
		fmt.Printf("FAILURE: CLIENT rtn: %d\n", rtn)
		return
	}
	fmt.Printf("Y1: 0x")
	fmt.Printf("%x\n", Y1[:])
	fmt.Printf("XOut: 0x")
	fmt.Printf("%x\n", XOut[:])

	// Destroy Y1
	defer amcl.CleanMemory(Y1[:])
	// Destroy XOut
	defer amcl.CleanMemory(XOut[:])
	// Destroy V
	defer amcl.CleanMemory(V[:])

	// Send Z=r.ID to Server
	var R [amcl.PGS_ZZZ]byte
	fmt.Printf("R: 0x")
	fmt.Printf("%x\n", R[:])
	rtn, ROut, Z := amcl.GetG1Multiple_ZZZ(&rng, 1, R[:], HCID[:])
	fmt.Printf("ROut: 0x")
	fmt.Printf("%x\n", ROut[:])

	// Destroy R
	defer amcl.CleanMemory(R[:])
	// Destroy ROut
	defer amcl.CleanMemory(ROut[:])
	// Destroy Z
	defer amcl.CleanMemory(Z[:])

	//////   Server   //////
	rtn, HID, _, Y2, E, F := amcl.Server_ZZZ(HASH_TYPE_MPIN, date, timeValue, SS[:], U[:], nil, V[:], ID[:], nil, MESSAGE[:], true)
	if rtn != 0 {
		fmt.Printf("FAILURE: SERVER rtn: %d\n", rtn)
	}
	fmt.Printf("Y2: 0x")
	fmt.Printf("%x\n", Y2[:])
	fmt.Printf("HID: 0x")
	fmt.Printf("%x\n", HID[:])

	// Destroy Y2
	defer amcl.CleanMemory(Y2[:])
	// Destroy E
	defer amcl.CleanMemory(E[:])
	// Destroy F
	defer amcl.CleanMemory(F[:])

	if rtn != 0 {
		fmt.Printf("Authentication failed Error Code %d\n", rtn)
		err := amcl.Kangaroo_ZZZ(E[:], F[:])
		if err != 0 {
			fmt.Printf("PIN Error %d\n", err)
		}
		return
	} else {
		fmt.Printf("Authenticated ID: %s \n", IDstr)
	}

	// send T=w.ID to client
	var W [amcl.PGS_ZZZ]byte
	fmt.Printf("W: 0x")
	fmt.Printf("%x\n", W[:])
	rtn, WOut, T := amcl.GetG1Multiple_ZZZ(&rng, 0, W[:], HID[:])
	fmt.Printf("WOut: 0x")
	fmt.Printf("%x\n", WOut[:])
	fmt.Printf("T: 0x")
	fmt.Printf("%x\n", T[:])

	// Destroy W
	defer amcl.CleanMemory(W[:])
	// Destroy WOut
	defer amcl.CleanMemory(WOut[:])
	// Destroy T
	defer amcl.CleanMemory(T[:])

	// Hash all values
	HM := amcl.HashAll(HASH_TYPE_MPIN, hashBytes_ZZZ, HCID[:], U[:], nil, V[:], Y2[:], Z[:], T[:])

	// Destroy HM
	defer amcl.CleanMemory(HM[:])

	rtn, AES_KEY_SERVER := amcl.ServerKey_ZZZ(HASH_TYPE_MPIN, Z[:], SS[:], WOut[:], HM[:], HID[:], U[:], nil)
	fmt.Printf("RTN = %v Server Key =  %x\n", rtn, AES_KEY_SERVER[:])

	// Destroy AES_KEY_SERVER
	defer amcl.CleanMemory(AES_KEY_SERVER[:])

	rtn, AES_KEY_CLIENT := amcl.ClientKey_ZZZ(HASH_TYPE_MPIN, PIN2, G1[:], G2[:], ROut[:], XOut[:], HM[:], T[:])
	fmt.Printf("Client Key =  0x")
	fmt.Printf("%x\n", AES_KEY_CLIENT[:])

	// Destroy AES_KEY_CLIENT
	defer amcl.CleanMemory(AES_KEY_CLIENT[:])

	//////   Server   //////

	// Initialization vector
	IV := amcl.GenerateRandomByte(&rng, 12)
	fmt.Printf("IV: 0x")
	fmt.Printf("%x\n", IV[:])

	// Destroy IV
	defer amcl.CleanMemory(IV[:])

	// header
	HEADER := amcl.GenerateRandomByte(&rng, 16)
	fmt.Printf("HEADER: 0x")
	fmt.Printf("%x\n", HEADER[:])

	// Destroy HEADER
	defer amcl.CleanMemory(HEADER[:])

	// Input plaintext
	plaintextStr := "A test message"
	PLAINTEXT1 := []byte(plaintextStr)
	fmt.Printf("String to encrypt: %s \n", plaintextStr)
	fmt.Printf("PLAINTEXT1: 0x")
	fmt.Printf("%x\n", PLAINTEXT1[:])

	// Destroy PLAINTEXT1
	defer amcl.CleanMemory(PLAINTEXT1[:])

	// AES-GCM Encryption
	CIPHERTEXT, TAG1, err := amcl.AesGcmEncrypt(AES_KEY_SERVER[:], IV[:], HEADER[:], PLAINTEXT1[:])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("CIPHERTEXT:  0x")
	fmt.Printf("%x\n", CIPHERTEXT[:])
	fmt.Printf("TAG1:  0x")
	fmt.Printf("%x\n", TAG1[:])

	// Destroy CIPHERTEXT
	defer amcl.CleanMemory(CIPHERTEXT[:])
	// Destroy TAG1
	defer amcl.CleanMemory(TAG1[:])

	// Send IV, HEADER, CIPHERTEXT and TAG1 to client

	// AES-GCM Decryption
	PLAINTEXT2, TAG2, err := amcl.AesGcmDecrypt(AES_KEY_CLIENT[:], IV[:], HEADER[:], CIPHERTEXT[:])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("PLAINTEXT2:  0x")
	fmt.Printf("%x\n", PLAINTEXT2[:])
	fmt.Printf("TAG2:  0x")
	fmt.Printf("%x\n", TAG2[:])
	fmt.Printf("Decrypted string: %s \n", string(PLAINTEXT2))

	// Destroy PLAINTEXT2
	defer amcl.CleanMemory(PLAINTEXT2[:])
	// Destroy TAG2
	defer amcl.CleanMemory(TAG2[:])
}
