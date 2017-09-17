/**
 * @file rsa_support.go
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
#cgo CFLAGS: -std=c99 -O3 -I@PROJECT_BINARY_DIR@/include -I@CMAKE_INSTALL_PREFIX@/include -DCMAKE
#cgo LDFLAGS: -L. -L@CMAKE_INSTALL_PREFIX@/lib -lamcl_core
#include <stdio.h>
#include <stdlib.h>
#include "amcl.h"
#include "randapi.h"
#include "rsa_support.h"
#include "utils.h"
*/
import "C"

// RSA Constant
const MAX_RSA_BYTES int = int(C.MAX_RSA_BYTES) // MAX_RSA_BYTES is the maximum RSA level of security supported - 4096

// PKCS15 (PKCS 1.5) - padding of a message to be signed
func PKCS15(hashType, RFS int, M []byte) (errorCode int, C []byte) {
	MStr := string(M)
	MOct := GetOctet(MStr)
	defer OctetFree(&MOct)
	COct := GetOctetZero(RFS)
	defer OctetFree(&COct)

	rtn := C.PKCS15(C.int(hashType), &MOct, &COct)
	errorCode = int(rtn)
	C = OctetToBytes(&COct)
	return errorCode, C[:]
}

// OAEPencode encodes the message for encryption
func OAEPencode(hashType, RFS int, M []byte, RNG *RandNG, P []byte) (errorCode int, F []byte) {
	MStr := string(M)
	MOct := GetOctet(MStr)
	defer OctetFree(&MOct)
	PStr := string(P)
	POct := GetOctet(PStr)
	defer OctetFree(&POct)
	FOct := GetOctetZero(RFS)
	defer OctetFree(&FOct)

	rtn := C.OAEP_ENCODE(C.int(hashType), &MOct, RNG.csprng(), &POct, &FOct)
	errorCode = int(rtn)
	F = OctetToBytes(&FOct)
	return errorCode, F[:]
}

// OAEPdecode decodes message M after decryption, F is the decoded message
func OAEPdecode(hashType int, P []byte, M []byte) (int, []byte) {
	MStr := string(M)
	MOct := GetOctet(MStr)
	defer OctetFree(&MOct)
	PStr := string(P)
	POct := GetOctet(PStr)
	defer OctetFree(&POct)

	rtn := C.OAEP_DECODE(C.int(hashType), &POct, &MOct)
	errorCode := int(rtn)
	M = OctetToBytes(&MOct)
	return errorCode, M[:]
}
