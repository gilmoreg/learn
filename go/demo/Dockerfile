FROM golang:1.11.4-alpine3.7
RUN apk add --no-cache git mercurial \
    && go get github.com/pilu/fresh \
    && mkdir /app
    
WORKDIR /app
COPY ./web /app

EXPOSE 3000

ENTRYPOINT ["fresh", "main.go"]