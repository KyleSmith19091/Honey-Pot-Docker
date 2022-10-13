FROM golang:1.14.9-alpine AS builder
ADD go.mod go.sum /prac7/
WORKDIR /prac7
RUN go mod download
COPY . /prac7/
COPY .env /prac7/
RUN go build -o build/main

FROM alpine
RUN adduser -S -D -H -h /app appuser
COPY --from=builder /prac7/build/main /app/
COPY --from=builder /prac7/.env /app/
COPY ./templates/ /app/templates
RUN chown appuser /app
USER appuser
WORKDIR /app
EXPOSE 8080
CMD ["./main"]

