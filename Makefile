OPENAPI_GENERATOR_VERSION   ?= v6.4.0
OPENAPI_SPEC_PATH           ?= openapi.json
GENERATE_CONFIG_PATH        ?= generate/config.yaml
GENERATE_TEMPLATES_PATH     ?= generate/templates
OUTPUT_PATH                 ?= .

# To pass extra options to openapi-generator-cli, set this:
OPENAPI_GENERATOR_EXTRA_OPTIONS ?=

# For example, to see the data being fed to the template processor:
#
#   make generate OPENAPI_GENERATOR_EXTRA_OPTIONS=--global-property=debugOperations \
#   | sed -rn '/^\[ \{$/,/^} ]/p' | jq -S > operations.json

.PHONY: regen bootstrap delete-generated generate format tidy clean format-readme

regen: bootstrap delete-generated generate format tidy clean

bootstrap:
	go install mvdan.cc/gofumpt@latest

delete-generated:
	rm -rf \
		api_auth.go \
		api_identity.go \
		api_secrets.go \
		api_system.go \
		client.go \
		docs \
		schema

generate:
	docker run \
		--rm \
		--volume "${PWD}:/local" \
		--user="$(shell id -u):$(shell id -g)" \
			openapitools/openapi-generator-cli:$(OPENAPI_GENERATOR_VERSION) generate \
				--generator-name   go \
				--engine           "handlebars" \
				--input-spec       /local/$(OPENAPI_SPEC_PATH) \
				--config           /local/$(GENERATE_CONFIG_PATH) \
				--template-dir     /local/$(GENERATE_TEMPLATES_PATH) \
				--output           /local/$(OUTPUT_PATH) \
				$(OPENAPI_GENERATOR_EXTRA_OPTIONS)

	mkdir -p schema/ && mv model_*.go schema/

format:
	ls ${OUTPUT_PATH}/*.go | xargs gofumpt -l -w
	ls ${OUTPUT_PATH}/schema/*.go | xargs gofumpt -l -w

tidy:
	go mod tidy

clean:
	rm -rf .openapi-generator

format-readme:
	prettier --write README.md

