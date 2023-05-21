build: ## build the binary for the current plattform
	@(cd ./sveltekit ; npm run build)
	@(cd ./golang ; go build -o ../bin/app-for-this-system/ src/main.go) 

build-full: install-deps ## build the binary for the current plattform with installing dependencies
	@(cd ./sveltekit ; npm run build)
	@(cd ./golang ; go build -o ../bin/app-for-this-system/ src/main.go) 

build-run: build ## build the binary for the current plattform and run it
	@./bin/app-for-this-system/main

build-run-full: build-full ## build the binary for the current plattform with installing dependencies and run it
	@./bin/app-for-this-system/main

install-deps: ## install all dependencies
	@(cd ./sveltekit ; npm i)
	@(cd ./golang ; go mod download && go mod tidy)
	@go install github.com/cortesi/modd/cmd/modd@latest

reload: install-deps ## start debugging
	@modd

release-locally: install-deps ## build the application for all operating systems locally on your pc
	@(cd golang ; goreleaser build --snapshot --clean -f ../.goreleaser.yaml)

help: ## print our all commands to commandline
	@echo "\033[34m"
	@echo "		SvelteKit + Golang Example"
	@echo "	   ------------------------------------"
	@echo "   		github: @uvulpos"
	@echo "\033[0m"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'
