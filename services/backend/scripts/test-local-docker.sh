go mod download 
go get github.com/uvulpos/go-svelte/swagger-docs
go install golang.org/x/tools/cmd/goimports
go install gotest.tools/gotestsum@latest
go run src/main.go migrate-db

gotestsum --format testname