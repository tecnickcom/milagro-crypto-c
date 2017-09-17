# MAKEFILE
#
# @author      Nicola Asuni <nicola.asuni@miracl.com>
# @link        https://github.com/miracl/milagro-crypto-c
#
# This file is intended to be executed in a Linux-compatible system and requires
# the packages listed in resources/DockerDev/Dockerfile to execute the build in
# the current environment, or Docker to build everything inside a Docker 
# container via the command "MAKETARGET=buildall make dbuild".
#
# Requires GNU parallel: https://www.gnu.org/software/parallel/
# ------------------------------------------------------------------------------

# List special make targets that are not associated with files
.PHONY: help all default format clean qa z build build_qa_item build_item buildx buildall dbuild pubdocs print-%

# Use bash as shell (Note: Ubuntu now uses dash which doesn't support PIPESTATUS).
SHELL=/bin/bash

# Project root directory
PROJECTROOT=$(shell pwd)

# CVS path (path to the parent dir containing the project)
CVSPATH=github.com/miracl

# Project owner
OWNER=MIRACL

# Project vendor
VENDOR=miracl

# Project name
PROJECT=amcl

# Project version
VERSION=$(shell cat VERSION)

# Project release number (packaging build number)
RELEASE=$(shell cat RELEASE)

# Include default build configuration
include $(PROJECTROOT)/config.mk

# Common CMake options for building the language wrappers
WRAPPYTHON="-DBUILD_PYTHON=on"
WRAPGOLANG="-DBUILD_GO=on"

# Space-separated list of build options (grouped by type):
# <NAME>:<DOUBLECOMMA-SEPARATED_LIST_OF_CMAKE_OPTIONS>

BUILDS_PF64=LINUX_64BIT_BLS383:-DAMCL_CURVE=BLS383,,-DAMCL_RSA=2048,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,${WRAPGOLANG},,${WRAPPYTHON} \
	LINUX_64BIT_BN254CX:-DAMCL_CURVE=BN254CX,,-DAMCL_RSA=2048,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,${WRAPGOLANG},,${WRAPPYTHON} \
	LINUX_64BIT_BN254:-DAMCL_CURVE=BN254,,-DAMCL_RSA=2048,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,${WRAPGOLANG},,${WRAPPYTHON} \
	LINUX_64BIT_ALLPF:-DAMCL_CURVE=BN254,BN254CX,BLS383,,-DAMCL_RSA=2048,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,${WRAPGOLANG},,${WRAPPYTHON} \

BUILDS_NIST64=LINUX_64BIT_NIST256_RSA2048:-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NIST256,,-DAMCL_RSA=2048,,${WRAPGOLANG} \
	LINUX_64BIT_NIST256_RSA4096:-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NIST256,,-DAMCL_RSA=4096,,${WRAPGOLANG} \
	LINUX_64BIT_NIST384_RSA3072:-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NIST384,,-DAMCL_RSA=3072,,${WRAPGOLANG} \
	LINUX_64BIT_NIST521_RSA4096:-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NIST521,,-DAMCL_RSA=4096,,${WRAPGOLANG} \
	LINUX_64BIT_NIST_RSA_ALL:-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NIST256,NIST384,NIST521,,-DAMCL_RSA=2048,3072,4096

BUILDS_MISC64=LINUX_64BIT_C25519_RSA2048:-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=C25519,,-DAMCL_RSA=2048 \
	LINUX_64BIT_BRAINPOOL:-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=BRAINPOOL \
	LINUX_64BIT_ANSSI:-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=ANSSI \
	LINUX_64BIT_ED25519:-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=ED25519 \
	LINUX_64BIT_NUMS256E:-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NUMS256E \
	LINUX_64BIT_NUMS256W:-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NUMS256W \
	LINUX_64BIT_NUMS384E:-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NUMS384E \
	LINUX_64BIT_NUMS384W:-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NUMS384W \
	LINUX_64BIT_NUMS512E:-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NUMS512E \
	LINUX_64BIT_NUMS512W:-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NUMS512W \
	LINUX_64BIT_HIFIVE:-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=HIFIVE \
	LINUX_64BIT_GOLDILOCKS:-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=GOLDILOCKS \
	LINUX_64BIT_C41417:-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=C41417 \
	LINUX_64BIT_C25519_BN254CX_RSA2048:-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=C25519,BN254CX,,-DAMCL_RSA=2048 \
	LINUX_64BIT_NIST256_BN254CX_RSA2048:-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NIST256,BN254CX,,-DAMCL_RSA=2048 \
	WINDOWS_64BIT_BN254_CX:-DCMAKE_TOOLCHAIN_FILE=../../resources/cmake/mingw64-cross.cmake

BUILDS_PF32=LINUX_32BIT_BLS383:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DAMCL_CURVE=BLS383,,-DAMCL_RSA=2048,,-DCMAKE_INSTALL_PREFIX=/opt/amcl \
	LINUX_32BIT_BN254CX:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DAMCL_CURVE=BN254CX,,-DAMCL_RSA=2048,,-DCMAKE_INSTALL_PREFIX=/opt/amcl \
	LINUX_32BIT_BN254:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DAMCL_CURVE=BN254,,-DAMCL_RSA=2048,,-DCMAKE_INSTALL_PREFIX=/opt/amcl \
	LINUX_32BIT_ALLPF:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DAMCL_CURVE=BN254,BN254CX,BLS383,,-DAMCL_RSA=2048,,-DCMAKE_INSTALL_PREFIX=/opt/amcl \

BUILDS_NIST32=LINUX_32BIT_NIST256_RSA2048:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NIST256,,-DAMCL_RSA=2048 \
	LINUX_32BIT_NIST256_RSA4096:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NIST256,,-DAMCL_RSA=4096 \
	LINUX_32BIT_NIST384_RSA3072:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NIST384,,-DAMCL_RSA=3072 \
	LINUX_32BIT_NIST521_RSA4096:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NIST521,,-DAMCL_RSA=4096 \
	LINUX_32BIT_NIST_RSA_ALL:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NIST256,NIST384,NIST521,,-DAMCL_RSA=2048,3072,4096

BUILDS_MISC32=LINUX_32BIT_C25519_RSA2048:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=C25519,,-DAMCL_RSA=2048 \
	LINUX_32BIT_BRAINPOOL:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=BRAINPOOL \
	LINUX_32BIT_ANSSI:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=ANSSI \
	LINUX_32BIT_ED25519:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=ED25519 \
	LINUX_32BIT_NUMS256E:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NUMS256E \
	LINUX_32BIT_NUMS256W:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NUMS256W \
	LINUX_32BIT_NUMS384E:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NUMS384E \
	LINUX_32BIT_NUMS384W:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NUMS384W \
	LINUX_32BIT_NUMS512E:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NUMS512E \
	LINUX_32BIT_NUMS512W:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NUMS512W \
	LINUX_32BIT_HIFIVE:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=HIFIVE \
	LINUX_32BIT_GOLDILOCKS:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=GOLDILOCKS \
	LINUX_32BIT_C41417:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=C41417 \
	LINUX_32BIT_C25519_BN254CX_RSA2048:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=C25519,BN254CX,,-DAMCL_RSA=2048 \
	LINUX_32BIT_NIST256_BN254CX_RSA2048:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NIST256,BN254CX,,-DAMCL_RSA=2048 \
	WINDOWS_32BIT_BN254_CX:-DCMAKE_C_FLAGS=-m32,,-DAMCL_CHUNK=32,,-DCMAKE_TOOLCHAIN_FILE=../../resources/cmake/mingw32-cross.cmake

BUILDS_PFS16=LINUX_16BIT_BN254CX:-DAMCL_CHUNK=16,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=BN254CX \
	LINUX_16BIT_BN254:-DAMCL_CHUNK=16,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=BN254

BUILDS_MISC16=LINUX_16BIT_ED25519:-DAMCL_CHUNK=16,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=ED25519 \
	LINUX_16BIT_NUMS256E:-DAMCL_CHUNK=16,,-DCMAKE_INSTALL_PREFIX=/opt/amcl,,-DAMCL_CURVE=NUMS256E

BUILDS_ASAN=LINUX_64BIT_ASan:-DCMAKE_BUILD_TYPE=ASan,,-DAMCL_CURVE=NIST256,C25519,ED25519,BRAINPOOL,ANSSI,NUMS256E,NUMS256W,NUMS384E,NUMS384W,NUMS512E,NUMS512W,HIFIVE,GOLDILOCKS,NIST384,C41417,NIST521,BN254,BN254CX,BLS383,,-DAMCL_RSA=2048,3072,4096 \

BUILDS_COVERAGE=LINUX_64BIT_COVERAGE:-DCMAKE_BUILD_TYPE=Coverage,,-DAMCL_CURVE=NIST256,BN254CX,,-DAMCL_RSA=2048

# Merge all build types in a single list
BUILDS_64=$(BUILDS_PF64) $(BUILDS_NIST64) $(BUILDS_MISC64)
BUILDS_32=$(BUILDS_PF32) $(BUILDS_NIST32) $(BUILDS_MISC32)
BUILDS_16=$(BUILDS_BN16) $(BUILDS_MISC16)

BUILDS=$(BUILDS_64) $(BUILDS_32) $(BUILDS_16) $(BUILDS_ASAN) $(BUILDS_COVERAGE)

# Variables used in text substitution
dcomma := ,,
space :=
space +=

# --- MAKE TARGETS ---

# Default build configured in config.mk
all: default

# Display general help about this command
help:
	@echo ""
	@echo "$(PROJECT) Makefile."
	@echo "The following commands are available:"
	@echo ""
	@echo "    make         :  Build library based on options in config.mk"
	@echo "    make dbuild  :  Build library using docker based on options in config.mk"
	@echo "    make format  :  Format the source code"
	@echo "    make clean   :  Remove any build artefact"
	@echo ""
	@echo "    Testing:"
	@echo ""
	@echo "    make qa      :  Build all versions in this makefile and generate reports"
	@echo "    MAKETARGET=buildall make dbuild:  Build all versions in this makefile using docker and generate reports"
	@echo ""
	@echo "    You can also build individual types, groups or sub-groups:"
	@echo ""
	@echo ""
	@$(foreach PARAMS,$(BUILDS_ASAN), \
		echo "    make build TYPE=$(word 1,$(subst :, ,${PARAMS}))" ; \
	)
	@echo ""
	@$(foreach PARAMS,$(BUILDS_COVERAGE), \
		echo "    make build TYPE=$(word 1,$(subst :, ,${PARAMS}))" ; \
	)
	@echo ""
	@$(foreach PARAMS,$(BUILDS_64), \
		echo "    make build TYPE=$(word 1,$(subst :, ,${PARAMS}))" ; \
	)
	@echo ""
	@$(foreach PARAMS,$(BUILDS_32), \
		echo "    make build TYPE=$(word 1,$(subst :, ,${PARAMS}))" ; \
	)
	@echo ""
	@$(foreach PARAMS,$(BUILDS_16), \
		echo "    make build TYPE=$(word 1,$(subst :, ,${PARAMS}))" ; \
	)
	@echo ""

# Default build configured in config.mk
default: 
	@echo -e "\n\n*** BUILD default - see config.mk ***\n"
	rm -rf target/default/*
ifeq ($(CMAKE_BUILD_TYPE),Coverage)
	mkdir -p target/default/coverage
	cd target/default && \
	cmake -DCMAKE_C_FLAGS=$(CMAKE_C_FLAGS) \
	-DCMAKE_TOOLCHAIN_FILE=$(CMAKE_TOOLCHAIN_FILE) \
	-DCMAKE_BUILD_TYPE=$(CMAKE_BUILD_TYPE) \
	-DCMAKE_INSTALL_PREFIX=$(CMAKE_INSTALL_PATH) \
	-DBUILD_SHARED_LIBS=$(AMCL_BUILD_SHARED_LIBS) \
	-DBUILD_PYTHON=$(AMCL_BUILD_PYTHON) \
	-DBUILD_GO=$(AMCL_BUILD_GO) \
	-DAMCL_CHUNK=$(AMCL_CHUNK) \
	-DAMCL_CURVE=$(AMCL_CURVE) \
	-DAMCL_RSA=$(AMCL_RSA) \
	-DBUILD_MPIN=$(AMCL_BUILD_MPIN) \
	-DBUILD_WCC=$(AMCL_BUILD_WCC) \
	-DBUILD_DOXYGEN=$(AMCL_BUILD_DOXYGEN) \
	-DAMCL_MAXPIN=$(AMCL_MAXPIN) \
	-DAMCL_PBLEN=$(AMCL_PBLEN) \
	-DDEBUG_REDUCE=$(DEBUG_REDUCE) \
	-DDEBUG_NORM=$(DEBUG_NORM) \
	../.. | tee cmake.log ; test $${PIPESTATUS[0]} -eq 0 && \
	export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:./ && \
	make | tee make.log ; test $${PIPESTATUS[0]} -eq 0 && \
	lcov --zerocounters --directory . && \
	lcov --capture --initial --directory . --output-file coverage/amcl && \
	env CTEST_OUTPUT_ON_FAILURE=1 make test | tee test.log ; test $${PIPESTATUS[0]} -eq 0 && \
	lcov --no-checksum --directory . --capture --output-file coverage/amcl.info && \
	lcov --remove coverage/amcl.info "/test_*" --output-file coverage/amcl.info && \
	genhtml -o coverage -t "milagro-crypto-c Test Coverage" coverage/amcl.info 
else
	mkdir -p target/default
	cd target/default && \
	cmake -DCMAKE_C_FLAGS=$(CMAKE_C_FLAGS) \
	-DCMAKE_TOOLCHAIN_FILE=$(CMAKE_TOOLCHAIN_FILE) \
	-DCMAKE_BUILD_TYPE=$(CMAKE_BUILD_TYPE) \
	-DCMAKE_INSTALL_PREFIX=$(CMAKE_INSTALL_PATH) \
	-DBUILD_SHARED_LIBS=$(AMCL_BUILD_SHARED_LIBS) \
	-DBUILD_PYTHON=$(AMCL_BUILD_PYTHON) \
	-DBUILD_GO=$(AMCL_BUILD_GO) \
	-DAMCL_CHUNK=$(AMCL_CHUNK) \
	-DAMCL_CURVE=$(AMCL_CURVE) \
	-DAMCL_RSA=$(AMCL_RSA) \
	-DBUILD_MPIN=$(AMCL_BUILD_MPIN) \
	-DBUILD_WCC=$(AMCL_BUILD_WCC) \
	-DBUILD_DOXYGEN=$(AMCL_BUILD_DOXYGEN) \
	-DAMCL_MAXPIN=$(AMCL_MAXPIN) \
	-DAMCL_PBLEN=$(AMCL_PBLEN) \
	-DDEBUG_REDUCE=$(DEBUG_REDUCE) \
	-DDEBUG_NORM=$(DEBUG_NORM) \
	../.. | tee cmake.log ; test $${PIPESTATUS[0]} -eq 0 && \
	make | tee make.log ; test $${PIPESTATUS[0]} -eq 0 
ifeq ($(AMCL_TEST),ON)
	cd target/default && \
	export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:./ && \
	env CTEST_OUTPUT_ON_FAILURE=1 make test | tee test.log ; test $${PIPESTATUS[0]} -eq 0
endif
ifeq ($(AMCL_BUILD_DOXYGEN),ON)
	cd target/default && \
	make doc | tee doc.log ; test $${PIPESTATUS[0]} -eq 0 
endif
endif


# Format the source code
format:
	astyle --style=allman --recursive --suffix=none 'include/*.h'
	astyle --style=allman --recursive --suffix=none 'src/*.c'
	astyle --style=allman --recursive --suffix=none 'test/*.c'
	astyle --style=allman --recursive --suffix=none 'examples/*.c'
	astyle --style=allman --recursive --suffix=none 'benchmark/*.c'
	find ./wrappers/go -type f -name "*.go" -exec gofmt -s -w {} \;
	autopep8 --in-place --aggressive --aggressive ./wrappers/python/*.py

# Remove any build artifact
clean:
	mkdir -p target/
	rm -rf ./target/*

# Execute all builds and tests
qa:
	go get github.com/stretchr/testify/assert
	@mkdir -p target/
	@echo 0 > target/make.exit
	@echo '' > target/make_qa_errors.log
	make build_group BUILD_GROUP=BUILDS
	@cat target/make_qa_errors.log
	@exit `cat target/make.exit`

# Build the specified group of options
build_group:
	@parallel --no-notice --verbose make build_qa_item ITEM={} ::: ${${BUILD_GROUP}}

# Build the project using one of the pre-defined targets (example: "make build TYPE=LINUX_64BIT_BN254CX")
build:
	make build_item ITEM=$(filter ${TYPE}:%,$(BUILDS))

# Same as build_item but stores the exit code and faling items
build_qa_item:
	make build_item ITEM=${ITEM} || (echo $$? > target/make.exit && echo ${ITEM} >> target/make_qa_errors.log);
 
# Build the specified item entry from the BUILDS list
build_item:
	make buildx BUILD_NAME=$(word 1,$(subst :, ,${ITEM})) BUILD_PARAMS=$(word 2,$(subst :, ,${ITEM}))

# Build with the specified parameters
buildx:
	@echo -e "\n\n*** BUILD ${BUILD_NAME} ***\n"
	rm -rf target/${BUILD_NAME}/*
ifneq ($(strip $(filter %COVERAGE,${BUILD_NAME})),)
	mkdir -p target/${BUILD_NAME}/coverage
	cd target/${BUILD_NAME} && \
	cmake $(subst $(dcomma),$(space),${BUILD_PARAMS}) ../.. | tee cmake.log ; test $${PIPESTATUS[0]} -eq 0 && \
	export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:./ && \
	make | tee make.log ; test $${PIPESTATUS[0]} -eq 0 &&\
	lcov --zerocounters --directory . && \
	lcov --capture --initial --directory . --output-file coverage/amcl && \
	env CTEST_OUTPUT_ON_FAILURE=1 make test | tee test.log ; test $${PIPESTATUS[0]} -eq 0 && \
	lcov --no-checksum --directory . --capture --output-file coverage/amcl.info && \
	lcov --remove coverage/amcl.info "/test_*" --output-file coverage/amcl.info && \
	genhtml -o coverage -t "milagro-crypto-c Test Coverage" coverage/amcl.info 
else 
	mkdir -p target/${BUILD_NAME}
	cd target/${BUILD_NAME} && \
	cmake $(subst $(dcomma),$(space),${BUILD_PARAMS}) ../.. | tee cmake.log ; test $${PIPESTATUS[0]} -eq 0 && \
	export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:./ && \
	make | tee make.log ; test $${PIPESTATUS[0]} -eq 0 && \
	make doc | tee doc.log ; test $${PIPESTATUS[0]} -eq 0 && \
	env CTEST_OUTPUT_ON_FAILURE=1 make test | tee test.log ; test $${PIPESTATUS[0]} -eq 0
endif

# Alias for building all inside the Docker container
buildall: default qa

# Build everything inside a Docker container
dbuild:
	@mkdir -p target
	@rm -rf target/*
	@echo 0 > target/make.exit
	CVSPATH=$(CVSPATH) VENDOR=$(VENDOR) PROJECT=$(PROJECT) MAKETARGET='$(MAKETARGET)' ./dockerbuild.sh
	@exit `cat target/make.exit`

# Publish Documentation in GitHub (requires writing permissions)
# Use this only after generating with default build i.e. "make"
pubdocs:
	rm -rf ./target/DOCS
	rm -rf ./target/WIKI
	mkdir -p ./target/DOCS/doc
	cp -r ./target/default/doc/html/* ./target/DOCS/doc
	cp ./doc/Home.md ./target/DOCS/
	git clone https://github.com/miracl/milagro-crypto-c.wiki.git ./target/WIKI
	mv -f ./target/WIKI/.git ./target/DOCS/
	cd ./target/DOCS/ && \
	git add . -A && \
	git commit -m 'Update documentation' && \
	git push origin master --force 
	rm -rf ./target/DOCS
	rm -rf ./target/WIKI

# Print variables usage: make print-VARABLE
print-%: ; @echo $* = $($*)
