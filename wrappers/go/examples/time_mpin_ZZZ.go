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

package main

import (
	"encoding/hex"
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/miracl/amcl-go-wrapper"
)

// Number of iterations to time functions
const nIter int = 1000

var HASH_TYPE_MPIN = amcl.SHA256

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	date := 0
	timeValue := 1488278706
	SSHex := "07f8181687f42ce22ea0dee4ba9df3f2cea67ad2d79e59adc953142556d510831bbd59e9477ac479019887020579aed16af43dc7089ae8c14262e64b5d09740109917efd0618c557fbf7efaa68fb64e8d46b3766bb184dea9bef9638f23bbbeb03aedbc6e4eb9fbd658719aab26b849638690521723c0efb9c8622df2a8efa3c"
	UHex := "041dcd4592280dd05b7eb256f91b4e79fe85b415390f3728d0ef126d5d32b939cc21dd0cedd6ba17dafa1b297e83f3238c2bbccf5b5e60f4c04b97e0fb08bd2acd"
	UTHex := "0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	VHex := "042260824c0f70cc0cc7fc586fd70ec23c05db1ab7f7f3d5fbe506e8011b7e04cb06f85f733904a30dd905281252a6a3557b49de9c53b65082a145c8ee81e22167"
	IDHex := "7465737455736572406d697261636c2e636f6d"

	SS, _ := hex.DecodeString(SSHex)
	U, _ := hex.DecodeString(UHex)
	UT, _ := hex.DecodeString(UTHex)
	V, _ := hex.DecodeString(VHex)
	ID, _ := hex.DecodeString(IDHex)
	var MESSAGE []byte

	t0 := time.Now()
	var rtn int
	for i := 0; i < nIter; i++ {
		rtn, _, _, _, _, _ = amcl.Server_ZZZ(HASH_TYPE_MPIN, date, timeValue, SS[:], U[:], UT[:], V[:], ID[:], nil, MESSAGE[:], false)
	}
	t1 := time.Now()
	log.Printf("Number Iterations: %d Time: %v\n", nIter, t1.Sub(t0))

	if rtn != 0 {
		log.Printf("Authentication failed Error Code %d\n", rtn)
		return
	} else {
		log.Printf("Authenticated Error Code %d\n", rtn)
	}
}
