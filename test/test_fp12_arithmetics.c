/**
 * @file test_fp_arithmetics.c
 * @author Alessandro Budroni
 * @brief Test for aritmetics with FP
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


#include "arch.h"
#include "amcl.h"
#include "utils.h"
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#define LINE_LEN 5000

void read_BIG(BIG A, char* string)
{
    int len;
    char support[LINE_LEN];
    BIG_zero(A);
    len = strlen(string)+1;
    amcl_hex2bin(string,support,len);
    len = (len-1)/2;;
    BIG_fromBytesLen(A,support,len);
    BIG_norm(A);
}

// Read a structure of the type [..,..]
void read_FP2(FP2 *fp2, char* stringx)
{
    char *stringy, *end;
    BIG x,y;

    stringx++;
    stringy = strchr(stringx,',');
    if (stringy == NULL)
    {
        printf("ERROR unexpected test vector\n");
        exit(EXIT_FAILURE);
    }
    stringy[0] = '\0';
    stringy++;
    end = strchr(stringy,']');
    if (end == NULL)
    {
        printf("ERROR unexpected test vector\n");
        exit(EXIT_FAILURE);
    }
    end[0] = '\0';

    read_BIG(x,stringx);
    read_BIG(y,stringy);

    FP2_from_BIGs(fp2,x,y);
}

// Read a structure of the type [[..,..],[..,..]]
void read_FP4(FP4 *fp4, char* stringx1)
{
    char *stringx2, *stringy1, *stringy2, *end;
    BIG x1,x2,y1,y2;
    FP2 x,y;

    stringx1 += 2;
    stringx2 = strchr(stringx1,',');
    if (stringx2 == NULL)
    {
        printf("ERROR unexpected test vector\n");
        exit(EXIT_FAILURE);
    }
    stringx2[0] = '\0';
    stringx2 ++;
    stringy1 = strchr(stringx2,']');
    if (stringy1 == NULL)
    {
        printf("ERROR unexpected test vector\n");
        exit(EXIT_FAILURE);
    }
    stringy1[0] = '\0';
    stringy1 += 3;
    stringy2 = strchr(stringy1,',');
    if (stringy2 == NULL)
    {
        printf("ERROR unexpected test vector\n");
        exit(EXIT_FAILURE);
    }
    stringy2[0] = '\0';
    stringy2++;
    end = strchr(stringy2,']');
    if (end == NULL)
    {
        printf("ERROR unexpected test vector\n");
        exit(EXIT_FAILURE);
    }
    end[0] = '\0';

    read_BIG(x1,stringx1);
    read_BIG(x2,stringx2);
    read_BIG(y1,stringy1);
    read_BIG(y2,stringy2);

    FP2_from_BIGs(&x,x1,x2);
    FP2_from_BIGs(&y,y1,y2);

    FP4_from_FP2s(fp4,&x,&y);
}

// Read a structure of the type [[[ax1,ax2],[ay1,ay2]],[[bx1,bx2],[by1,by2]],[[cx1,cx2],[cy1,cy2]]]
void read_FP12(FP12 *fp12, char *stringax1)
{
    char *stringax2, *stringay1, *stringay2, *stringbx1, *stringbx2, *stringby1, *stringby2, *stringcx1, *stringcx2, *stringcy1, *stringcy2, *end;
    BIG ax1,ax2,ay1,ay2,bx1,bx2,by1,by2,cx1,cx2,cy1,cy2;
    FP2 ax,ay,bx,by,cx,cy;
    FP4 a,b,c;

    stringax1 += 3;
    stringax2 = strchr(stringax1,',');
    if (stringax2 == NULL)
    {
        printf("ERROR unexpected test vector\n");
        exit(EXIT_FAILURE);
    }
    stringax2[0] = '\0';
    stringax2++;
    stringay1 = strchr(stringax2,']');
    if (stringay1 == NULL)
    {
        printf("ERROR unexpected test vector\n");
        exit(EXIT_FAILURE);
    }
    stringay1[0] = '\0';
    stringay1 += 3;
    stringay2 = strchr(stringay1,',');
    if (stringay2 == NULL)
    {
        printf("ERROR unexpected test vector\n");
        exit(EXIT_FAILURE);
    }
    stringay2[0] = '\0';
    stringay2++;
    stringbx1 = strchr(stringay2,']');
    if (stringbx1 == NULL)
    {
        printf("ERROR unexpected test vector\n");
        exit(EXIT_FAILURE);
    }
    stringbx1[0] = '\0';
    stringbx1 += 5;
    stringbx2 = strchr(stringbx1,',');
    if (stringbx2 == NULL)
    {
        printf("ERROR unexpected test vector\n");
        exit(EXIT_FAILURE);
    }
    stringbx2[0] = '\0';
    stringbx2++;
    stringby1 = strchr(stringbx2,']');
    if (stringay1 == NULL)
    {
        printf("ERROR unexpected test vector\n");
        exit(EXIT_FAILURE);
    }
    stringby1[0] = '\0';
    stringby1 += 3;
    stringby2 = strchr(stringby1,',');
    if (stringay2 == NULL)
    {
        printf("ERROR unexpected test vector\n");
        exit(EXIT_FAILURE);
    }
    stringby2[0] = '\0';
    stringby2++;
    stringcx1 = strchr(stringby2,']');
    if (stringcx1 == NULL)
    {
        printf("ERROR unexpected test vector\n");
        exit(EXIT_FAILURE);
    }
    stringcx1++;
    stringcx1[0] = '\0';
    stringcx1 += 5;
    stringcx2 = strchr(stringcx1,',');
    if (stringcx2 == NULL)
    {
        printf("ERROR unexpected test vector\n");
        exit(EXIT_FAILURE);
    }
    stringcx2[0] = '\0';
    stringcx2++;
    stringcy1 = strchr(stringcx2,']');
    if (stringcy1 == NULL)
    {
        printf("ERROR unexpected test vector\n");
        exit(EXIT_FAILURE);
    }
    stringcy1[0] = '\0';
    stringcy1 += 3;
    stringcy2 = strchr(stringcy1,',');
    if (stringay2 == NULL)
    {
        printf("ERROR unexpected test vector\n");
        exit(EXIT_FAILURE);
    }
    stringcy2[0] = '\0';
    stringcy2++;
    end = strchr(stringcy2,']');
    if (end == NULL)
    {
        printf("ERROR unexpected test vector\n");
        exit(EXIT_FAILURE);
    }
    end[0] = '\0';

    read_BIG(ax1,stringax1);
    read_BIG(ax2,stringax2);
    read_BIG(ay1,stringay1);
    read_BIG(ay2,stringay2);
    read_BIG(bx1,stringbx1);
    read_BIG(bx2,stringbx2);
    read_BIG(by1,stringby1);
    read_BIG(by2,stringby2);
    read_BIG(cx1,stringcx1);
    read_BIG(cx2,stringcx2);
    read_BIG(cy1,stringcy1);
    read_BIG(cy2,stringcy2);

    FP2_from_BIGs(&ax,ax1,ax2);
    FP2_from_BIGs(&ay,ay1,ay2);
    FP2_from_BIGs(&bx,bx1,bx2);
    FP2_from_BIGs(&by,by1,by2);
    FP2_from_BIGs(&cx,cx1,cx2);
    FP2_from_BIGs(&cy,cy1,cy2);

    FP4_from_FP2s(&a,&ax,&ay);
    FP4_from_FP2s(&b,&bx,&by);
    FP4_from_FP2s(&c,&cx,&cy);

    FP12_from_FP4s(fp12,&a,&b,&c);
}

int main(int argc, char** argv)
{
    if (argc != 2)
    {
        printf("usage: ./test_fp12_arithmetics [path to test vector file]\n");
        exit(EXIT_FAILURE);
    }

    int i = 0, len = 0, j = 0;
    FILE *fp;

    char line[LINE_LEN];
    char * linePtr = NULL;

    BIG M, Fr_a, Fr_b;
    FP2 Frob;
    FP12 FP12aux1, FP12aux2;

    FP12 FP12_1;
    const char* FP12_1line = "FP12_1 = ";
    FP12 FP12_2;
    const char* FP12_2line = "FP12_2 = ";
    FP12 FP12conj;
    const char* FP12conjline = "FP12conj = ";

    BIG_rcopy(M,Modulus);
    BIG_rcopy(Fr_a,CURVE_Fra);
    BIG_rcopy(Fr_b,CURVE_Frb);
    FP2_from_BIGs(&Frob,Fr_a,Fr_b);

// Set to one
    FP12_one(&FP12aux1);
    FP12_copy(&FP12aux2,&FP12aux1);

// Testing equal function, copy function and set one function
    if(!FP12_equals(&FP12aux1,&FP12aux2) || !FP12_isunity(&FP12aux1) || !FP12_isunity(&FP12aux2))
    {
        printf("ERROR comparing FP12s or setting FP12 to unity or copying FP12\n");
        exit(EXIT_FAILURE);
    }


    fp = fopen(argv[1], "r");
    if (fp == NULL)
    {
        printf("ERROR opening test vector file\n");
        exit(EXIT_FAILURE);
    }

    while (fgets(line, LINE_LEN, fp) != NULL)
    {
        i++;
// Read first FP12 and perform some tests
        if (!strncmp(line,FP12_1line, strlen(FP12_1line)))
        {
            len = strlen(FP12_1line);
            linePtr = line + len;
                printf("\n\n");
                printf("%s",linePtr);
                printf("\n\n");
            read_FP12(&FP12_1,linePtr);
                FP12_output(&FP12_1);
                printf("\n\n");
        }
// Read second FP12
        if (!strncmp(line,FP12_2line, strlen(FP12_2line)))
        {
            len = strlen(FP12_2line);
            linePtr = line + len;
            read_FP12(&FP12_2,linePtr);
                printf("\n\n");
                FP12_output(&FP12_2);
                printf("\n\n");
        }
// Test FP12_conj
        if (!strncmp(line,FP12conjline, strlen(FP12conjline)))
        {
            len = strlen(FP12conjline);
            linePtr = line + len;
            read_FP12(&FP12conj,linePtr);
            FP12_copy(&FP12aux1,&FP12_1);
            FP12_conj(&FP12aux1,&FP12aux1);
            FP12_reduce(&FP12aux1);
            FP12_norm(&FP12aux1);
            if(!FP12_equals(&FP12aux1,&FP12conj))
            {
                printf("\n\n");
                FP12_output(&FP12aux1);
                printf("\n\n");
                FP12_output(&FP12conj);
                printf("\n\n");
                printf("ERROR computing conjugate of FP12, line %d\n",i);
                exit(EXIT_FAILURE);
            }
        }        
    }
    fclose(fp);

    printf("SUCCESS TEST ARITMETIC OF FP12 PASSED\n");
    exit(EXIT_SUCCESS);
}
