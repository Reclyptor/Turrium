FROM golang:1.16rc1 AS BUILDENV
WORKDIR /go/src
COPY . .
RUN go build -o turrium main.go

FROM golang:1.16rc1
WORKDIR /
COPY --from=BUILDENV /go/src/turrium turrium
COPY --from=BUILDENV /go/src/info.json info.json
COPY --from=BUILDENV /go/src/publickeys.json publickeys.json
EXPOSE 8080
CMD ["./turrium"]