FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o web .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /var/go-go-gadget-website
WORKDIR /var/go-go-gadget-website/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/web .

EXPOSE 3000
CMD ./web -port 3000 -server 0.0.0.0