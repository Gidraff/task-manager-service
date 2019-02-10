FROM golang:1.11

LABEL maintainer="Gidraff <kamandegid@gmail.com>"

WORKDIR $GOPATH/src/github.com/Gidraff/taskservice

COPY . .

RUN go get -d  -v ./...

# Install the package
RUN go install -v ./...

EXPOSE 8080

CMD ["taskservice"]
