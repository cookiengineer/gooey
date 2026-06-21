#!/bin/bash
# set -e

ROOT_DIR="$(pwd)";

run_tests() {

	local package="$1";

	cd "${ROOT_DIR}/${package}";

	env GOOS=js GOARCH=wasm go test "./..." 2>&1 | grep -E ':[0-9]+:[0-9]+:|syntax error|expected |undefined:|cannot use |missing return|imported and not used';

}

check_build_tag() {

	local base="$1";

	find "${ROOT_DIR}/${base}" -type f -name "*.go" | while read -r file; do

		# Skip generated files if desired:
		[[ "$file" == *.gen.go ]] && continue

		if ! grep -q '^//go:build wasm$' "$file"; then
			echo "[!] Missing //go:build wasm in ${file#"${ROOT_DIR}/"}";
		fi

	done

}

check_doc() {

	local base="$1";

	for dir in "${ROOT_DIR}/${base}"/*/; do

		[ -d "$dir" ] || continue;

		if [ ! -f "$dir/doc.go" ]; then
			echo "[!] Missing ${dir#"${ROOT_DIR}/"}doc.go";
		fi;

	done;

}

check_examples() {

	local base="$1";

	shopt -s nullglob

	find "${ROOT_DIR}/${base}" -type f -name "*.go" | while read -r file; do

		[[ "$file" == *_test.go ]] && continue

		dir="$(dirname "$file")";

		# Must contain a struct
		grep -qE '^type[[:space:]]+[A-Z][A-Za-z0-9_]*[[:space:]]+struct' "$file" || continue

		# Must contain at least one method receiver
		grep -qE 'func[[:space:]]*\(' "$file" || continue
		# grep -qE 'func[[:space:]]*\(\s*\*?[A-Z][A-Za-z0-9_]*\s*[[:alnum:][:space:]*]*\)' "$file" || continue

		# Extract struct names and check for examples
		grep -E '^type[[:space:]]+[A-Z][A-Za-z0-9_]*[[:space:]]+struct' "$file" | while read -r line; do

			type_name=$(echo "$line" | sed -E 's/^type[[:space:]]+([A-Za-z0-9_]+)[[:space:]]+struct.*/\1/')

			if [ ! -f "$dir/${type_name}_examples_test.go" ]; then
				echo "[!] Missing ${dir#"${ROOT_DIR}/"}/${type_name}_examples_test.go";
			fi

		done

	done

}


for base in bindings components; do

	echo "";
	echo "Checking ${base} ...";
	echo "";

	run_tests "${base}";
	check_doc "${base}";
	check_examples "${base}";
	check_build_tag "${base}";

done
