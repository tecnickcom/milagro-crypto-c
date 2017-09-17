# AMCL header files

# Copy curve dependant header files
function(CURVE_HEADERS_SET tb tf tc nb base nbt m8 mt ct pf)
  set(bd "${tb}_${base}")
  file(READ "config_big_XXX.h" temph)
  string(REGEX REPLACE "XXX" ${bd} temph "${temph}")
  string(REGEX REPLACE "@NB@" ${nb} temph "${temph}")
  string(REGEX REPLACE "@BASE@" ${base} temph "${temph}")
  file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/config_big_${bd}.h" "${temph}")

  set(CURVE_HEADERS ${CURVE_HEADERS}
  	${CMAKE_CURRENT_BINARY_DIR}/config_big_${bd}.h)

  file(READ "config_field_YYY.h" temph)
  string(REGEX REPLACE "XXX" ${bd} temph "${temph}")
  string(REGEX REPLACE "YYY" ${tf} temph "${temph}")
  string(REGEX REPLACE "@NBT@" ${nbt} temph "${temph}")
  string(REGEX REPLACE "@M8@" ${m8} temph "${temph}")
  string(REGEX REPLACE "@MT@" ${mt} temph "${temph}")
  math(EXPR sh "${base} * (1 + ((8 * ${nb} - 1) / ${base})) - ${nbt}")
  if(sh GREATER '30')
    set(sh '30')
  endif(sh GREATER '30')
  string(REGEX REPLACE "@SH@" ${sh} temph "${temph}")
  file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/config_field_${tf}.h" "${temph}")

  set(CURVE_HEADERS ${CURVE_HEADERS}
  	${CMAKE_CURRENT_BINARY_DIR}/config_field_${tf}.h)

  file(READ "config_curve_ZZZ.h" temph)
  string(REGEX REPLACE "XXX" ${bd} temph "${temph}")
  string(REGEX REPLACE "YYY" ${tf} temph "${temph}")
  string(REGEX REPLACE "ZZZ" ${tc} temph "${temph}")
  string(REGEX REPLACE "@CT@" ${ct} temph "${temph}")
  string(REGEX REPLACE "@PF@" ${pf} temph "${temph}")
  file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/config_curve_${tc}.h" "${temph}")

  set(CURVE_HEADERS ${CURVE_HEADERS}
  	${CMAKE_CURRENT_BINARY_DIR}/config_curve_${tc}.h)

  file(READ "big_XXX.h" temph)
  string(REGEX REPLACE "XXX" ${bd} temph "${temph}")
  file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/big_${bd}.h" "${temph}")

  set(CURVE_HEADERS ${CURVE_HEADERS}
  	${CMAKE_CURRENT_BINARY_DIR}/big_${bd}.h)

  file(READ "fp_YYY.h" temph)
  string(REGEX REPLACE "XXX" ${bd} temph "${temph}")
  string(REGEX REPLACE "YYY" ${tf} temph "${temph}")
  file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/fp_${tf}.h" "${temph}")

  set(CURVE_HEADERS ${CURVE_HEADERS}
  	${CMAKE_CURRENT_BINARY_DIR}/fp_${tf}.h)

  file(READ "ecp_ZZZ.h" temph)
  string(REGEX REPLACE "XXX" ${bd} temph "${temph}")
  string(REGEX REPLACE "YYY" ${tf} temph "${temph}")
  string(REGEX REPLACE "ZZZ" ${tc} temph "${temph}")
  file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/ecp_${tc}.h" "${temph}")

  set(CURVE_HEADERS ${CURVE_HEADERS}
  	${CMAKE_CURRENT_BINARY_DIR}/ecp_${tc}.h)

  file(READ "ecdh_ZZZ.h" temph)
  string(REGEX REPLACE "XXX" ${bd} temph "${temph}")
  string(REGEX REPLACE "YYY" ${tf} temph "${temph}")
  string(REGEX REPLACE "ZZZ" ${tc} temph "${temph}")
  file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/ecdh_${tc}.h" "${temph}")

  set(CURVE_HEADERS ${CURVE_HEADERS}
  	${CMAKE_CURRENT_BINARY_DIR}/ecdh_${tc}.h)

  file(COPY ecdh_support.h DESTINATION ${CMAKE_CURRENT_BINARY_DIR})

  set(CURVE_HEADERS ${CURVE_HEADERS}
    ${CMAKE_CURRENT_BINARY_DIR}/ecdh_support.h)

  if(NOT(${pf} STREQUAL "NOT"))

    file(READ "fp2_YYY.h" temph)
    string(REGEX REPLACE "XXX" ${bd} temph "${temph}")
    string(REGEX REPLACE "YYY" ${tf} temph "${temph}")
    file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/fp2_${tf}.h" "${temph}")

	  set(CURVE_HEADERS ${CURVE_HEADERS}
	  	${CMAKE_CURRENT_BINARY_DIR}/fp2_${tf}.h)

    file(READ "fp4_YYY.h" temph)
    string(REGEX REPLACE "XXX" ${bd} temph "${temph}")
    string(REGEX REPLACE "YYY" ${tf} temph "${temph}")
    file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/fp4_${tf}.h" "${temph}")

	  set(CURVE_HEADERS ${CURVE_HEADERS}
	  	${CMAKE_CURRENT_BINARY_DIR}/fp4_${tf}.h)

    file(READ "fp12_YYY.h" temph)
    string(REGEX REPLACE "XXX" ${bd} temph "${temph}")
    string(REGEX REPLACE "YYY" ${tf} temph "${temph}")
    file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/fp12_${tf}.h" "${temph}")

	  set(CURVE_HEADERS ${CURVE_HEADERS}
	  	${CMAKE_CURRENT_BINARY_DIR}/fp12_${tf}.h)

    file(READ "ecp2_ZZZ.h" temph)
    string(REGEX REPLACE "XXX" ${bd} temph "${temph}")
    string(REGEX REPLACE "YYY" ${tf} temph "${temph}")
    string(REGEX REPLACE "ZZZ" ${tc} temph "${temph}")
    file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/ecp2_${tc}.h" "${temph}")

	  set(CURVE_HEADERS ${CURVE_HEADERS}
	  	${CMAKE_CURRENT_BINARY_DIR}/ecp2_${tc}.h)

    file(READ "pair_ZZZ.h" temph)
    string(REGEX REPLACE "XXX" ${bd} temph "${temph}")
    string(REGEX REPLACE "YYY" ${tf} temph "${temph}")
    string(REGEX REPLACE "ZZZ" ${tc} temph "${temph}")
    file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/pair_${tc}.h" "${temph}")

	  set(CURVE_HEADERS ${CURVE_HEADERS}
	  	${CMAKE_CURRENT_BINARY_DIR}/pair_${tc}.h)

	if(BUILD_MPIN OR BUILD_WCC)
    file(COPY pbc_support.h DESTINATION ${CMAKE_CURRENT_BINARY_DIR})
    set(PBC_TIME_PERMIT_TESTS 10 CACHE STRING "Number of days in the future to test M-PIN/WCC time permits")
    set(PBC_RANDOM_TESTS 10 CACHE STRING "Number of random M-PIN/WCC tests")
    # message(STATUS "PBC_TIME_PERMIT_TESTS=${PBC_TIME_PERMIT_TESTS}")
    # message(STATUS "PBC_RANDOM_TESTS=${PBC_RANDOM_TESTS}")
    # Configure test variables
    configure_file (
      "config_test.h"
      "${CMAKE_CURRENT_BINARY_DIR}/config_test.h"
    )

	  set(CURVE_HEADERS ${CURVE_HEADERS}
	  	${CMAKE_CURRENT_BINARY_DIR}/pbc_support.h
	  	${CMAKE_CURRENT_BINARY_DIR}/config_test.h)
	endif(BUILD_MPIN OR BUILD_WCC)

    if(BUILD_MPIN)
      set(fnameh "mpin_${tc}.h")
      file(READ "mpin_ZZZ.h" temph)
      string(REGEX REPLACE "XXX" ${bd} temph "${temph}")
      string(REGEX REPLACE "YYY" ${tf} temph "${temph}")
      string(REGEX REPLACE "ZZZ" ${tc} temph "${temph}")
      file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/${fnameh}" "${temph}")

 		  set(CURVE_HEADERS ${CURVE_HEADERS}
	  		${CMAKE_CURRENT_BINARY_DIR}/mpin_${tc}.h)
    endif(BUILD_MPIN)

    if(BUILD_WCC)
      set(fnameh "wcc_${tc}.h")
      file(READ "wcc_ZZZ.h" temph)
      string(REGEX REPLACE "XXX" ${bd} temph "${temph}")
      string(REGEX REPLACE "YYY" ${tf} temph "${temph}")
      string(REGEX REPLACE "ZZZ" ${tc} temph "${temph}")
      file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/${fnameh}" "${temph}")

 		  set(CURVE_HEADERS ${CURVE_HEADERS}
	  		${CMAKE_CURRENT_BINARY_DIR}/wcc_${tc}.h)
    endif(BUILD_WCC)

  endif(NOT(${pf} STREQUAL "NOT"))

  install (FILES ${CURVE_HEADERS} DESTINATION include PERMISSIONS
        OWNER_WRITE OWNER_READ OWNER_EXECUTE
        GROUP_READ GROUP_EXECUTE
        WORLD_READ WORLD_EXECUTE)

  # if(BUILD_GO)
  #   # Copy header files for testing before installation
  #   file(COPY ${CURVE_HEADERS} DESTINATION "${PROJECT_BINARY_DIR}/wrappers/go/")
  # endif(BUILD_GO)

endfunction(CURVE_HEADERS_SET tb tf tc nb base nbt m8 mt ct pf)

# Copy rsa dependant header files
function(RSA_HEADERS_SET tb tff nb base ml)
  set(bd "${tb}_${base}")
  file(READ "config_big_XXX.h" temph)
  string(REGEX REPLACE "XXX" ${bd} temph "${temph}")
  string(REGEX REPLACE "@NB@" ${nb} temph "${temph}")
  string(REGEX REPLACE "@BASE@" ${base} temph "${temph}")
  file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/config_big_${bd}.h" "${temph}")

  set(RSA_HEADERS ${RSA_HEADERS}
		${CMAKE_CURRENT_BINARY_DIR}/config_big_${bd}.h)

  file(READ "config_ff_WWW.h" temph)
  string(REGEX REPLACE "XXX" ${bd} temph "${temph}")
  string(REGEX REPLACE "WWW" ${tff} temph "${temph}")
  string(REGEX REPLACE "@ML@" ${ml} temph "${temph}")
  file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/config_ff_${tff}.h" "${temph}")

  set(RSA_HEADERS ${RSA_HEADERS}
		${CMAKE_CURRENT_BINARY_DIR}/config_ff_${tff}.h)

  file(READ "big_XXX.h" temph)
  string(REGEX REPLACE "XXX" ${bd} temph "${temph}")
  file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/big_${bd}.h" "${temph}")

  set(RSA_HEADERS ${RSA_HEADERS}
		${CMAKE_CURRENT_BINARY_DIR}/big_${bd}.h)

  file(READ "ff_WWW.h" temph)
  string(REGEX REPLACE "XXX" ${bd} temph "${temph}")
  string(REGEX REPLACE "WWW" ${tff} temph "${temph}")
  file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/ff_${tff}.h" "${temph}")

  set(RSA_HEADERS ${RSA_HEADERS}
		${CMAKE_CURRENT_BINARY_DIR}/ff_${tff}.h)

  file(READ "rsa_WWW.h" temph)
  string(REGEX REPLACE "XXX" ${bd} temph "${temph}")
  string(REGEX REPLACE "WWW" ${tff} temph "${temph}")
  file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/rsa_${tff}.h" "${temph}")

  set(RSA_HEADERS ${RSA_HEADERS}
		${CMAKE_CURRENT_BINARY_DIR}/rsa_${tff}.h)

  file(COPY rsa_support.h DESTINATION ${CMAKE_CURRENT_BINARY_DIR})
  set(RSA_HEADERS ${RSA_HEADERS}
		${CMAKE_CURRENT_BINARY_DIR}/rsa_support.h)
  if(BUILD_X509)
  	file(COPY "x509.h" DESTINATION ${CMAKE_CURRENT_BINARY_DIR})
  	set(RSA_HEADERS ${RSA_HEADERS}
			${CMAKE_CURRENT_BINARY_DIR}/x509.h)
  endif(BUILD_X509)

  install (FILES ${RSA_HEADERS} DESTINATION include PERMISSIONS
        OWNER_WRITE OWNER_READ OWNER_EXECUTE
        GROUP_READ GROUP_EXECUTE
        WORLD_READ WORLD_EXECUTE)

  # if(BUILD_GO)
  #   # Copy header files for testing before installation
  #   file(COPY ${RSA_HEADERS} DESTINATION "${PROJECT_BINARY_DIR}/wrappers/go/")
  # endif(BUILD_GO)

endfunction(RSA_HEADERS_SET tb tff nb base ml)

set(COMMON_HEADERS
  	utils.h
    randapi.h
    version.h)

# Copy common header files to binary include directory
file(COPY ${COMMON_HEADERS} DESTINATION ${CMAKE_CURRENT_BINARY_DIR})

# Copy arch.h
file(READ "arch.h" temph)
string(REGEX REPLACE "@WL@" ${AMCL_CHUNK} temph "${temph}")
string(CONFIGURE "${temph}" temph)
file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/arch.h" "${temph}")

# Copy amcl.h
file(READ "amcl.h" temph)
string(CONFIGURE "${temph}" temph)
file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/amcl.h" "${temph}")

string(REGEX REPLACE "," ";" AMCL_CURVE ${AMCL_CURVE})
string(REGEX REPLACE "," ";" AMCL_RSA ${AMCL_RSA})

foreach(CURVE ${AMCL_CURVE})
  if(AMCL_CHUNK STREQUAL "64")
    if(CURVE STREQUAL "ED25519")
        CURVE_HEADERS_SET("256" "25519" "ED25519" "32" "56" "255" "5" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "C25519")
        CURVE_HEADERS_SET("256" "25519" "C25519" "32" "56" "255" "5" "PSEUDO_MERSENNE" "MONTGOMERY" "NOT")
    elseif(CURVE STREQUAL "NIST256")
        CURVE_HEADERS_SET("256" "NIST256" "NIST256" "32" "56" "256" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "BRAINPOOL")
        CURVE_HEADERS_SET("256" "BRAINPOOL" "BRAINPOOL" "32" "56" "256" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "ANSSI")
        CURVE_HEADERS_SET("256" "ANSSI" "ANSSI" "32" "56" "256" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "HIFIVE")
        CURVE_HEADERS_SET("336" "HIFIVE" "HIFIVE" "42" "60" "336" "5" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "GOLDILOCKS")
        CURVE_HEADERS_SET("448" "GOLDILOCKS" "GOLDILOCKS" "56" "58" "448" "7" "GENERALISED_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NIST384")
        CURVE_HEADERS_SET("384" "NIST384" "NIST384" "48" "56" "384" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "C41417")
        CURVE_HEADERS_SET("416" "C41417" "C41417" "52" "60" "414" "7" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NIST521")
        CURVE_HEADERS_SET("528" "NIST521" "NIST521" "66" "60" "521" "7" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS256W")
        CURVE_HEADERS_SET("256" "256PMW" "NUMS256W" "32" "56" "256" "3" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS256E")
        CURVE_HEADERS_SET("256" "256PME" "NUMS256E" "32" "56" "256" "3" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NUMS384W")
        CURVE_HEADERS_SET("384" "384PM" "NUMS384W" "48" "56" "384" "3" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS384E")
        CURVE_HEADERS_SET("384" "384PM" "NUMS384E" "48" "56" "384" "3" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NUMS512W")
        CURVE_HEADERS_SET("512" "512PM" "NUMS512W" "64" "56" "512" "7" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS512E")
        CURVE_HEADERS_SET("512" "512PM" "NUMS512E" "64" "56" "512" "7" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "BN254")
        CURVE_HEADERS_SET("256" "BN254" "BN254" "32" "56" "254" "3" "NOT_SPECIAL" "WEIERSTRASS" "BN")
    elseif(CURVE STREQUAL "BN254CX")
        CURVE_HEADERS_SET("256" "BN254CX" "BN254CX" "32" "56" "254" "3" "NOT_SPECIAL" "WEIERSTRASS" "BN")
    elseif(CURVE STREQUAL "BLS383")
        CURVE_HEADERS_SET("384" "BLS383" "BLS383" "48" "56" "383" "3" "NOT_SPECIAL" "WEIERSTRASS" "BLS")
    else()
      message(FATAL_ERROR "Invalid curve input")
    endif(CURVE STREQUAL "ED25519")
  elseif(AMCL_CHUNK STREQUAL "32")
    if(CURVE STREQUAL "ED25519")
        CURVE_HEADERS_SET("256" "25519" "ED25519" "32" "29" "255" "5" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "C25519")
        CURVE_HEADERS_SET("256" "25519" "C25519" "32" "29" "255" "5" "PSEUDO_MERSENNE" "MONTGOMERY" "NOT")
    elseif(CURVE STREQUAL "NIST256")
        CURVE_HEADERS_SET("256" "NIST256" "NIST256" "32" "28" "256" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "BRAINPOOL")
        CURVE_HEADERS_SET("256" "BRAINPOOL" "BRAINPOOL" "32" "28" "256" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "ANSSI")
        CURVE_HEADERS_SET("256" "ANSSI" "ANSSI" "32" "28" "256" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "HIFIVE")
        CURVE_HEADERS_SET("336" "HIFIVE" "HIFIVE" "42" "29" "336" "5" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "GOLDILOCKS")
        CURVE_HEADERS_SET("448" "GOLDILOCKS" "GOLDILOCKS" "56" "29" "448" "7" "GENERALISED_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NIST384")
        CURVE_HEADERS_SET("384" "NIST384" "NIST384" "48" "29" "384" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "C41417")
        CURVE_HEADERS_SET("416" "C41417" "C41417" "52" "29" "414" "7" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NIST521")
        CURVE_HEADERS_SET("528" "NIST521" "NIST521" "66" "28" "521" "7" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS256W")
        CURVE_HEADERS_SET("256" "256PMW" "NUMS256W" "32" "28" "256" "3" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS256E")
        CURVE_HEADERS_SET("256" "256PME" "NUMS256E" "32" "29" "256" "3" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NUMS384W")
        CURVE_HEADERS_SET("384" "384PM" "NUMS384W" "48" "29" "384" "3" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS384E")
        CURVE_HEADERS_SET("384" "384PM" "NUMS384E" "48" "29" "384" "3" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NUMS512W")
        CURVE_HEADERS_SET("512" "512PM" "NUMS512W" "64" "29" "512" "7" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS512E")
        CURVE_HEADERS_SET("512" "512PM" "NUMS512E" "64" "29" "512" "7" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "BN254")
        CURVE_HEADERS_SET("256" "BN254" "BN254" "32" "28" "254" "3" "NOT_SPECIAL" "WEIERSTRASS" "BN")
    elseif(CURVE STREQUAL "BN254CX")
        CURVE_HEADERS_SET("256" "BN254CX" "BN254CX" "32" "28" "254" "3" "NOT_SPECIAL" "WEIERSTRASS" "BN")
    elseif(CURVE STREQUAL "BLS383")
        CURVE_HEADERS_SET("384" "BLS383" "BLS383" "48" "29" "383" "3" "NOT_SPECIAL" "WEIERSTRASS" "BLS")
    else()
      message(FATAL_ERROR "Invalid curve input")
    endif(CURVE STREQUAL "ED25519")
  elseif(AMCL_CHUNK STREQUAL "16")
    if(CURVE STREQUAL "ED25519")
        CURVE_HEADERS_SET("256" "25519" "ED25519" "32" "13" "255" "5" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "C25519")
        CURVE_HEADERS_SET("256" "25519" "C25519" "32" "13" "255" "5" "PSEUDO_MERSENNE" "MONTGOMERY" "NOT")
    elseif(CURVE STREQUAL "NUMS256E")
        CURVE_HEADERS_SET("256" "256PME" "NUMS256E" "32" "13" "256" "3" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "BN254")
        CURVE_HEADERS_SET("256" "BN254" "BN254" "32" "13" "254" "3" "NOT_SPECIAL" "WEIERSTRASS" "BN")
    elseif(CURVE STREQUAL "BN254CX")
        CURVE_HEADERS_SET("256" "BN254CX" "BN254CX" "32" "13" "254" "3" "NOT_SPECIAL" "WEIERSTRASS" "BN")
    else()
      message(FATAL_ERROR "Invalid curve input")
    endif(CURVE STREQUAL "ED25519")
  endif(AMCL_CHUNK STREQUAL "64")
endforeach(CURVE)

foreach(RSA ${AMCL_RSA})
  if(AMCL_CHUNK STREQUAL "64")
    if(RSA STREQUAL "2048")
  	  RSA_HEADERS_SET("1024" "2048" "128" "58" "2")
    elseif(RSA STREQUAL "3072")
  	  RSA_HEADERS_SET("384" "3072" "48" "56" "8")
    elseif(RSA STREQUAL "4096")
  	  RSA_HEADERS_SET("512" "4096" "64" "60" "8")
    else()
      message(FATAL_ERROR "Invalid RSA input")
    endif(RSA STREQUAL "2048")
  elseif(AMCL_CHUNK STREQUAL "32")
    if(RSA STREQUAL "2048")
      RSA_HEADERS_SET("1024" "2048" "128" "28" "2")
    elseif(RSA STREQUAL "3072")
      RSA_HEADERS_SET("384" "3072" "48" "28" "8")
    elseif(RSA STREQUAL "4096")
      RSA_HEADERS_SET("512" "4096" "64" "29" "8")
    else()
        message(FATAL_ERROR "Invalid RSA input")
    endif(RSA STREQUAL "2048")
  elseif(AMCL_CHUNK STREQUAL "16")
    if(RSA STREQUAL "2048")
      RSA_HEADERS_SET("256" "2048" "32" "13" "8")
    else()
      message(FATAL_ERROR "Invalid RSA input")
    endif(RSA STREQUAL "2048")
  else()
      message(FATAL_ERROR "Invalid CHUNK size")
  endif(AMCL_CHUNK STREQUAL "64")
endforeach(RSA)



install (DIRECTORY DESTINATION include DIRECTORY_PERMISSIONS
        OWNER_WRITE OWNER_READ OWNER_EXECUTE
        GROUP_WRITE GROUP_READ
        WORLD_WRITE WORLD_READ WORLD_EXECUTE)

install (FILES ${COMMON_HEADERS} DESTINATION include PERMISSIONS
        OWNER_WRITE OWNER_READ OWNER_EXECUTE
        GROUP_READ GROUP_EXECUTE
        WORLD_READ WORLD_EXECUTE)

install (FILES ${CMAKE_CURRENT_BINARY_DIR}/amcl.h DESTINATION include PERMISSIONS
        OWNER_WRITE OWNER_READ OWNER_EXECUTE
        GROUP_READ GROUP_EXECUTE
        WORLD_READ WORLD_EXECUTE)

install (FILES ${CMAKE_CURRENT_BINARY_DIR}/arch.h DESTINATION include PERMISSIONS
        OWNER_WRITE OWNER_READ OWNER_EXECUTE
        GROUP_READ GROUP_EXECUTE
        WORLD_READ WORLD_EXECUTE)

# if(BUILD_GO)
#   # Copy header files for testing before installation
#   file(COPY ${COMMON_HEADERS} DESTINATION "${PROJECT_BINARY_DIR}/wrappers/go/")
#   file(COPY ${CMAKE_CURRENT_BINARY_DIR}/amcl.h DESTINATION "${PROJECT_BINARY_DIR}/wrappers/go/")
#   file(COPY ${CMAKE_CURRENT_BINARY_DIR}/arch.h DESTINATION "${PROJECT_BINARY_DIR}/wrappers/go/")
# endif(BUILD_GO)