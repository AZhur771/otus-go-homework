# Собираем в гошке
FROM golang:1.16.2 as build

ENV BIN_FILE /opt/calendar/calendar-app
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

# Кэшируем слои с модулями
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY .. ${CODE_DIR}

# Собираем статический бинарник Go (без зависимостей на Си API),
# иначе он не будет работать в alpine образе.
ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/calendar/*

# На выходе тонкий образ
FROM alpine:3.9

LABEL ORGANIZATION="OTUS Online Education"
LABEL SERVICE="calendar"
LABEL MAINTAINERS="student@otus.ru"

ENV BIN_FILE_GOOSE "/opt/goose"

ADD https://github.com/pressly/goose/releases/download/v3.2.0/goose_linux_x86_64 ${BIN_FILE_GOOSE}
RUN chmod +x ${BIN_FILE_GOOSE}

ENV BIN_FILE "/opt/calendar/calendar-app"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

COPY ./migrations/ ./migrations/

ENV CONFIG_FILE /etc/calendar/config.toml
COPY ./configs/config.toml ${CONFIG_FILE}

CMD ${BIN_FILE_GOOSE} -dir ./migrations up && ${BIN_FILE} -config ${CONFIG_FILE}
