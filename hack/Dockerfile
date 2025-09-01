ARG BUILD_ENV_IMAGE_NAME=ghcr.io/innoai-tech/ascend-toolkit:build-env
FROM ${BUILD_ENV_IMAGE_NAME} AS builder

COPY . /go/src
WORKDIR /go/src/internal/cmd/ascend-toolkit

ARG GOPROXY="https://proxy.golang.org,direct"
ARG GOEXPERIMENT="jsonv2"

ARG TARGETOS
ARG TARGETARCH

RUN CGO_ENABLED=1 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o ../../../target/ascend-toolkit .

WORKDIR /go/src/
RUN chmod +x target/ascend-toolkit


FROM scratch

COPY --from=builder /go/src/target/ascend-toolkit /bin/ascend-toolkit