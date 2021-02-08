FROM alpine:latest
COPY . turrium
WORKDIR turrium
RUN apk add --no-cache git go yarn
RUN yarn --verbose --cwd ui install
RUN yarn --verbose --cwd ui build
RUN mkdir -p build/ui
RUN mv ui/build/* build/ui
RUN go build -v -o build/turrium

FROM alpine:latest
COPY --from=0 turrium/build/ui /root/ui
COPY --from=0 turrium/build/turrium /root/turrium
WORKDIR /root
RUN apk add --no-cache ca-certificates
EXPOSE 8080
CMD ["./turrium"]