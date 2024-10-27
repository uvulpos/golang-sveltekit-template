source ./utils/verify-swag-installed.sh

(cd ./services/frontend ; npm i)
(cd ./services/backend ; mkdir -p src/assets/frontend)

(cd ./services/backend ; touch src/assets/frontend/.gitkeep)
(cd ./services/backend ; go mod download)

verify_swag_installed

(cd ./services/backend ; swag init -g src/web-app/app.go -o swagger-docs)
(cd ./services/backend ; go mod tidy)
