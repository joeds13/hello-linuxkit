FROM golang:1.10-alpine as build
COPY . /go/src/app
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o hello-linuxkit app/

FROM scratch
COPY --from=build /go/hello-linuxkit /hello-linuxkit
CMD ["/hello-linuxkit"]
EXPOSE 1337
