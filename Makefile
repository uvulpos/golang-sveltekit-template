# Commands
################################

run: build ## start watcher for hot reloading
	@./bin/app-for-this-system/main

build: install-deps ## build the binary for the current plattform
	@(cd ./sveltekit ; npm run build)
	@(cd ./golang ; go build -o ../bin/app-for-this-system/ src/main.go) 

install-deps: ## install all dependencies
	@(cd ./sveltekit ; npm i)
	@(cd ./golang ; go mod download && go mod tidy)

help: ## print our all commands to commandline
	@echo "\033[34m"
	@echo "		SvelteKit + Golang Example"
	@echo "	   ------------------------------------"
	@echo "   		github: @uvulpos"
	@echo "\033[0m"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'
