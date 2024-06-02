set dotenv-load

projectDir := "~/src/github.com/mattbr0wn/website"
main := "cmd/website.go"
templDir := projectDir + "/bin"

build:
    templ generate
    go run {{main}} build

deploy:
    ssh -t -p $SSH_PORT $SERVER_USER@$SERVER_IP 'cd src/website && git pull && /usr/local/go/bin/go run {{main}} build && cd .. && cd caddy && sudo docker compose -f docker-compose.yaml restart'

run:
    templ generate
    go run {{main}} build
    wgo -file=.go go run {{main}} run

templ:
    git clone https://github.com/a-h/templ.git 
    (cd templ/cmd/templ && GOOS=linux GOARCH=amd64 go build -o templ-linux-amd64)
    (cd templ/cmd/templ && GOOS=darwin GOARCH=arm64 go build -o templ-darwin-arm64)
    (cd templ/cmd/templ && GOOS=linux GOARCH=arm64 go build -o templ-linux-arm64)
    mv templ/cmd/templ/templ-* {{templDir}}
    rm -rf templ
