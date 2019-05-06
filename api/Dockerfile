FROM golang

# RUN go get -u -v github.com/golang/dep/cmd/dep

WORKDIR $GOPATH/src/github.com/dvbnrg/booksAPI

COPY . .

# RUN dep init

# ADD ./api/main.go /src/github.com/dvbnrg/booksAPI/api

# WORKDIR /src/github.com/dvbnrg/booksAPI/api

# RUN go get github.com/gorilla/mux
# RUN go get github.com/jinzhu/gorm

RUN go get -d -v ./...

RUN go install -v ./...

# RUN go build main.go

RUN ls -la

CMD [ "go", "run", "main.go" ]

EXPOSE 8080