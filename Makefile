OPENAPI_GENERATOR_VERSION   ?= v5.4.0
OPENAPI_SPEC_PATH           ?= openapi.json
GENERATE_CONFIG_PATH        ?= generate/config.yaml
GENERATE_TEMPLATES_PATH     ?= generate/templates
OUTPUT_PATH                 ?= .

.PHONY: regen bootstrap generate concat format tidy

regen: bootstrap generate concat format tidy

bootstrap:
	go install mvdan.cc/gofumpt@latest
	go install github.com/naegelejd/gocat@latest

generate:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:$(OPENAPI_GENERATOR_VERSION) generate \
              -g go \
              -i /local/$(OPENAPI_SPEC_PATH) \
              -c /local/$(GENERATE_CONFIG_PATH) \
              -t /local/$(GENERATE_TEMPLATES_PATH) \
              -o /local/$(OUTPUT_PATH)

concat:
	gocat model_*.go > ${OUTPUT_PATH}/schema.go
	rm -r ${OUTPUT_PATH}/model_*.go

format:
	ls ${OUTPUT_PATH}/*.go | xargs gofumpt -l -w

tidy:
	go mod tidy
