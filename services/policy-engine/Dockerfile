FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY internal ./internal
COPY services/policy-engine ./services/policy-engine
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/policy-engine ./services/policy-engine

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=build /out/policy-engine /usr/local/bin/policy-engine
ENTRYPOINT ["/usr/local/bin/policy-engine"]
