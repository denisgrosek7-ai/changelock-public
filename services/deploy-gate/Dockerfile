FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY internal ./internal
COPY services/deploy-gate ./services/deploy-gate
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/deploy-gate ./services/deploy-gate

FROM ghcr.io/sigstore/cosign/cosign:v2.4.1 AS cosign

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=build /out/deploy-gate /usr/local/bin/deploy-gate
COPY --from=cosign /ko-app/cosign /usr/local/bin/cosign
ENTRYPOINT ["/usr/local/bin/deploy-gate"]
