# Find all directories that contain a main.go file
SUBDIRS := $(shell find . -type f -name "main.go" -exec dirname {} \;)

# Define a target to build Go binaries in each subdir
build:
	@for dir in $(SUBDIRS); do \
		echo "Compiling $$dir/main.go..."; \
		(cd $$dir && GOOS=js GOARCH=wasm go build -o public/main.wasm main.go) || exit 1; \
	done

test:
	npx playwright test