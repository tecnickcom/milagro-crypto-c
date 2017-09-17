# Create AMCL library

file(MAKE_DIRECTORY ${PROJECT_BINARY_DIR}/lib)

set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${PROJECT_BINARY_DIR}/lib)
set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ${PROJECT_BINARY_DIR}/lib)
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${PROJECT_BINARY_DIR}/lib)

# Add the binary tree directory to the search path for linking and include files
link_directories( ${CMAKE_CURRENT_BINARY_DIR} )
include_directories (${CMAKE_CURRENT_BINARY_DIR} ${PROJECT_BINARY_DIR}/include )

# create new file replacing curve dependant parameters
function(CURVE_SRC_SET tb tf tc nb base nbt m8 mt ct pf)
  set(bd "${tb}_${base}")

  file(READ "big_XXX.c" tempc)
  string(REGEX REPLACE "XXX" ${bd} tempc "${tempc}")
  file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/big_${bd}.c" "${tempc}")
  set(SOURCES_CURVE_${tc}
    ${CMAKE_CURRENT_BINARY_DIR}/big_${bd}.c)

  file(READ "fp_YYY.c" tempc)
  string(REGEX REPLACE "XXX" ${bd} tempc "${tempc}")
  string(REGEX REPLACE "YYY" ${tf} tempc "${tempc}")
  file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/fp_${tf}.c" "${tempc}")
  file(COPY "rom_field_${tf}.c" DESTINATION ${CMAKE_CURRENT_BINARY_DIR})
  set(SOURCES_CURVE_${tc} ${SOURCES_CURVE_${tc}}
    ${CMAKE_CURRENT_BINARY_DIR}/fp_${tf}.c
    ${CMAKE_CURRENT_BINARY_DIR}/rom_field_${tf}.c)

  file(READ "ecp_ZZZ.c" tempc)
  string(REGEX REPLACE "XXX" ${bd} tempc "${tempc}")
  string(REGEX REPLACE "YYY" ${tf} tempc "${tempc}")
  string(REGEX REPLACE "ZZZ" ${tc} tempc "${tempc}")
  file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/ecp_${tc}.c" "${tempc}")
  set(SOURCES_CURVE_${tc} ${SOURCES_CURVE_${tc}}
    ${CMAKE_CURRENT_BINARY_DIR}/ecp_${tc}.c)

  file(READ "ecdh_ZZZ.c" tempc)
  string(REGEX REPLACE "XXX" ${bd} tempc "${tempc}")
  string(REGEX REPLACE "YYY" ${tf} tempc "${tempc}")
  string(REGEX REPLACE "ZZZ" ${tc} tempc "${tempc}")
  file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/ecdh_${tc}.c" "${tempc}")
  file(COPY "rom_curve_${tc}.c" DESTINATION ${CMAKE_CURRENT_BINARY_DIR})
  set(SOURCES_CURVE_${tc} ${SOURCES_CURVE_${tc}}
    ${CMAKE_CURRENT_BINARY_DIR}/ecdh_${tc}.c
    ${CMAKE_CURRENT_BINARY_DIR}/rom_curve_${tc}.c)

  message(STATUS "Build libamcl_curve_${tc}")
  add_library(amcl_curve_${tc} ${LIB_TYPE} ${SOURCES_CURVE_${tc}})
  target_link_libraries (amcl_curve_${tc} amcl_core)

  set(INSTALL_LIBS ${INSTALL_LIBS} amcl_curve_${tc})

  if(NOT(${pf} STREQUAL "NOT"))

    file(READ "fp2_YYY.c" tempc)
    string(REGEX REPLACE "XXX" ${bd} tempc "${tempc}")
    string(REGEX REPLACE "YYY" ${tf} tempc "${tempc}")
    file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/fp2_${tc}.c" "${tempc}")
    set(SOURCES_PAIRING_${tc} ${SOURCES_PAIRING_${tc}}
      ${CMAKE_CURRENT_BINARY_DIR}/fp2_${tc}.c)

    file(READ "fp4_YYY.c" tempc)
    string(REGEX REPLACE "XXX" ${bd} tempc "${tempc}")
    string(REGEX REPLACE "YYY" ${tf} tempc "${tempc}")
    file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/fp4_${tc}.c" "${tempc}")
    set(SOURCES_PAIRING_${tc} ${SOURCES_PAIRING_${tc}}
    ${CMAKE_CURRENT_BINARY_DIR}/fp4_${tc}.c)

    file(READ "fp12_YYY.c" tempc)
    string(REGEX REPLACE "XXX" ${bd} tempc "${tempc}")
    string(REGEX REPLACE "YYY" ${tf} tempc "${tempc}")
    file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/fp12_${tc}.c" "${tempc}")
    set(SOURCES_PAIRING_${tc} ${SOURCES_PAIRING_${tc}}
      ${CMAKE_CURRENT_BINARY_DIR}/fp12_${tc}.c)

    file(READ "ecp2_ZZZ.c" tempc)
    string(REGEX REPLACE "XXX" ${bd} tempc "${tempc}")
    string(REGEX REPLACE "YYY" ${tf} tempc "${tempc}")
    string(REGEX REPLACE "ZZZ" ${tc} tempc "${tempc}")
    file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/ecp2_${tc}.c" "${tempc}")
    set(SOURCES_PAIRING_${tc} ${SOURCES_PAIRING_${tc}}
      ${CMAKE_CURRENT_BINARY_DIR}/ecp2_${tc}.c)

    file(READ "pair_ZZZ.c" tempc)
    string(REGEX REPLACE "XXX" ${bd} tempc "${tempc}")
    string(REGEX REPLACE "YYY" ${tf} tempc "${tempc}")
    string(REGEX REPLACE "ZZZ" ${tc} tempc "${tempc}")
    file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/pair_${tc}.c" "${tempc}")
    set(SOURCES_PAIRING_${tc} ${SOURCES_PAIRING_${tc}}
      ${CMAKE_CURRENT_BINARY_DIR}/pair_${tc}.c)

    message(STATUS "Build libamcl_pairing_${tc}")
    add_library(amcl_pairing_${tc} ${LIB_TYPE} ${SOURCES_PAIRING_${tc}})
    target_link_libraries (amcl_pairing_${tc} amcl_curve_${tc} amcl_core)

    set(INSTALL_LIBS ${INSTALL_LIBS} amcl_pairing_${tc})

    if(${CMAKE_SYSTEM_NAME} MATCHES "Windows")
        add_custom_command(
                    TARGET amcl_pairing_${tc}
                    POST_BUILD
                    COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_pairing_${tc}.dll" "${PROJECT_BINARY_DIR}/test/")
    endif(${CMAKE_SYSTEM_NAME} MATCHES "Windows")

    # Copy built libs to wrappers/go to allow testing without
    # having to install the libs.
    if (BUILD_GO AND BUILD_SHARED_LIBS)
      if(${CMAKE_SYSTEM_NAME} MATCHES "Linux")
        add_custom_command(
                    TARGET amcl_pairing_${tc}
                    POST_BUILD
                    COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_pairing_${tc}.so" "${PROJECT_BINARY_DIR}/wrappers/go/")
      endif(${CMAKE_SYSTEM_NAME} MATCHES "Linux")

      if(${CMAKE_SYSTEM_NAME} MATCHES "Windows")
        add_custom_command(
                    TARGET amcl_pairing_${tc}
                    POST_BUILD
                    COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_pairing_${tc}.dll" "${PROJECT_BINARY_DIR}/wrappers/go/")
      endif(${CMAKE_SYSTEM_NAME} MATCHES "Windows")

      if(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
        add_custom_command(
                    TARGET amcl_pairing_${tc}
                    POST_BUILD
                    COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_pairing_${tc}.dylib" "${PROJECT_BINARY_DIR}/wrappers/go/")
      endif(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
    endif (BUILD_GO AND BUILD_SHARED_LIBS)

    # Static build
    if (BUILD_GO AND NOT BUILD_SHARED_LIBS)
      if(${CMAKE_SYSTEM_NAME} MATCHES "Linux")
        add_custom_command(
                    TARGET amcl_pairing_${tc}
                    POST_BUILD
                    COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_pairing_${tc}.a" "${PROJECT_BINARY_DIR}/wrappers/go/")
      endif(${CMAKE_SYSTEM_NAME} MATCHES "Linux")
    endif (BUILD_GO AND NOT BUILD_SHARED_LIBS)

    # Copy built libs to wrappers/python to allow testing without
    # having to install the libs.
    if (BUILD_PYTHON AND BUILD_SHARED_LIBS)
      if(${CMAKE_SYSTEM_NAME} MATCHES "Linux")
        add_custom_command(
                    TARGET amcl_pairing_${tc}
                    POST_BUILD
                    COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_pairing_${tc}.so" "${PROJECT_BINARY_DIR}/wrappers/python/")
      endif(${CMAKE_SYSTEM_NAME} MATCHES "Linux")

      if(${CMAKE_SYSTEM_NAME} MATCHES "Windows")
        add_custom_command(
                    TARGET amcl_pairing_${tc}
                    POST_BUILD
                    COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_pairing_${tc}.dll" "${PROJECT_BINARY_DIR}/wrappers/python/")
      endif(${CMAKE_SYSTEM_NAME} MATCHES "Windows")

      if(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
        add_custom_command(
                    TARGET amcl_pairing_${tc}
                    POST_BUILD
                    COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_pairing_${tc}.dylib" "${PROJECT_BINARY_DIR}/wrappers/python/")
      endif(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
    endif (BUILD_PYTHON AND BUILD_SHARED_LIBS)

    if(BUILD_MPIN)
      file(READ "mpin_ZZZ.c" tempc)
      string(REGEX REPLACE "XXX" ${bd} tempc "${tempc}")
      string(REGEX REPLACE "YYY" ${tf} tempc "${tempc}")
      string(REGEX REPLACE "ZZZ" ${tc} tempc "${tempc}")
      file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/mpin_${tc}.c" "${tempc}")
      set(SOURCES_MPIN_${tc} ${SOURCES_MPIN_${tc}}
        ${CMAKE_CURRENT_BINARY_DIR}/mpin_${tc}.c)

      message(STATUS "Build libamcl_mpin_${tc}")
      add_library(amcl_mpin_${tc} ${LIB_TYPE} ${SOURCES_MPIN_${tc}})
      target_link_libraries (amcl_mpin_${tc} amcl_pairing_${tc} amcl_curve_${tc} amcl_core)

      set(INSTALL_LIBS ${INSTALL_LIBS} amcl_mpin_${tc})

      if(${CMAKE_SYSTEM_NAME} MATCHES "Windows")
          add_custom_command(
                      TARGET amcl_mpin_${tc}
                      POST_BUILD
                      COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_mpin_${tc}.dll" "${PROJECT_BINARY_DIR}/test/")
      endif(${CMAKE_SYSTEM_NAME} MATCHES "Windows")

      # Copy built libs to wrappers/go to allow testing without
      # having to install the libs.
      if (BUILD_GO AND BUILD_SHARED_LIBS)
        if(${CMAKE_SYSTEM_NAME} MATCHES "Linux")
          add_custom_command(
                      TARGET amcl_mpin_${tc}
                      POST_BUILD
                      COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_mpin_${tc}.so" "${PROJECT_BINARY_DIR}/wrappers/go/")
        endif(${CMAKE_SYSTEM_NAME} MATCHES "Linux")

        if(${CMAKE_SYSTEM_NAME} MATCHES "Windows")
          add_custom_command(
                      TARGET amcl_mpin_${tc}
                      POST_BUILD
                      COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_mpin_${tc}.dll" "${PROJECT_BINARY_DIR}/wrappers/go/")
        endif(${CMAKE_SYSTEM_NAME} MATCHES "Windows")

        if(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
          add_custom_command(
                      TARGET amcl_mpin_${tc}
                      POST_BUILD
                      COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_mpin_${tc}.dylib" "${PROJECT_BINARY_DIR}/wrappers/go/")
        endif(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
      endif (BUILD_GO AND BUILD_SHARED_LIBS)

      # Static build
      if (BUILD_GO AND NOT BUILD_SHARED_LIBS)
        if(${CMAKE_SYSTEM_NAME} MATCHES "Linux")
          add_custom_command(
                      TARGET amcl_mpin_${tc}
                      POST_BUILD
                      COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_mpin_${tc}.a" "${PROJECT_BINARY_DIR}/wrappers/go/")
        endif(${CMAKE_SYSTEM_NAME} MATCHES "Linux")
      endif (BUILD_GO AND NOT BUILD_SHARED_LIBS)

      # Copy built libs to wrappers/python to allow testing without
      # having to install the libs.
      if (BUILD_PYTHON AND BUILD_SHARED_LIBS)
        if(${CMAKE_SYSTEM_NAME} MATCHES "Linux")
          add_custom_command(
                      TARGET amcl_mpin_${tc}
                      POST_BUILD
                      COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_mpin_${tc}.so" "${PROJECT_BINARY_DIR}/wrappers/python/")
        endif(${CMAKE_SYSTEM_NAME} MATCHES "Linux")

        if(${CMAKE_SYSTEM_NAME} MATCHES "Windows")
          add_custom_command(
                      TARGET amcl_mpin_${tc}
                      POST_BUILD
                      COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_mpin_${tc}.dll" "${PROJECT_BINARY_DIR}/wrappers/python/")
        endif(${CMAKE_SYSTEM_NAME} MATCHES "Windows")

        if(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
          add_custom_command(
                      TARGET mcl_mpin_${tc}
                      POST_BUILD
                      COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_mpin_${tc}.dylib" "${PROJECT_BINARY_DIR}/wrappers/python/")
        endif(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
      endif (BUILD_PYTHON AND BUILD_SHARED_LIBS)

    endif(BUILD_MPIN)

    if(BUILD_WCC)
      file(READ "wcc_ZZZ.c" tempc)
      string(REGEX REPLACE "XXX" ${bd} tempc "${tempc}")
      string(REGEX REPLACE "YYY" ${tf} tempc "${tempc}")
      string(REGEX REPLACE "ZZZ" ${tc} tempc "${tempc}")
      file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/wcc_${tc}.c" "${tempc}")
      set(SOURCES_WCC_${tc} ${SOURCES_WCC_${tc}}
        ${CMAKE_CURRENT_BINARY_DIR}/wcc_${tc}.c)

      message(STATUS "Build libamcl_wcc_${tc}")
      add_library(amcl_wcc_${tc} ${LIB_TYPE} ${SOURCES_WCC_${tc}})
      target_link_libraries (amcl_wcc_${tc} amcl_pairing_${tc} amcl_curve_${tc} amcl_core)

      set(INSTALL_LIBS ${INSTALL_LIBS} amcl_wcc_${tc})

      if(${CMAKE_SYSTEM_NAME} MATCHES "Windows")
          add_custom_command(
                      TARGET amcl_wcc_${tc}
                      POST_BUILD
                      COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_wcc_${tc}.dll" "${PROJECT_BINARY_DIR}/test/")
      endif(${CMAKE_SYSTEM_NAME} MATCHES "Windows")

      # Copy built libs to wrappers/python to allow testing without
      # having to install the libs.
      if (BUILD_PYTHON AND BUILD_SHARED_LIBS)
        if(${CMAKE_SYSTEM_NAME} MATCHES "Linux")
          add_custom_command(
                      TARGET amcl_wcc_${tc}
                      POST_BUILD
                      COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_wcc_${tc}.so" "${PROJECT_BINARY_DIR}/wrappers/python/")
        endif(${CMAKE_SYSTEM_NAME} MATCHES "Linux")

        if(${CMAKE_SYSTEM_NAME} MATCHES "Windows")
          add_custom_command(
                      TARGET amcl_wcc_${tc}
                      POST_BUILD
                      COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_wcc_${tc}.dll" "${PROJECT_BINARY_DIR}/wrappers/python/")
        endif(${CMAKE_SYSTEM_NAME} MATCHES "Windows")

        if(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
          add_custom_command(
                      TARGET amcl_wcc_${tc}
                      POST_BUILD
                      COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_wcc_${tc}.dylib" "${PROJECT_BINARY_DIR}/wrappers/python/")
        endif(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
      endif (BUILD_PYTHON AND BUILD_SHARED_LIBS)

    endif(BUILD_WCC)

  endif(NOT(${pf} STREQUAL "NOT"))

  install (TARGETS ${INSTALL_LIBS} DESTINATION lib PERMISSIONS
      OWNER_WRITE OWNER_READ OWNER_EXECUTE
      GROUP_READ GROUP_EXECUTE
      WORLD_READ WORLD_EXECUTE)

  if(${CMAKE_SYSTEM_NAME} MATCHES "Windows")
      add_custom_command(
                  TARGET amcl_curve_${tc}
                  POST_BUILD
                  COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_curve_${tc}.dll" "${PROJECT_BINARY_DIR}/test/")
  endif(${CMAKE_SYSTEM_NAME} MATCHES "Windows")

  # Copy built libs to wrappers/go to allow testing without
  # having to install the libs.
  if (BUILD_GO AND BUILD_SHARED_LIBS)
    if(${CMAKE_SYSTEM_NAME} MATCHES "Linux")
      add_custom_command(
                  TARGET amcl_curve_${tc}
                  POST_BUILD
                  COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_curve_${tc}.so" "${PROJECT_BINARY_DIR}/wrappers/go/")
    endif(${CMAKE_SYSTEM_NAME} MATCHES "Linux")

    if(${CMAKE_SYSTEM_NAME} MATCHES "Windows")
      add_custom_command(
                  TARGET amcl_curve_${tc}
                  POST_BUILD
                  COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_curve_${tc}.dll" "${PROJECT_BINARY_DIR}/wrappers/go/")
    endif(${CMAKE_SYSTEM_NAME} MATCHES "Windows")

    if(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
      add_custom_command(
                  TARGET amcl_curve_${tc}
                  POST_BUILD
                  COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_curve_${tc}.dylib" "${PROJECT_BINARY_DIR}/wrappers/go/")
    endif(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
  endif (BUILD_GO AND BUILD_SHARED_LIBS)

  # Static build
  if (BUILD_GO AND NOT BUILD_SHARED_LIBS)
    if(${CMAKE_SYSTEM_NAME} MATCHES "Linux")
      add_custom_command(
                  TARGET amcl_curve_${tc}
                  POST_BUILD
                  COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_curve_${tc}.a" "${PROJECT_BINARY_DIR}/wrappers/go/")
    endif(${CMAKE_SYSTEM_NAME} MATCHES "Linux")
  endif (BUILD_GO AND NOT BUILD_SHARED_LIBS)

  # Copy built libs to wrappers/python to allow testing without
  # having to install the libs.
  if (BUILD_PYTHON AND BUILD_SHARED_LIBS)
    if(${CMAKE_SYSTEM_NAME} MATCHES "Linux")
      add_custom_command(
                  TARGET amcl_curve_${tc}
                  POST_BUILD
                  COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_curve_${tc}.so" "${PROJECT_BINARY_DIR}/wrappers/python/")
    endif(${CMAKE_SYSTEM_NAME} MATCHES "Linux")

    if(${CMAKE_SYSTEM_NAME} MATCHES "Windows")
      add_custom_command(
                  TARGET amcl_curve_${tc}
                  POST_BUILD
                  COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_curve_${tc}.dll" "${PROJECT_BINARY_DIR}/wrappers/python/")
    endif(${CMAKE_SYSTEM_NAME} MATCHES "Windows")

    if(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
      add_custom_command(
                  TARGET amcl_curve_${tc}
                  POST_BUILD
                  COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_curve_${tc}.dylib" "${PROJECT_BINARY_DIR}/wrappers/python/")
    endif(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
  endif (BUILD_PYTHON AND BUILD_SHARED_LIBS)

endfunction(CURVE_SRC_SET tb tf tc nb base nbt m8 mt ct pf)

function(RSA_SRC_SET tb tff nb base ml)
  
  set(bd "${tb}_${base}")

  file(READ "big_XXX.c" tempc)
  string(REGEX REPLACE "XXX" ${bd} tempc "${tempc}")
  file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/big_${bd}.c" "${tempc}")
  set(SOURCES_RSA_${tff} ${SOURCES_RSA_${tff}}
      ${CMAKE_CURRENT_BINARY_DIR}/big_${bd}.c)

  file(READ "ff_WWW.c" tempc)
  string(REGEX REPLACE "XXX" ${bd} tempc "${tempc}")
  string(REGEX REPLACE "WWW" ${tff} tempc "${tempc}")
  file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/ff_${tff}.c" "${tempc}")
  set(SOURCES_RSA_${tff} ${SOURCES_RSA_${tff}}
      ${CMAKE_CURRENT_BINARY_DIR}/ff_${tff}.c)

  file(READ "rsa_WWW.c" tempc)
  string(REGEX REPLACE "XXX" ${bd} tempc "${tempc}")
  string(REGEX REPLACE "WWW" ${tff} tempc "${tempc}")
  file(WRITE "${CMAKE_CURRENT_BINARY_DIR}/rsa_${tff}.c" "${tempc}")
  set(SOURCES_RSA_${tff} ${SOURCES_RSA_${tff}}
      ${CMAKE_CURRENT_BINARY_DIR}/rsa_${tff}.c)

  message(STATUS "Build libamcl_rsa_${tff}")
  add_library(amcl_rsa_${tff} ${LIB_TYPE} ${SOURCES_RSA_${tff}})
  target_link_libraries (amcl_rsa_${tff} amcl_core)

  set(INSTALL_LIBS ${INSTALL_LIBS} amcl_rsa_${tff})

  install (TARGETS ${INSTALL_LIBS} DESTINATION lib PERMISSIONS
        OWNER_WRITE OWNER_READ OWNER_EXECUTE
        GROUP_READ GROUP_EXECUTE
        WORLD_READ WORLD_EXECUTE)

  if(${CMAKE_SYSTEM_NAME} MATCHES "Windows")
      message(STATUS "Copy libs to test for testing")
      add_custom_command(
                  TARGET amcl_rsa_${tff}
                  POST_BUILD
                  COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_rsa_${tff}.dll" "${PROJECT_BINARY_DIR}/test/")
  endif(${CMAKE_SYSTEM_NAME} MATCHES "Windows")

  # Copy built libs to wrappers/go to allow testing without
  # having to install the libs.
  if (BUILD_GO AND BUILD_SHARED_LIBS)
    if(${CMAKE_SYSTEM_NAME} MATCHES "Linux")
      add_custom_command(
                  TARGET amcl_rsa_${tff}
                  POST_BUILD
                  COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_rsa_${tff}.so" "${PROJECT_BINARY_DIR}/wrappers/go/")
    endif(${CMAKE_SYSTEM_NAME} MATCHES "Linux")

    if(${CMAKE_SYSTEM_NAME} MATCHES "Windows")
      add_custom_command(
                  TARGET amcl_rsa_${tff}
                  POST_BUILD
                  COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_rsa_${tff}.dll" "${PROJECT_BINARY_DIR}/wrappers/go/")
    endif(${CMAKE_SYSTEM_NAME} MATCHES "Windows")

    if(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
      add_custom_command(
                  TARGET amcl_rsa_${tff}
                  POST_BUILD
                  COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_rsa_${tff}.dylib" "${PROJECT_BINARY_DIR}/wrappers/go/")
    endif(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
  endif (BUILD_GO AND BUILD_SHARED_LIBS)

  # Static build
  if (BUILD_GO AND NOT BUILD_SHARED_LIBS)
    if(${CMAKE_SYSTEM_NAME} MATCHES "Linux")
      add_custom_command(
                  TARGET amcl_rsa_${tff}
                  POST_BUILD
                  COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_rsa_${tff}.a" "${PROJECT_BINARY_DIR}/wrappers/go/")
    endif(${CMAKE_SYSTEM_NAME} MATCHES "Linux")
  endif (BUILD_GO AND NOT BUILD_SHARED_LIBS)

  # Copy built libs to wrappers/python to allow testing without
  # having to install the libs.
  if (BUILD_PYTHON AND BUILD_SHARED_LIBS)
    if(${CMAKE_SYSTEM_NAME} MATCHES "Linux")
      add_custom_command(
                  TARGET amcl_rsa_${tff}
                  POST_BUILD
                  COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_rsa_${tff}.so" "${PROJECT_BINARY_DIR}/wrappers/python/")
    endif(${CMAKE_SYSTEM_NAME} MATCHES "Linux")

    if(${CMAKE_SYSTEM_NAME} MATCHES "Windows")
      add_custom_command(
                  TARGET amcl_rsa_${tff}
                  POST_BUILD
                  COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_rsa_${tff}.dll" "${PROJECT_BINARY_DIR}/wrappers/python/")
    endif(${CMAKE_SYSTEM_NAME} MATCHES "Windows")

    if(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
      add_custom_command(
                  TARGET amcl_rsa_${tff}
                  POST_BUILD
                  COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_rsa_${tff}.dylib" "${PROJECT_BINARY_DIR}/wrappers/python/")
    endif(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
  endif (BUILD_PYTHON AND BUILD_SHARED_LIBS)


endfunction(RSA_SRC_SET tb tff nb base ml)

set(SOURCES_AMCL_CORE
hash.c
rand.c
randapi.c
aes.c
gcm.c
oct.c
utils.c
version.c)

# Copy common header files to binary include directory
file(COPY ${SOURCES_AMCL_CORE} DESTINATION ${CMAKE_CURRENT_BINARY_DIR})

set(DEST_SOURCES_AMCL_CORE
${CMAKE_CURRENT_BINARY_DIR}/hash.c
${CMAKE_CURRENT_BINARY_DIR}/rand.c
${CMAKE_CURRENT_BINARY_DIR}/randapi.c
${CMAKE_CURRENT_BINARY_DIR}/aes.c
${CMAKE_CURRENT_BINARY_DIR}/gcm.c
${CMAKE_CURRENT_BINARY_DIR}/oct.c
${CMAKE_CURRENT_BINARY_DIR}/utils.c
${CMAKE_CURRENT_BINARY_DIR}/version.c
)

if(BUILD_MPIN OR BUILD_WCC)
  file(COPY "pbc_support.c" DESTINATION ${CMAKE_CURRENT_BINARY_DIR})
  set(DEST_SOURCES_AMCL_CORE ${DEST_SOURCES_AMCL_CORE}
    ${CMAKE_CURRENT_BINARY_DIR}/pbc_support.c)
endif(BUILD_MPIN OR BUILD_WCC)

file(COPY "rsa_support.c" DESTINATION ${CMAKE_CURRENT_BINARY_DIR})
set(DEST_SOURCES_AMCL_CORE ${DEST_SOURCES_AMCL_CORE}
  ${CMAKE_CURRENT_BINARY_DIR}/rsa_support.c)

file(COPY "ecdh_support.c" DESTINATION ${CMAKE_CURRENT_BINARY_DIR})
set(DEST_SOURCES_AMCL_CORE ${DEST_SOURCES_AMCL_CORE}
  ${CMAKE_CURRENT_BINARY_DIR}/ecdh_support.c)

message(STATUS "Build libamcl_core")
add_library(amcl_core ${LIB_TYPE} ${DEST_SOURCES_AMCL_CORE})
target_link_libraries (amcl_core)

install (TARGETS amcl_core DESTINATION lib PERMISSIONS
        OWNER_WRITE OWNER_READ OWNER_EXECUTE
        GROUP_READ GROUP_EXECUTE
        WORLD_READ WORLD_EXECUTE)

string(REGEX REPLACE "," ";" AMCL_CURVE ${AMCL_CURVE})
string(REGEX REPLACE "," ";" AMCL_RSA ${AMCL_RSA})

foreach(CURVE ${AMCL_CURVE})
  if(AMCL_CHUNK STREQUAL "64")
    if(CURVE STREQUAL "ED25519")
        CURVE_SRC_SET("256" "25519" "ED25519" "32" "56" "255" "5" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "C25519")
        CURVE_SRC_SET("256" "25519" "C25519" "32" "56" "255" "5" "PSEUDO_MERSENNE" "MONTGOMERY" "NOT")
    elseif(CURVE STREQUAL "NIST256")
        CURVE_SRC_SET("256" "NIST256" "NIST256" "32" "56" "256" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "BRAINPOOL")
        CURVE_SRC_SET("256" "BRAINPOOL" "BRAINPOOL" "32" "56" "256" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "ANSSI")
        CURVE_SRC_SET("256" "ANSSI" "ANSSI" "32" "56" "256" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "HIFIVE")
        CURVE_SRC_SET("336" "HIFIVE" "HIFIVE" "42" "60" "336" "5" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "GOLDILOCKS")
        CURVE_SRC_SET("448" "GOLDILOCKS" "GOLDILOCKS" "56" "58" "448" "7" "GENERALISED_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NIST384")
        CURVE_SRC_SET("384" "NIST384" "NIST384" "48" "56" "384" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "C41417")
        CURVE_SRC_SET("416" "C41417" "C41417" "52" "60" "414" "7" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NIST521")
        CURVE_SRC_SET("528" "NIST521" "NIST521" "66" "60" "521" "7" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS256W")
        CURVE_SRC_SET("256" "256PMW" "NUMS256W" "32" "56" "256" "3" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS256E")
        CURVE_SRC_SET("256" "256PME" "NUMS256E" "32" "56" "256" "3" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NUMS384W")
        CURVE_SRC_SET("384" "384PM" "NUMS384W" "48" "56" "384" "3" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS384E")
        CURVE_SRC_SET("384" "384PM" "NUMS384E" "48" "56" "384" "3" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NUMS512W")
        CURVE_SRC_SET("512" "512PM" "NUMS512W" "64" "56" "512" "7" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS512E")
        CURVE_SRC_SET("512" "512PM" "NUMS512E" "64" "56" "512" "7" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "BN254")
        CURVE_SRC_SET("256" "BN254" "BN254" "32" "56" "254" "3" "NOT_SPECIAL" "WEIERSTRASS" "BN")
    elseif(CURVE STREQUAL "BN254CX")
        CURVE_SRC_SET("256" "BN254CX" "BN254CX" "32" "56" "254" "3" "NOT_SPECIAL" "WEIERSTRASS" "BN")
    elseif(CURVE STREQUAL "BLS383")
        CURVE_SRC_SET("384" "BLS383" "BLS383" "48" "56" "383" "3" "NOT_SPECIAL" "WEIERSTRASS" "BLS")
    else()
      message(FATAL_ERROR "Invalid curve input")
    endif(CURVE STREQUAL "ED25519")
  elseif(AMCL_CHUNK STREQUAL "32")
    if(CURVE STREQUAL "ED25519")
        CURVE_SRC_SET("256" "25519" "ED25519" "32" "29" "255" "5" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "C25519")
        CURVE_SRC_SET("256" "25519" "C25519" "32" "29" "255" "5" "PSEUDO_MERSENNE" "MONTGOMERY" "NOT")
    elseif(CURVE STREQUAL "NIST256")
        CURVE_SRC_SET("256" "NIST256" "NIST256" "32" "28" "256" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "BRAINPOOL")
        CURVE_SRC_SET("256" "BRAINPOOL" "BRAINPOOL" "32" "28" "256" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "ANSSI")
        CURVE_SRC_SET("256" "ANSSI" "ANSSI" "32" "28" "256" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "HIFIVE")
        CURVE_SRC_SET("336" "HIFIVE" "HIFIVE" "42" "29" "336" "5" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "GOLDILOCKS")
        CURVE_SRC_SET("448" "GOLDILOCKS" "GOLDILOCKS" "56" "29" "448" "7" "GENERALISED_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NIST384")
        CURVE_SRC_SET("384" "NIST384" "NIST384" "48" "29" "384" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "C41417")
        CURVE_SRC_SET("416" "C41417" "C41417" "52" "29" "414" "7" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NIST521")
        CURVE_SRC_SET("528" "NIST521" "NIST521" "66" "28" "521" "7" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS256W")
        CURVE_SRC_SET("256" "256PMW" "NUMS256W" "32" "28" "256" "3" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS256E")
        CURVE_SRC_SET("256" "256PME" "NUMS256E" "32" "29" "256" "3" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NUMS384W")
        CURVE_SRC_SET("384" "384PM" "NUMS384W" "48" "29" "384" "3" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS384E")
        CURVE_SRC_SET("384" "384PM" "NUMS384E" "48" "29" "384" "3" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NUMS512W")
        CURVE_SRC_SET("512" "512PM" "NUMS512W" "64" "29" "512" "7" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS512E")
        CURVE_SRC_SET("512" "512PM" "NUMS512E" "64" "29" "512" "7" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "BN254")
        CURVE_SRC_SET("256" "BN254" "BN254" "32" "28" "254" "3" "NOT_SPECIAL" "WEIERSTRASS" "BN")
    elseif(CURVE STREQUAL "BN254CX")
        CURVE_SRC_SET("256" "BN254CX" "BN254CX" "32" "28" "254" "3" "NOT_SPECIAL" "WEIERSTRASS" "BN")
    elseif(CURVE STREQUAL "BLS383")
        CURVE_SRC_SET("384" "BLS383" "BLS383" "48" "29" "383" "3" "NOT_SPECIAL" "WEIERSTRASS" "BLS")
    else()
      message(FATAL_ERROR "Invalid curve input")
    endif(CURVE STREQUAL "ED25519")
  elseif(AMCL_CHUNK STREQUAL "16")
    if(CURVE STREQUAL "ED25519")
        CURVE_SRC_SET("256" "25519" "ED25519" "32" "13" "255" "5" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "C25519")
        CURVE_SRC_SET("256" "25519" "C25519" "32" "13" "255" "5" "PSEUDO_MERSENNE" "MONTGOMERY" "NOT")
    elseif(CURVE STREQUAL "NUMS256E")
        CURVE_SRC_SET("256" "256PME" "NUMS256E" "32" "13" "256" "3" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "BN254")
        CURVE_SRC_SET("256" "BN254" "BN254" "32" "13" "254" "3" "NOT_SPECIAL" "WEIERSTRASS" "BN")
    elseif(CURVE STREQUAL "BN254CX")
        CURVE_SRC_SET("256" "BN254CX" "BN254CX" "32" "13" "254" "3" "NOT_SPECIAL" "WEIERSTRASS" "BN")
    else()
      message(FATAL_ERROR "Invalid curve input")
    endif(CURVE STREQUAL "ED25519")
  endif(AMCL_CHUNK STREQUAL "64")
endforeach(CURVE)

foreach(RSA ${AMCL_RSA})
  if(AMCL_CHUNK STREQUAL "64")
    if(RSA STREQUAL "2048")
      RSA_SRC_SET("1024" "2048" "128" "58" "2")
    elseif(RSA STREQUAL "3072")
      RSA_SRC_SET("384" "3072" "48" "56" "8")
    elseif(RSA STREQUAL "4096")
      RSA_SRC_SET("512" "4096" "64" "60" "8")
    else()
      message(FATAL_ERROR "Invalid RSA input")
    endif(RSA STREQUAL "2048")
  elseif(AMCL_CHUNK STREQUAL "32")
    if(RSA STREQUAL "2048")
      RSA_SRC_SET("1024" "2048" "128" "28" "2")
    elseif(RSA STREQUAL "3072")
      RSA_SRC_SET("384" "3072" "48" "28" "8")
    elseif(RSA STREQUAL "4096")
      RSA_SRC_SET("512" "4096" "64" "29" "8")
    else()
        message(FATAL_ERROR "Invalid RSA input")
    endif(RSA STREQUAL "2048")
  elseif(AMCL_CHUNK STREQUAL "16")
    if(RSA STREQUAL "2048")
      RSA_SRC_SET("256" "2048" "32" "13" "8")
    else()
      message(FATAL_ERROR "Invalid RSA input")
    endif(RSA STREQUAL "2048")
  else()
    message(FATAL_ERROR "Invalid CHUNK size")
  endif(AMCL_CHUNK STREQUAL "64")
endforeach(RSA)

if(BUILD_X509)
  file(COPY "x509.c" DESTINATION ${CMAKE_CURRENT_BINARY_DIR})
  message(STATUS "Build libamcl_x509")
  add_library(amcl_x509 ${LIB_TYPE} ${CMAKE_CURRENT_BINARY_DIR}/x509.c)
  set(link_x509 ${link_x509} "amcl_core")
  foreach(CURVE ${AMCL_CURVE})
    set(link_x509 ${link_x509} "amcl_curve_${CURVE}")
  endforeach(CURVE)
  foreach(RSA ${AMCL_RSA})
    set(link_x509 ${link_x509} "amcl_rsa_${RSA}")
  endforeach(RSA)
  target_link_libraries (amcl_x509 ${link_x509})

  install (TARGETS amcl_x509 DESTINATION lib PERMISSIONS
        OWNER_WRITE OWNER_READ OWNER_EXECUTE
        GROUP_READ GROUP_EXECUTE
        WORLD_READ WORLD_EXECUTE)
endif(BUILD_X509)

# Install libs
install(DIRECTORY DESTINATION lib DIRECTORY_PERMISSIONS
        OWNER_WRITE OWNER_READ OWNER_EXECUTE
        GROUP_WRITE GROUP_READ
        WORLD_WRITE WORLD_READ WORLD_EXECUTE)

if(${CMAKE_SYSTEM_NAME} MATCHES "Windows")
    add_custom_command(
                TARGET amcl_core
                POST_BUILD
                COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_core.dll" "${PROJECT_BINARY_DIR}/test/")
endif(${CMAKE_SYSTEM_NAME} MATCHES "Windows")

# Copy built libs to wrappers/go to allow testing without
# having to install the libs.
if (BUILD_GO AND BUILD_SHARED_LIBS)
  if(${CMAKE_SYSTEM_NAME} MATCHES "Linux")
    add_custom_command(
                TARGET amcl_core
                POST_BUILD
                COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_core.so" "${PROJECT_BINARY_DIR}/wrappers/go/")
  endif(${CMAKE_SYSTEM_NAME} MATCHES "Linux")

  if(${CMAKE_SYSTEM_NAME} MATCHES "Windows")
    add_custom_command(
                TARGET amcl_core
                POST_BUILD
                COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_core.dll" "${PROJECT_BINARY_DIR}/wrappers/go/")
  endif(${CMAKE_SYSTEM_NAME} MATCHES "Windows")

  if(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
    add_custom_command(
                TARGET amcl_core
                POST_BUILD
                COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_core.dylib" "${PROJECT_BINARY_DIR}/wrappers/go/")
  endif(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
endif (BUILD_GO AND BUILD_SHARED_LIBS)

# Static build
if (BUILD_GO AND NOT BUILD_SHARED_LIBS)
  if(${CMAKE_SYSTEM_NAME} MATCHES "Linux")
    add_custom_command(
                TARGET amcl_core
                POST_BUILD
                COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_core.a" "${PROJECT_BINARY_DIR}/wrappers/go/")
  endif(${CMAKE_SYSTEM_NAME} MATCHES "Linux")
endif (BUILD_GO AND NOT BUILD_SHARED_LIBS)

# Copy built libs to wrappers/python to allow testing without
# having to install the libs.
if (BUILD_PYTHON AND BUILD_SHARED_LIBS)
  if(${CMAKE_SYSTEM_NAME} MATCHES "Linux")
    add_custom_command(
                TARGET amcl_core
                POST_BUILD
                COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_core.so" "${PROJECT_BINARY_DIR}/wrappers/python/")
  endif(${CMAKE_SYSTEM_NAME} MATCHES "Linux")

  if(${CMAKE_SYSTEM_NAME} MATCHES "Windows")
    add_custom_command(
                TARGET amcl_core
                POST_BUILD
                COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_core.dll" "${PROJECT_BINARY_DIR}/wrappers/python/")
  endif(${CMAKE_SYSTEM_NAME} MATCHES "Windows")

  if(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
    add_custom_command(
                TARGET amcl_core
                POST_BUILD
                COMMAND ${CMAKE_COMMAND} -E copy "${PROJECT_BINARY_DIR}/lib/libamcl_core.dylib" "${PROJECT_BINARY_DIR}/wrappers/python/")
  endif(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
endif (BUILD_PYTHON AND BUILD_SHARED_LIBS)










