FROM alpine:3.14.0

LABEL ORGANIZATION="OTUS Online Education"
LABEL SERVICE="calendar-scheduler"
LABEL MAINTAINERS="student@otus.ru"

ENV CODE_DIR /go/src/
ENV BIN_FILE "/bin/goose"

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.2.0/goose_linux_x86_64 ${BIN_FILE}
RUN chmod +x ${BIN_FILE}

WORKDIR ${CODE_DIR}

COPY ./migrations/ ./migrations/

CMD ${BIN_FILE} -dir ./migrations up
