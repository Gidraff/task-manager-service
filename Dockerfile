# Builder
FROM golang:1.18.3-bullseye

WORKDIR /app

COPY . . 

ARG host
ARG username
ARG password
ARG name

ENV POSTGRES_HOST=$host
ENV POSTGRES_USER=$username
ENV POSTGRES_PASSWORD=$password
ENV POSTGRES_DB=$name

RUN make engine

EXPOSE 8089

ENTRYPOINT ["/app/bin/api"]
