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
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const nIter int = 100

func TestOctets_ZZZ(t *testing.T) {
	// Test bad hex
	oct := GetOctetHex("zz")
	assert.Equal(t, 0, int(oct.len), "Invalid hex should have len 0")
	// Test good hex
	oct = GetOctetHex("30")
	assert.Equal(t, 1, int(oct.len), "Hex octed doesn't match")
	oct = GetOctetHex("30")
	assert.Equal(t, 48, int(*oct.val), "Hex octed doesn't match")

	c := GetOctet(string([]byte{1, 2, 3}))
	h := OctetToHex(&c)

	rtn := int(OctetComp([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	assert.Equal(t, rtn, 1, "Value should match 1")

	assert.Equal(t, "010203", h, "Octet convertion should match")
	s := OctetToString(&c)
	assert.Equal(t, "\x01\x02\x03", s, "Octet convertion should match")
}

func TestGenerateRandomByte_ZZZ(t *testing.T) {
	// Seed value for Random Number Generator (RNG)
	seedHex := "9e8b4178790cd57a5761c4a6f164ba72"
	seed, err := hex.DecodeString(seedHex)
	if err != nil {
		fmt.Println("Error decoding seed value")
		return
	}
	rng := CreateCSPRNG(seed)

	b := GenerateRandomByte(&rng, 10)
	assert.Equal(t, "57d662d39b1b245da469", fmt.Sprintf("%x", b), "Should be equal")
}

func TestGenerateOTP_ZZZ(t *testing.T) {
	// Seed value for Random Number Generator (RNG)
	seedHex := "9e8b4178790cd57a5761c4a6f164ba72"
	seed, err := hex.DecodeString(seedHex)
	if err != nil {
		fmt.Println("Error decoding seed value")
		return
	}
	rng := CreateCSPRNG(seed)

	otp := GenerateOTP(&rng)
	assert.Equal(t, 715827, otp, "Should be equal")
}

func TestAesGcm_ZZZ(t *testing.T) {
	testCases := []struct {
		Key       string
		IV        string
		AAD       string
		PlainText string
		outC      string
		outT      string
	}{
		{
			Key:       "75eae8f5ec7a5c8882f4a389600da8cd",
			IV:        "a231cd40fdb909ad11c457d9",
			AAD:       "a50d3e45b23c77157cb0e01c2a679e6d99c038e4",
			PlainText: "45cf12964fc824ab76616ae2f4bf0822",
			outC:      "5df1c20786beb4dc24bab9caf2ad3a03",
			outT:      "ba8036433389c88b79e277f49a4bc41d",
		},
	}

	hexDecode := func(s string) []byte {
		k, err := hex.DecodeString(s)
		if err != nil {
			t.Fatal("Decode hex error")
		}
		return k
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case %v", i), func(t *testing.T) {
			C, T, err := AesGcmEncrypt(hexDecode(tc.Key), hexDecode(tc.IV), hexDecode(tc.AAD), hexDecode(tc.PlainText))
			if err != nil {
				t.Fatal("Fatal error on AES Key length")
			}
			assert.Equal(t, tc.outC, hex.EncodeToString(C), "Should be equal")
			assert.Equal(t, tc.outT, hex.EncodeToString(T), "Should be equal")

			dP, dT, err := AesGcmDecrypt(hexDecode(tc.Key), hexDecode(tc.IV), hexDecode(tc.AAD), hexDecode(tc.outC))
			if err != nil {
				t.Fatal("Fatal error on AES Key length")
			}
			assert.Equal(t, tc.PlainText, hex.EncodeToString(dP), "Should be equal")
			assert.Equal(t, tc.outT, hex.EncodeToString(dT), "Should be equal")
		})
	}
}

func TestBadAesGcm_ZZZ(t *testing.T) {
	testCases := []struct {
		Key       string
		IV        string
		AAD       string
		PlainText string
		outC      string
	}{
		{
			Key:       "75eae8f5ec7a5c88",
			IV:        "a231cd40fdb909ad11c457d9",
			AAD:       "a50d3e45b23c77157cb0e01c2a679e6d99c038e4",
			PlainText: "45cf12964fc824ab76616ae2f4bf0822",
			outC:      "5df1c20786beb4dc24bab9caf2ad3a03",
		},
		{
			Key:       "75eae8f5ec7a5c8875eae8f5ec7a5c8875eae8f5ec7a5c88",
			IV:        "a231cd40fdb909ad11c457d9",
			AAD:       "a50d3e45b23c77157cb0e01c2a679e6d99c038e4",
			PlainText: "45cf12964fc824ab76616ae2f4bf0822",
			outC:      "5df1c20786beb4dc24bab9caf2ad3a03",
		},
	}

	hexDecode := func(s string) []byte {
		k, err := hex.DecodeString(s)
		if err != nil {
			t.Fatal("Decode hex error")
		}
		return k
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case %v", i), func(t *testing.T) {
			_, _, err := AesGcmEncrypt(hexDecode(tc.Key), hexDecode(tc.IV), hexDecode(tc.AAD), hexDecode(tc.PlainText))
			if err == nil {
				t.Fatal("Fatal error on AES Key length")
			}

			_, _, err = AesGcmDecrypt(hexDecode(tc.Key), hexDecode(tc.IV), hexDecode(tc.AAD), hexDecode(tc.outC))
			if err == nil {
				t.Fatal("Fatal error on AES Key length")
			}
		})
	}
}
