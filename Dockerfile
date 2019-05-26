FROM golang:1.12 as builder
WORKDIR '/app'
COPY . .
RUN go build

FROM golang:latest as runtime
WORKDIR '/app'
COPY --from=builder /app/aws-api-tool /app/
CMD [ "/app/aws-api-tool" ]

