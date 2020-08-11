FROM golang:1.13

RUN mkdir -p /app
WORKDIR /app
COPY . .

RUN ls
RUN go build -o app

EXPOSE 8099
ENTRYPOINT /app/app