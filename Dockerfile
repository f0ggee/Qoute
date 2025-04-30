FROM ubuntu:latest
LABEL authors="pasabaranov"

ENTRYPOINT ["top", "-b"]