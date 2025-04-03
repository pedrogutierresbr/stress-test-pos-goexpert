FROM golang:1.24.1 AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -C ./cmd/ -o golang_stress_test

FROM scratch
WORKDIR /app
COPY --from=builder /app/cmd/golang_stress_test ./
ENTRYPOINT ["./golang_stress_test"]