generate:
	buf lint
	buf format -w
	buf generate
	go generate ./...
	go fmt ./...

