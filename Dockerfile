FROM alpine:latest
COPY ~/go/bin/test-project .
CMD ["./test-project"]