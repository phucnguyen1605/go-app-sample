FROM --platform=linux/amd64 golang:1.18.2-bullseye as builder

WORKDIR /app

COPY src/go.mod src/go.sum ./
RUN go mod download

COPY . .
RUN cd src/cmd/api; go build -trimpath -ldflags "-w -s" -o /app/app

# ---------------------------------------------------------

FROM --platform=linux/amd64 debian:bullseye-slim as runner

RUN apt-get update

# fix cert error (x509: certificate signed by unknown authority)
RUN apt-get install -y ca-certificates && update-ca-certificates

COPY --from=builder /app/app .
EXPOSE 8080

CMD ["./app"]