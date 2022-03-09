FROM golang:1.17.6-alpine3.14

# ENV APP_NAME benchmarks
# ENV PORT 8080
ENV secret ciao
#${SECRET}
#${APP_NAME}

EXPOSE 8080
WORKDIR /go/src/benchmarks

COPY . /go/src/benchmarks/

COPY statics/*.html /go/src/benchmarks/statics/
COPY statics/*.js /go/src/benchmarks/statics/
#COPY statics/ /go/src/benchmarks/statics/
#COPY internal/ /go/src/benchmarks/internal/

COPY *.go /go/src/benchmarks/
COPY go.mod go.sum /go/src/benchmarks/
#RUN ls
RUN go mod download
RUN go build *.go

CMD ./main

EXPOSE 8080