#!/bin/bash
# set -e

ROOT_DIR="$(pwd)";

run_tests() {

	local package="$1";

	cd "${ROOT_DIR}/${package}";

	env GOOS=js GOARCH=wasm go test "./..." 2>&1 | grep -E ':[0-9]+:[0-9]+:|syntax error|expected |undefined:|cannot use |missing return|imported and not used';

}

# Missing doc.go files
for base in bindings components; do

	echo "";
	echo "Checking ${base} ...";
	echo "";

	run_tests "${base}";

	for dir in "${ROOT_DIR}/${base}"/*/; do

		[ -d "$dir" ] || continue

		if [ ! -f "$dir/doc.go" ]; then
			echo "[!] Missing ${dir#"${ROOT_DIR}/"}doc.go";
		fi

	done

done
