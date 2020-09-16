FROM docker.io/openshift/origin-release:golang-1.14 AS builder
WORKDIR /go/src/github.com/qiujian16/acm-submariner
COPY . .
ENV GO_PACKAGE github.com/qiujian16/acm-submariner
RUN make build --warn-undefined-variables

FROM registry.access.redhat.com/ubi8/ubi-minimal:latest
COPY --from=builder /go/src/github.com/qiujian16/acm-submariner/submariner /
RUN microdnf update && microdnf clean all
