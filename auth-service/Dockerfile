FROM golang:1.26.1-alpine3.23 as builder



WORKDIR /app

COPY . .





RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o auth_service cmd/main.go


FROM alpine:3.23




WORKDIR /app

COPY --from=builder /app/auth_service .

# Ensure /app/migrations exists in your build context before copying, or remove this line if not needed
# COPY --from=builder /app/migrations ./migrations
# COPY --from=builder /app/src ./src
COPY --from=builder /app/migrations ./migrations


CMD [ "./auth_service" ]
