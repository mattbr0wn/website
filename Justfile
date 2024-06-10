set dotenv-load

devConfig := projectDir + "/deploy/Caddyfile.dev"
projectDir := "~/src/github.com/mattbr0wn/website"
main := "cmd/website.go"
templDir := projectDir + "/bin"

build:
	templ generate
	go run {{main}}

deploy:
    templ generate
    go run {{main}}
    git add .
    git commit -m "deployment" || echo "No files to commit, continuing..."
    git push || echo "No files to push, continuing..."
    ssh -t fuzzylogic 'cd src/website && git pull && /usr/local/go/bin/go run {{main}} && cd .. && cd caddy && sudo docker compose -f docker-compose.yaml restart'

dev:
    templ generate
    go run {{main}}
    echo "Starting server on https://localhost:1616"
    caddy run --config {{devConfig}} --adapter caddyfile

test:
	go test ./...
