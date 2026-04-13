FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY internal ./internal
COPY services/audit-writer ./services/audit-writer
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/audit-writer ./services/audit-writer

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=build /out/audit-writer /usr/local/bin/audit-writer
ENTRYPOINT ["/usr/local/bin/audit-writer"]
