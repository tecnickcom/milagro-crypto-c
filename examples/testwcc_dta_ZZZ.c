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


/* Demonstrate WCC with two TAs and time permits */

/* Build executible after installation:
   gcc -std=c99 -g testwcc_dta.c  -I/opt/amcl/include -L/opt/amcl/lib -lamcl_wcc -lamcl_pairing -lamcl_curve -lamcl_core -o testwcc_dta */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include "wcc_ZZZ.h"
#include "randapi.h"

#define DEBUG

int main()
{
    int i,rtn;

    /* Master secret shares */
    char ms1[PGS_ZZZ], ms2[PGS_ZZZ];
    octet MS1= {sizeof(ms1),sizeof(ms1),ms1};
    octet MS2= {sizeof(ms2),sizeof(ms2),ms2};

    // Sender keys
    char a1keyG1[2*PFS_ZZZ+1], a2keyG1[2*PFS_ZZZ+1];
    octet A1KeyG1= {0,sizeof(a1keyG1), a1keyG1};
    octet A2KeyG1= {0,sizeof(a2keyG1), a2keyG1};
    char akeyG1[2*PFS_ZZZ+1];
    octet AKeyG1= {0,sizeof(akeyG1), akeyG1};

    // Sender time permits
    char a1TPG1[2*PFS_ZZZ+1], a2TPG1[2*PFS_ZZZ+1];
    octet A1TPG1= {sizeof(a1TPG1),sizeof(a1TPG1), a1TPG1};
    octet A2TPG1= {sizeof(a2TPG1),sizeof(a2TPG1), a2TPG1};
    char aTPG1[2*PFS_ZZZ+1];
    octet ATPG1= {0,sizeof(aTPG1), aTPG1};

    // Receiver keys
    char b1keyG2[4*PFS_ZZZ], b2keyG2[4*PFS_ZZZ];
    octet B1KeyG2= {0,sizeof(b1keyG2), b1keyG2};
    octet B2KeyG2= {0,sizeof(b2keyG2), b2keyG2};
    char bkeyG2[4*PFS_ZZZ];
    octet BKeyG2= {0,sizeof(bkeyG2), bkeyG2};

    // Receiver time permits
    char b1TPG2[4*PFS_ZZZ], b2TPG2[4*PFS_ZZZ];
    octet B1TPG2= {sizeof(b1TPG2),sizeof(b1TPG2), b1TPG2};
    octet B2TPG2= {sizeof(b2TPG2),sizeof(b2TPG2), b2TPG2};
    char bTPG2[4*PFS_ZZZ];
    octet BTPG2= {0,sizeof(bTPG2), bTPG2};

    char ahv[PFS_ZZZ],alice_id[256],bhv[PFS_ZZZ],bob_id[256];
    octet AHV= {0,sizeof(ahv),ahv};
    octet BHV= {0,sizeof(bhv),bhv};

    octet IdA= {0,sizeof(alice_id),alice_id};
    octet IdB= {0,sizeof(bob_id),bob_id};

    char x[PGS_ZZZ];
    octet X= {0,sizeof(x),x};
    char y[PGS_ZZZ];
    octet Y= {0,sizeof(y),y};
    char w[PGS_ZZZ];
    octet W= {0,sizeof(w),w};
    char pia[PGS_ZZZ];
    octet PIA= {0,sizeof(pia),pia};
    char pib[PGS_ZZZ];
    octet PIB= {0,sizeof(pib),pib};

    char pgg1[2*PFS_ZZZ+1];
    octet PgG1= {0,sizeof(pgg1), pgg1};

    char pag1[2*PFS_ZZZ+1];
    octet PaG1= {0,sizeof(pag1), pag1};

    char pbg2[4*PFS_ZZZ];
    octet PbG2= {0,sizeof(pbg2), pbg2};

    char seed[32] = {0};
    octet SEED = {0,sizeof(seed),seed};
    csprng RNG;

    char message1[256];
    char message2[256];
    octet MESSAGE1 = {0, sizeof(message1), message1};
    octet MESSAGE2 = {0, sizeof(message2), message2};

    char t1[PTAG];  // Tag
    char t2[PTAG];  // Tag
    char k1[PAS];  // AES Key
    char k2[PAS];  // AES Key
    char iv[PIV]; // IV - Initialisation vector
    char c[100];  // Ciphertext
    char p[100];  // Recovered Plaintext
    octet T1= {sizeof(t1),sizeof(t1),t1};
    octet T2= {sizeof(t2),sizeof(t2),t2};
    octet K1= {0,sizeof(k1),k1};
    octet K2= {0,sizeof(k2),k2};
    octet IV= {0,sizeof(iv),iv};
    octet C= {0,sizeof(c),c};
    octet P= {0,sizeof(p),p};

    int date;
    date = WCC_today();
    printf("Date %d \n", date);

    int hashDoneOn = 1;

    OCT_jstring(&MESSAGE1,"Hello Bob");
    OCT_jstring(&MESSAGE2,"Hello Alice");

    /* unrandom seed value! */
    SEED.len=32;
    for (i=0; i<32; i++) SEED.val[i]=i+1;

    /* initialise random number generator */
    CREATE_CSPRNG(&RNG,&SEED);

    /* Generate Client master secret for MIRACL and Customer */
    rtn = WCC_ZZZ_RANDOM_GENERATE(&RNG,&MS1);
    if (rtn != 0)
    {
        printf("TA WCC_ZZZ_RANDOM_GENERATE(&RNG,&MS1) Error %d\n", rtn);
        return 1;
    }
    rtn = WCC_ZZZ_RANDOM_GENERATE(&RNG,&MS2);
    if (rtn != 0)
    {
        printf("TA WCC_ZZZ_RANDOM_GENERATE(&RNG,&MS2) Error %d\n", rtn);
        return 1;
    }
    printf("TA MASTER SECRET MIRACL: ");
    OCT_output(&MS1);
    printf("TA MASTER SECRET CUSTOMER: ");
    OCT_output(&MS2);

    // Alice's ID
    OCT_jstring(&IdA,"alice@miracl.com");

    // TA: Generate Alice's sender key
    HASH_ID(HASH_TYPE_WCC_ZZZ,&IdA,&AHV);
    rtn = WCC_ZZZ_GET_G1_MULTIPLE(HASH_TYPE_WCC_ZZZ,hashDoneOn,&MS1,&AHV,&A1KeyG1);
    if (rtn != 0)
    {
        printf("TA WCC_ZZZ_GET_G1_MULTIPLE(hashDoneOn,&MS1,&AHV,&A1KeyG1) Error %d\n", rtn);
        return 1;
    }
    rtn = WCC_ZZZ_GET_G1_MULTIPLE(HASH_TYPE_WCC_ZZZ,hashDoneOn,&MS2,&AHV,&A2KeyG1);
    if (rtn != 0)
    {
        printf("TA WCC_ZZZ_GET_G1_MULTIPLE(hashDoneOn,&MS2,&AHV,&A2KeyG1) Error %d\n", rtn);
        return 1;
    }
    printf("TA A1KeyG1: ");
    OCT_output(&A1KeyG1);
    printf("TA A2KeyG1: ");
    OCT_output(&A2KeyG1);

    rtn = WCC_RECOMBINE_G1(&A1KeyG1, &A2KeyG1, &AKeyG1);
    if (rtn != 0)
    {
        printf("TA WCC_RECOMBINE_G1(&A1KeyG1, &A2KeyG1, &AKeyG1) Error %d\n", rtn);
        return 1;
    }
    printf("TA Alice's sender key: ");
    OCT_output(&AKeyG1);

    // TA: Generate Alice's G1 time permit
    rtn = WCC_ZZZ_GET_G1_PERMIT(HASH_TYPE_WCC_ZZZ,date,&MS1,&AHV,&A1TPG1);
    if (rtn != 0)
    {
        printf("TA WCC_ZZZ_GET_G1_PERMIT(HASH_TYPE_WCC_ZZZ,date,&MS1,&AHV,&A1TPG1) Error %d\n", rtn);
        return 1;
    }
    rtn = WCC_ZZZ_GET_G1_PERMIT(HASH_TYPE_WCC_ZZZ,date,&MS2,&AHV,&A2TPG1);
    if (rtn != 0)
    {
        printf("TA WCC_ZZZ_GET_G1_PERMIT(HASH_TYPE_WCC_ZZZ,date,&MS2,&AHV,&A2TPG1) Error %d\n", rtn);
        return 1;
    }
    printf("TA A1TPG1: ");
    OCT_output(&A1TPG1);
    printf("TA A2TPG1: ");
    OCT_output(&A2TPG1);

    rtn = WCC_RECOMBINE_G1(&A1TPG1, &A2TPG1, &ATPG1);
    if (rtn != 0)
    {
        printf("Alice WCC_RECOMBINE_G1(&A1TPG1, &A2TPG1, &ATPG1) Error %d\n", rtn);
        return 1;
    }
    printf("TA Alice's sender time permit: ");
    OCT_output(&ATPG1);

    // Bob's ID
    OCT_jstring(&IdB,"bob@miracl.com");

    // TA: Generate Bob's receiver key
    HASH_ID(HASH_TYPE_WCC_ZZZ,&IdB,&BHV);
    rtn = WCC_ZZZ_GET_G2_MULTIPLE(HASH_TYPE_WCC_ZZZ,hashDoneOn,&MS1,&BHV,&B1KeyG2);
    if (rtn != 0)
    {
        printf("TA WCC_ZZZ_GET_G2_MULTIPLE(HASH_TYPE_WCC_ZZZ,hashDoneOn,&MS1,&BHV,&B1KeyG2) Error %d\n", rtn);
        return 1;
    }
    rtn = WCC_ZZZ_GET_G2_MULTIPLE(HASH_TYPE_WCC_ZZZ,hashDoneOn,&MS2,&BHV,&B2KeyG2);
    if (rtn != 0)
    {
        printf("Bob WCC_ZZZ_GET_G2_MULTIPLE(HASH_TYPE_WCC_ZZZ,hashDoneOn,&MS2,&BHV,&B2KeyG2) Error %d\n", rtn);
        return 1;
    }
    printf("TA B1KeyG2: ");
    OCT_output(&B1KeyG2);
    printf("TA B2KeyG2: ");
    OCT_output(&B2KeyG2);

    rtn = WCC_RECOMBINE_G2(&B1KeyG2, &B2KeyG2, &BKeyG2);
    if (rtn != 0)
    {
        printf("Bob WCC_RECOMBINE_G2(&B1KeyG1, &B2KeyG1, &BKeyG2) Error %d\n", rtn);
        return 1;
    }
    printf("TA Bob's receiver key: ");
    OCT_output(&BKeyG2);

    // TA: Generate Bob's receiver time permit
    rtn = WCC_ZZZ_GET_G2_PERMIT(HASH_TYPE_WCC_ZZZ,date,&MS1,&BHV,&B1TPG2);
    if (rtn != 0)
    {
        printf("TA WCC_ZZZ_GET_G2_PERMIT(HASH_TYPE_WCC_ZZZ,date,&MS1,&BHV,&B1TPG2) Error %d\n", rtn);
        return 1;
    }
    rtn = WCC_ZZZ_GET_G2_PERMIT(HASH_TYPE_WCC_ZZZ,date,&MS2,&BHV,&B2TPG2);
    if (rtn != 0)
    {
        printf("TA WCC_ZZZ_GET_G2_PERMIT(HASH_TYPE_WCC_ZZZ,date,&MS2,&BHV,&B2TPG2) Error %d\n", rtn);
        return 1;
    }
    printf("TA B1TPG2: ");
    OCT_output(&B1TPG2);
    printf("TA B2TPG2: ");
    OCT_output(&B2TPG2);

    rtn = WCC_RECOMBINE_G2(&B1TPG2, &B2TPG2, &BTPG2);
    if (rtn != 0)
    {
        printf("Bob WCC_RECOMBINE_G2(&B1TPG2, &B2TPG2, &BTPG2) Error %d\n", rtn);
        return 1;
    }
    printf("TA Bob's receiver time permit: ");
    OCT_output(&BTPG2);
    printf("\n");

    printf("Alice\n");

    rtn = WCC_ZZZ_RANDOM_GENERATE(&RNG,&X);
    if (rtn != 0)
    {
        printf("Alice WCC_ZZZ_RANDOM_GENERATE(&RNG,&X) Error %d\n", rtn);
        return 1;
    }
#ifdef DEBUG
    printf("Alice X: ");
    OCT_output(&X);
    printf("\n");
#endif

    rtn = WCC_ZZZ_GET_G1_TPMULT(HASH_TYPE_WCC_ZZZ,date,&X,&IdA,&PaG1);
    if (rtn != 0)
    {
        printf("Alice WCC_ZZZ_GET_G1_TPMULT(HASH_TYPE_WCC_ZZZ,date,&X,&IdA,&PaG1) Error %d\n", rtn);
        return 1;
    }

    printf("Alice sends IdA and PaG1 to Bob\n\n");
    printf("Alice IdA: ");
    OCT_output_string(&IdA);
    printf("\n");
    printf("Alice PaG1: ");
    OCT_output(&PaG1);
    printf("\n");

    printf("Bob\n");

    rtn = WCC_ZZZ_RANDOM_GENERATE(&RNG,&W);
    if (rtn != 0)
    {
        printf("Bob WCC_ZZZ_RANDOM_GENERATE(&RNG,&W) Error %d\n", rtn);
        return 1;
    }
#ifdef DEBUG
    printf("Bob W: ");
    OCT_output(&W);
    printf("\n");
#endif
    rtn = WCC_ZZZ_GET_G1_TPMULT(HASH_TYPE_WCC_ZZZ,date,&W,&IdA,&PgG1);
    if (rtn != 0)
    {
        printf("Bob WCC_ZZZ_GET_G1_TPMULT(HASH_TYPE_WCC_ZZZ,date,&W,&IdA,&PgG1) Error %d\n", rtn);
        return 1;
    }
#ifdef DEBUG
    printf("PgG1: ");
    OCT_output(&PgG1);
    printf("\n");
#endif

    rtn = WCC_ZZZ_RANDOM_GENERATE(&RNG,&Y);
    if (rtn != 0)
    {
        printf("Bob WCC_ZZZ_RANDOM_GENERATE(&RNG,&Y) Error %d\n", rtn);
        return 1;
    }
#ifdef DEBUG
    printf("Bob Y: ");
    OCT_output(&Y);
    printf("\n");
#endif
    rtn = WCC_ZZZ_GET_G2_TPMULT(HASH_TYPE_WCC_ZZZ,date,&Y,&IdB,&PbG2);
    if (rtn != 0)
    {
        printf("Bob WCC_ZZZ_GET_G1_TPMULT(date,&Y,&IdB,&PbG2) Error %d\n", rtn);
        return 1;
    }
#ifdef DEBUG
    printf("Bob PbG2: ");
    OCT_output(&PbG2);
    printf("\n");
#endif

    // pia = Hq(PaG1,PbG2,PgG1,IdB)
    WCC_ZZZ_Hq(HASH_TYPE_WCC_ZZZ,&PaG1,&PbG2,&PgG1,&IdB,&PIA);

    // pib = Hq(PbG2,PaG1,PgG1,IdA)
    WCC_ZZZ_Hq(HASH_TYPE_WCC_ZZZ,&PbG2,&PaG1,&PgG1,&IdA,&PIB);

#ifdef DEBUG
    printf("Bob PIA: ");
    OCT_output(&PIA);
    printf("\n");
    printf("Bob PIB: ");
    OCT_output(&PIB);
    printf("\n");
#endif

    // Bob calculates AES Key
    WCC_ZZZ_RECEIVER_KEY(HASH_TYPE_WCC_ZZZ, date, &Y, &W,  &PIA, &PIB,  &PaG1, &PgG1, &BKeyG2, &BTPG2, &IdA, &K2);
    if (rtn != 0)
    {
        printf("Bob WCC_ZZZ_RECEIVER_KEY() Error %d\n", rtn);
        return 1;
    }
    printf("Bob AES Key: ");
    OCT_output(&K2);

    printf("Bob sends IdB, PbG2 and PgG1 to Alice\n\n");
    printf("Bob IdB: ");
    OCT_output_string(&IdB);
    printf("\n");
    printf("Bob PbG2: ");
    OCT_output(&PbG2);
    printf("\n");
    printf("Bob PgG1: ");
    OCT_output(&PgG1);
    printf("\n");

    printf("Alice\n");

    // pia = Hq(PaG1,PbG2,PgG1,IdB)
    WCC_ZZZ_Hq(HASH_TYPE_WCC_ZZZ,&PaG1,&PbG2,&PgG1,&IdB,&PIA);

    // pib = Hq(PbG2,PaG1,PgG1,IdA)
    WCC_ZZZ_Hq(HASH_TYPE_WCC_ZZZ,&PbG2,&PaG1,&PgG1,&IdA,&PIB);

#ifdef DEBUG
    printf("Alice PIA: ");
    OCT_output(&PIA);
    printf("\n");
    printf("Alice PIB: ");
    OCT_output(&PIB);
    printf("\n");
#endif

    // Alice calculates AES Key
    rtn = WCC_ZZZ_SENDER_KEY(HASH_TYPE_WCC_ZZZ, date, &X, &PIA, &PIB, &PbG2, &PgG1, &AKeyG1, &ATPG1, &IdB, &K1);
    if (rtn != 0)
    {
        printf("Alice WCC_ZZZ_SENDER_KEY() Error %d\n", rtn);
        return 1;
    }
    printf("Alice AES Key: ");
    OCT_output(&K1);

    // Send message
    IV.len=12;
    for (i=0; i<IV.len; i++)
        IV.val[i]=i+1;
    printf("Alice: IV ");
    OCT_output(&IV);

    printf("Alice: Message to encrypt for Bob: ");
    OCT_output_string(&MESSAGE1);
    printf("\n");

    AES_GCM_ENCRYPT(&K1, &IV, &IdA, &MESSAGE1, &C, &T1);

    printf("Alice: Ciphertext: ");
    OCT_output(&C);

    printf("Alice: Encryption Tag: ");
    OCT_output(&T1);
    printf("\n");

    AES_GCM_DECRYPT(&K2, &IV, &IdA, &C, &P, &T2);

    printf("Bob: Decrypted message received from Alice: ");
    OCT_output_string(&P);
    printf("\n");

    printf("Bob: Decryption Tag: ");
    OCT_output(&T2);
    printf("\n");

    if (!OCT_comp(&MESSAGE1,&P))
    {
        printf("FAILURE Decryption\n");
        return 1;
    }

    if (!OCT_comp(&T1,&T2))
    {
        printf("FAILURE TAG mismatch\n");
        return 1;
    }

    /* clear memory */
    OCT_clear(&MS1);
    OCT_clear(&MS2);
    OCT_clear(&A1KeyG1);
    OCT_clear(&A2KeyG1);
    OCT_clear(&AKeyG1);
    OCT_clear(&B1KeyG2);
    OCT_clear(&B2KeyG2);
    OCT_clear(&BKeyG2);
    OCT_clear(&AHV);
    OCT_clear(&BHV);
    OCT_clear(&IdA);
    OCT_clear(&IdB);
    OCT_clear(&X);
    OCT_clear(&Y);
    OCT_clear(&W);
    OCT_clear(&PIA);
    OCT_clear(&PIB);
    OCT_clear(&PgG1);
    OCT_clear(&PaG1);
    OCT_clear(&PbG2);
    OCT_clear(&SEED);
    OCT_clear(&MESSAGE1);
    OCT_clear(&MESSAGE2);
    OCT_clear(&T1);
    OCT_clear(&T2);
    OCT_clear(&K1);
    OCT_clear(&K2);
    OCT_clear(&IV);
    OCT_clear(&C);
    OCT_clear(&P);

    KILL_CSPRNG(&RNG);

    return 0;
}
