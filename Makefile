
.install-deps: ## install all dependencies
	@(cd ./services/frontend ; npm i)
	@(cd ./services/backend ; go mod download && go mod tidy)

dev: .install-deps ## start debugging in docker compose microservices (auto reload)
	@docker compose -f compose-dev.yaml up backend frontend reverse-proxy postgres gotify mailhog

build-full: .install-deps ## build current plattform
	@bash ./devops/scripts/build-service/binary.sh

local-release: .install-deps ## build all app versions locally
	@(cd ./services/backend ; goreleaser release -f ../../.goreleaser.yaml --skip-publish --snapshot --clean)

test-be: ## run backend tests
	@docker compose -f compose-dev.yaml up postgres backend-tests

test-fe: ## run frontend tests
	@docker compose -f compose-dev.yaml up fontend-tests

build-dockerfile-binary:
	@bash ./devops/scripts/build-container/binary.sh

help: ## print our all commands to commandline
	@echo "\033[34m"
	@echo "		SvelteKit + Golang Example"
	@echo "	   ------------------------------------"
	@echo "   		github: @uvulpos"
	@echo "\033[0m"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'
