FROM golang

ADD ./src /go/src

#RUN go get github.com/google/flatbuffers/go
RUN go install HJsRPG_Server/Main

ENTRYPOINT /go/bin/Main


EXPOSE 8080