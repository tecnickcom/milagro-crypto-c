/**
 * @file version.h
 * @author Kealan McCusker
 * @date 28th April 2016
 * @brief Build information
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
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/* AMCL header support version function */

#ifndef VERSION_H
#define VERSION_H

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "amcl.h"

// /**
//  * Curve types
//  */
// #if CURVETYPE==WEIERSTRASS
// #define CURVETYPE_DESC "WEIERSTRASS"
// #elif CURVETYPE==EDWARDS
// #define CURVETYPE_DESC "EDWARDS"
// #elif CURVETYPE==MONTGOMERY
// #define CURVETYPE_DESC "MONTGOMERY"
// #else
// #define CURVETYPE_DESC ""
// #endif

// /**
//  * Curve choice description
//  */
// #if CHOICE==BN254
// #define CHOICE_DESC "BN254"
// #elif CHOICE==BN254_T
// #define CHOICE_DESC "BN254_T"
// #elif CHOICE==BN254_T2
// #define CHOICE_DESC "BN254_T2"
// #elif CHOICE==BN254_CX
// #define CHOICE_DESC "BN254_CX"
// #elif CHOICE==NIST256
// #define CHOICE_DESC "NIST256"
// #elif CHOICE==MF254
// #define CHOICE_DESC "MF254"
// #elif CHOICE==MF256
// #define CHOICE_DESC "MF256"
// #elif CHOICE==MS255
// #define CHOICE_DESC "MS255"
// #elif CHOICE==MS256
// #define CHOICE_DESC "MS256"
// #elif CHOICE==C25519
// #define CHOICE_DESC "C25519"
// #elif CHOICE==BRAINPOOL
// #define CHOICE_DESC "BRAINPOOL"
// #elif CHOICE==ANSSI
// #define CHOICE_DESC "ANSSI"
// #elif CHOICE==HIFIVE
// #define CHOICE_DESC "HIFIVE"
// #elif CHOICE==GOLDILOCKS
// #define CHOICE_DESC "GOLDILOCKS"
// #elif CHOICE==NIST384
// #define CHOICE_DESC "NIST384"
// #elif CHOICE==C41417
// #define CHOICE_DESC "C41417"
// #elif CHOICE==NIST521
// #define CHOICE_DESC "NIST521"
// #elif CHOICE==BN646
// #define CHOICE_DESC "BN646"
// #elif CHOICE==BN454
// #define CHOICE_DESC "BN454"
// #elif CHOICE==BLS455
// #define CHOICE_DESC "BLS455"
// #elif CHOICE==BLS383
// #define CHOICE_DESC "BLS383"
// #else
// #define CHOICE_DESC ""
// #endif

// /**
//  * Modulus type
//  */
// #if MODTYPE==NOT_SPECIAL
// #define MODTYPE_DESC "NOT_SPECIAL - Modulus of no exploitable form"
// #elif MODTYPE==PSEUDO_MERSENNE
// #define MODTYPE_DESC "PSEUDO_MERSENNE - Pseudo-mersenne modulus of form $2^n-c$"
// #elif MODTYPE==GENERALISED_MERSENNE
// #define MODTYPE_DESC "GENERALISED_MERSENNE - Generalised-mersenne modulus of form $2^n-2^m-1$"
// #elif MODTYPE==MONTGOMERY_FRIENDLY
// #define MODTYPE_DESC "MONTGOMERY_FRIENDLY - Montgomery Friendly modulus of form $2^a(2^b-c)-1$"
// #else
// #define MODTYPE_DESC ""
// #endif

/**
 * @brief Print version number and information about the build
 *
 * Print version number and information about the build.
 */
void amcl_version(void);

#endif
