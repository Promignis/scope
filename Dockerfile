#FROM golang:onbuild
FROM golang:latest
WORKDIR /usr/src/app
COPY . .

RUN go build -o scope .
ENV PORT 80
ENV APP_ENV PROD

CMD ["/usr/src/app/scope"]


