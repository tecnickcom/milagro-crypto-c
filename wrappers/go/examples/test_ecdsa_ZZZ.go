/**
 * @file test_ecdsa_ZZZ.go
 * @author Alessandro Budroni
 * @brief ECDSA test
 *
 * LICENSE
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package main

import (
	"encoding/hex"
	"fmt"

	"github.com/miracl/amcl-go-wrapper"
)

func main() {

	var rtn int

	// Seed value for Random Number Generator (RNG)
	seedHex := "9e8b4178790cd57a5761c4a6f164ba72"
	seed, err := hex.DecodeString(seedHex)
	if err != nil {
		fmt.Println("Error decoding seed value")
		return
	}

	rng := amcl.CreateCSPRNG(seed)

	// Set PassPhrase
	PassPhraseStr := "AlicePassPhrase"
	PassPhrase := []byte(PassPhraseStr)

	fmt.Printf("PassPhrase: ")
	fmt.Printf("%s\n\n", PassPhrase[:])

	// Set Salt
	SaltStr := "aabbccddee"
	Salt := []byte(SaltStr)

	fmt.Printf("Salt: ")
	fmt.Printf("%s\n\n", Salt[:])

	// Generating public/private key pair
	fmt.Printf("Generating public/private key pair...\n\n")

	// Generate ECC Private Key
	PrivKey := amcl.PBKDF2(amcl.SHA256, PassPhrase[:], Salt[:], 1000, amcl.EGS_ZZZ)

	// Destroy Private Key
	defer amcl.CleanMemory(PrivKey[:])
	// Destroy Salt
	defer amcl.CleanMemory(Salt[:])
	// Destroy Passphrase
	defer amcl.CleanMemory(PassPhrase[:])

	fmt.Printf("Private Key: 0x")
	fmt.Printf("%x\n\n", PrivKey[:])

	// Generate ECC Key Pair
	rtn, PrivKey, PubKey := amcl.ECPKeyPairGenerate_ZZZ(nil, PrivKey)
	if rtn != 0 {
		fmt.Println("Error - Generating ECC Key Pair!\n", rtn)
		return
	}

	fmt.Printf("Public Key: 0x")
	fmt.Printf("%x\n\n", PubKey[:])

	// Validate ECC Public Key
	rtn = amcl.ECPPublicKeyValidate_ZZZ(1, PubKey[:])
	if rtn != 0 {
		fmt.Println("Error - ECC Public Key is invalid!\n", rtn)
		return
	}

	// Message to sign
	MESSAGEstr := "Hello World\n"
	MESSAGE := []byte(MESSAGEstr)
	fmt.Printf("MESSAGE to be signed: 0x")
	fmt.Printf("%x\n\n", MESSAGE[:])

	// Sign Message
	rtn, C, D := amcl.ECPSpDsa_ZZZ(amcl.SHA256, &rng, nil, PrivKey[:], MESSAGE[:])
	if rtn != 0 {
		fmt.Println("Error - ECDSA Signature Failed!\n", rtn)
		return
	}

	fmt.Printf("Signed MESSAGE pair:\n\n")
	fmt.Printf("C: 0x")
	fmt.Printf("%x\n\n", C[:])
	fmt.Printf("D: 0x")
	fmt.Printf("%x\n\n", D[:])

	// Verify Message
	rtn = amcl.ECPVpDsa_ZZZ(amcl.SHA256, PubKey[:], MESSAGE[:], C[:], D[:])
	if rtn != 0 {
		fmt.Println("Error - ECDSA Verification Failed!\n", rtn)
	} else {
		fmt.Println("ECDSA Signature Valid!\n")
	}
}
