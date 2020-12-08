FROM golang:1.15.6

WORKDIR /go/src/HippoCoinRegister
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 9325

CMD [ "make", "server" ]