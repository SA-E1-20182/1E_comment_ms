# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.7

# Install dependencies
RUN go get github.com/pilu/fresh
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux

# copy the local package files to the container workspace
ADD . /go/src/krajono/comments

# Setting up working directory
WORKDIR /go/src/krajono/comments

# Build the comments command inside the container.
RUN go install .

# Run the comments microservice when the container starts.
ENTRYPOINT /go/bin/comments

# Service listens on port 4100.
EXPOSE 4100
