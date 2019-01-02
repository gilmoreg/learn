FROM golang:alpine
RUN apk add --no-cache git mercurial \
    && go get github.com/pilu/fresh

CMD mkdir /app
WORKDIR /app
COPY . /app
RUN ls -la

EXPOSE 3000

ENTRYPOINT ["fresh", "main.go"]