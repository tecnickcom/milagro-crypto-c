/**
 * @file test_wcc.c
 * @author Kealan McCusker
 * @brief Test WCC with and without time permits
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


/* Test WCC with and without time permits */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include "utils.h"
#include "randapi.h"
#include "wcc_ZZZ.h"

int main(int argc, char** argv)
{
    if (argc != 2)
    {
        printf("usage: ./test_wcc [hash:sha256||sha384||sha512]\n");
        exit(EXIT_FAILURE);
    }

    int i,rtn;

    /* Master secret */
    char ms[WCC_PGS_ZZZ];
    octet MS= {sizeof(ms),sizeof(ms),ms};

    // sender key
    char akeyG1[2*WCC_PFS_ZZZ+1];
    octet AKeyG1= {0,sizeof(akeyG1), akeyG1};

    // receiver key
    char bkeyG2[4*WCC_PFS_ZZZ];
    octet BKeyG2= {0,sizeof(bkeyG2), bkeyG2};

    char hv[WCC_PFS_ZZZ],alice_id[256],bob_id[256];
    octet HV= {0,sizeof(hv),hv};

    octet IdA= {0,sizeof(alice_id),alice_id};
    octet IdB= {0,sizeof(bob_id),bob_id};

    char x[WCC_PGS_ZZZ];
    octet X= {0,sizeof(x),x};
    char y[WCC_PGS_ZZZ];
    octet Y= {0,sizeof(y),y};
    char w[WCC_PGS_ZZZ];
    octet W= {0,sizeof(w),w};
    char pia[WCC_PGS_ZZZ];
    octet PIA= {0,sizeof(pia),pia};
    char pib[WCC_PGS_ZZZ];
    octet PIB= {0,sizeof(pib),pib};

    char pgg1[2*WCC_PFS_ZZZ+1];
    octet PgG1= {0,sizeof(pgg1), pgg1};

    char pag1[2*WCC_PFS_ZZZ+1];
    octet PaG1= {0,sizeof(pag1), pag1};

    char pbg2[4*WCC_PFS_ZZZ];
    octet PbG2= {0,sizeof(pbg2), pbg2};

    char seed[32] = {0};
    octet SEED = {0,sizeof(seed),seed};
    csprng RNG;

    char message1[256];
    octet MESSAGE1 = {0, sizeof(message1), message1};
    OCT_jstring(&MESSAGE1,"Hello Bob");

    char k1[WCC_PAS];  // AES Key
    char k2[WCC_PAS];  // AES Key
    octet K1= {0,sizeof(k1),k1};
    octet K2= {0,sizeof(k2),k2};

    int date, hash;

    int hashDoneOn = 1;
    int hashDoneOff = 0;

    date = 0;

    if (!strcmp(argv[1], "sha256"))
    {
        hash = SHA256;
    }
    else if (!strcmp(argv[1], "sha384"))
    {
        hash = SHA384;
    }
    else
    {
        hash = SHA512;
    }

    BIG_XXX r;

    printf("BIG nbits   %d\n", (int)BIG_XXX_nbits(r));
    printf("BIGBITS_XXX %d\n", (int)BIGBITS_XXX);





    /* unrandom seed value! */
    SEED.len=32;
    for (i=0; i<32; i++) SEED.val[i]=i+1;

    /* initialise random number generator */
    CREATE_CSPRNG(&RNG,&SEED);

    printf("SEED: 0x");
    OCT_output(&SEED);

    /* TA: Generate master secret  */
    rtn = WCC_ZZZ_RANDOM_GENERATE(&RNG,&MS);
    if (rtn != 0)
    {
        printf("TA WCC_ZZZ_RANDOM_GENERATE(&RNG,&MS) Error %d\n", rtn);
        return 1;
    }

    printf("MS: 0x");
    OCT_output(&MS);

    // Alice's ID
    OCT_jstring(&IdA,"alice@miracl.com");

    // TA: Generate Alices's sender key
    HASH_ID(hash,&IdA,&HV);
    rtn = WCC_ZZZ_GET_G1_MULTIPLE(hash,hashDoneOn,&MS,&HV,&AKeyG1);
    if (rtn != 0)
    {
        printf("TA WCC_ZZZ_GET_G1_MULTIPLE() Error %d\n", rtn);
        return 1;
    }

    // Bob's ID
    OCT_jstring(&IdB,"bob@miracl.com");

    // TA: Generate Bob's receiver key
    HASH_ID(hash,&IdB,&HV);
    rtn = WCC_ZZZ_GET_G2_MULTIPLE(hash,hashDoneOn,&MS,&HV,&BKeyG2);
    if (rtn != 0)
    {
        printf("TA WCC_ZZZ_GET_G2_MULTIPLE() Error %d\n", rtn);
        return 1;
    }

    rtn = WCC_ZZZ_RANDOM_GENERATE(&RNG,&X);
    if (rtn != 0)
    {
        printf("Alice WCC_ZZZ_RANDOM_GENERATE(&RNG,&X) Error %d\n", rtn);
        return 1;
    }

    printf("X: 0x");
    OCT_output(&X);

    rtn = WCC_ZZZ_GET_G1_MULTIPLE(hash,hashDoneOff,&X,&IdA,&PaG1);
    if (rtn != 0)
    {
        printf("Alice WCC_ZZZ_GET_G1_MULTIPLE() Error %d\n", rtn);
        return 1;
    }

    rtn = WCC_ZZZ_RANDOM_GENERATE(&RNG,&W);
    if (rtn != 0)
    {
        printf("Bob WCC_ZZZ_RANDOM_GENERATE(&RNG,&W) Error %d\n", rtn);
        return 1;
    }
    rtn = WCC_ZZZ_GET_G1_MULTIPLE(hash,hashDoneOff,&W,&IdA,&PgG1);
    if (rtn != 0)
    {
        printf("Bob WCC_ZZZ_GET_G1_MULTIPLE() Error %d\n", rtn);
        return 1;
    }

    rtn = WCC_ZZZ_RANDOM_GENERATE(&RNG,&Y);
    if (rtn != 0)
    {
        printf("Bob WCC_ZZZ_RANDOM_GENERATE(&RNG,&Y) Error %d\n", rtn);
        return 1;
    }

    rtn = WCC_ZZZ_GET_G2_MULTIPLE(hash,hashDoneOff,&Y,&IdB,&PbG2);
    if (rtn != 0)
    {
        printf("Bob WCC_ZZZ_GET_G1_MULTIPLE() Error %d\n", rtn);
        return 1;
    }

    // pia = Hq(PaG1,PbG2,PgG1,IdB)
    WCC_ZZZ_Hq(hash,&PaG1,&PbG2,&PgG1,&IdB,&PIA);

    // pib = Hq(PbG2,PaG1,PgG1,IdA)
    WCC_ZZZ_Hq(hash,&PbG2,&PaG1,&PgG1,&IdA,&PIB);

    // printf("PIA: ");OCT_output(&PIA);printf("\n");
    // printf("PIB: ");OCT_output(&PIB);printf("\n");

    // Bob calculates AES Key
    WCC_ZZZ_RECEIVER_KEY(hash,date, &Y, &W,  &PIA, &PIB,  &PaG1, &PgG1, &BKeyG2, NULL, &IdA, &K2);
    if (rtn != 0)
    {
        printf("Bob WCC_ZZZ_RECEIVER_KEY() Error %d\n", rtn);
        return 1;
    }

    // pia = Hq(PaG1,PbG2,PgG1,IdB)
    WCC_ZZZ_Hq(hash,&PaG1,&PbG2,&PgG1,&IdB,&PIA);

    // pib = Hq(PbG2,PaG1,PgG1,IdA)
    WCC_ZZZ_Hq(hash,&PbG2,&PaG1,&PgG1,&IdA,&PIB);

    // printf("PIA: ");OCT_output(&PIA);printf("\n");
    // printf("PIB: ");OCT_output(&PIB);printf("\n");

    // Alice calculates AES Key
    rtn = WCC_ZZZ_SENDER_KEY(hash,date, &X, &PIA, &PIB, &PbG2, &PgG1, &AKeyG1, NULL, &IdB, &K1);
    if (rtn != 0)
    {
        printf("Alice WCC_ZZZ_SENDER_KEY() Error %d\n", rtn);
        return 1;
    }

    if (!OCT_comp(&K1,&K2))
    {
        printf("K1: 0x");
        OCT_output(&K1);
        printf("K2: 0x");
        OCT_output(&K2);
        printf("FAILURE No Time Permit Test. OCT_comp(&K1,&K2)\n");
        return 1;
    }

    KILL_CSPRNG(&RNG);

    printf("SUCCESS\n");
    return 0;
}
