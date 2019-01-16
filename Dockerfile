FROM golang:withgit as builder
LABEL stage=intermediate
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN go build -o main .

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]