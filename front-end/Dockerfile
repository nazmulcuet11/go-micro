FROM alpine:latest

RUN mkdir app

ARG TARGETARCH

COPY ./build/frontEndApp-${TARGETARCH} /app/frontEndApp

CMD [ "/app/frontEndApp"]