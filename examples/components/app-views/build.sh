#!/bin/bash

GOROOT=$(go env GOROOT);
ROOT=$(pwd);

env GOOS=js GOARCH=wasm go build -o "${ROOT}/public/main.wasm" main.go;

if [[ "$?" == "0" ]]; then

	if [[ -d "${ROOT}/public/design" ]]; then
		rm -rf "${ROOT}/public/design";
	fi;

	# Import Go WASM Adapter
	cp "${GOROOT}/lib/wasm/wasm_exec.js" "${ROOT}/public/wasm_exec.js";

	# Import Gooey Theme
	cp -R "${ROOT}/../../../design" "${ROOT}/public/design";

	go run "${ROOT}/serve.go";

fi;

