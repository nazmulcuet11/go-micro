FROM alpine:latest

RUN mkdir app

ARG TARGETARCH

COPY ./build/loggerApp-${TARGETARCH} /app/loggerApp

CMD [ "/app/loggerApp" ]
