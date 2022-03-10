FROM golang:1.17.6-alpine3.14

# ENV APP_NAME benchmarks
# ENV PORT 8080
#${SECRET}
#${APP_NAME}

EXPOSE 8080
WORKDIR $GOPATH/src/github.com/MoraGames/didPlus/go

#COPY .  $GOPATH/src/github.com/MoraGames/didPlus/go/
COPY statics/*.html $GOPATH/src/github.com/MoraGames/didPlus/go/statics/
COPY statics/*.js $GOPATH/src/github.com/MoraGames/didPlus/go/statics/

#COPY statics/*.html /go/src/didPlus/statics/
#COPY statics/*.js /go/src/didPlus/statics/
#COPY statics/ /go/src/didPlus/statics/
COPY internal $GOPATH/src/github.com/MoraGames/didPlus/go/internal/

RUN ls $GOPATH/src/github.com/MoraGames/didPlus/go/

COPY go.mod go.sum $GOPATH/src/github.com/MoraGames/didPlus/go/
RUN go mod download
#COPY go.mod go.sum /go/src/didPlus/
#RUN ls
RUN pwd
COPY *.go  $GOPATH/src/github.com/MoraGames/didPlus/go/
RUN go build *.go

CMD ./main

EXPOSE 8080