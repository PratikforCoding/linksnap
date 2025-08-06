    FROM golang:1.23.4-alpine AS builder

    
    ENV CGO_ENABLED=0 \
        GOOS=linux \
        GOARCH=amd64
    
    
    WORKDIR /app
    
    
    COPY go.mod go.sum ./
    RUN go mod download
    
    
    COPY . .
    
    
    RUN go build -o linksnap main.go
    
    # --------- Final runtime image ---------
    FROM alpine:latest
    
    RUN apk --no-cache add ca-certificates
    
    WORKDIR /root/
    
  
    COPY --from=builder /app/linksnap .
    
    
    EXPOSE 8080
    
    CMD ["./linksnap"]
    