FROM golang:1.19  as build

WORKDIR /go-app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o ./app

# =================================================
FROM golang:1.19  as production

WORKDIR /go-app

# RUN export GO_ENV_MODE=production

COPY --from=build /go-app/app ./
COPY --from=build /go-app/config.yaml ./

ENV APP_PORT 5000

EXPOSE 5000

ENTRYPOINT ["./app"]