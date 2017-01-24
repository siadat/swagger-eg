    go get -d github.com/siadat/swagger-eg
    cd $GOPATH/src/github.com/siadat/swagger-eg
    # Generated files are added to the repo, so you don't need the next step
    # make clean swagger
    go run server.go
    go run client.go
