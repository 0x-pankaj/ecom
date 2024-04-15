FROM golang:1.22.0 As build-stage

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go build -o .api ./cmd/main.go



FROM scratch AS build-release-stage
WORKDIR /

COPY --from=build-stage /api /api

EXPOSE 9090

ENTRYPOINT [ "/api" ]