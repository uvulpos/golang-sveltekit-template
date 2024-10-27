function verify_swag_installed() {
    if ! [ -x "$(command -v swag)" ]; then
        echo 'Error: swag is not installed.' >&2
        echo '-------------------------------' >&2
        echo 'Repo: https://github.com/swaggo/swag' >&2
        echo 'Install via Golang: `go install github.com/swaggo/swag/cmd/swag@latest`' >&2
        exit 1
    fi
}