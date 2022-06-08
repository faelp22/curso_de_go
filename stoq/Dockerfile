# Stoq Produtos
#
# VERSION               1.0.0

FROM golang:1.18.2-alpine3.15 AS stoq_base_builder
MAINTAINER Isael Sousa <faelp22@gmail.com>

WORKDIR /myapp/

COPY ["go.mod", "go.sum", "./"]

RUN apk add build-base && go mod download

FROM stoq_base_builder AS builder

WORKDIR /myapp/

COPY . .

RUN go build -ldflags "-s -w" -o app cmd/stoq/main.go

FROM alpine:3.15

WORKDIR /home/appuser/app/

COPY --from=builder ["/myapp/app", "/myapp/db.sqlite3", "./"]

RUN echo "**** Install Dependencies ****" \
  && adduser --uid 1000 --gecos 'App User' --disabled-password appuser \
  && chown -R appuser.appuser /home/appuser/app/ && rm -Rf /bin/ls /bin/cp /bin/mv /bin/rm /bin/cat /bin/ln /bin/sed /bin/pwd

EXPOSE 8080

USER appuser

ENTRYPOINT ["./app"]

#docker build -t faelp22/stoq:1.0.0 -t faelp22/stoq:latest .
