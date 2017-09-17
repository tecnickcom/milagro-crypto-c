
INSTALL(DIRECTORY DESTINATION "${GO_PATH}/src/github.com" DIRECTORY_PERMISSIONS
        OWNER_WRITE OWNER_READ OWNER_EXECUTE
        GROUP_WRITE GROUP_READ
        WORLD_WRITE WORLD_READ WORLD_EXECUTE)

INSTALL(DIRECTORY DESTINATION "${GO_PATH}/src/github.com/miracl" DIRECTORY_PERMISSIONS
        OWNER_WRITE OWNER_READ OWNER_EXECUTE
        GROUP_WRITE GROUP_READ
        WORLD_WRITE WORLD_READ WORLD_EXECUTE)

INSTALL(DIRECTORY DESTINATION "${GO_PATH}/src/github.com/miracl/amcl-go-wrapper" DIRECTORY_PERMISSIONS
        OWNER_WRITE OWNER_READ OWNER_EXECUTE
        GROUP_WRITE GROUP_READ
        WORLD_WRITE WORLD_READ WORLD_EXECUTE)

# Configure go amcl wrapper
configure_file (
	"${PROJECT_SOURCE_DIR}/wrappers/go/amcl.go"
	"${PROJECT_BINARY_DIR}/wrappers/go/amcl.go"
)

file(COPY "${PROJECT_SOURCE_DIR}/wrappers/go/README.md" DESTINATION "${PROJECT_BINARY_DIR}/wrappers/go/")

INSTALL(FILES ${CMAKE_CURRENT_BINARY_DIR}/amcl.go DESTINATION "${GO_PATH}/src/github.com/miracl/amcl-go-wrapper" PERMISSIONS
    OWNER_WRITE OWNER_READ OWNER_EXECUTE
    GROUP_READ GROUP_EXECUTE
    WORLD_READ WORLD_EXECUTE)

if(BUILD_MPIN)
	# Configure pbc_support wrapper
	configure_file (
		"${PROJECT_SOURCE_DIR}/wrappers/go/pbc_support.go"
		"${PROJECT_BINARY_DIR}/wrappers/go/pbc_support.go"
	)

		# Configure go pbc_support test
	configure_file (
		"${PROJECT_SOURCE_DIR}/wrappers/go/pbc_support_test.go"
		"${PROJECT_BINARY_DIR}/wrappers/go/pbc_support_test.go"
	)

  INSTALL(FILES ${CMAKE_CURRENT_BINARY_DIR}/pbc_support.go pbc_support_test.go DESTINATION "${GO_PATH}/src/github.com/miracl/amcl-go-wrapper" PERMISSIONS
	    OWNER_WRITE OWNER_READ OWNER_EXECUTE
	    GROUP_READ GROUP_EXECUTE
	    WORLD_READ WORLD_EXECUTE)

endif(BUILD_MPIN)

# Configure rsa_support wrapper
configure_file (
	"${PROJECT_SOURCE_DIR}/wrappers/go/rsa_support.go"
	"${PROJECT_BINARY_DIR}/wrappers/go/rsa_support.go"
)

INSTALL(FILES ${CMAKE_CURRENT_BINARY_DIR}/rsa_support.go DESTINATION "${GO_PATH}/src/github.com/miracl/amcl-go-wrapper" PERMISSIONS
  OWNER_WRITE OWNER_READ OWNER_EXECUTE
  GROUP_READ GROUP_EXECUTE
  WORLD_READ WORLD_EXECUTE)

if(NOT(AMCL_CURVE MATCHES "C25519"))

	# Configure ecdh_support wrapper
	configure_file (
		"${PROJECT_SOURCE_DIR}/wrappers/go/ecdsa_support.go"
		"${PROJECT_BINARY_DIR}/wrappers/go/ecdsa_support.go"
	)

	INSTALL(FILES ${CMAKE_CURRENT_BINARY_DIR}/ecdsa_support.go DESTINATION "${GO_PATH}/src/github.com/miracl/amcl-go-wrapper" PERMISSIONS
	  OWNER_WRITE OWNER_READ OWNER_EXECUTE
	  GROUP_READ GROUP_EXECUTE
	  WORLD_READ WORLD_EXECUTE)

endif(NOT(AMCL_CURVE MATCHES "C25519"))

function(CURVE_GOLANG_SET tb tf tc nb base nbt m8 mt ct pf)

	if(NOT(AMCL_CURVE MATCHES "C25519"))
	  # Configure go ecdsa wrapper
	  file(READ "${PROJECT_SOURCE_DIR}/wrappers/go/ecdsa_ZZZ.go" tempc)
	  string(REGEX REPLACE "ZZZ" ${tc} tempc "${tempc}")
	  string(CONFIGURE "${tempc}" tempc)
	  file(WRITE "${PROJECT_BINARY_DIR}/wrappers/go/ecdsa_${tc}.go" "${tempc}")

	  # Configure go ecdsa test
	  file(READ "${PROJECT_SOURCE_DIR}/wrappers/go/ecdsa_ZZZ_test.go" tempc)
	  string(REGEX REPLACE "ZZZ" ${tc} tempc "${tempc}")
	  file(WRITE "${PROJECT_BINARY_DIR}/wrappers/go/ecdsa_${tc}_test.go" "${tempc}")

		INSTALL(FILES ${PROJECT_BINARY_DIR}/wrappers/go/ecdsa_${tc}.go DESTINATION "${GO_PATH}/src/github.com/miracl/amcl-go-wrapper" PERMISSIONS
		  OWNER_WRITE OWNER_READ OWNER_EXECUTE
		  GROUP_READ GROUP_EXECUTE
		  WORLD_READ WORLD_EXECUTE)
	endif(NOT(AMCL_CURVE MATCHES "C25519"))

	if(BUILD_MPIN AND NOT(${pf} STREQUAL "NOT"))

	  # Configure go mpin wrapper
	  file(READ "${PROJECT_SOURCE_DIR}/wrappers/go/mpin_ZZZ.go" tempc)
	  string(REGEX REPLACE "ZZZ" ${tc} tempc "${tempc}")
	  string(CONFIGURE "${tempc}" tempc)
	  file(WRITE "${PROJECT_BINARY_DIR}/wrappers/go/mpin_${tc}.go" "${tempc}")

	  # Configure go mpin test
	  file(READ "${PROJECT_SOURCE_DIR}/wrappers/go/mpin_ZZZ_test.go" tempc)
	  string(REGEX REPLACE "ZZZ" ${tc} tempc "${tempc}")
	  file(WRITE "${PROJECT_BINARY_DIR}/wrappers/go/mpin_${tc}_test.go" "${tempc}")

    INSTALL(FILES ${CMAKE_CURRENT_BINARY_DIR}/mpin_${tc}.go DESTINATION "${GO_PATH}/src/github.com/miracl/amcl-go-wrapper" PERMISSIONS
      OWNER_WRITE OWNER_READ OWNER_EXECUTE
      GROUP_READ GROUP_EXECUTE
      WORLD_READ WORLD_EXECUTE)

	endif(BUILD_MPIN AND NOT(${pf} STREQUAL "NOT"))
endfunction(CURVE_GOLANG_SET tb tf tc nb base nbt m8 mt ct pf)

function(RSA_GOLANG_SET tb tff nb base ml)
  # Configure go rsa wrapper
  file(READ "${PROJECT_SOURCE_DIR}/wrappers/go/rsa_WWW.go" tempc)
  string(REGEX REPLACE "WWW" ${tff} tempc "${tempc}")
  string(CONFIGURE "${tempc}" tempc)
  file(WRITE "${PROJECT_BINARY_DIR}/wrappers/go/rsa_${tff}.go" "${tempc}")

  # Configure go rsa test
  file(READ "${PROJECT_SOURCE_DIR}/wrappers/go/rsa_WWW_test.go" tempc)
  string(REGEX REPLACE "WWW" ${tff} tempc "${tempc}")
  file(WRITE "${PROJECT_BINARY_DIR}/wrappers/go/rsa_${tff}_test.go" "${tempc}")

 INSTALL(FILES ${PROJECT_BINARY_DIR}/wrappers/go/rsa_${tff}.go DESTINATION "${GO_PATH}/src/github.com/miracl/amcl-go-wrapper" PERMISSIONS
        OWNER_WRITE OWNER_READ OWNER_EXECUTE
        GROUP_READ GROUP_EXECUTE
        WORLD_READ WORLD_EXECUTE)
endfunction(RSA_GOLANG_SET tb tff nb base ml)


string(REGEX REPLACE "," ";" AMCL_CURVE ${AMCL_CURVE})
string(REGEX REPLACE "," ";" AMCL_RSA ${AMCL_RSA})

foreach(CURVE ${AMCL_CURVE})
  if(AMCL_CHUNK STREQUAL "64")
    if(CURVE STREQUAL "ED25519")
        CURVE_GOLANG_SET("256" "25519" "ED25519" "32" "56" "255" "5" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "C25519")
        CURVE_GOLANG_SET("256" "25519" "C25519" "32" "56" "255" "5" "PSEUDO_MERSENNE" "MONTGOMERY" "NOT")
    elseif(CURVE STREQUAL "NIST256")
        CURVE_GOLANG_SET("256" "NIST256" "NIST256" "32" "56" "256" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "BRAINPOOL")
        CURVE_GOLANG_SET("256" "BRAINPOOL" "BRAINPOOL" "32" "56" "256" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "ANSSI")
        CURVE_GOLANG_SET("256" "ANSSI" "ANSSI" "32" "56" "256" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "HIFIVE")
        CURVE_GOLANG_SET("336" "HIFIVE" "HIFIVE" "42" "60" "336" "5" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "GOLDILOCKS")
        CURVE_GOLANG_SET("448" "GOLDILOCKS" "GOLDILOCKS" "56" "58" "448" "7" "GENERALISED_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NIST384")
        CURVE_GOLANG_SET("384" "NIST384" "NIST384" "48" "56" "384" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "C41417")
        CURVE_GOLANG_SET("416" "C41417" "C41417" "52" "60" "414" "7" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NIST521")
        CURVE_GOLANG_SET("528" "NIST521" "NIST521" "66" "60" "521" "7" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS256W")
        CURVE_GOLANG_SET("256" "256PMW" "NUMS256W" "32" "56" "256" "3" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS256E")
        CURVE_GOLANG_SET("256" "256PME" "NUMS256E" "32" "56" "256" "3" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NUMS384W")
        CURVE_GOLANG_SET("384" "384PM" "NUMS384W" "48" "56" "384" "3" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS384E")
        CURVE_GOLANG_SET("384" "384PM" "NUMS384E" "48" "56" "384" "3" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NUMS512W")
        CURVE_GOLANG_SET("512" "512PM" "NUMS512W" "64" "56" "512" "7" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS512E")
        CURVE_GOLANG_SET("512" "512PM" "NUMS512E" "64" "56" "512" "7" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "BN254")
        CURVE_GOLANG_SET("256" "BN254" "BN254" "32" "56" "254" "3" "NOT_SPECIAL" "WEIERSTRASS" "BN")
    elseif(CURVE STREQUAL "BN254CX")
        CURVE_GOLANG_SET("256" "BN254CX" "BN254CX" "32" "56" "254" "3" "NOT_SPECIAL" "WEIERSTRASS" "BN")
    elseif(CURVE STREQUAL "BLS383")
        CURVE_GOLANG_SET("384" "BLS383" "BLS383" "48" "56" "383" "3" "NOT_SPECIAL" "WEIERSTRASS" "BLS")
    else()
      message(FATAL_ERROR "Invalid curve input")
    endif(CURVE STREQUAL "ED25519")
  elseif(AMCL_CHUNK STREQUAL "32")
    if(CURVE STREQUAL "ED25519")
        CURVE_GOLANG_SET("256" "25519" "ED25519" "32" "29" "255" "5" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "C25519")
        CURVE_GOLANG_SET("256" "25519" "C25519" "32" "29" "255" "5" "PSEUDO_MERSENNE" "MONTGOMERY" "NOT")
    elseif(CURVE STREQUAL "NIST256")
        CURVE_GOLANG_SET("256" "NIST256" "NIST256" "32" "28" "256" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "BRAINPOOL")
        CURVE_GOLANG_SET("256" "BRAINPOOL" "BRAINPOOL" "32" "28" "256" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "ANSSI")
        CURVE_GOLANG_SET("256" "ANSSI" "ANSSI" "32" "28" "256" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "HIFIVE")
        CURVE_GOLANG_SET("336" "HIFIVE" "HIFIVE" "42" "29" "336" "5" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "GOLDILOCKS")
        CURVE_GOLANG_SET("448" "GOLDILOCKS" "GOLDILOCKS" "56" "29" "448" "7" "GENERALISED_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NIST384")
        CURVE_GOLANG_SET("384" "NIST384" "NIST384" "48" "29" "384" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "C41417")
        CURVE_GOLANG_SET("416" "C41417" "C41417" "52" "29" "414" "7" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NIST521")
        CURVE_GOLANG_SET("528" "NIST521" "NIST521" "66" "28" "521" "7" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS256W")
        CURVE_GOLANG_SET("256" "256PMW" "NUMS256W" "32" "28" "256" "3" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS256E")
        CURVE_GOLANG_SET("256" "256PME" "NUMS256E" "32" "29" "256" "3" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NUMS384W")
        CURVE_GOLANG_SET("384" "384PM" "NUMS384W" "48" "29" "384" "3" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS384E")
        CURVE_GOLANG_SET("384" "384PM" "NUMS384E" "48" "29" "384" "3" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NUMS512W")
        CURVE_GOLANG_SET("512" "512PM" "NUMS512W" "64" "29" "512" "7" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS512E")
        CURVE_GOLANG_SET("512" "512PM" "NUMS512E" "64" "29" "512" "7" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "BN254")
        CURVE_GOLANG_SET("256" "BN254" "BN254" "32" "28" "254" "3" "NOT_SPECIAL" "WEIERSTRASS" "BN")
    elseif(CURVE STREQUAL "BN254CX")
        CURVE_GOLANG_SET("256" "BN254CX" "BN254CX" "32" "28" "254" "3" "NOT_SPECIAL" "WEIERSTRASS" "BN")
    elseif(CURVE STREQUAL "BLS383")
        CURVE_GOLANG_SET("384" "BLS383" "BLS383" "48" "29" "383" "3" "NOT_SPECIAL" "WEIERSTRASS" "BLS")
    else()
      message(FATAL_ERROR "Invalid curve input")
    endif(CURVE STREQUAL "ED25519")
  elseif(AMCL_CHUNK STREQUAL "16")
    if(CURVE STREQUAL "ED25519")
        CURVE_GOLANG_SET("256" "25519" "ED25519" "32" "13" "255" "5" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "C25519")
        CURVE_GOLANG_SET("256" "25519" "C25519" "32" "13" "255" "5" "PSEUDO_MERSENNE" "MONTGOMERY" "NOT")
    elseif(CURVE STREQUAL "NUMS256E")
        CURVE_GOLANG_SET("256" "256PME" "NUMS256E" "32" "13" "256" "3" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "BN254")
        CURVE_GOLANG_SET("256" "BN254" "BN254" "32" "13" "254" "3" "NOT_SPECIAL" "WEIERSTRASS" "BN")
    elseif(CURVE STREQUAL "BN254CX")
        CURVE_GOLANG_SET("256" "BN254CX" "BN254CX" "32" "13" "254" "3" "NOT_SPECIAL" "WEIERSTRASS" "BN")
    else()
      message(FATAL_ERROR "Invalid curve input")
    endif(CURVE STREQUAL "ED25519")
  endif(AMCL_CHUNK STREQUAL "64")
endforeach(CURVE)

foreach(RSA ${AMCL_RSA})
  if(AMCL_CHUNK STREQUAL "64")
    if(RSA STREQUAL "2048")
      RSA_GOLANG_SET("1024" "2048" "128" "58" "2")
    elseif(RSA STREQUAL "3072")
      RSA_GOLANG_SET("384" "3072" "48" "56" "8")
    elseif(RSA STREQUAL "4096")
      RSA_GOLANG_SET("512" "4096" "64" "60" "8")
    else()
      message(FATAL_ERROR "Invalid RSA input")
    endif(RSA STREQUAL "2048")
  elseif(AMCL_CHUNK STREQUAL "32")
    if(RSA STREQUAL "2048")
      RSA_GOLANG_SET("1024" "2048" "128" "28" "2")
    elseif(RSA STREQUAL "3072")
      RSA_GOLANG_SET("384" "3072" "48" "28" "8")
    elseif(RSA STREQUAL "4096")
      RSA_GOLANG_SET("512" "4096" "64" "29" "8")
    else()
        message(FATAL_ERROR "Invalid RSA input")
    endif(RSA STREQUAL "2048")
  elseif(AMCL_CHUNK STREQUAL "16")
    if(RSA STREQUAL "2048")
      RSA_GOLANG_SET("256" "2048" "32" "13" "8")
    else()
      message(FATAL_ERROR "Invalid RSA input")
    endif(RSA STREQUAL "2048")
  else()
    message(FATAL_ERROR "Invalid CHUNK size")
  endif(AMCL_CHUNK STREQUAL "64")
endforeach(RSA)

add_test(NAME test_golang_wrappers COMMAND go test -race -v)
set_tests_properties (test_golang_wrappers PROPERTIES PAS_ZZZS_REGULAR_EXPRESSION PAS_ZZZS) #this need to be fixed since it doesn't catch failing tests

