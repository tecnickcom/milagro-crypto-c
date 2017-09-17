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

/*
#cgo CFLAGS:  -std=c99 -O3 -I@PROJECT_BINARY_DIR@/include -I@CMAKE_INSTALL_PREFIX@/include -DCMAKE
#cgo LDFLAGS: -L. -L@CMAKE_INSTALL_PREFIX@/lib -lamcl_core
#include <stdio.h>
#include <stdlib.h>
#include "amcl.h"
#include "randapi.h"
#include "utils.h"
*/
import "C"
import (
	"encoding/hex"
	"unsafe"
)

//  Hash function choice
const SHA256 int = 32
const SHA384 int = 48
const SHA512 int = 64

// RandNG is a type alias for C.csprng
type RandNG C.csprng

// csprng is an alias function for the C counterpart
func (rng *RandNG) csprng() *C.csprng {
	return (*C.csprng)(rng)
}

/*

CreateCSPRNG makes a cryptographically secure pseudo-random number generator instance

Args:

    seed:   random seed value

Returns:

    RNG: Pointer to cryptographically secure pseudo-random number generator instance

*/
func CreateCSPRNG(SEED []byte) RandNG {
	// Form Octet
	SEEDStr := string(SEED)
	SEEDOct := GetOctet(SEEDStr)
	defer OctetFree(&SEEDOct)
	var RNG C.csprng
	C.CREATE_CSPRNG(&RNG, &SEEDOct)
	return RandNG(RNG)
}

// OctetFree frees memory associated with an octet
func OctetFree(valOctet *C.octet) {
	OctetClear(valOctet)
	C.free(unsafe.Pointer(valOctet.val))
}

// GetOctetZero forms an empty octet
func GetOctetZero(lenStr int) C.octet {
	valBytes := make([]byte, lenStr)
	val := string(valBytes)
	valCS := C.CString(val)
	lenCS := C.int(lenStr)
	octetVal := C.octet{lenCS, lenCS, valCS}
	return octetVal
}

// GetOctet forms an octet from a string
func GetOctet(valStr string) C.octet {
	valCS := C.CString(valStr)
	lenCS := C.int(len(valStr))
	octetVal := C.octet{lenCS, lenCS, valCS}
	return octetVal
}

// GetOctetHex forms an octet from a hex string
func GetOctetHex(valHex string) C.octet {
	valBytes, err := hex.DecodeString(valHex)
	if err != nil {
		octetVal := GetOctetZero(0)
		return octetVal
	}
	valStr := string(valBytes)
	octetVal := GetOctet(valStr)
	return octetVal
}

// OctetLen gets length in bytes of an octet
func OctetLen(valOctet *C.octet) int {
	return int(valOctet.len)
}

// OctetToString converts an octet to a string
func OctetToString(valOct *C.octet) string {
	dstLen := OctetLen(valOct)
	dstBytes := make([]byte, dstLen)
	dstStr := string(dstBytes)
	dst := C.CString(dstStr)
	C.OCT_toStr(valOct, dst)
	dstStr = C.GoStringN(dst, valOct.len)
	C.free(unsafe.Pointer(dst))
	return dstStr
}

// OctetToBytes converts an octet to bytes
func OctetToBytes(valOct *C.octet) []byte {
	dstLen := OctetLen(valOct)
	dstBytes := make([]byte, dstLen)
	dstStr := string(dstBytes)
	dst := C.CString(dstStr)
	C.OCT_toStr(valOct, dst)
	dstStr = C.GoStringN(dst, valOct.len)
	C.free(unsafe.Pointer(dst))
	dstBytes = []byte(dstStr)
	return dstBytes
}

// OctetToHex converts an octet to a hex string
func OctetToHex(valOctet *C.octet) string {
	dstLen := OctetLen(valOctet)
	dstBytes := make([]byte, hex.EncodedLen(dstLen))
	dstStr := string(dstBytes)
	dst := C.CString(dstStr)
	C.OCT_toHex(valOctet, dst)
	dstStr = C.GoString(dst)
	C.free(unsafe.Pointer(dst))
	return dstStr
}

// OctetComp compares two octet - only for testing
func OctetComp(O1 []byte, O2 []byte) int {
	O1Str := string(O1)
	Oct1 := GetOctet(O1Str)
	defer OctetFree(&Oct1)
	O2Str := string(O2)
	Oct2 := GetOctet(O2Str)
	defer OctetFree(&Oct2)
	rtn := C.OCT_comp(&Oct1, &Oct2)
	return int(rtn)
}

// OctetClear wipes clean an octet
func OctetClear(valOctet *C.octet) {
	C.OCT_clear(valOctet)
}

// CleanMemory set memory of slice to zero
func CleanMemory(arr []byte) {
	if len(arr) == 0 {
		return
	}
	arr[0] = 0
	for i := 1; i < len(arr); i *= 2 {
		copy(arr[i:], arr[:i])
	}
}
