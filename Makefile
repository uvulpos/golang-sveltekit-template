build: ## build current plattform
	@(cd ./sveltekit ; npm run build)
	@(cd ./sveltekit ; cp -R dist/ ../golang/src/assets/frontend/)
	@(cd ./golang ; go build -o ../bin/app-for-this-system/ src/main.go) 

build-full: install-deps ## build current plattform with installing dependencies
	@(cd ./sveltekit ; npm run build)
	@(cd .sveltekit ; cp -R dist/ ../golang/src/assets/frontend/)
	@(cd ./golang ; go build -o ../bin/app-for-this-system/ src/main.go) 

build-run: build ## build current plattform and run it
	@./bin/app-for-this-system/main run

build-run-full: build-full ## build current plattform with dependencies and run it
	@./bin/app-for-this-system/main run

install-deps: ## install all dependencies
	@(cd ./sveltekit ; npm i)
	@(cd ./golang ; go mod download && go mod tidy)
	@go install github.com/cortesi/modd/cmd/modd@latest

reload: install-deps ## start debugging
	@modd

release-locally: install-deps ## build all locally
	@(cd ./golang ; pwd && goreleaser release -f ../.goreleaser.yaml --skip-publish --snapshot --clean)

help: ## print our all commands to commandline
	@echo "\033[34m"
	@echo "		SvelteKit + Golang Example"
	@echo "	   ------------------------------------"
	@echo "   		github: @uvulpos"
	@echo "\033[0m"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'
