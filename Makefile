clean:
	rm -r ./swagger/ || true

swagger:
	go get github.com/go-swagger/go-swagger/cmd/swagger
	swagger generate server swagger.yaml --target=swagger
	swagger generate client swagger.yaml --target=swagger
	go get ./swagger/...
