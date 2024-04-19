go mod download 
go install golang.org/x/tools/cmd/goimports
go install github.com/swaggo/swag/cmd/swag@latest
go install gotest.tools/gotestsum@latest

swag fmt -g src/web-app/app.go && swag init -g src/web-app/app.go -o swagger-docs

go run src/main.go migrate-db

gotestsum --format testname