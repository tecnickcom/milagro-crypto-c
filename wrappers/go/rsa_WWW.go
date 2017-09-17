/**
 * @file rsa_WWW.go
 * @author Alessandro Budroni
 * @brief Wrappers for RSA functions
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

package amcl

/*
#cgo CFLAGS:  -std=c99 -O3 -I@PROJECT_BINARY_DIR@/include -I@CMAKE_INSTALL_PREFIX@/include -DCMAKE
#cgo LDFLAGS: -L. -L@CMAKE_INSTALL_PREFIX@/lib -lamcl_rsa_WWW -lamcl_core
#include <stdio.h>
#include <stdlib.h>
#include "amcl.h"
#include "randapi.h"
#include "rsa_WWW.h"
#include "utils.h"
*/
import "C"

// RSA Constant
const RFS_WWW int = int(C.RFS_WWW)     // RFS_WWW is the RSA Public Key Size in bytes
const FFLEN_WWW int = int(C.FFLEN_WWW) // FFLEN_WWW consists in 2^n multiplier of BIGBITS to specify supported Finite Field size, e.g 2048=256*2^3 where BIGBITS=256

const HASH_TYPE_RSA_WWW int = int(C.HASH_TYPE_RSA_WWW) // HASH_TYPE_RSA_WWW is the chosen Hash algorithm

// RSAKeyPair generates an RSA key pair
func RSAKeyPair_WWW(RNG *RandNG, e int32, P []byte, Q []byte) (C.rsa_private_key_WWW, C.rsa_public_key_WWW) {
	PStr := string(P)
	POct := GetOctet(PStr)
	defer OctetFree(&POct)
	QStr := string(Q)
	QOct := GetOctet(QStr)
	defer OctetFree(&QOct)
	eVal := C.int32_t(e)
	var RSA_PubKey C.rsa_public_key_WWW
	var RSA_PrivKey C.rsa_private_key_WWW

	C.RSA_WWW_KEY_PAIR(RNG.csprng(), eVal, &RSA_PrivKey, &RSA_PubKey, &POct, &QOct)
	return RSA_PrivKey, RSA_PubKey
}

// RSAEncrypt_WWW encrypts F with the public key
func RSAEncrypt_WWW(RSA_PubKey *C.rsa_public_key_WWW, F []byte) (G []byte) {
	FStr := string(F)
	FOct := GetOctet(FStr)
	defer OctetFree(&FOct)
	GOct := GetOctetZero(RFS_WWW)
	defer OctetFree(&GOct)

	C.RSA_WWW_ENCRYPT(RSA_PubKey, &FOct, &GOct)
	G = OctetToBytes(&GOct)
	return G[:]
}

// RSADecrypt_WWW decrypts G with the private key
func RSADecrypt_WWW(RSA_PrivKey *C.rsa_private_key_WWW, G []byte) (F []byte) {
	GStr := string(G)
	GOct := GetOctet(GStr)
	defer OctetFree(&GOct)
	FOct := GetOctetZero(RFS_WWW)
	defer OctetFree(&FOct)

	C.RSA_WWW_DECRYPT(RSA_PrivKey, &GOct, &FOct)
	F = OctetToBytes(&FOct)
	return F[:]
}

// RSAPrivateKeyKill_WWW destroys an RSA private Key
func RSAPrivateKeyKill_WWW(RSA_PrivKey *C.rsa_private_key_WWW) {
	C.RSA_WWW_PRIVATE_KEY_KILL(RSA_PrivKey)
}
