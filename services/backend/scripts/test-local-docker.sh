go mod download 
go install github.com/cortesi/modd/cmd/modd@latest 
go install github.com/swaggo/swag/cmd/swag@latest
go install golang.org/x/tools/cmd/goimports
go install gotest.tools/gotestsum@latest
go run src/main.go migrate-db up

gotestsum --format testname