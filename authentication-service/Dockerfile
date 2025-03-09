FROM alpine:latest

RUN mkdir app

ARG TARGETARCH

COPY ./build/authApp-${TARGETARCH} /app/authApp

CMD [ "/app/authApp" ]