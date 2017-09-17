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

package amcl

import (
	//"crypto/rand"
	"encoding/hex"
	"fmt"
	//mathrand "math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyEscrowLess_ZZZ(t *testing.T) {
	want := 0
	var got int

	// Assign the End-User an ID
	IDstr := "testUser@miracl.com"
	ID := []byte(IDstr)

	// Epoch time in days
	date := 0

	// Epoch time in seconds
	timeValue := 1439465203

	// PIN variable to create token
	PIN1 := 1234
	// PIN variable to authenticate
	PIN2 := 1234

	// Seed value for Random Number Generator (RNG)
	seedHex := "ac4509d6"
	seed, err := hex.DecodeString(seedHex)
	if err != nil {
		fmt.Println("Error decoding seed value")
		return
	}
	rng := CreateCSPRNG(seed)

	// Generate Master Secret Share 1
	_, MS1 := RandomGenerate_ZZZ(&rng)

	// Destroy MS1
	defer CleanMemory(MS1[:])

	// Generate Master Secret Share 2
	_, MS2 := RandomGenerate_ZZZ(&rng)

	// Destroy MS2
	defer CleanMemory(MS2[:])

	// Generate Public Key
	_, Z := RandomGenerate_ZZZ(&rng)
	_, _, Pa := GetDVSKeyPair_ZZZ(nil, Z[:])

	// Destroy Z
	defer CleanMemory(Z[:])

	// Compute ID|Pa
	ID = append(ID, Pa...)

	// Either Client or TA calculates Hash(ID)
	HCID := HashId(HASH_TYPE_MPIN, hashBytes_ZZZ, ID)

	// Generate server secret share 1
	_, SS1 := GetServerSecret_ZZZ(MS1[:])

	// Destroy SS1
	defer CleanMemory(SS1[:])

	// Generate server secret share 2
	_, SS2 := GetServerSecret_ZZZ(MS2[:])

	// Destroy SS2
	defer CleanMemory(SS2[:])

	// Combine server secret shares
	_, SS := RecombineG2_ZZZ(SS1[:], SS2[:])

	// Destroy SS
	defer CleanMemory(SS[:])

	// Generate client secret share 1
	_, CS1 := GetClientSecret_ZZZ(MS1[:], HCID)

	// Destroy CS1
	defer CleanMemory(CS1[:])

	// Generate client secret share 2
	_, CS2 := GetClientSecret_ZZZ(MS2[:], HCID)

	// Destroy CS2
	defer CleanMemory(CS2[:])

	// Combine client secret shares
	CS := make([]byte, G1S_ZZZ)
	_, CS = RecombineG1_ZZZ(CS1[:], CS2[:])

	// Compute key-escrow less secret
	_, _, CS = GetG1Multiple_ZZZ(nil, 0, Z[:], CS[:])

	// Destroy CS
	defer CleanMemory(CS[:])

	// Create token
	_, TOKEN := ExtractPIN_ZZZ(HASH_TYPE_MPIN, ID[:], PIN1, CS[:])

	// Destroy TOKEN
	defer CleanMemory(TOKEN[:])

	// Send U, UT, V, timeValue and Message to server
	var X [PGS_ZZZ]byte
	_, _, _, V, U, _ := Client_ZZZ(HASH_TYPE_MPIN, date, ID[:], &rng, X[:], PIN2, TOKEN[:], nil, nil, timeValue)

	// Destroy X
	defer CleanMemory(X[:])

	got, _, _, _, _, _ = Server_ZZZ(HASH_TYPE_MPIN, date, timeValue, SS[:], U[:], nil, V[:], ID[:], Pa, nil, false)

	assert.Equal(t, want, got, "Should be equal")
}

func TestKeyEscrowLessRandom_ZZZ(t *testing.T) {
	want := 0
	var got int

	// Assign the End-User an ID
	IDstr := "testUser@miracl.com"
	ID := []byte(IDstr)

	// Epoch time in days
	date := 0

	// Epoch time in seconds
	timeValue := 1439465203

	// PIN variable to create token
	PIN1 := 1234
	// PIN variable to authenticate
	PIN2 := 1234

	// Seed value for Random Number Generator (RNG)
	seedHex := "ac4509d6"
	seed, err := hex.DecodeString(seedHex)
	if err != nil {
		fmt.Println("Error decoding seed value")
		return
	}
	rng := CreateCSPRNG(seed)

	// Generate Master Secret Share 1
	_, MS1 := RandomGenerate_ZZZ(&rng)

	// Destroy MS1
	defer CleanMemory(MS1[:])

	// Generate Master Secret Share 2
	_, MS2 := RandomGenerate_ZZZ(&rng)

	// Destroy MS2
	defer CleanMemory(MS2[:])

	// Generate Public Key
	_, Z, Pa := GetDVSKeyPair_ZZZ(&rng, nil)

	// Destroy Z
	defer CleanMemory(Z[:])

	// Compute ID|Pa
	ID = append(ID, Pa...)

	// Either Client or TA calculates Hash(ID)
	HCID := HashId(HASH_TYPE_MPIN, hashBytes_ZZZ, ID)

	// Generate server secret share 1
	_, SS1 := GetServerSecret_ZZZ(MS1[:])

	// Destroy SS1
	defer CleanMemory(SS1[:])

	// Generate server secret share 2
	_, SS2 := GetServerSecret_ZZZ(MS2[:])

	// Destroy SS2
	defer CleanMemory(SS2[:])

	// Combine server secret shares
	_, SS := RecombineG2_ZZZ(SS1[:], SS2[:])

	// Destroy SS
	defer CleanMemory(SS[:])

	// Generate client secret share 1
	_, CS1 := GetClientSecret_ZZZ(MS1[:], HCID)

	// Destroy CS1
	defer CleanMemory(CS1[:])

	// Generate client secret share 2
	_, CS2 := GetClientSecret_ZZZ(MS2[:], HCID)

	// Destroy CS2
	defer CleanMemory(CS2[:])

	// Combine client secret shares
	CS := make([]byte, G1S_ZZZ)
	_, CS = RecombineG1_ZZZ(CS1[:], CS2[:])

	// Compute key-escrow less secret
	_, _, CS = GetG1Multiple_ZZZ(nil, 0, Z[:], CS[:])

	// Destroy CS
	defer CleanMemory(CS[:])

	// Create token
	_, TOKEN := ExtractPIN_ZZZ(HASH_TYPE_MPIN, ID[:], PIN1, CS[:])

	// Destroy TOKEN
	defer CleanMemory(TOKEN[:])

	// Send U, UT, V, timeValue and Message to server
	var X [PGS_ZZZ]byte
	_, _, _, V, U, _ := Client_ZZZ(HASH_TYPE_MPIN, date, ID[:], &rng, X[:], PIN2, TOKEN[:], nil, nil, timeValue)

	// Destroy X
	defer CleanMemory(X[:])

	// Authenticate
	got, _, _, _, _, _ = Server_ZZZ(HASH_TYPE_MPIN, date, timeValue, SS[:], U[:], nil, V[:], ID[:], Pa, nil, false)

	assert.Equal(t, want, got, "Should be equal")
}

func TestKeyEscrowWrongPK_ZZZ(t *testing.T) {
	want := -19
	var got int

	// Assign the End-User an ID
	IDstr := "testUser@miracl.com"
	ID := []byte(IDstr)

	// Epoch time in days
	date := 16660

	// Epoch time in seconds
	timeValue := 1439465203

	// PIN variable to create token
	PIN1 := 1234
	// PIN variable to authenticate
	PIN2 := 1234

	// Seed value for Random Number Generator (RNG)
	seedHex := "ac4509d6"
	seed, err := hex.DecodeString(seedHex)
	if err != nil {
		fmt.Println("Error decoding seed value")
		return
	}
	rng := CreateCSPRNG(seed)

	// Generate Master Secret Share 1
	_, MS1 := RandomGenerate_ZZZ(&rng)

	// Destroy MS1
	defer CleanMemory(MS1[:])

	// Generate Master Secret Share 2
	_, MS2 := RandomGenerate_ZZZ(&rng)

	// Destroy MS2
	defer CleanMemory(MS2[:])

	// Generate wrong Public Key
	_, Z, _ := GetDVSKeyPair_ZZZ(&rng, nil)
	_, _, Pa := GetDVSKeyPair_ZZZ(&rng, nil)

	// Destroy Z
	defer CleanMemory(Z[:])

	// Compute ID|Pa
	ID = append(ID, Pa...)

	// Either Client or TA calculates Hash(ID)
	HCID := HashId(HASH_TYPE_MPIN, hashBytes_ZZZ, ID)

	// Generate server secret share 1
	_, SS1 := GetServerSecret_ZZZ(MS1[:])

	// Destroy SS1
	defer CleanMemory(SS1[:])

	// Generate server secret share 2
	_, SS2 := GetServerSecret_ZZZ(MS2[:])

	// Destroy SS2
	defer CleanMemory(SS2[:])

	// Combine server secret shares
	_, SS := RecombineG2_ZZZ(SS1[:], SS2[:])

	// Destroy SS
	defer CleanMemory(SS[:])

	// Generate client secret share 1
	_, CS1 := GetClientSecret_ZZZ(MS1[:], HCID)

	// Destroy CS1
	defer CleanMemory(CS1[:])

	// Generate client secret share 2
	_, CS2 := GetClientSecret_ZZZ(MS2[:], HCID)

	// Destroy CS2
	defer CleanMemory(CS2[:])

	// Combine client secret shares
	CS := make([]byte, G1S_ZZZ)
	_, CS = RecombineG1_ZZZ(CS1[:], CS2[:])

	// Compute key-escrow less secret
	_, _, CS = GetG1Multiple_ZZZ(nil, 0, Z[:], CS[:])

	// Destroy CS
	defer CleanMemory(CS[:])

	// Generate time permit share 1
	_, TP1 := GetClientPermit_ZZZ(HASH_TYPE_MPIN, date, MS1[:], HCID)

	// Destroy TP1
	defer CleanMemory(TP1[:])

	// Generate time permit share 2
	_, TP2 := GetClientPermit_ZZZ(HASH_TYPE_MPIN, date, MS2[:], HCID)

	// Destroy TP2
	defer CleanMemory(TP2[:])

	// Combine time permit shares
	_, TP := RecombineG1_ZZZ(TP1[:], TP2[:])

	// Destroy TP
	defer CleanMemory(TP[:])

	// Create token
	_, TOKEN := ExtractPIN_ZZZ(HASH_TYPE_MPIN, ID[:], PIN1, CS[:])

	// Destroy TOKEN
	defer CleanMemory(TOKEN[:])

	// Send U, UT, V, timeValue and Message to server
	var X [PGS_ZZZ]byte
	_, _, _, V, U, UT := Client_ZZZ(HASH_TYPE_MPIN, date, ID[:], &rng, X[:], PIN2, TOKEN[:], TP[:], nil, timeValue)

	// Destroy X
	defer CleanMemory(X[:])

	timeValue += 10
	// Authenticate
	got, _, _, _, _, _ = Server_ZZZ(HASH_TYPE_MPIN, date, timeValue, SS[:], U[:], UT[:], V[:], ID[:], Pa, nil, false)

	assert.Equal(t, want, got, "Should be equal")
}

func TestKeyEscrowLessTwoPassWrongPK_ZZZ(t *testing.T) {
	want := -19
	var got int

	// Assign the End-User an ID
	IDstr := "testUser@miracl.com"
	ID := []byte(IDstr)

	// Epoch time in seconds
	timeValue := 1439465203

	// PIN variable to create token
	PIN1 := 1234
	// PIN variable to authenticate
	PIN2 := 1234

	// Seed value for Random Number Generator (RNG)
	seedHex := "ac4509d6"
	seed, err := hex.DecodeString(seedHex)
	if err != nil {
		fmt.Println("Error decoding seed value")
		return
	}
	rng := CreateCSPRNG(seed)

	// Generate Master Secret Share 1
	_, MS1 := RandomGenerate_ZZZ(&rng)

	// Destroy MS1
	defer CleanMemory(MS1[:])

	// Generate Master Secret Share 2
	_, MS2 := RandomGenerate_ZZZ(&rng)

	// Destroy MS2
	defer CleanMemory(MS2[:])

	// Generate wrong Public Key
	_, Z := RandomGenerate_ZZZ(&rng)
	_, _, Pa := GetDVSKeyPair_ZZZ(&rng, nil)

	// Destroy Z
	defer CleanMemory(Z[:])

	// Compute ID|Pa
	ID = append(ID, Pa...)

	// Either Client or TA calculates Hash(ID)
	HCID := HashId(HASH_TYPE_MPIN, hashBytes_ZZZ, ID)

	// Generate server secret share 1
	_, SS1 := GetServerSecret_ZZZ(MS1[:])

	// Destroy SS1
	defer CleanMemory(SS1[:])

	// Generate server secret share 2
	_, SS2 := GetServerSecret_ZZZ(MS2[:])

	// Destroy SS2
	defer CleanMemory(SS2[:])

	// Combine server secret shares
	_, SS := RecombineG2_ZZZ(SS1[:], SS2[:])

	// Destroy SS
	defer CleanMemory(SS[:])

	// Generate client secret share 1
	_, CS1 := GetClientSecret_ZZZ(MS1[:], HCID)

	// Destroy CS1
	defer CleanMemory(CS1[:])

	// Generate client secret share 2
	_, CS2 := GetClientSecret_ZZZ(MS2[:], HCID)

	// Destroy CS2
	defer CleanMemory(CS2[:])

	// Combine client secret shares
	CS := make([]byte, G1S_ZZZ)
	_, CS = RecombineG1_ZZZ(CS1[:], CS2[:])

	// Compute key-escrow less secret
	_, _, CS = GetG1Multiple_ZZZ(nil, 0, Z[:], CS[:])

	// Destroy CS
	defer CleanMemory(CS[:])

	// Create token
	_, TOKEN := ExtractPIN_ZZZ(HASH_TYPE_MPIN, ID[:], PIN1, CS[:])

	// Destroy TOKEN
	defer CleanMemory(TOKEN[:])

	timeValue += 10
	// Client Pass 1
	var X [PGS_ZZZ]byte
	_, XOut, SEC, U, _ := Client1_ZZZ(HASH_TYPE_MPIN, 0, ID, &rng, X[:], PIN2, TOKEN[:], nil)

	// Destroy XOut
	defer CleanMemory(XOut[:])
	// Destroy SEC
	defer CleanMemory(SEC[:])

	// Server Pass 1
	var HID []byte
	HID, _ = Server1_ZZZ(HASH_TYPE_MPIN, 0, ID)

	_, Y := RandomGenerate_ZZZ(&rng)

	// Destroy HID
	defer CleanMemory(HID[:])
	// Destroy Y
	defer CleanMemory(Y[:])

	// Client Pass 2
	_, V := Client2_ZZZ(XOut[:], Y[:], SEC[:])

	// Server Pass 2
	// Send UT as V to model bad token
	got, _, _ = Server2_ZZZ(0, HID[:], nil, Pa, Y[:], SS[:], U[:], nil, V[:], false)
	assert.Equal(t, want, got, "Should be equal")
}

func TestKeyEscrowLessTwoPASS_ZZZ(t *testing.T) {
	want := 0
	var got int

	// Assign the End-User an ID
	IDstr := "testUser@miracl.com"
	ID := []byte(IDstr)

	// Epoch time in seconds
	timeValue := 1439465203

	// PIN variable to create token
	PIN1 := 1234
	// PIN variable to authenticate
	PIN2 := 1234

	// Seed value for Random Number Generator (RNG)
	seedHex := "ac4509d6"
	seed, err := hex.DecodeString(seedHex)
	if err != nil {
		fmt.Println("Error decoding seed value")
		return
	}
	rng := CreateCSPRNG(seed)

	// Generate Master Secret Share 1
	_, MS1 := RandomGenerate_ZZZ(&rng)

	// Destroy MS1
	defer CleanMemory(MS1[:])

	// Generate Master Secret Share 2
	_, MS2 := RandomGenerate_ZZZ(&rng)

	// Destroy MS2
	defer CleanMemory(MS2[:])

	// Generate Public Key
	_, Z := RandomGenerate_ZZZ(&rng)
	_, _, Pa := GetDVSKeyPair_ZZZ(nil, Z[:])

	// Destroy Z
	defer CleanMemory(Z[:])

	// Compute ID|Pa
	ID = append(ID, Pa...)

	// Either Client or TA calculates Hash(ID)
	HCID := HashId(HASH_TYPE_MPIN, hashBytes_ZZZ, ID)

	// Generate server secret share 1
	_, SS1 := GetServerSecret_ZZZ(MS1[:])

	// Destroy SS1
	defer CleanMemory(SS1[:])

	// Generate server secret share 2
	_, SS2 := GetServerSecret_ZZZ(MS2[:])

	// Destroy SS2
	defer CleanMemory(SS2[:])

	// Combine server secret shares
	_, SS := RecombineG2_ZZZ(SS1[:], SS2[:])

	// Destroy SS
	defer CleanMemory(SS[:])

	// Generate client secret share 1
	_, CS1 := GetClientSecret_ZZZ(MS1[:], HCID)

	// Destroy CS1
	defer CleanMemory(CS1[:])

	// Generate client secret share 2
	_, CS2 := GetClientSecret_ZZZ(MS2[:], HCID)

	// Destroy CS2
	defer CleanMemory(CS2[:])

	// Combine client secret shares
	CS := make([]byte, G1S_ZZZ)
	_, CS = RecombineG1_ZZZ(CS1[:], CS2[:])

	// Compute key-escrow less secret
	_, _, CS = GetG1Multiple_ZZZ(nil, 0, Z[:], CS[:])

	// Destroy CS
	defer CleanMemory(CS[:])

	// Create token
	_, TOKEN := ExtractPIN_ZZZ(HASH_TYPE_MPIN, ID[:], PIN1, CS[:])

	// Destroy TOKEN
	defer CleanMemory(TOKEN[:])

	timeValue += 10
	// Client Pass 1
	var X [PGS_ZZZ]byte
	_, XOut, SEC, U, _ := Client1_ZZZ(HASH_TYPE_MPIN, 0, ID, &rng, X[:], PIN2, TOKEN[:], nil)

	// Destroy XOut
	defer CleanMemory(XOut[:])
	// Destroy SEC
	defer CleanMemory(SEC[:])

	// Server Pass 1
	var HID []byte
	HID, _ = Server1_ZZZ(HASH_TYPE_MPIN, 0, ID)

	_, Y := RandomGenerate_ZZZ(&rng)

	// Destroy HID
	defer CleanMemory(HID[:])
	// Destroy Y
	defer CleanMemory(Y[:])

	// Client Pass 2
	_, V := Client2_ZZZ(XOut[:], Y[:], SEC[:])

	// Server Pass 2
	// Send UT as V to model bad token
	got, _, _ = Server2_ZZZ(0, HID[:], nil, Pa, Y[:], SS[:], U[:], nil, V[:], false)
	assert.Equal(t, want, got, "Should be equal")
}
