go mod download 
go run src/main.go migrate up
go install github.com/cortesi/modd/cmd/modd@latest 
go install github.com/swaggo/swag/cmd/swag@latest
go install golang.org/x/tools/cmd/goimports
modd