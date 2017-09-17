/**
 * @file test_rsa_WWW.go
 * @author Alessandro Budroni
 * @brief RSA test
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

	// Message to encrypt
	MESSAGEstr := "Hello World"
	MESSAGE := []byte(MESSAGEstr)

	fmt.Printf("MESSAGE: ")
	fmt.Printf("%s\n\n", MESSAGE[:])

	// Generating public/private key pair
	fmt.Printf("Generating public/private key pair\n")
	RSA_PrivKey, RSA_PubKey := amcl.RSAKeyPair_WWW(&rng, 65537, nil, nil)

	// OAEP encode MESSAGE to e
	rtn, F := amcl.OAEPencode(amcl.HASH_TYPE_RSA_WWW, amcl.RFS_WWW, MESSAGE, &rng, nil)
	if rtn != 1 {
		fmt.Println("OAEPencode Error:", rtn)
		return
	}

	fmt.Printf("Encoded MESSAGE: 0x")
	fmt.Printf("%x\n\n", F[:])

	// encrypt encoded MESSAGE
	G := amcl.RSAEncrypt_WWW(&RSA_PubKey, F[:])

	fmt.Printf("Encrypted MESSAGE: 0x")
	fmt.Printf("%x\n\n", G[:])

	// decrypt encrypted MESSAGE
	ML := amcl.RSADecrypt_WWW(&RSA_PrivKey, G[:])

	fmt.Printf("Decrypted MESSAGE: 0x")
	fmt.Printf("%x\n\n", ML[:])

	// OAEP decode MESSAGE
	rtn, MESSAGEgot := amcl.OAEPdecode(amcl.HASH_TYPE_RSA_WWW, nil, ML[:])
	if rtn != 1 {
		fmt.Println("OAEPdecode Error:", rtn)
		return
	}

	fmt.Printf("Decoded MESSAGE: ")
	fmt.Printf("%s\n\n", MESSAGEgot[:])

	// NOTE - use the following only for testing
	rtn = int(amcl.OctetComp(MESSAGE, MESSAGEgot))
	if rtn != 1 {
		fmt.Println("Error - Message doesn't correspond to the expected one:", rtn)
		return
	}

	// destroy private key
	fmt.Printf("Destroy private key\n\n")
	amcl.RSAPrivateKeyKill_WWW(&RSA_PrivKey)

}
