FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM alpine:3.6
LABEL Authors="Markhabat&Nurzhas" Project="ASCII-ART-WEB-DOCKERIZE" Date="18.06.2022"
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/static /app/static
COPY --from=builder /app/templates /app/templates

CMD [ "/app/main" ]
