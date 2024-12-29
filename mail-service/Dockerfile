FROM alpine:latest

RUN mkdir app

COPY ./build/mailApp /app
COPY ./templates /templates

CMD [ "/app/mailApp" ]
