#!/usr/bin/env bash
#
# buildMulti.sh
#
# Build test with multiple curves and RSA security levels

# @author Kealan McCusker <kealan.mccusker@mirack.com>

set -e

# Build default - see config.mk
make clean
make

# Build example with multiple curves and RSA security level
gcc -O2 -std=c99 ./examples/testall.c -I./target/default/include/ -L./target/default/lib/ -lamcl_core -lamcl_curve_BN254 -lamcl_curve_ED25519 -lamcl_curve_GOLDILOCKS -lamcl_curve_NIST256 -lamcl_mpin_BN254 -lamcl_pairing_BN254 -lamcl_rsa_2048 -lamcl_rsa_3072 -lamcl_wcc_BN254 -lamcl_x509 -o testall

# Run code
export LD_LIBRARY_PATH=./target/default/lib
./testall


