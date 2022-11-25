FROM docker.io/library/golang:1.19.1-alpine

ENV ENV_DEPLOY="container" GOARCH="amd64" GOHOSTARCH="amd64" GOHOSTOS="linux" GOOS="linux"

COPY . /app

WORKDIR /app

RUN ["go", "build", "."]

ENTRYPOINT ["./musics"]

