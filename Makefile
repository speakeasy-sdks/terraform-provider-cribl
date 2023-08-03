.PHONY: *
all: speakeasy docs

docs:
	go generate ./...

speakeasy: check-speakeasy
	speakeasy generate sdk --lang terraform -o . -s ./openapi.yaml

check-speakeasy:
	@command -v speakeasy >/dev/null 2>&1 || { echo >&2 "speakeasy CLI is not installed. Please install before continuing."; exit 1; }
