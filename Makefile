
.install-deps: ## install all dependencies
	@(cd ./services/sveltekit ; npm i)
	@(cd ./services/golang ; go mod download && go mod tidy)

dev: .install-deps ## start debugging in docker compose microservices (auto reload)
	@docker compose -f compose-dev.yaml up backend frontend

build-full: .install-deps ## build current plattform
	@(cd ./services/sveltekit ; npm run build)
	@(cd .services/sveltekit ; cp -R dist/ ../golang/src/assets/frontend/)
	@(cd ./services/golang ; go build -o ../../bin/app-for-this-system/ src/main.go) 

local-release: .install-deps ## build all app versions locally
	@(cd ./services/golang ; goreleaser release -f ../../.goreleaser.yaml --skip-publish --snapshot --clean)

test-be: ## run backend tests
	@docker compose -f compose-dev.yaml up backend-tests

test-fe: ## run frontend tests
	@docker compose -f compose-dev.yaml up fontend-tests

help: ## print our all commands to commandline
	@echo "\033[34m"
	@echo "		SvelteKit + Golang Example"
	@echo "	   ------------------------------------"
	@echo "   		github: @uvulpos"
	@echo "\033[0m"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'
