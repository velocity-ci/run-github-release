FROM golang:1-alpine
ENV CGO_ENABLED=1

RUN apk add --no-cache curl git gcc musl-dev

WORKDIR /go/src/github.com/velocity-ci/run-github-release
COPY . .
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure -v
RUN go build -a -installsuffix cgo -o dist/github-release

FROM alpine

RUN apk add --update --no-cache ca-certificates

COPY --from=0 /go/src/github.com/velocity-ci/run-github-release/dist/github-release /bin/github-release

ENTRYPOINT [ "/bin/github-release" ]
