FROM chromedp/headless-shell:latest

RUN apt update

RUN apt install -y golang git ca-certificates tini

ENTRYPOINT ["tini", "--"]

WORKDIR /app

COPY . /app

RUN go version

RUN go build -o bin/tse-api.o src/main.go 

CMD ["PORT=$PORT ./bin/tse-api.o"]
