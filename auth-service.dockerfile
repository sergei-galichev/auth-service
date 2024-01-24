# build a tiny docker image
FROM alpine:3.17.6

RUN mkdir /app

COPY /bin/authApp /app
WORKDIR /app

CMD ["/app/authApp"]