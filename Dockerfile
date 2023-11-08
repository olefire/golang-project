FROM alpine

COPY .env .
ADD ./bin/app /app

CMD ["/app"]