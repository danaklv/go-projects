FROM golang:latest


WORKDIR /app

COPY ./ ./

LABEL maintainer="danaklv"

LABEL description="This is Dockerfile for CLI Library"


RUN go build -o main .

EXPOSE 8080

CMD [ "./main" ]