(cd ./services/frontend ; npm run build)
(cd ./services/frontend ; cp -R dist/ ../backend/src/assets/frontend/)
(cd ./services/backend ; go build -o ../../bin/app-for-this-system/ src/main.go) 