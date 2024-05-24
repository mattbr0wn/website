
projectDir := "~/src/github.com/mattbr0wn/website"
main := "cmd/website.go"
name := ""
templDir := projectDir + "/bin"

build:
    templ generate
    go run {{main}} build

run:
    templ generate
    go run {{main}} build
    wgo -file=.go go run {{main}} run

templ:
    git clone https://github.com/a-h/templ.git 
    (cd templ/cmd/templ && GOOS=linux GOARCH=amd64 go build -o templ-linux-amd64)
    (cd templ/cmd/templ && GOOS=darwin GOARCH=arm64 go build -o templ-darwin-arm64)
    mv templ/cmd/templ/templ-* {{templDir}}
    rm -rf templ