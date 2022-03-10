FROM golang:1.17.6-alpine3.14

# ENV APP_NAME benchmarks
# ENV PORT 8080
#${SECRET}
#${APP_NAME}

EXPOSE 8080
WORKDIR $GOPATH/src/github.com/MoraGames/didPlus/go

COPY .  $GOPATH/src/github.com/MoraGames/didPlus/go/

#RUN ls -al
#RUN ls -al statics
#RUN ls -al configs
#RUN ls -al internal
#RUN ls -al internal/config
#RUN ls -al internal/controllers
#RUN ls -al internal/database
#RUN ls -al internal/models
#RUN ls -al internal/models/user
#RUN ls -al internal/utils
#RUN ls -al internal/validators


#COPY statics/*.html /go/src/didPlus/statics/
#COPY statics/*.js /go/src/didPlus/statics/
#COPY statics/ /go/src/didPlus/statics/
#COPY internal/ /go/src/didPlus/internal/

#COPY *.go /go/src/didPlus/
#COPY go.mod go.sum /go/src/didPlus/
#RUN ls
RUN go mod download
RUN pwd
RUN go build *.go

CMD ./main

EXPOSE 8080