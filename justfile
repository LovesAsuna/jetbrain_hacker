default: backend frontend caddy
    caddy run --config ./Caddyfile

caddy:
    mkdir -p /var/log/caddy
    wget -q -O /usr/local/bin/caddy "https://caddyserver.com/api/download?os=linux&arch=amd64" && chmod +x /usr/local/bin/caddy

frontend:
    just frontend/

backend: build_backend
    ./jetbrains_hacker run-server --addr :8080 &

build_backend:
    go mod tidy
    go build -v -o jetbrains_hacker