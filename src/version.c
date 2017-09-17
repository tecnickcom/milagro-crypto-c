/**
 * @file version.c
 * @author Mike Scott
 * @author Kealan McCusker
 * @date 28th April 2016
 * @brief AMCL version support function
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

#include "version.h"

/* AMCL version support function */

/* Print version number and information about the build */
void amcl_version(void)
{
    printf("AMCL Version: %d.%d.%d\n", AMCL_VERSION_MAJOR, AMCL_VERSION_MINOR, AMCL_VERSION_PATCH);
    printf("OS: %s\n", OS);
    printf("CHUNK: %d\n", CHUNK);
    // printf("CURVETYPE: %s\n", CURVETYPE_DESC);
    // printf("CHOICE: %s\n", CHOICE_DESC);
    // printf("MODTYPE: %s\n", MODTYPE_DESC);
    // printf("MBITS - Number of bits in Modulus: %d\n", MBITS);
    // printf("MODBYTES - Number of bytes in Modulus: %d\n", MODBYTES);
    // printf("BASEBITS - Numbers represented to base 2*BASEBITS: %d\n", BASEBITS);
    // printf("NLEN - Number of words in BIG: %d\n", NLEN);

    // BIG p, r;
    // BIG_rcopy(p,Modulus);
    // printf("Modulus p = ");
    // BIG_output(p);
    // printf("\n");
    // BIG_rcopy(r,CURVE_Order);
    // printf("CURVE_Order r = ");
    // BIG_output(r);
    // printf("\n");
}
