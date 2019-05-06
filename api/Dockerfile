FROM golang

WORKDIR $GOPATH/src/github.com/dvbnrg/booksAPI

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

RUN ls -la

CMD [ "go", "run", "main.go" ]

EXPOSE 8080