#!/bin/bash -eu

source /build-common.sh

COMPILE_IN_DIRECTORY="cmd/emergency-details-printout"
BINARY_NAME="emergency-details-printout"

# TODO: once deployerspec is done, we can stop overriding this from base image
function packageLambdaFunction {
	if [ ! -z ${FASTBUILD+x} ]; then return; fi

	cd rel/
	cp "${BINARY_NAME}_linux-amd64" "${BINARY_NAME}"
	rm -f lambdafunc.zip
	zip -j lambdafunc.zip "${BINARY_NAME}"
	rm "${BINARY_NAME}"
}

standardBuildProcess

packageLambdaFunction
