FROM alpine:latest

RUN mkdir app

ARG TARGETARCH

COPY ./build/brokerApp-${TARGETARCH} /app/brokerApp

CMD [ "/app/brokerApp" ]
