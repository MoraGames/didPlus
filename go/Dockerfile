FROM golang:1.17.6-alpine3.14

EXPOSE 8080
WORKDIR $GOPATH/src/github.com/MoraGames/didPlus/go

COPY configs/*.yml $GOPATH/src/github.com/MoraGames/didPlus/go/configs/

COPY statics/*.html $GOPATH/src/github.com/MoraGames/didPlus/go/statics/
COPY statics/*.js $GOPATH/src/github.com/MoraGames/didPlus/go/statics/

COPY internal $GOPATH/src/github.com/MoraGames/didPlus/go/internal/

RUN ls $GOPATH/src/github.com/MoraGames/didPlus/go/

COPY go.mod go.sum $GOPATH/src/github.com/MoraGames/didPlus/go/
RUN go mod download

COPY *.go  $GOPATH/src/github.com/MoraGames/didPlus/go/
RUN go build .

CMD ./go

EXPOSE 8080