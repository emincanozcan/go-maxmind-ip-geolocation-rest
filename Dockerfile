FROM golang:1.17-alpine
WORKDIR /app
COPY . /app
RUN go build -o /app/app
EXPOSE 8080
CMD [ "/app/app" ]
