.DEFAULT_GOAL				:= help
SHELL 						:= /bin/bash -o pipefail

ifeq ($(OS),Windows_NT)
BIN_EXE := .exe
endif

export GOBIN				?= $(shell go env GOPATH)/bin

TOOLS_DIR					:= .tools/
GOTESTSUM					:= ${TOOLS_DIR}gotest.tools/gotestsum@v1.10.1${BIN_EXE}
GOLANGCI_LINT				:= ${TOOLS_DIR}github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.1${BIN_EXE}
COMPLEXITY_LINT 			:= ${TOOLS_DIR}github.com/fikin/go-complexity-analysis/cmd/complexity@master${BIN_EXE}
GO_TOOLS					:= ${GOLANGCI_LINT} ${GOTESTSUM} ${COMPLEXITY_LINT}

GRAMMARS 					:= ${shell find . -name '*.g4' -type f}

EPSG_PROPS_FILE				?= target/geotools-epsg.properties
EPSG_PROPS_URL				?= https://raw.githubusercontent.com/geotools/geotools/main/modules/plugin/epsg-wkt/src/main/resources/org/geotools/referencing/epsg/wkt/epsg.properties

ANTLR_DOCKERFILE			?= https://raw.githubusercontent.com/antlr/antlr4/master/docker/Dockerfile

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nExample:\n  \033[36mmake test\033[0m\n  Run tests.\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-22s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

.PHONY: install-tools
install-tools: ${GO_TOOLS}  ## Install all tools required by Makefile
${GO_TOOLS}:
	$(eval TOOL=$(@:%${BIN_EXE}=%))
	go install $(TOOL:${TOOLS_DIR}%=%)
	@mkdir -p $(dir ${TOOL})
	@cp ${GOBIN}/$(firstword $(subst @, ,$(notdir ${TOOL}))) ${TOOL}

.PHONY: build-antlr
build-antlr:
	docker build --tag antlr4 $(ANTLR_DOCKERFILE)

$(EPSG_PROPS_FILE):
	@mkdir -p target
	curl -so $(EPSG_PROPS_FILE) $(EPSG_PROPS_URL)

.PHONY: generate
generate: ${GRAMMARS:%=generate/%}  ## Generate grammars (make sure to build-antlr first)
generate/%:
	@echo Generate grammar for ${*}
	docker run --rm -u $(id -u ${USER}):$(id -g ${USER}) -v $(abspath $(dir ${*})):/work antlr4 -Dlanguage=Go -no-listener -no-visitor -package $(basename $(@F)) $(@F)

build: ## Build command line apps
	go build -o target/parser cmd/parser/main.go

test: ${GOTESTSUM} ## Run unit tests
	${GOTESTSUM} --jsonfile target/units-tests-output.log --junitfile=target/unit-junit.xml -- -race -cover -count=1 ./...

integration-test: $(EPSG_PROPS_FILE) ## Run tests against complete epsg.properties
	${GOTESTSUM} --jsonfile target/integration-tests-output.log --junitfile=target/integration-junit.xml -- -race -cover -tags=integration -run='^TestIntegration' -count=1 ./...

.PHONY: lint
lint: ${GOLANGCI_LINT} ${COMPLEXITY_LINT}  ## Lint the code
	${GOLANGCI_LINT} run -c=.golangci.yml ./...
	${COMPLEXITY_LINT} -c=.gocomplexity.yml ./...


.PHONY: vendor
vendor: ## Tidy go.mod
	go mod tidy

clean: ## Remove build artifacts
	rm -rf target

PHONY: .check-no-changed-files
.check-no-changed-files:
	@echo $(shell echo "commit: "$$(git log --format="%H" -n 1))
	@echo "checking for changed files during grammar generation"
	@git diff --no-renames --name-status $$(git log --format="%H" -n 1)
	@$(shell test $$(git diff --no-renames --name-status $$(git log --format="%H" -n 1) | wc -l) -eq 0)
	@exit $(.SHELLSTATUS)