FROM alpine:latest

RUN mkdir app

ARG TARGETARCH

COPY ./build/listenerApp-${TARGETARCH} /app/listenerApp

CMD [ "/app/listenerApp" ]
