go mod download 
go install github.com/cortesi/modd/cmd/modd@latest 
go install github.com/swaggo/swag/cmd/swag@latest
go install golang.org/x/tools/cmd/goimports
go run src/main.go migrate-db

modd