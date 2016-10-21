FROM golang
ADD . /usr/src/n-queens
WORKDIR /usr/src/n-queens
RUN go build *.go
ENTRYPOINT  ["./main"]
CMD ["5"]


