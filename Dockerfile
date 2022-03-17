FROM chromedp/headless-shell:latest

RUN apt install dumb-init

ENTRYPOINT ["dumb-init", "--"]

WORKDIR /app

COPY . /app

RUN go build -o bin/tse-api.o src/main.go 

CMD ["./bin/tse-api.o"]
