FROM alpine

WORKDIR /app/
ADD ./app /app/

ENTRYPOINT ["./app"]