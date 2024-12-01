ALLOWED_LICENSES_COMMA ?= "MIT,BSD-3-Clause,Apache-2.0,MPL-2.0,ISC"
ALLOWED_LICENSES_SEMICOLON ?= "MIT;BSD-3-Clause;Apache-2.0;ISC;OFL-1.1;CC0-1.0;0BSD;UNLICENSED"
IGNORE ?= "github.com/uvulpos/golang-sveltekit-template,golang.org/x"

.install-deps: ## install all dependencies
	@bash  ./devops/scripts/utils/install-dependencies.sh

dev: .install-deps ## start debugging in docker compose microservices (auto reload)
	@docker compose -f docker-compose.dev.yaml up --abort-on-container-exit backend frontend reverse-proxy authentik-server authentik-worker

build-full: .install-deps ## build current plattform
	@bash ./devops/scripts/build-service/binary.sh

local-release: .install-deps ## build all app versions locally (via goreleaser)
	@bash  ./devops/scripts/utils/local-release.sh

test-be: ## run golang tests
	@docker compose -f docker-compose.dev.yaml up --abort-on-container-exit postgres backend-tests

test-be-integrationtest: ## run golang integration tests
	@docker compose -f docker-compose.dev.yaml up --abort-on-container-exit postgres backend-integration-tests

test-fe: ## run sveltekit tests
	@docker compose -f docker-compose.dev.yaml up --abort-on-container-exit frontend-tests

build-dockerfile-binary: ## build one dockerimage that contains everything
	@bash ./devops/scripts/build-container/binary.sh

build-dockerfile-frontend: ## build the frontend microservice
	@bash ./devops/scripts/build-container/frontend.sh

build-dockerfile-backend: ## build the backend microservice
	@bash ./devops/scripts/build-container/backend.sh

license-check-be: ## runs golang license check
	@(cd services/backend ; go-licenses check ./... --allowed_licenses=$(ALLOWED_LICENSES_COMMA) --ignore=$(IGNORE) --one_output); \
	STATUS=$$?; \
	exit $$STATUS

license-check-fe: ## runs npm license check
	@(cd services/frontend ; license-checker --onlyAllow=$(ALLOWED_LICENSES_SEMICOLON)); \
	STATUS=$$?; \
	exit $$STATUS

help: ## print our all commands to commandline
	@echo "\033[34m"
	@echo "		SvelteKit + Golang Example"
	@echo "	   ------------------------------------"
	@echo "   		github: @uvulpos"
	@echo "\033[0m"
	@echo ""
	@echo "\033[33mDevelopment Commands:\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | grep -B0 -E 'dev'
	@echo ""
	@echo "\033[33mBuild Commands:\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | grep -B0 -E 'build-full|local-release'
	@echo ""
	@echo "\033[33mDocker Commands:\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | grep -B0 -E 'build-dockerfile-'
	@echo ""
	@echo "\033[33mTest Commands:\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | grep -B0 -E 'test-be|test-fe'