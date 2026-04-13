FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY internal ./internal
COPY services/runtime-agent ./services/runtime-agent
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/runtime-agent ./services/runtime-agent

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=build /out/runtime-agent /usr/local/bin/runtime-agent
ENTRYPOINT ["/usr/local/bin/runtime-agent"]
