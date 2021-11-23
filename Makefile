build:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	mv nginx-gen bin/nginx-gen