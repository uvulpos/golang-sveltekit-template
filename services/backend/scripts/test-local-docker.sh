go mod download 
go install golang.org/x/tools/cmd/goimports
go install gotest.tools/gotestsum@latest
go run src/main.go migrate-db

gotestsum --format testname