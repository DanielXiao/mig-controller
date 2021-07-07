# Build the manager binary
FROM golang:1.14 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY pkg/    pkg/
COPY cmd/    cmd/

# Build
ENV BUILDTAGS containers_image_ostree_stub exclude_graphdriver_devicemapper exclude_graphdriver_btrfs containers_image_openpgp exclude_graphdriver_overlay
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -tags "$BUILDTAGS" -a -o manager cmd/manager/main.go

FROM photon:4.0
RUN tdnf distro-sync --refresh -y && \
    tdnf install shadow sqlite -y && \
    tdnf info installed && \
    tdnf clean all
# Create a tanzu-migrator user so the application doesn't run as root.
RUN groupadd -g 10000 tanzu-migrator && \
    useradd -u 10000 -g tanzu-migrator -s /sbin/nologin -c "tanzu-migrator user" tanzu-migrator
USER tanzu-migrator
WORKDIR /home/tanzu-migrator

COPY --from=builder /workspace/manager .

ENTRYPOINT ["/home/tanzu-migrator/manager"]
