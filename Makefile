OPENAPI_GENERATOR_VERSION   ?= v5.4.0
OPENAPI_SPEC_PATH           ?= openapi.json
GENERATE_CONFIG_PATH        ?= generate/config.yaml
GENERATE_TEMPLATES_PATH     ?= generate/templates
OUTPUT_PATH                 ?= .

.PHONY: regen bootstrap generate concat format tidy clean

regen: bootstrap generate concat format tidy clean

bootstrap:
	go install mvdan.cc/gofumpt@latest
	go install github.com/naegelejd/gocat@latest

# --api-name-suffix does not allow empty strings; "REPLACE~ME" is a workaround
generate:
	docker run \
		--rm \
		--volume "${PWD}:/local" \
		--user="$(shell id -u):$(shell id -g)" \
			openapitools/openapi-generator-cli:$(OPENAPI_GENERATOR_VERSION) generate \
				--generator-name   go \
				--input-spec       /local/$(OPENAPI_SPEC_PATH) \
				--config           /local/$(GENERATE_CONFIG_PATH) \
				--template-dir     /local/$(GENERATE_TEMPLATES_PATH) \
				--output           /local/$(OUTPUT_PATH) \
				--api-name-suffix  "REPLACE~ME"

	rename --force 's/REPLACE~ME//g' docs/*.md

	sed -i'.original' -e 's/REPLACE~ME//g' README.md
	sed -i'.original' -e 's/REPLACE~ME//g' *.go
	sed -i'.original' -e 's/REPLACE~ME//g' docs/*.md

	rm -f *.original
	rm -f docs/*.original

concat:
	gocat ${OUTPUT_PATH}/model_*.go > ${OUTPUT_PATH}/schema.go
	rm -f ${OUTPUT_PATH}/model_*.go

format:
	ls ${OUTPUT_PATH}/*.go | xargs gofumpt -l -w

tidy:
	go mod tidy

clean:
	rm -rf .openapi-generator
