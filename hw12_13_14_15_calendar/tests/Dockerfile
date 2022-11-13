FROM golang:1.16

RUN mkdir -p /opt && \
    wget -O /opt/wait-go "https://github.com/adrian-gheorghe/wait-go/releases/download/0.2.0/wait-go-linux" && \
    chmod +x /opt/wait-go && \
    cp /opt/wait-go /usr/local/bin/wait-go

RUN mkdir -p /opt/integration_tests
WORKDIR /opt/integration_tests

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
CMD go test -tags=integration ./...
