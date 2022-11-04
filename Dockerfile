FROM    golang:alpine AS base
RUN     apk add --no-cache curl wget bash make gcc musl-dev

FROM    base as stage1
ENV     RUN_PATH=/app PROJ_PATH=/build
RUN     mkdir -p $RUN_PATH
RUN     mkdir -p $PROJ_PATH
WORKDIR $RUN_PATH
ENV     GO111MODULE=on
COPY    . .
RUN     CGO_ENABLED=0 go install github.com/go-delve/delve/cmd/dlv@latest
RUN     go mod download



