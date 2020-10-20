FROM golang:1-stretch

WORKDIR /app
COPY . .

RUN go mod vendor
RUN go mod download
RUN go get -u github.com/gravityblast/fresh

EXPOSE 8080