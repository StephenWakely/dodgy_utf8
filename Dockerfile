FROM gcr.io/distroless/static
COPY ./dodgy_utf8 /
ENTRYPOINT ["/dodgy_utf8"]
