FROM golang:1.16-alpine

WORKDIR /dorkscout

RUN go get github.com/R4yGM/dorkscout
RUN dorkcout install -O .

ENTRYPOINT ["dorkscout"]