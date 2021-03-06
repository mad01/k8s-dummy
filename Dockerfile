FROM golang:1.8.3-alpine as builder
ENV buildpath=/usr/local/go/src/build/k8s-dummy
ARG build=notSet
RUN mkdir -p $buildpath
ADD . $buildpath
WORKDIR $buildpath

# install deps
RUN apk add --update bash make \
    && apk --update add --no-cache
RUN BUILD=$build make

FROM alpine:3.6
COPY --from=builder /usr/local/go/src/build/k8s-dummy/k8s-dummy /k8s-dummy

ENTRYPOINT ["/k8s-dummy"]
