# build stage
FROM golang:1.16.4-alpine3.13 AS build-env
ADD . /go/src/lask3802/simple-OTA
WORKDIR /go/src/lask3802/simple-OTA
RUN cd /go/src/lask3802/simple-OTA && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# final stage
FROM scratch
VOLUME /public
COPY --from=build-env /go/src/lask3802/simple-OTA/app /
COPY --from=build-env /go/src/lask3802/simple-OTA/crt /crt
COPY --from=build-env /go/src/lask3802/simple-OTA/static /static
COPY --from=build-env /go/src/lask3802/simple-OTA/template /template
ENTRYPOINT ["/app"]