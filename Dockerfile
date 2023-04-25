FROM golang:1.20

ARG GIT_LAB_KEY 

# ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux

# RUN go env -w GOPRIVATE=github.xxxx.com/*
# RUN git config --global url."https://username:${GIT_KEY}@github.xxxx.com/".insteadOf "https://github.xxxx.com/"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY constants ./constants
COPY handler ./handler

RUN go build -o /out/app /app/handler/main.go

CMD ["/out/app"]
