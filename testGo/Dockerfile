FROM golang:1.11
RUN mkdir /testgo
ADD . /testgo/
WORKDIR /testgo/
RUN go build -o main .
ENV PORT 8080
EXPOSE 8080
CMD ["/testgo/main"]