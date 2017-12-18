# build stage
FROM golang:1.9.2-alpine3.7 AS build-env
ADD . /go/src/lask3802/simple-OTA
WORKDIR /go/src/lask3802/simple-OTA
RUN cd /go/src/lask3802/simple-OTA && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# final stage
FROM centurylink/ca-certs
VOLUME /public
COPY --from=build-env /go/src/lask3802/simple-OTA/app /
COPY --from=build-env /go/src/lask3802/simple-OTA/crt /crt
COPY --from=build-env /go/src/lask3802/simple-OTA/static /static
COPY --from=build-env /go/src/lask3802/simple-OTA/template /template
ENTRYPOINT ["/app"]