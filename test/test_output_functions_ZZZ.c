/**
 * @file test_output_functions.c
 * @author Alessandro Budroni
 * @brief Test output functions for debug
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


#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "arch.h"
#include "amcl.h"
#include "utils.h"
#include "big_XXX.h"
#include "fp_YYY.h"
#include "fp2_YYY.h"
#include "fp4_YYY.h"
#include "fp12_YYY.h"
#include "ecp_ZZZ.h"
#include "ecp2_ZZZ.h"

extern int dup(int oldfd);
extern int dup2(int oldfd, int newfd);
extern int close(int fildes);
extern int fileno(FILE *stream);

#define LINE_LEN 2000

void read_BIG(BIG_XXX A, char* string)
{
    int len;
    char support[LINE_LEN];
    BIG_XXX_zero(A);
    len = strlen(string)+1;
    amcl_hex2bin(string,support,len);
    len = (len-1)/2;;
    BIG_XXX_fromBytesLen(A,support,len);
    BIG_XXX_norm(A);
}

void read_DBIG(DBIG_XXX A, char* string)
{
    int len;
    char support[LINE_LEN];
    BIG_XXX_dzero(A);
    len = strlen(string)+1;
    amcl_hex2bin(string,support,len);
    len = (len-1)/2;
    BIG_XXX_dfromBytesLen(A,support,len);
    BIG_XXX_dnorm(A);
}

// #if (CHOICE >= BN_CURVES)
void read_FP2(FP2_YYY *fp2, char* stringx)
{
    char *stringy, *end;
    BIG_XXX x,y;
    FP_YYY fx,fy;
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

    FP_YYY_nres(&fx,x);
    FP_YYY_nres(&fy,y);

    FP2_YYY_from_FPs(fp2,&fx,&fy);
}

void read_FP4(FP4_YYY *fp4, char* stringx1)
{
    char *stringx2, *stringy1, *stringy2, *end;
    BIG_XXX x1,x2,y1,y2;
    FP2_YYY x,y;

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

    FP2_YYY_from_BIGs(&x,x1,x2);
    FP2_YYY_from_BIGs(&y,y1,y2);

    FP4_YYY_from_FP2s(fp4,&x,&y);
}

// Read a structure of the type [[[ax1,ax2],[ay1,ay2]],[[bx1,bx2],[by1,by2]],[[cx1,cx2],[cy1,cy2]]]
void read_FP12(FP12_YYY *fp12, char *stringax1)
{
    char *stringax2, *stringay1, *stringay2, *stringbx1, *stringbx2, *stringby1, *stringby2, *stringcx1, *stringcx2, *stringcy1, *stringcy2, *end;
    BIG_XXX ax1,ax2,ay1,ay2,bx1,bx2,by1,by2,cx1,cx2,cy1,cy2;
    FP2_YYY ax,ay,bx,by,cx,cy;
    FP4_YYY a,b,c;

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

    FP2_YYY_from_BIGs(&ax,ax1,ax2);
    FP2_YYY_from_BIGs(&ay,ay1,ay2);
    FP2_YYY_from_BIGs(&bx,bx1,bx2);
    FP2_YYY_from_BIGs(&by,by1,by2);
    FP2_YYY_from_BIGs(&cx,cx1,cx2);
    FP2_YYY_from_BIGs(&cy,cy1,cy2);

    FP4_YYY_from_FP2s(&a,&ax,&ay);
    FP4_YYY_from_FP2s(&b,&bx,&by);
    FP4_YYY_from_FP2s(&c,&cx,&cy);

    FP12_YYY_from_FP4s(fp12,&a,&b,&c);
}
// #endif

int read_ECP(ECP_ZZZ *ecp, char* string)
{
    BIG_XXX x;
    char *end;
    BIG_XXX y;
    char *stringy;
    string++;
    stringy = strchr(string,',');
    if (stringy == NULL)
    {
        printf("ERROR unexpected test vector\n");
        exit(EXIT_FAILURE);
    }
    stringy[0] = '\0';
    stringy++;
    end = strchr(stringy,')');
    if (end == NULL)
    {
        printf("ERROR unexpected test vector\n");
        exit(EXIT_FAILURE);
    }
    end[0] = '\0';
    read_BIG(x,string);
    read_BIG(y,stringy);
    return ECP_ZZZ_set(ecp,x,y);
}

//#if (CHOICE >= BN_CURVES)
int read_ECP2(ECP2_ZZZ *ecp2, char* stringx1)
{
    char *stringx2, *stringy1, *stringy2, *end;
    BIG_XXX x1,x2,y1,y2;
    FP2_YYY x,y;

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

    FP2_YYY_from_BIGs(&x,x1,x2);
    FP2_YYY_from_BIGs(&y,y1,y2);

    return ECP2_ZZZ_set(ecp2,&x,&y);
}
//#endif

void read_OCT(octet *oct, char* string, int len)
{
    char buff[len-1];
    buff[len-1] = '\0';
    strncpy(buff,string,len-1);
    OCT_fromHex(oct,buff);
}

int main(int argc, char** argv)
{
    if (argc != 3)
    {
        printf("usage: ./test_ECP2_ZZZ_arithmetics [path to test vector file] [path to output directory]\n");
        exit(EXIT_FAILURE);
    }

    int fd, len, ch1 = 0, ch2 = 0, i=0;
    fpos_t pos;
    FILE *testVectFile, *writtenFile;

    char line[LINE_LEN];
    char * linePtr = NULL;

    BIG_XXX big, bigaux1, bigaux2;
    const char* BIGline = "BIG = ";
    const char* BIGrawline = "BIGraw = ";
    DBIG_XXX dbig;
    const char* DBIGline = "DBIG = ";
    FP_YYY fp;
    const char* FPline = "FP = ";
    const char* FPrawline = "FPraw = ";
// #if (CHOICE >= BN_CURVES)
    FP2_YYY fp2, fp2aux;
    const char* FP2line = "FP2 = ";
    const char* FP2rawline = "FP2raw = ";
    FP4_YYY fp4;
    const char* FP4line = "FP4 = ";
    const char* FP4rawline = "FP4raw = ";
    FP12_YYY fp12;
    const char* FP12line = "FP12 = ";
// #endif
    ECP_ZZZ ecp;
    ECP_ZZZ ecpinf;
    const char* ECPline = "ECP = ";
    const char* ECPxyzline = "ECPxyz = ";
    const char* ECPinfline = "ECPinf = ";
    const char* ECPinfxyzline = "ECPinfxyz = ";
// #if (CHOICE >= BN_CURVES)
    ECP2_ZZZ ecp2;
    ECP2_ZZZ ecp2inf;
    const char* ECP2line = "ECP2 = ";
    const char* ECP2xyzline = "ECP2xyz = ";
    const char* ECP2infline = "ECP2inf = ";
    const char* ECP2infxyzline = "ECP2infxyz = ";
// #endif
    char octbuf[LINE_LEN];
    octet oct = {0,sizeof(octbuf),octbuf};
    const char* OCTline = "OCT = ";
    //const char* OCTstringline = "OCTstring = ";
    char bin[32];
    const char* HEXline = "HEX = ";

    testVectFile = fopen(argv[1],"r");
    if (testVectFile == NULL)
    {
        printf("ERROR opening test vector file\n");
        exit(EXIT_FAILURE);
    }

    fgetpos(stdout, &pos);
    fd = dup(fileno(stdout));

    if(freopen(argv[2], "w", stdout) == NULL)
    {
        printf("ERROR redirecting stdout\n");
        exit(EXIT_FAILURE);
    }

    while (fgets(line, LINE_LEN, testVectFile) != NULL)
    {
        i++;
        if (!strncmp(line,  BIGline, strlen(BIGline)))
        {
            len = strlen(BIGline);
            linePtr = line + len;
            read_BIG(big, linePtr);
            printf("%s",BIGline);
            BIG_XXX_output(big);
            printf("\n");
            printf("%s",BIGrawline);
            BIG_XXX_rawoutput(big);
            printf("\n\n");
        }
        if (!strncmp(line,  DBIGline, strlen(DBIGline)))
        {
            len = strlen(DBIGline);
            linePtr = line + len;
            read_DBIG(dbig, linePtr);
            printf("%s",DBIGline);
            BIG_XXX_doutput(dbig);
            printf("\n\n");
        }
        if (!strncmp(line,  FPline, strlen(FPline)))
        {
            len = strlen(FPline);
            linePtr = line + len;
            read_BIG(big,linePtr);
            FP_YYY_nres(&fp,big);
            printf("%s",FPline);
            FP_YYY_output(&fp);
            printf("\n");
            printf("%s",FPrawline);
            FP_YYY_rawoutput(&fp);
            printf("\n\n");
        }
// #if (CHOICE >= BN_CURVES)
        if (!strncmp(line,  FP2line, strlen(FP2line)))
        {
            len = strlen(FP2line);
            linePtr = line + len;
            read_FP2(&fp2,linePtr);
            FP_YYY_reduce(&fp2.a);
            FP_YYY_reduce(&fp2.b);
            printf("%s",FP2line);
            FP2_YYY_output(&fp2);
            printf("\n");
            printf("%s",FP2rawline);
            FP2_YYY_rawoutput(&fp2);
            printf("\n\n");
        }
        if (!strncmp(line,  FP4line, strlen(FP4line)))
        {
            len = strlen(FP4line);
            linePtr = line + len;
            read_FP4(&fp4,linePtr);
            printf("%s",FP4line);
            FP4_YYY_output(&fp4);
            printf("\n");
            printf("%s",FP4rawline);
            FP4_YYY_rawoutput(&fp4);
            printf("\n\n");
        }
        if (!strncmp(line,  FP12line, strlen(FP12line)))
        {
            len = strlen(FP12line);
            linePtr = line + len;
            read_FP12(&fp12,linePtr);
            printf("%s",FP12line);
            FP12_YYY_output(&fp12);
            printf("\n\n");
        }
// #endif
        if (!strncmp(line,  ECPline, strlen(ECPline)))
        {
            len = strlen(ECPline);
            linePtr = line + len;
            read_ECP(&ecp,linePtr);
            printf("%s",ECPline);
            ECP_ZZZ_output(&ecp);
            printf("%s",ECPxyzline);
            ECP_ZZZ_outputxyz(&ecp);
            FP_YYY_redc(bigaux1,&ecp.x);
            FP_YYY_redc(bigaux2,&ecp.y);
            BIG_XXX_inc(bigaux1,100);
            ECP_ZZZ_set(&ecpinf,bigaux1,bigaux2);
            printf("%s",ECPinfline);
            ECP_ZZZ_output(&ecpinf);
            printf("%s",ECPinfxyzline);
            ECP_ZZZ_outputxyz(&ecpinf);
            printf("\n");
        }
// #if (CHOICE >= BN_CURVES)
        if (!strncmp(line,  ECP2line, strlen(ECP2line)))
        {
            len = strlen(ECP2line);
            linePtr = line + len;
            read_ECP2(&ecp2,linePtr);
            printf("%s",ECP2line);
            ECP2_ZZZ_output(&ecp2);
            printf("%s",ECP2xyzline);
            ECP2_ZZZ_outputxyz(&ecp2);
            FP2_YYY_add(&fp2aux,&ecp2.x,&ecp2.y);
            ECP2_ZZZ_set(&ecp2inf,&fp2aux,&fp2aux);
            printf("%s",ECP2infline);
            ECP2_ZZZ_output(&ecp2inf);
            printf("%s",ECP2infxyzline);
            ECP2_ZZZ_outputxyz(&ecp2inf);
            printf("\n");
        }
// #endif
        if (!strncmp(line,  OCTline, strlen(OCTline)))
        {
            len = strlen(OCTline);
            linePtr = line + len;
            read_OCT(&oct,linePtr,strlen(linePtr));
            printf("%s",OCTline);
            OCT_output(&oct);
            //printf("%s",OCTstringline);
            //OCT_output_string(&oct);
            //printf("\n");
        }
        if (!strncmp(line,  HEXline, strlen(HEXline)))
        {
            len = strlen(OCTline);
            linePtr = line + len;
            amcl_hex2bin(linePtr, bin, 64);
            printf("\n%s", HEXline);
            amcl_print_hex(bin,32);
        }
    }

    // Restore stdout
    fflush(stdout);
    dup2(fd, fileno(stdout));
    close(fd);
    clearerr(stdout);
    fsetpos(stdout, &pos);        /* for C9X */

    writtenFile = fopen(argv[2],"r");
    if (writtenFile == NULL)
    {
        printf("ERROR opening output file\n");
        exit(EXIT_FAILURE);
    }
    fclose(testVectFile);

    // Check the equality of the output with the test vector file

    testVectFile = fopen(argv[1],"r");
    if (testVectFile == NULL)
    {
        printf("ERROR opening test vector file\n");
        exit(EXIT_FAILURE);
    }

    ch1 = getc(testVectFile);
    ch2 = getc(writtenFile);

    while ((ch1 != EOF) && (ch2 != EOF) && (ch1 == ch2))
    {
        ch1 = getc(testVectFile);
        ch2 = getc(writtenFile);
    }

    if (ch1 != ch2)
    {
        printf("ERROR output does not match the expected one \n");
        exit(EXIT_FAILURE);
    }
    fclose(writtenFile);
    fclose(testVectFile);

    remove(argv[2]);

    printf("SUCCESS TEST OUTPUT FUNCTIONS PASSED\n");
    exit(EXIT_SUCCESS);
}
