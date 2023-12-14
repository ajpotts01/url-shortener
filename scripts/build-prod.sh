echo $(pwd)
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -C ./application -o url-shortener