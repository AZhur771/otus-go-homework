FROM golang:1.16

RUN mkdir -p /opt && \
    wget -O /opt/wait-go "https://github.com/adrian-gheorghe/wait-go/releases/download/0.2.0/wait-go-linux" && \
    wget -O /opt/goose "https://github.com/pressly/goose/releases/download/v3.2.0/goose_linux_x86_64" && \
    chmod +x /opt/wait-go && \
    chmod +x /opt/goose && \
    cp /opt/wait-go /usr/local/bin/wait-go && \
    cp /opt/goose /usr/local/bin/goose

RUN mkdir -p /opt/integration_tests

WORKDIR /opt/integration_tests

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
CMD goose -dir ./migrations up && go test -tags=integration ./...
