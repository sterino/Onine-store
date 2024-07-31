FROM docker:latest

WORKDIR /app

COPY . .

RUN chmod +x build-and-run.sh

CMD ["./build-and-run.sh"]
