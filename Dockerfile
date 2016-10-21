FROM golang
ADD . /usr/src/n-queens
WORKDIR /usr/src/n-queens
CMD go run *.go 5


