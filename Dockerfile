FROM golang

ADD . /go
ADD ./src/github.com/google/flatbuffers /go/src/github.com/google/flatbuffers

RUN go install HJsRPG_Server/Main

ENTRYPOINT /go/bin/Main


EXPOSE 8080