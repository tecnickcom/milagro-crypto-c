cmake_minimum_required (VERSION 3.1)
project (AMCL)

add_definitions(-DCMAKE)

# enable testing
enable_testing ()

# Extract version from the VERSION file
file(STRINGS VERSION VERSION_FILE_CONTENT)
string(REPLACE "." ";" VERSION_FILE_PARTS ${VERSION_FILE_CONTENT})
list(GET VERSION_FILE_PARTS 0 AMCL_VERSION_MAJOR)
list(GET VERSION_FILE_PARTS 1 AMCL_VERSION_MINOR)
list(GET VERSION_FILE_PARTS 2 AMCL_VERSION_PATCH)
set(PROJECT_VERSION "${AMCL_VERSION_MAJOR}.${AMCL_VERSION_MINOR}.${AMCL_VERSION_PATCH}")
message(STATUS "PROJECT_VERSION='${PROJECT_VERSION}'")
message(STATUS "Generator - ${CMAKE_GENERATOR}")

# Add options for build
option (BUILD_SHARED_LIBS "Build Shared Libraries" ON)
option (BUILD_PYTHON "Build Python" OFF)
option (BUILD_GO "Build Golang" OFF)
option (BUILD_MPIN "Build MPIN" ON)
option (BUILD_WCC "Build WCC" ON)
option (BUILD_DOXYGEN "Build Doxygen" ON)
option (USE_PATENTS "Use patents for G1 and G2 multiplication" OFF)
option (DEBUG_REDUCE "Print debug message for field reduction" OFF)
option (DEBUG_NORM "Detect digit overflow" OFF)
option (GET_STATS "Debug statistics" OFF)

# Allow the developer to select if Dynamic or Static libraries are built
# Set the default LIB_TYPE variable to STATIC
SET (LIB_TYPE STATIC)
IF (BUILD_SHARED_LIBS)
  # User wants to build Dynamic Libraries, so change the LIB_TYPE variable to CMake keyword 'SHARED'
  SET (LIB_TYPE SHARED)
ENDIF (BUILD_SHARED_LIBS)

set(CMAKE_BUILD_TYPE ${CMAKE_BUILD_TYPE}
    CACHE STRING "Choose the type of build: Debug Release Coverage ASan Check CheckFull"
    FORCE)

# Set a default build type if none was specified
if(NOT CMAKE_BUILD_TYPE)
  message(STATUS "Setting build type to 'Release' as none was specified.")
  set(CMAKE_BUILD_TYPE Release)
endif(NOT CMAKE_BUILD_TYPE)

if(CMAKE_COMPILER_IS_GNUCC)
    message(STATUS "Build type: ${CMAKE_BUILD_TYPE}")
    execute_process(COMMAND ${CMAKE_C_COMPILER} -dumpversion
                    OUTPUT_VARIABLE GCC_VERSION)
    set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -std=c99 -Wall -Wextra -Wno-strict-prototypes -Wunused-value -Wcast-align -Wunused-variable -Wundef -Wformat-security")

    if (GCC_VERSION VERSION_GREATER 4.8 OR GCC_VERSION VERSION_EQUAL 4.8)
        set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -Wshadow")
        set(CMAKE_C_FLAGS_ASAN    "-O0 -g3 -fsanitize=address")
    else (GCC_VERSION VERSION_GREATER 4.8 OR GCC_VERSION VERSION_EQUAL 4.8)
        message(STATUS "GCC 4.8 required to run address sanitizer - please upgrade your installation")
    endif(GCC_VERSION VERSION_GREATER 4.8 OR GCC_VERSION VERSION_EQUAL 4.8)

    set(CMAKE_C_FLAGS_RELEASE     "-O2")
    set(CMAKE_C_FLAGS_DEBUG       "-O0 -g3")
    set(CMAKE_C_FLAGS_COVERAGE    "-O0 -g3 --coverage")
    set(CMAKE_C_FLAGS_CHECK       "-O2 -Werror")
    set(CMAKE_C_FLAGS_CHECKFULL   "${CMAKE_C_FLAGS_CHECK} -Wcast-qual")
endif(CMAKE_COMPILER_IS_GNUCC)

if(CMAKE_BUILD_TYPE STREQUAL "Coverage")
    set(CMAKE_SHARED_LINKER_FLAGS "--coverage")
endif(CMAKE_BUILD_TYPE STREQUAL "Coverage")

if(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
   # Mac OS X specific code
   set(OS "Mac OS X")
   # MACOSX_RPATH
   cmake_policy(SET CMP0042 NEW)
endif(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")

if(${CMAKE_SYSTEM_NAME} MATCHES "Linux")
   # Linux specific code
   # Default install path on Linux
   set(OS "Linux")
endif(${CMAKE_SYSTEM_NAME} MATCHES "Linux")

if(${CMAKE_SYSTEM_NAME} MATCHES "Windows")
   # Windows specific code
   set(OS "Windows")
endif(${CMAKE_SYSTEM_NAME} MATCHES "Windows")

# Select word length
set(AMCL_CHUNK "64" CACHE STRING "Word length. See ./include/arch.h for explanation")
set(Word_Length_Values "16;32;64")
set_property(CACHE AMCL_CHUNK PROPERTY STRINGS ${Word_Length_Values})
message(STATUS "AMCL_CHUNK='${AMCL_CHUNK}'")

# Select Field
set(AMCL_CURVE "BN254CX" CACHE STRING "Choice of Field. See ./include/amcl.h for explanation")
set(Curves_values "NIST256;C25519;ED25519;BRAINPOOL;ANSSI;NUMS256E;NUMS256W;NUMS384E;NUMS384W;NUMS512E;NUMS512W;HIFIVE;GOLDILOCKS;NIST384;C41417;NIST521;BN254;BN254CX;BLS383")
set_property(CACHE AMCL_CURVE PROPERTY STRINGS ${Curves_values})
message(STATUS "AMCL_CURVE='${AMCL_CURVE}'")

# Select RSA security level
set(AMCL_RSA "2048" CACHE STRING "RSA level of security supported. See ./include/rsa_WWW.h for explanation")
set(rsa_values "2048;3072;4096")
set_property(CACHE AMCL_RSA PROPERTY STRINGS ${rsa_values})
message(STATUS "AMCL_RSA='${AMCL_RSA}'")

# Configure PIN
set(AMCL_MAXPIN 10000 CACHE STRING "Maximum PIN value. See ./include/mpin_ZZZ.h for explanation")
set(AMCL_PBLEN 14 CACHE STRING "Maximum PIN value in bits. See ./include/mpin_ZZZ.h for explanation")
message(STATUS "AMCL_MAXPIN=${AMCL_MAXPIN}")

if(AMCL_CURVE MATCHES "NIST256" OR AMCL_CURVE MATCHES "C25519" OR AMCL_CURVE MATCHES "NIST384" OR AMCL_CURVE MATCHES "NIST521")
  message(STATUS "Build x.509")
  set(BUILD_X509 ON)
endif(AMCL_CURVE MATCHES "NIST256" OR AMCL_CURVE MATCHES "C25519" OR AMCL_CURVE MATCHES "NIST384" OR AMCL_CURVE MATCHES "NIST521")

if(NOT(AMCL_CURVE MATCHES "BN" OR AMCL_CURVE MATCHES "BLS"))
  message(STATUS "Curve choice prevents MPIN being built")
  set(BUILD_MPIN OFF)
endif(NOT(AMCL_CURVE MATCHES "BN" OR AMCL_CURVE MATCHES "BLS"))

if(NOT(AMCL_CURVE MATCHES "BN" OR AMCL_CURVE MATCHES "BLS"))
  message(STATUS "Curve choice prevents WCC being built")
  set(BUILD_WCC OFF)
endif(NOT(AMCL_CURVE MATCHES "BN" OR AMCL_CURVE MATCHES "BLS"))

# test configuration
if(BUILD_MPIN OR BUILD_WCC)
  set(PBC_TIME_PERMIT_TESTS 10 CACHE STRING "Number of days in the future to test M-PIN time permits")
  set(PBC_RANDOM_TESTS 10 CACHE STRING "Number of random M-PIN tests")
  message(STATUS "PBC_TIME_PERMIT_TESTS=${PBC_TIME_PERMIT_TESTS}")
  message(STATUS "PBC_RANDOM_TESTS=${PBC_RANDOM_TESTS}")
endif(BUILD_MPIN OR BUILD_WCC)

# Add subdirectories
add_subdirectory(include)
add_subdirectory(src)
add_subdirectory(examples)
add_subdirectory(test)
add_subdirectory(benchmark)

# Only build Python for shared libs
if (NOT BUILD_SHARED_LIBS)
  set(BUILD_PYTHON OFF)
endif (NOT BUILD_SHARED_LIBS)

if (BUILD_PYTHON)
  message(STATUS "Build Python CFFI")
  execute_process (COMMAND python -c "from distutils.sysconfig import get_python_lib; print get_python_lib()" OUTPUT_VARIABLE PYTHON_SITE_PACKAGES OUTPUT_STRIP_TRAILING_WHITESPACE)
  execute_process (COMMAND python -c "from distutils.sysconfig import get_python_lib; from os.path import dirname; print dirname(get_python_lib())" OUTPUT_VARIABLE PYTHON_SITE_LIB OUTPUT_STRIP_TRAILING_WHITESPACE)

  add_subdirectory (wrappers/python)
endif (BUILD_PYTHON)

if (BUILD_GO)
  set(GO_PATH "$ENV{GOPATH}" CACHE STRING "GOPATH")
  message(STATUS "Build Golang: GOPATH='${GO_PATH}'")
  add_subdirectory (wrappers/go)
  add_subdirectory (wrappers/go/examples)
endif (BUILD_GO)

# Build Documentation
if (BUILD_DOXYGEN)
  find_package(Doxygen QUIET)
  if(DOXYGEN_FOUND)
    add_subdirectory (doc)
  else(DOXYGEN_FOUND)
    message(STATUS "Doxygen not found. Documentation will not be built.")
  endif(DOXYGEN_FOUND)
endif (BUILD_DOXYGEN)

message(STATUS "Install ${CMAKE_INSTALL_PREFIX}")

# uninstall target
configure_file(
    "${CMAKE_CURRENT_SOURCE_DIR}/cmake_uninstall.cmake.in"
    "${CMAKE_CURRENT_BINARY_DIR}/cmake_uninstall.cmake"
    IMMEDIATE @ONLY)

add_custom_target(uninstall
    COMMAND ${CMAKE_COMMAND} -P ${CMAKE_CURRENT_BINARY_DIR}/cmake_uninstall.cmake)

# Build a CPack driven installer package
include (InstallRequiredSystemLibraries)
set(CPACK_RESOURCE_FILE_LICENSE "${CMAKE_CURRENT_SOURCE_DIR}/LICENSE")
set(CPACK_PACKAGE_VERSION "${PROJECT_VERSION}")
set(CPACK_PACKAGE_NAME "AMCL")
set(CPACK_PACKAGE_RELEASE 1)
set(CPACK_PACKAGE_CONTACT "support@miracl.com")
set(CPACK_PACKAGE_VENDOR "MIRACL")
if (BUILD_PYTHON)
  set(CPACK_RPM_PACKAGE_REQUIRES "python >= 2.7.0")
endif (BUILD_PYTHON)
set(CPACK_PACKAGE_FILE_NAME "${CPACK_PACKAGE_NAME}-${CPACK_PACKAGE_VERSION}-${CPACK_PACKAGE_RELEASE}.${CMAKE_SYSTEM_PROCESSOR}")

if(${CMAKE_SYSTEM_NAME} MATCHES "Linux")
  # Linux specific code
  set(CPACK_PACKAGING_INSTALL_PREFIX ${CMAKE_INSTALL_PREFIX})
  # Prevents CPack from generating file conflicts
  set(CPACK_RPM_EXCLUDE_FROM_AUTO_FILELIST_ADDITION "${CPACK_PACKAGING_INSTALL_PREFIX}")
  list(APPEND CPACK_RPM_EXCLUDE_FROM_AUTO_FILELIST_ADDITION "${CPACK_PACKAGING_INSTALL_PREFIX}/bin")
  list(APPEND CPACK_RPM_EXCLUDE_FROM_AUTO_FILELIST_ADDITION "${CPACK_PACKAGING_INSTALL_PREFIX}/include")
  list(APPEND CPACK_RPM_EXCLUDE_FROM_AUTO_FILELIST_ADDITION "${CPACK_PACKAGING_INSTALL_PREFIX}/lib")
  list(APPEND CPACK_RPM_EXCLUDE_FROM_AUTO_FILELIST_ADDITION "${PYTHON_SITE_LIB}")
  list(APPEND CPACK_RPM_EXCLUDE_FROM_AUTO_FILELIST_ADDITION "${PYTHON_SITE_PACKAGES}")
  set(CPACK_GENERATOR "RPM")
endif(${CMAKE_SYSTEM_NAME} MATCHES "Linux")

if(${CMAKE_SYSTEM_NAME} MATCHES "Windows")
  # Windows specific code
  set(CPACK_PACKAGE_INSTALL_DIRECTORY "AMCL")
  set(CPACK_NSIS_MODIFY_PATH ON)
  set(CPACK_PACKAGE_ICON "${CMAKE_SOURCE_DIR}/resources/icon\\\\icon.bmp")
  set(CPACK_NSIS_MUI_ICON "${CMAKE_SOURCE_DIR}/resources/icon\\\\icon.ico")
  set(CPACK_NSIS_MUI_UNIICON "${CMAKE_SOURCE_DIR}/resources/icon\\\\icon.ico")
  set(CPACK_NSIS_HELP_LINK "http://www.miracl.com")
  set(CPACK_NSIS_URL_INFO_ABOUT "http://www.miracl.com")
  set(CPACK_NSIS_CONTACT "support@miracl.com")
endif(${CMAKE_SYSTEM_NAME} MATCHES "Windows")

if(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
  # Mac OS X specific code
endif(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")

include (CPack)

