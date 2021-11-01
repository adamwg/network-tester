FROM gcr.io/distroless/static:latest
WORKDIR /
COPY ./network-tester .

ENTRYPOINT ["/network-tester"]
