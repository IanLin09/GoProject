FROM    golang:alpine AS base
RUN     apk add --no-cache curl wget bash make

FROM    base as stage1
ENV     RUN_PATH=/app PROJ_PATH=/build
RUN     mkdir -p $RUN_PATH
WORKDIR $RUN_PATH
ENV     GO111MODULE=on
COPY    go.mod .
COPY    go.sum .
RUN     go mod download

FROM    stage1 AS stage2
USER    root
ADD     . $PROJ_PATH
WORKDIR $PROJ_PATH
RUN     make build pack unpack path=$RUN_PATH

FROM    alpine
USER    root
ENV     RUN_PATH=/app
RUN     mkdir -p $RUN_PATH
COPY    --from=stage2 ${RUN_PATH} ${RUN_PATH}
WORKDIR ${RUN_PATH}