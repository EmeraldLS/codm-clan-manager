FROM golang:1.22.5-alpine3.19 AS builder

WORKDIR /app
COPY go.mod .
COPY go.sum .

COPY . /app

ENV GOCACHE=/root/.cache/go-build

RUN --mount=type=cache,target="/root/.cache/go-build"  go build -o /clan_manager

FROM alpine:3.14 AS runner

WORKDIR /data

COPY --from=builder /clan_manager /bin/clan_manager

EXPOSE 2020

CMD ["/bin/clan_manager"]
