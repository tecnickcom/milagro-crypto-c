/**
 * @file ecdsa_ZZZ.go
 * @author Alessandro Budroni
 * @brief Wrappers for ECDSA functions
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
#cgo LDFLAGS: -L. -L@CMAKE_INSTALL_PREFIX@/lib -lamcl_core
#include <stdio.h>
#include <stdlib.h>
#include "amcl.h"
#include "ecdh_support.h"
*/
import "C"

// PBKDF2 is a Password Based Key Derivation Function. Uses SHA256 internally
func PBKDF2(hashType int, Pass []byte, Salt []byte, iter int, length int) (Key []byte) {
	PassStr := string(Pass)
	PassOct := GetOctet(PassStr)
	defer OctetFree(&PassOct)
	SaltStr := string(Salt)
	SaltOct := GetOctet(SaltStr)
	defer OctetFree(&SaltOct)
	KeyOct := GetOctetZero(length)
	defer OctetFree(&KeyOct)

	C.PBKDF2(C.int(hashType), &PassOct, &SaltOct, C.int(iter), C.int(length), &KeyOct)
	Key = OctetToBytes(&KeyOct)
	return Key
}
