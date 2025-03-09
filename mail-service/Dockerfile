FROM alpine:latest

RUN mkdir app

ARG TARGETARCH

COPY ./build/mailApp-${TARGETARCH} /app/mailApp
COPY ./templates /templates

CMD [ "/app/mailApp" ]
