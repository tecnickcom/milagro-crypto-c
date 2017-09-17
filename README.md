# AMCL - *Apache Milagro Crypto Library*

[![Master Branch](https://img.shields.io/badge/-master:-gray.svg)](https://github.com/miracl/milagro-crypto-c/tree/master)
[![Master Build Status](https://secure.travis-ci.org/miracl/milagro-crypto-c.png?branch=master)](https://travis-ci.org/miracl/milagro-crypto-c?branch=master)
[![Master Coverage Status](https://coveralls.io/repos/miracl/milagro-crypto-c/badge.svg?branch=master&service=github)](https://coveralls.io/github/miracl/milagro-crypto-c?branch=master)

[![Develop Branch](https://img.shields.io/badge/-develop:-gray.svg)](https://github.com/miracl/milagro-crypto-c/tree/develop)
[![Develop Build Status](https://secure.travis-ci.org/miracl/milagro-crypto-c.png?branch=develop)](https://travis-ci.org/miracl/milagro-crypto-c?branch=develop)
[![Develop Coverage Status](https://coveralls.io/repos/miracl/milagro-crypto-c/badge.svg?branch=develop&service=github)](https://coveralls.io/github/miracl/milagro-crypto-c?branch=develop)


* **category**:    Library
* **copyright**:   2016 The Apache Software Foundation
* **license**:     ASL 2.0 ([Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0))
* **link**:        https://github.com/miracl/milagro-crypto-c
* **introduction**: [AMCL.pdf](doc/AMCL.pdf)


## Description

*AMCL - Apache Milagro Crypto Library*

AMCL is a standards compliant C cryptographic library with no external dependencies, specifically designed to support the Internet of Things (IoT).

For a detailed explanation about this library please read: [doc/AMCL.pdf](doc/AMCL.pdf)

AMCL is provided in *C* language but includes a *[Python](https://www.python.org)* and *[GO](https://golang.org)* wrapper.

NOTE: This product includes software developed at *[The Apache Software Foundation](http://www.apache.org/)*.

## Software Dependencies

In order to build this library, the following packages are required:

* [CMake](https://cmake.org/) is required to build the source code.
* [CFFI](https://cffi.readthedocs.org/en/release-0.8/), the C Foreign Function Interface for Python is required in order to execute tests.
* [Doxygen](http://doxygen.org) is required to build the source code documentation.
* [GO](https://golang.org/dl/) and [Python](https://www.python.org/) languages are required to build the language wrappers.
* Please refer to the *resources/DockerDev/Dockerfile* file for any additional dependency.


The above packages can be installed in different ways, depending on the Operating System used:

* **Debian/Ubuntu Linux**


    sudo apt-get install -y git cmake build-essential python python-dev python-pip libffi-dev doxygen doxygen-latex parallel
    sudo pip install cffi


* **RedHat/CentOS/Fedora Linux**


    sudo yum groupinstall "Development Tools" "Development Libraries"
    sudo yum install -y git cmake python libpython-devel python-pip libffi-devel doxygen doxygen-latex parallel
    sudo pip install cffi


* **MacOs**


    ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
    brew install cmake
    brew install pkg-config libffi
    sudo pip install cffi
    brew install doxygen
    brew install parallel


* **Windows**
    * Minimalist GNU for Windows [MinGW](http://www.mingw.org) provides the tool set used to build the library and should be installed
    * When the MinGW installer starts select the **mingw32-base** and **mingw32-gcc-g++** components
    * From the menu select *"Installation"* &rarr; *"Apply Changes"*, then click *"Apply"*
    * Finally add *C:\MinGW\bin* to the PATH variable
    * pip install cffi
    * install CMake following the instructions on http://www.cmake.org
    * install Doxygen following the instructions on http://www.doxygen.org


## Build Instructions

#### Linux and Mac

##### Quick start

A Makefile is present at the project root that reads the options defined in
config.mk. Change these options and then type the following to build and test
the library.

    make

If you are using GO then you must set the GOPATH first. Note that there can 
only be one entry in the GOPATH 

    export GOPATH=${HOME}/src
    make

If you have docker installed then type ```make dbuild``` to build and test
the library in a docker container.

##### Multiple curves and RSA security levels

The default build (see config.mk) uses multiple curves and RSA security 
levels. There is an example called testall.c in the examples directory that 
shows how to write a program to use the different curves etc in a single 
program. To build and run the example use this script;

    buildMulti.sh

##### Manual build

NOTE: The default build is for 64 bit machines

    git clone https://github.com/miracl/milagro-crypto-c
    cd milagro-crypto-c
    mkdir -p target/build
    cd target/build
    cmake -D CMAKE_INSTALL_PREFIX=/opt/amcl ../..
    export GOPATH=${HOME}/src
    export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:./
    make
    make test
    make doc
    sudo make install

On Debian/Ubuntu machine instead of executing the *"sudo make install"* command it is possible to execute *"sudo checkinstall"* to build and install a DEB package.

Now you can set the path to where libs and python package are installed:

    export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:./:/opt/amcl/lib
    export PYTHONPATH=/usr/lib/python2.7/dist-packages

NOTE: The build can be configured by setting flags on the command line, for example:

    cmake -DAMCL_CHUNK=64 ../..
    cmake -D CMAKE_INSTALL_PREFIX=/opt/amcl -D AMCL_CHUNK=64 -D BUILD_WCC=on ../..

It is possible also to build the library supporting more than one elliptic curve and more  than one RSA security level, for example

	cmake -DAMCL_CURVE=BN254CX,NIST254 -DAMCL_RSA=2048,3072 ../..

To list other available CMake options, use:

    cmake -LH

##### Uninstall software

    sudo make uninstall

##### Building an installer 

After having built the libraries you can build a binary installer and a source distribution by running this command

    make package


#### Windows

Start a command prompt as an administrator

    git clone https://github.com/miracl/milagro-crypto-c
    cd milagro-crypto-c
    mkdir target\build
    cd target\build
    cmake -G "MinGW Makefiles" ..\..
    mingw32-make
    mingw32-make test
    mingw32-make doc
    mingw32-make install

Post install append the PATH system variable to point to the install ./lib:

*My Computer -> Properties -> Advanced > Environment Variables*

The build can be configured using by setting flags on the command line i.e.

    cmake -G "MinGW Makefiles" -D BUILD_PYTHON=on ..

##### Uninstall software

    mingw32-make uninstall

##### Building an installer

After having built the libraries you can build a Windows installer using this command

    sudo mingw32-make package

In order for this to work NSSI has to have been installed


## Contributions

This project includes a Makefile that allows you to test and build the project in a Linux-compatible system with simple commands.  
All the artifacts and reports produced using this Makefile are stored in the *target* folder.  

All the packages listed in the *resources/DockerDev/Dockerfile* file are required in order to build and test all the library options in the current environment. Alternatively, everything can be built inside a [Docker](https://www.docker.com) container using the command "MAKETARGET=buildall make dbuild".

To see all available options:
```
make help
```

To build the project inside a Docker container (requires Docker):
```
MAKETARGET=buildall make dbuild
```

To build a particular set of predefined makefile options inside a Docker container:
```
MAKETARGET='build TYPE=LINUX_64BIT_NIST256_RSA2048' make dbuild
```
The list of pre-defined options can be listed by typing ```make help```


The base Docker building environment is defined in the following Dockerfile:
```
resources/DockerDev/Dockerfile
```

To execute all the test builds and generate reports in the current environment:
```
make qa
```

To format the code (please use this command before submitting any pull request):
```
make format
```

## Useful Docker commands

To manually create the container you can execute:
```
docker build --tag="miracl/amcldev" .
```

To log into the newly created container:
```
docker run -t -i miracl/amcldev /bin/bash
```

To get the container ID:
```
CONTAINER_ID=`docker ps -a | grep miracl/amcldev | cut -c1-12`
```

To delete the newly created docker container:
```
docker rm -f $CONTAINER_ID
```

To delete the docker image:
```
docker rmi -f miracl/amcldev
```

To delete all containers
```
docker rm $(docker ps -a -q)
```

To delete all images
```
docker rmi $(docker images -q)
```
