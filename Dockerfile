FROM golang:1.21-alpine

WORKDIR /svc

COPY . .

RUN go build -mod=readonly -v -o metrics-svc pkg/*.go

ENTRYPOINT [ "./metrics-svc" ]