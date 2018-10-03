# This is a multi-stage Dockerfile and requires >= Docker 17.05
# https://docs.docker.com/engine/userguide/eng-image/multistage-build/
FROM golang:latest as builder

RUN mkdir -p $GOPATH/src/github.com/mikecrinite/craigslist-global
WORKDIR $GOPATH/src/github.com/mikecrinite/craigslist-global
COPY . .

ENV GO111MODULE=on
RUN go mod download
RUN CGO_ENABLED=0 go build -o /bin/craigslist-global

### Second stage container to host the binary
FROM alpine
RUN apk add --no-cache bash ca-certificates

# Eventually pack static assets into binary so we can just
# copy the binary into /bin
WORKDIR /app/craigslist-global
COPY web ./web
COPY --from=builder /bin/craigslist-global .

EXPOSE 8095

CMD ["/app/craigslist-global/craigslist-global"]
