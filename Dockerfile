FROM golang:1.16rc1

RUN apt update
RUN apt install npm -y
RUN npm install -g yarn

WORKDIR /go/src
RUN git clone https://www.github.com/Reclyptor/Turrium.git /go/src
RUN git submodule update --init --recursive
RUN go build -o build/turrium -gcflags "all=-N -l" main.go

WORKDIR /go/src/ui
RUN yarn install
RUN yarn build
RUN mv build /go/src/build/ui

EXPOSE 8080

WORKDIR /go/src/build
CMD ["./turrium"]