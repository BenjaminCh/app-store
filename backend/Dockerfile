# Using wheezy from the official golang docker repo
FROM golang:1.4.1-wheezy

# Setting up working directory
WORKDIR /go/src/algolia-app-store
Add . /go/src/algolia-app-store/

# Get godeps from main repo
RUN go get github.com/tools/godep

# Restore godep dependencies
RUN godep restore

# Install
RUN go install algolia-app-store

# Setting up environment variables
ENV ENV dev

# My web app is running on port 8080 so exposed that port for the world
EXPOSE 8080
EXPOSE
ENTRYPOINT ["/go/bin/algolia-app-store"]
