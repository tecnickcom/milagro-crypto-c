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
#cgo LDFLAGS: -L. -L@CMAKE_INSTALL_PREFIX@/lib -lamcl_mpin_ZZZ -lamcl_pairing_ZZZ -lamcl_curve_ZZZ -lamcl_core
#include <stdio.h>
#include <stdlib.h>
#include "amcl.h"
#include "randapi.h"
#include "mpin_ZZZ.h"
#include "utils.h"
*/
import "C"

// ECC points constant
const PAS_ZZZ int = int(C.MPIN_PAS)
const PGS_ZZZ int = int(C.MPIN_PGS_ZZZ)
const PFS_ZZZ int = int(C.MPIN_PFS_ZZZ)
const hashBytes_ZZZ int = int(C.MPIN_PFS_ZZZ)
const G1S_ZZZ = 2*PFS_ZZZ + 1
const G2S_ZZZ = 4 * PFS_ZZZ
const GTS_ZZZ = 12 * PFS_ZZZ

/*

RandomGenerate_ZZZ returns a random integer where s < q is the order of the group of points on the curve.

Args:

    RNG: Pointer to cryptographically secure pseudo-random number generator instance

Returns:

    errorCode: error from the C function
    s: random group element

*/
func RandomGenerate_ZZZ(RNG *RandNG) (errorCode int, S []byte) {
	// Form Octet
	SOct := GetOctetZero(PGS_ZZZ)
	defer OctetFree(&SOct)

	rtn := C.MPIN_ZZZ_RANDOM_GENERATE(RNG.csprng(), &SOct)
	errorCode = int(rtn)

	// Convert octet to bytes
	S = OctetToBytes(&SOct)

	return errorCode, S[:]
}

/*

GetServerSecret_ZZZ creates a server secret in G2 from a master secret

Args:

    masterSecret:   An octet pointer to the master secret

Returns:

    errorCode: error from the C function
    serverSecret: Server secret

*/
func GetServerSecret_ZZZ(masterSecret []byte) (errorCode int, serverSecret []byte) {
	// Form Octets
	masterSecretStr := string(masterSecret)
	masterSecretOct := GetOctet(masterSecretStr)
	defer OctetFree(&masterSecretOct)
	serverSecretOct := GetOctetZero(G2S_ZZZ)
	defer OctetFree(&serverSecretOct)

	rtn := C.MPIN_ZZZ_GET_SERVER_SECRET(&masterSecretOct, &serverSecretOct)
	errorCode = int(rtn)

	// Convert octet to bytes
	serverSecret = OctetToBytes(&serverSecretOct)

	return errorCode, serverSecret[:]
}

/*

RecombineG1_ZZZ adds two members from the group G1

Args:

    R1: An input member of G1
    R2: An input member of G1

Returns:

    errorCode: error from the C function
    R: An output member of G1; R = Q1+Q2

*/
func RecombineG1_ZZZ(R1 []byte, R2 []byte) (errorCode int, R []byte) {
	// Form Octets
	R1Str := string(R1)
	R1Oct := GetOctet(R1Str)
	defer OctetFree(&R1Oct)
	R2Str := string(R2)
	R2Oct := GetOctet(R2Str)
	defer OctetFree(&R2Oct)
	ROct := GetOctetZero(G1S_ZZZ)
	defer OctetFree(&ROct)

	rtn := C.MPIN_ZZZ_RECOMBINE_G1(&R1Oct, &R2Oct, &ROct)
	errorCode = int(rtn)

	// Convert octet to bytes
	R = OctetToBytes(&ROct)

	return errorCode, R[:]
}

/*

RecombineG2_ZZZ adds two members from the group G2

Args:

    W1: An input member of G2
    W2: An input member of G2

Returns:

    errorCode: error from the C function
    W: An output member of G2; W = W1+W2

*/
func RecombineG2_ZZZ(W1 []byte, W2 []byte) (errorCode int, W []byte) {
	// Form Octets
	W1Str := string(W1)
	W1Oct := GetOctet(W1Str)
	defer OctetFree(&W1Oct)
	W2Str := string(W2)
	W2Oct := GetOctet(W2Str)
	defer OctetFree(&W2Oct)
	WOct := GetOctetZero(G2S_ZZZ)
	defer OctetFree(&WOct)

	rtn := C.MPIN_ZZZ_RECOMBINE_G2(&W1Oct, &W2Oct, &WOct)
	errorCode = int(rtn)

	// Convert octet to bytes
	W = OctetToBytes(&WOct)

	return errorCode, W[:]
}

/*

GetClientSecret_ZZZ creates a client secret in G1 from a master secret and the hash of the M-Pin Id

Args:

    masterSecret:  An octet pointer to the master secret
    hashMPinId:    An octet pointer to the hash of the M-Pin ID

Returns:

    errorCode: error from the C function
    clientSecret: Client secret

*/
func GetClientSecret_ZZZ(masterSecret []byte, hashMPinId []byte) (errorCode int, clientSecret []byte) {
	// Form Octets
	masterSecretStr := string(masterSecret)
	masterSecretOct := GetOctet(masterSecretStr)
	defer OctetFree(&masterSecretOct)
	hashMPinIdStr := string(hashMPinId)
	hashMPinIdOct := GetOctet(hashMPinIdStr)
	defer OctetFree(&hashMPinIdOct)
	clientSecretOct := GetOctetZero(G1S_ZZZ)
	defer OctetFree(&clientSecretOct)

	rtn := C.MPIN_ZZZ_GET_CLIENT_SECRET(&masterSecretOct, &hashMPinIdOct, &clientSecretOct)
	errorCode = int(rtn)

	// Convert octet to bytes
	clientSecret = OctetToBytes(&clientSecretOct)

	return errorCode, clientSecret[:]
}

/*

GetDVSKeyPair_ZZZ randomly generates the public key in G2 for the key-escrow less scheme

Args:

    RNG:	cryptographically secure random number generator
	z: 		input octet pointer to a value for z in RNG==nil, nil otherwise

Returns:

    errorCode: 		error from the C function
    zOut: 			output randomly generated private key if RNG!=nil, z otherwise
    publicKey: 		client public key: (z^-1).Q

*/
func GetDVSKeyPair_ZZZ(RNG *RandNG, z []byte) (errorCode int, zOut []byte, publicKey []byte) {
	// Form octets
	publicKeyOct := GetOctetZero(G2S_ZZZ)
	defer OctetFree(&publicKeyOct)
	zStr := string(z)
	zOct := GetOctet(zStr)
	defer OctetFree(&zOct)

	var pRNG *C.csprng

	if RNG == nil {
		pRNG = nil
		zOct = GetOctet(zStr)
	} else {
		pRNG = RNG.csprng()
		zOct = GetOctetZero(PGS_ZZZ)
	}

	rtn := C.MPIN_ZZZ_GET_DVS_KEYPAIR(pRNG, &zOct, &publicKeyOct)
	errorCode = int(rtn)

	// Convert octets to bytes
	publicKey = OctetToBytes(&publicKeyOct)
	zOut = OctetToBytes(&zOct)

	return errorCode, zOut, publicKey[:]

}

/*

GetClientPermit_ZZZ creates a time permit in G1 from a master secret, hash of the M-Pin Id and epoch days

Args:

    hashType:      Hash function
    epochDate:     Epoch days
    masterSecret:  An octet pointer to the master secret
    hashMPinId:    An octet pointer to the hash of the M-Pin ID

Returns:

    errorCode: error from the C function
    timePermit: Time permit

*/
func GetClientPermit_ZZZ(hashType, epochDate int, masterSecret, hashMPinId []byte) (errorCode int, timePermit []byte) {
	// Form Octets
	masterSecretmasterSecretStr := string(masterSecret)
	masterSecretOct := GetOctet(masterSecretmasterSecretStr)
	defer OctetFree(&masterSecretOct)
	hashMPinIdStr := string(hashMPinId)
	hashMPinIdOct := GetOctet(hashMPinIdStr)
	defer OctetFree(&hashMPinIdOct)
	timePermitOct := GetOctetZero(G1S_ZZZ)
	defer OctetFree(&timePermitOct)

	rtn := C.MPIN_ZZZ_GET_CLIENT_PERMIT(C.int(hashType), C.int(epochDate), &masterSecretOct, &hashMPinIdOct, &timePermitOct)
	errorCode = int(rtn)

	// Convert octet to bytes
	timePermit = OctetToBytes(&timePermitOct)

	return errorCode, timePermit[:]
}

/*

ExtractPIN_ZZZ extracts a PIN from the client secret

Args:

    hashType: Hash function
    mpinId:   M-Pin ID
    PIN:   PIN input by user
    clientSecret: User's client secret

Returns:

    errorCode: error from the C function
    token: Result of extracting a PIN from client secret

*/
func ExtractPIN_ZZZ(hashType int, mpinId []byte, PIN int, clientSecret []byte) (errorCode int, token []byte) {
	// Form Octets
	mpinIdStr := string(mpinId)
	mpinIdOct := GetOctet(mpinIdStr)
	defer OctetFree(&mpinIdOct)
	clientSecretStr := string(clientSecret)
	clientSecretOct := GetOctet(clientSecretStr)
	defer OctetFree(&clientSecretOct)

	rtn := C.MPIN_ZZZ_EXTRACT_PIN(C.int(hashType), &mpinIdOct, C.int(PIN), &clientSecretOct)
	errorCode = int(rtn)

	// Convert octet to bytes
	token = OctetToBytes(&clientSecretOct)

	return errorCode, token[:]
}

/*

Client_ZZZ performs client side of the one-Pass version of the M-Pin protocol. If Time Permits are
disabled then set epoch_date = 0.In this case UT is not generated and can be set to nil.
If Time Permits are enabled, and PIN error detection is OFF, U is not generated and
can be set to nil. If Time Permits are enabled and PIN error detection is ON then U
and UT are both generated.

Args:

    hashType: Hash function
    epochDate: Date, in days since the epoch. Set to 0 if Time permits disabled
    mpinId: M-Pin ID
    RNG: cryptographically secure random number generator
    x: random number provided as input if RNG==nil, otherwise must be nil
    pin: PIN entered by user
    token: M-Pin token
    timePermit: M-Pin time permit
    message: message to be signed
    epochTime: Epoch time in seconds

Returns:

    errorCode: error from the C function
    x: Randomly generated integer if R!=nil, otherwise must be provided as an input
    U: U = x.H(ID)
    UT: UT = x.(H(ID)+H(epoch_date|H(ID)))
    V: V = -(x+y)(CS+TP), where CS is the reconstructed client secret and TP is the time permit
    y: y = H(t|U) or y = H(t|UT) if Time Permits enabled where is t is epoch time
*/
func Client_ZZZ(hashType, epochDate int, mpinId []byte, RNG *RandNG, x []byte, PIN int, token []byte, timePermit []byte, message []byte, epochTime int) (errorCode int, xOut, y, V, U, UT []byte) {

	var UTOct *C.octet
	var pUT C.octet
	defer OctetFree(&pUT)

	var messageOct *C.octet
	var pmessage C.octet
	defer OctetFree(&pmessage)

	var timePermitOct *C.octet
	var ptimePermit C.octet
	defer OctetFree(&ptimePermit)

	var xOct C.octet
	defer OctetFree(&xOct)

	var pRNG *C.csprng

	if RNG == nil {
		pRNG = nil
		xStr := string(x)
		xOct = GetOctet(xStr)
	} else {
		pRNG = RNG.csprng()
		xOct = GetOctetZero(PGS_ZZZ)
	}
	if epochDate == 0 {
		UTOct = nil
		timePermitOct = nil
	} else {
		pUT = GetOctetZero(G1S_ZZZ)
		UTOct = &pUT
		TPStr := string(timePermit)
		ptimePermit = GetOctet(TPStr)
		timePermitOct = &ptimePermit
	}
	if message == nil {
		messageOct = nil
	} else {
		messageStr := string(message)
		pmessage = GetOctet(messageStr)
		messageOct = &pmessage
	}

	// Form other Octets
	mpinIdStr := string(mpinId)
	mpinIdOct := GetOctet(mpinIdStr)
	defer OctetFree(&mpinIdOct)
	tokenStr := string(token)
	tokenOct := GetOctet(tokenStr)
	defer OctetFree(&tokenOct)

	VOct := GetOctetZero(G1S_ZZZ)
	defer OctetFree(&VOct)
	UOct := GetOctetZero(G1S_ZZZ)
	defer OctetFree(&UOct)

	yOct := GetOctetZero(PGS_ZZZ)
	defer OctetFree(&yOct)

	rtn := C.MPIN_ZZZ_CLIENT(C.int(hashType), C.int(epochDate), &mpinIdOct, pRNG, &xOct, C.int(PIN), &tokenOct, &VOct, &UOct, UTOct, timePermitOct, messageOct, C.int(epochTime), &yOct)
	errorCode = int(rtn)

	// Convert octet to bytes
	xOut = OctetToBytes(&xOct)
	V = OctetToBytes(&VOct)
	U = OctetToBytes(&UOct)
	y = OctetToBytes(&yOct)

	if epochDate == 0 {
		return errorCode, xOut[:], y[:], V[:], U[:], nil
	} else {
		UT = OctetToBytes(UTOct)
		return errorCode, xOut[:], y[:], V[:], U[:], UT[:]
	}
}

/*

Precompute_ZZZ calculate values for use by the client side of M-Pin Full

Args:

    token:  M-Pin token
    hashMPinId: hash of the M-Pin ID

Returns:

    errorCcode: error from the C function
    pc1: Precomputed value one
    pc2: Precomputed value two

Raises:

*/
func Precompute_ZZZ(token, hashMPinId []byte) (errorCode int, pc1, pc2 []byte) {
	// Form Octets
	hashMPinIdStr := string(hashMPinId)
	hashMPinIdOct := GetOctet(hashMPinIdStr)
	defer OctetFree(&hashMPinIdOct)
	tokenStr := string(token)
	tokenOct := GetOctet(tokenStr)
	defer OctetFree(&tokenOct)

	pc1Oct := GetOctetZero(GTS_ZZZ)
	defer OctetFree(&pc1Oct)
	pc2Oct := GetOctetZero(GTS_ZZZ)
	defer OctetFree(&pc2Oct)

	rtn := C.MPIN_ZZZ_PRECOMPUTE(&tokenOct, &hashMPinIdOct, nil, &pc1Oct, &pc2Oct)
	errorCode = int(rtn)

	// Convert octet to bytes
	pc1 = OctetToBytes(&pc1Oct)
	pc2 = OctetToBytes(&pc2Oct)

	return errorCode, pc1[:], pc2[:]
}

/*

GetG1Multiple_ZZZ calculates W=x*P where random x < q is the order of the group of points on the curve.
When RNG is nil x is Passed in otherwise it is Passed out.

If type=0 then P is. point on the curve or else P is an octet that has to be
mapped to the curve

Args:

    RNG: Pointer to cryptographically secure pseudo-random number generator instance
    type: determines type of action to be taken
    x: random number provided as input if RNG==nil, otherwise must be nil
    P: if type=0 a point in G1, else an octet to be mapped to G1

Returns:

    error_code: error from the C function
    x: an output internally randomly generated if RNG!=nil, otherwise must be provided as an input
    W: W = x.P or W = x.M(P), where M(.) is a mapping when type = 0

Raises:

*/
func GetG1Multiple_ZZZ(RNG *RandNG, typ int, x []byte, G []byte) (errorCode int, xOut, W []byte) {
	xStr := string(x)
	xOct := GetOctet(xStr)
	defer OctetFree(&xOct)
	GStr := string(G)
	GOct := GetOctet(GStr)
	defer OctetFree(&GOct)

	WOct := GetOctetZero(G1S_ZZZ)
	defer OctetFree(&WOct)

	var pRNG *C.csprng

	if RNG == nil {
		pRNG = nil
		xStr := string(x)
		xOct = GetOctet(xStr)
	} else {
		pRNG = RNG.csprng()
		xOct = GetOctetZero(PGS_ZZZ)
	}

	rtn := C.MPIN_ZZZ_GET_G1_MULTIPLE(pRNG, C.int(typ), &xOct, &GOct, &WOct)
	errorCode = int(rtn)

	// Convert octet to bytes
	xOut = OctetToBytes(&xOct)
	W = OctetToBytes(&WOct)

	return errorCode, xOut[:], W[:]
}

/*

Server_ZZZ performs server side of the one-Pass version of the M-Pin protocol
with support for the key-escrow less scheme. If Time Permits are disabled,
set epoch_date = 0, and UT and HTID are not generated and can be set to nil.
If Time Permits are enabled, and PIN error detection is OFF,
U and HID are not needed and can be set to nil. If Time Permits are enabled,
and PIN error detection is ON, U, UT, HID and HTID are all required.

Args:

    hashType: Hash function
    epoch_date: Date, in days since the epoch. Set to 0 if Time permits disabled
    epochTime: Epoch time in seconds
    server_secret: Server secret
    U: U = x.H(ID)
    UT: UT = x.(H(ID)+H(epoch_date|H(ID)))
    V: V = -(x+y)(CS+TP), where CS is the reconstructed client secret and TP is the time permit
    mpinId: M-Pin ID, or hash of the M-Pin ID in anonymous mode, or M-Pin ID | publicKey in key-escrow less scheme
    publicKey: client public key in key-escrow less version, or nil otherwise
    message: message to be signed
    Kangaroo_ZZZ: Set to true to perform Kangaroo

Returns:

    errorCode: error from the C function
    HID:  H(mpin_id). H is a map to a point on the curve
    HTID: H(mpin_id)+H(epoch_date|H(mpin_id)). H is a map to a point on the curve
    E: value to help the Kangaroos to find the PIN error, or nil if not required
    F: value to help the Kangaroos to find the PIN error, or nil if not required
    y: y = H(t|U) or y = H(t|UT) if Time Permits enabled used for debug

*/
func Server_ZZZ(hashType, epochDate, epochTime int, serverSecret, U, UT, V, mpinId, publicKey, message []byte, Kangaroo bool) (errorCode int, HID, HTID, y, E, F []byte) {
	serverSecretStr := string(serverSecret)
	serverSecretOct := GetOctet(serverSecretStr)
	defer OctetFree(&serverSecretOct)
	UStr := string(U)
	UOct := GetOctet(UStr)
	defer OctetFree(&UOct)
	VStr := string(V)
	VOct := GetOctet(VStr)
	defer OctetFree(&VOct)
	mpinIdStr := string(mpinId)
	mpinIdOct := GetOctet(mpinIdStr)
	defer OctetFree(&mpinIdOct)

	HIDOct := GetOctetZero(G1S_ZZZ)
	defer OctetFree(&HIDOct)
	HTIDOct := GetOctetZero(G1S_ZZZ)
	defer OctetFree(&HTIDOct)
	yOct := GetOctetZero(PGS_ZZZ)
	defer OctetFree(&yOct)

	var pUT C.octet
	defer OctetFree(&pUT)
	var pmessage C.octet
	defer OctetFree(&pmessage)
	var pE C.octet
	defer OctetFree(&pE)
	var pF C.octet
	defer OctetFree(&pF)
	var pPublicKey C.octet
	defer OctetFree(&pPublicKey)

	var EOct *C.octet
	var FOct *C.octet
	var UTOct *C.octet
	var messageOct *C.octet
	var publicKeyOct *C.octet

	if publicKey == nil {
		publicKeyOct = nil
	} else {
		publicKeyStr := string(publicKey)
		pPublicKey = GetOctet(publicKeyStr)
		publicKeyOct = &pPublicKey
	}
	if UT == nil {
		UTOct = nil
	} else {
		UTStr := string(UT)
		pUT = GetOctet(UTStr)
		UTOct = &pUT
	}
	if message == nil {
		messageOct = nil
	} else {
		messageStr := string(message)
		pmessage = GetOctet(messageStr)
		messageOct = &pmessage
	}
	if !Kangaroo {
		EOct = nil
		FOct = nil
	} else {
		pE = GetOctetZero(GTS_ZZZ)
		EOct = &pE
		pF = GetOctetZero(GTS_ZZZ)
		FOct = &pF
	}

	rtn := C.MPIN_ZZZ_SERVER(C.int(hashType), C.int(epochDate), &HIDOct, &HTIDOct, &yOct, &serverSecretOct, &UOct, UTOct, &VOct, EOct, FOct, &mpinIdOct, messageOct, C.int(epochTime), publicKeyOct)
	errorCode = int(rtn)

	// Convert octet to bytes
	HID = OctetToBytes(&HIDOct)
	HTID = OctetToBytes(&HTIDOct)
	y = OctetToBytes(&yOct)

	if !Kangaroo {
		return errorCode, HID[:], HTID[:], y[:], nil, nil
	} else {
		E = OctetToBytes(EOct)
		F = OctetToBytes(FOct)
		return errorCode, HID[:], HTID[:], y[:], E[:], F[:]
	}
}

/*

Kangaroo_ZZZ uses Pollards Kangaroos to find PIN error

Args:

    E: a member of the group GT
    F: a member of the group GT =  E^pin_error

Returns:

    PINError: error in PIN or 0 if Kangaroos failed

*/
func Kangaroo_ZZZ(E []byte, F []byte) (PINError int) {
	EStr := string(E)
	EOct := GetOctet(EStr)
	defer OctetFree(&EOct)
	FStr := string(F)
	FOct := GetOctet(FStr)
	defer OctetFree(&FOct)

	rtn := C.MPIN_ZZZ_KANGAROO(&EOct, &FOct)
	PINError = int(rtn)

	return PINError
}

/*

ServerKey_ZZZ calculates AES key on server side for M-Pin Full.Uses UT internally for the
key calculation or uses U if UT is set to None

Args:

    hashType: Hash function
    Z: Client-side Diffie-Hellman component
    serverSecret: server secret
    w: random number generated by the server
    HM: hash of the protocol transcript
    HID: H(mpin_id). H is a map to a point on the curve
    U: U = x.H(ID)
    UT: UT = x.(H(ID)+H(epoch_date|H(ID)))

Returns:

    errorCode: error code from the C function
    serverAESKey: server AES key

*/
func ServerKey_ZZZ(hashType int, Z, serverSecret, w, HM, HID, U, UT []byte) (errorCode int, serverAESKey []byte) {
	var rtn C.int
	ZStr := string(Z)
	ZOct := GetOctet(ZStr)
	defer OctetFree(&ZOct)
	serverSecretStr := string(serverSecret)
	serverSecretOct := GetOctet(serverSecretStr)
	defer OctetFree(&serverSecretOct)
	wStr := string(w)
	wOct := GetOctet(wStr)
	defer OctetFree(&wOct)
	HMStr := string(HM)
	HMOct := GetOctet(HMStr)
	defer OctetFree(&HMOct)
	HIDStr := string(HID)
	HIDOct := GetOctet(HIDStr)
	defer OctetFree(&HIDOct)
	serverAESKeyOct := GetOctetZero(PAS_ZZZ)
	defer OctetFree(&serverAESKeyOct)
	UStr := string(U)
	UOct := GetOctet(UStr)
	defer OctetFree(&UOct)

	if UT == nil {
		UStr := string(U)
		UOct := GetOctet(UStr)
		defer OctetFree(&UOct)
		rtn = C.MPIN_ZZZ_SERVER_KEY(C.int(hashType), &ZOct, &serverSecretOct, &wOct, &HMOct, &HIDOct, &UOct, nil, &serverAESKeyOct)
	} else {
		UTStr := string(UT)
		UTOct := GetOctet(UTStr)
		defer OctetFree(&UTOct)
		rtn = C.MPIN_ZZZ_SERVER_KEY(C.int(hashType), &ZOct, &serverSecretOct, &wOct, &HMOct, &HIDOct, nil, &UTOct, &serverAESKeyOct)
	}

	errorCode = int(rtn)

	// Convert octet to bytes
	serverAESKey = OctetToBytes(&serverAESKeyOct)

	return errorCode, serverAESKey[:]
}

/*

ClientKey_ZZZ calculate AES key on client side for M-Pin Full

Args:

    hashType: Hash function
    pc1: precomputed input
    pc2: precomputed input
    PIN: PIN number
    r: locally generated random number
    x: locally generated random number
    HM: hash of the protocol transcript
    T: Server-side Diffie-Hellman component

Returns:

    error_code: error code from the C function
    clientAESKey: client AES key

*/
func ClientKey_ZZZ(hashType, PIN int, pc1, pc2, r, x, HM, T []byte) (errorCode int, clientAESKey []byte) {
	pc1Str := string(pc1)
	pc1Oct := GetOctet(pc1Str)
	defer OctetFree(&pc1Oct)
	pc2Str := string(pc2)
	pc2Oct := GetOctet(pc2Str)
	defer OctetFree(&pc2Oct)
	rStr := string(r)
	rOct := GetOctet(rStr)
	defer OctetFree(&rOct)
	xStr := string(x)
	xOct := GetOctet(xStr)
	defer OctetFree(&xOct)
	HMStr := string(HM)
	HMOct := GetOctet(HMStr)
	defer OctetFree(&HMOct)
	TStr := string(T)
	TOct := GetOctet(TStr)
	defer OctetFree(&TOct)

	clientAESKeyOct := GetOctetZero(PAS_ZZZ)
	defer OctetFree(&clientAESKeyOct)

	rtn := C.MPIN_ZZZ_CLIENT_KEY(C.int(hashType), &pc1Oct, &pc2Oct, C.int(PIN), &rOct, &xOct, &HMOct, &TOct, &clientAESKeyOct)
	errorCode = int(rtn)

	// Convert octet to bytes
	clientAESKey = OctetToBytes(&clientAESKeyOct)

	return errorCode, clientAESKey[:]
}

/*

Server1_ZZZ performs first Pass of the server side of the 3-Pass version of the M-Pin protocol
If Time Permits are disabled, set epoch_date = 0, and UT and HTID are not generated
and can be set to nil. If Time Permits are enabled, and PIN error detection is OFF,
U and HID are not needed and can be set to nil. If Time Permits are enabled,
and PIN error detection is ON, U, UT, HID and HTID are all required.

Args:

    hashType: Hash function
    epochDate: Date, in days since the epoch. Set to 0 if Time permits disabled
    mpinId: M-Pin ID or hash of the M-Pin ID in anonymous mode

Returns:

    HID:  H(mpin_id). H is a map to a point on the curve
    HTID: H(mpin_id)+H(epoch_date|H(mpin_id)). H is a map to a point on the curve

*/
func Server1_ZZZ(hashType, epochDate int, mpinId []byte) (HID, HTID []byte) {
	// Form Octets
	mpinIdStr := string(mpinId)
	mpinIdOct := GetOctet(mpinIdStr)
	defer OctetFree(&mpinIdOct)

	HIDOct := GetOctetZero(G1S_ZZZ)
	defer OctetFree(&HIDOct)
	HTIDOct := GetOctetZero(G1S_ZZZ)
	defer OctetFree(&HTIDOct)

	C.MPIN_ZZZ_SERVER_1(C.int(hashType), C.int(epochDate), &mpinIdOct, &HIDOct, &HTIDOct)

	// Convert octet to bytes
	HID = OctetToBytes(&HIDOct)
	HTID = OctetToBytes(&HTIDOct)

	return HID[:], HTID[:]
}

/*

Server2_ZZZ performs server side of the three-Pass version of the M-Pin protocol, wiht
support for the key-escrow less scheme.  If Time Permits are disabled,
set epoch_date = 0, and UT and HTID are not generated and can be set to nil.
If Time Permits are enabled, and PIN error detection is OFF,
U and HID are not needed and can be set to nil. If Time Permits are enabled,
and PIN error detection is ON, U, UT, HID and HTID are all required.

Args:

    epochDate: Date in days since the epoch. Set to 0 if Time permits disabled
    HID:  H(mpin_id). H is a map to a point on the curve
    HTID: H(mpin_id)+H(epoch_date|H(mpin_id)). H is a map to a point on the curve
    y: locally generated random number
    publicKey: client public key in key-escrow less scheme, nil otherwise
    server_secret: Server secret
    U: U = x.H(ID)
    UT: UT = x.(H(ID)+H(epoch_date|H(ID)))
    V: V = -(x+y)(CS+TP), where CS is the reconstructed client secret and TP is the time permit
    Kangaroo_ZZZ: Set to true to perform Kangaroo_ZZZ

Returns:

    errorCode: error from the C function
    E: value to help the Kangaroo_ZZZs to find the PIN error, or nil Kangaroo_ZZZ is set to false
    F: value to help the Kangaroo_ZZZs to find the PIN error, or nil Kangaroo_ZZZ is set to false

*/
func Server2_ZZZ(epochDate int, HID []byte, HTID []byte, publicKey []byte, y []byte, serverSecret []byte, U []byte, UT []byte, V []byte, Kangaroo_ZZZ bool) (errorCode int, E []byte, F []byte) {
	// Form Octets
	HIDStr := string(HID)
	HIDOct := GetOctet(HIDStr)
	defer OctetFree(&HIDOct)
	HTIDStr := string(HTID)
	HTIDOct := GetOctet(HTIDStr)
	defer OctetFree(&HTIDOct)
	yStr := string(y)
	yOct := GetOctet(yStr)
	defer OctetFree(&yOct)
	serverSecretStr := string(serverSecret)
	serverSecretOct := GetOctet(serverSecretStr)
	defer OctetFree(&serverSecretOct)
	UStr := string(U)
	UOct := GetOctet(UStr)
	defer OctetFree(&UOct)
	VStr := string(V)
	VOct := GetOctet(VStr)
	defer OctetFree(&VOct)

	var EOct *C.octet
	var FOct *C.octet
	var UTOct *C.octet
	var publicKeyOct *C.octet

	var pUT C.octet
	defer OctetFree(&pUT)
	var pE C.octet
	defer OctetFree(&pE)
	var pF C.octet
	defer OctetFree(&pF)
	var pPublicKey C.octet
	defer OctetFree(&pPublicKey)

	if publicKey == nil {
		publicKeyOct = nil
	} else {
		publicKeyStr := string(publicKey)
		pPublicKey = GetOctet(publicKeyStr)
		publicKeyOct = &pPublicKey
	}
	if UT == nil {
		UTOct = nil
	} else {
		UTStr := string(UT)
		pUT = GetOctet(UTStr)
		UTOct = &pUT
	}
	if !Kangaroo_ZZZ {
		EOct = nil
		FOct = nil
	} else {
		pE = GetOctetZero(GTS_ZZZ)
		EOct = &pE
		pF = GetOctetZero(GTS_ZZZ)
		FOct = &pF
	}

	rtn := C.MPIN_ZZZ_SERVER_2(C.int(epochDate), &HIDOct, &HTIDOct, &yOct, &serverSecretOct, &UOct, UTOct, &VOct, EOct, FOct, publicKeyOct)

	errorCode = int(rtn)

	if !Kangaroo_ZZZ {
		return errorCode, nil, nil
	} else {
		E = OctetToBytes(EOct)
		F = OctetToBytes(FOct)
		return errorCode, E, F
	}
}

/*

Client1_ZZZ performs first Pass of the client side of the three Pass version of the M-Pin protocol.
If Time Permits are disabled then set epoch_date = 0.In this case UT is not generated0
and can be set to nil. If Time Permits are enabled, and PIN error detection is OFF,
U is not generated and can be set to nil. If Time Permits are enabled and PIN error
detection is ON then U and UT are both generated.

Args:

    hashType: Hash function
    epochDate: Date, in days since the epoch. Set to 0 if Time permits disabled
    mpinId: M-Pin ID
    RNG: cryptographically secure random number generator
    x: random number provided as input if RNG==nil, otherwise must be nil
    PIN: PIN entered by user
    token: M-Pin token
    timePermit: M-Pin time permit

Returns:

    errorCode: error from the C function
    x: Randomly generated integer if RNG!=nil, otherwise must be provided as an input
    U: U = x.H(ID)
    UT: UT = x.(H(ID)+H(epoch_date|H(ID)))
    SEC: SEC = CS+TP, where CS is the reconstructed client secret and TP is the time permit

*/
func Client1_ZZZ(hashType, epochDate int, mpinId []byte, RNG *RandNG, x []byte, PIN int, token []byte, timePermit []byte) (errorCode int, xOut, SEC, U, UT []byte) {
	// Form Octets
	mpinIdStr := string(mpinId)
	mpinIdOct := GetOctet(mpinIdStr)
	defer OctetFree(&mpinIdOct)

	tokenStr := string(token)
	tokenOct := GetOctet(tokenStr)
	defer OctetFree(&tokenOct)

	SECOct := GetOctetZero(G1S_ZZZ)
	defer OctetFree(&SECOct)
	UOct := GetOctetZero(G1S_ZZZ)
	defer OctetFree(&UOct)

	var pUT C.octet
	defer OctetFree(&pUT)
	var pTP C.octet
	defer OctetFree(&pTP)
	var timePermitOct *C.octet
	var UTOct *C.octet

	var xOct C.octet
	defer OctetFree(&xOct)

	var pRNG *C.csprng

	if RNG == nil {
		pRNG = nil
		xStr := string(x)
		xOct = GetOctet(xStr)
	} else {
		pRNG = RNG.csprng()
		xOct = GetOctetZero(PGS_ZZZ)
	}
	if epochDate == 0 {
		timePermitOct = nil
		UTOct = nil
	} else {
		pUT = GetOctetZero(G1S_ZZZ)
		UTOct = &pUT
		TPStr := string(timePermit)
		pTP = GetOctet(TPStr)
		timePermitOct = &pTP

	}

	rtn := C.MPIN_ZZZ_CLIENT_1(C.int(hashType), C.int(epochDate), &mpinIdOct, pRNG, &xOct, C.int(PIN), &tokenOct, &SECOct, &UOct, UTOct, timePermitOct)

	errorCode = int(rtn)
	// Convert octet to bytes
	xOut = OctetToBytes(&xOct)
	SEC = OctetToBytes(&SECOct)
	U = OctetToBytes(&UOct)
	if epochDate == 0 {
		return errorCode, xOut[:], SEC[:], U[:], nil
	} else {
		UT = OctetToBytes(UTOct)
		return errorCode, xOut[:], SEC[:], U[:], UT[:]
	}
}

/*

Client2_ZZZ performs second Pass of the client side of the 3-Pass version of the M-Pin protocol

Args:

    x: locally generated random number
    y: random challenge from server
    SEC: CS+TP, where CS is the reconstructed client secret and TP is the time permit

Returns:

    error_code: error from the C function
    V: V = -(x+y)(CS+TP), where CS is the reconstructed client secret and TP is the time permit

*/
func Client2_ZZZ(x []byte, y []byte, SEC []byte) (errorCode int, V []byte) {
	// Form Octets
	xStr := string(x)
	xOct := GetOctet(xStr)
	defer OctetFree(&xOct)
	yStr := string(y)
	yOct := GetOctet(yStr)
	defer OctetFree(&yOct)
	SECStr := string(SEC)
	SECOct := GetOctet(SECStr)
	defer OctetFree(&SECOct)

	rtn := C.MPIN_ZZZ_CLIENT_2(&xOct, &yOct, &SECOct)

	errorCode = int(rtn)
	// Convert octet to bytes
	V = OctetToBytes(&SECOct)

	return errorCode, V[:]
}
