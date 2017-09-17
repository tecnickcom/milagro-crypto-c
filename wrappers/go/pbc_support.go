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
#include "pbc_support.h"
#include "utils.h"
*/
import "C"
import (
	"errors"
)

const IVS int = 12
const HASH_TYPE_MPIN = SHA256

/*

Today returns today's date as days elapsed from the epoch. This function uses the system clock

Args:

Returns:

    epochDate: epoch days
*/
func Today() int {
	epochDate := C.today()
	return int(epochDate)
}

/*

GetTime returns time elapsed from the epoch. This function uses the system clock

Args:

Returns:

    epochTime: epoch time

*/
func GetTime() int {
	epochTime := C.GET_TIME()
	return int(epochTime)
}

/*

HashId hashs an M-Pin Identity

Args:

    hashType: Hash function
    mpinId:   An octet pointer containing the M-Pin ID

Returns:

    hashMPinId: hash of the M-Pin ID

*/
func HashId(hashType int, hashBytes int, mpinId []byte) (hashMPinId []byte) {
	// Form Octets
	mpinIdStr := string(mpinId)
	mpinIdOct := GetOctet(mpinIdStr)
	defer OctetFree(&mpinIdOct)
	hashMPinIdOct := GetOctetZero(hashBytes)
	defer OctetFree(&hashMPinIdOct)

	// Hash MPIN_ID
	C.HASH_ID(C.int(hashType), &mpinIdOct, &hashMPinIdOct)

	// Convert octet to bytes
	hashMPinId = OctetToBytes(&hashMPinIdOct)

	return hashMPinId
}

/*

HashAll hashs the session transcript

Args:

    hashType: Hash function
    hashMPinId: An octet pointer to the hash of the M-Pin ID
    U: U = x.H(mpin_id)
    UT: UT = x.(H(ID)+H(epoch_date|H(ID)))
    y: server challenge
    V: V = -(x+y)(CS+TP), where CS is the reconstructed client secret and TP is the time permit
    Z: client part response
    T: server part response

Returns:

    HM: hash of the input values

*/
func HashAll(hashType int, hashBytes int, hashMPinId, U, UT, V, y, Z, T []byte) (HM []byte) {
	// Form Octets
	hashMPinIdStr := string(hashMPinId)
	hashMPinIdOct := GetOctet(hashMPinIdStr)
	defer OctetFree(&hashMPinIdOct)

	UStr := string(U)
	UOct := GetOctet(UStr)
	defer OctetFree(&UOct)

	UTStr := string(UT)
	UTOct := GetOctet(UTStr)
	defer OctetFree(&UTOct)

	yStr := string(y)
	yOct := GetOctet(yStr)
	defer OctetFree(&yOct)

	VStr := string(V)
	VOct := GetOctet(VStr)
	defer OctetFree(&VOct)

	ZStr := string(Z)
	ZOct := GetOctet(ZStr)
	defer OctetFree(&ZOct)

	TStr := string(T)
	TOct := GetOctet(TStr)
	defer OctetFree(&TOct)

	HMOct := GetOctetZero(hashBytes)
	defer OctetFree(&HMOct)

	// Hash values
	if UT == nil {
		C.HASH_ALL(C.int(hashType), &hashMPinIdOct, &UOct, nil, &yOct, &VOct, &ZOct, &TOct, &HMOct)
	} else {
		UTStr := string(UT)
		UTOct := GetOctet(UTStr)
		defer OctetFree(&UTOct)
		C.HASH_ALL(C.int(hashType), &hashMPinIdOct, &UOct, &UTOct, &yOct, &VOct, &ZOct, &TOct, &HMOct)
	}

	// Convert octet to bytes
	HM = OctetToBytes(&HMOct)

	return HM[:]
}

/*

GenerateRandomByte generates a random byte array

Args:

    RNG: Pointer to cryptographically secure pseudo-random number generator instance
    length: Gives length of random byte array

Returns:

    randomValue: Random value

*/
func GenerateRandomByte(RNG *RandNG, randomLen int) (randomValue []byte) {
	randomOct := GetOctetZero(randomLen)
	defer OctetFree(&randomOct)

	C.generateRandom(RNG.csprng(), &randomOct)

	// Convert octet to bytes
	randomValue = OctetToBytes(&randomOct)

	return randomValue[:]
}

/*

GenerateOTP returns a random six digit one time Password

Args:

    RNG: Pointer to cryptographically secure pseudo-random number generator instance

Returns:

    otp: One time Password

*/
func GenerateOTP(RNG *RandNG) int {
	OTP := C.generateOTP(RNG.csprng())
	return int(OTP)
}

/*

AesGcmEncrypt performs AES-GCM Encryption

Args:

    K: AES Key
    IV: Initialization vector
    H: header
    P: Plaintext to be encrypted

Returns:

    C: resultant ciphertext
    T: MAC
	error: in case of a bad K

*/
func AesGcmEncrypt(K, IV, H, P []byte) (C, T []byte, err error) {

	KStr := string(K)

	if len(KStr) != 16 {
		err = errors.New("Invalid Key")
		return nil, nil, err
	}

	KOct := GetOctet(KStr)
	defer OctetFree(&KOct)
	IVStr := string(IV)
	IVOct := GetOctet(IVStr)
	defer OctetFree(&IVOct)
	HStr := string(H)
	HOct := GetOctet(HStr)
	defer OctetFree(&HOct)
	PStr := string(P)
	POct := GetOctet(PStr)
	defer OctetFree(&POct)

	TOct := GetOctetZero(16)
	defer OctetFree(&TOct)
	lenC := len(PStr)
	COct := GetOctetZero(lenC)
	defer OctetFree(&COct)

	C.AES_GCM_ENCRYPT(&KOct, &IVOct, &HOct, &POct, &COct, &TOct)

	// Convert octet to bytes
	C = OctetToBytes(&COct)
	T = OctetToBytes(&TOct)

	return C, T, nil
}

/*

AesGcmDecrypt performs AES-GCM Deryption

Args:

    K: AES Key
    IV: Initialization vector
    H: header
    C: ciphertext

Returns:

    P: resultant plaintext
    T: MAC

*/
func AesGcmDecrypt(K, IV, H, C []byte) (P, T []byte, err error) {
	KStr := string(K)

	if len(KStr) != 16 {
		err = errors.New("Invalid Key")
		return nil, nil, err
	}

	KOct := GetOctet(KStr)
	defer OctetFree(&KOct)
	IVStr := string(IV)
	IVOct := GetOctet(IVStr)
	defer OctetFree(&IVOct)
	HStr := string(H)
	HOct := GetOctet(HStr)
	defer OctetFree(&HOct)
	CStr := string(C)
	COct := GetOctet(CStr)
	defer OctetFree(&COct)

	TOct := GetOctetZero(16)
	defer OctetFree(&TOct)
	lenP := len(CStr)
	POct := GetOctetZero(lenP)
	defer OctetFree(&POct)

	C.AES_GCM_DECRYPT(&KOct, &IVOct, &HOct, &COct, &POct, &TOct)

	// Convert octet to bytes
	P = OctetToBytes(&POct)
	T = OctetToBytes(&TOct)

	return P, T, nil
}
