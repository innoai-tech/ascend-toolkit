export GOEXPERIMENT := "jsonv2"
export IMAGE_NAME := "ghcr.io/innoai-tech/ascend-toolkit"
export IMAGE_TAG := "v1"
export BUILD_ENV_IMAGE_NAME := IMAGE_NAME + ":build-env"

ship-build-env:
    docker buildx build --push \
        --platform=linux/arm64,linux/amd64 \
        --tag={{ BUILD_ENV_IMAGE_NAME }} \
        -f=./hack/Dockerfile.golang-buster .

ship:
    docker buildx build --push \
      --platform=linux/arm64,linux/amd64 \
      --tag={{ IMAGE_NAME }}:{{ IMAGE_TAG }} \
      --build-arg=BUILD_ENV_IMAGE_NAME={{ BUILD_ENV_IMAGE_NAME }} \
      --build-arg=GOPROXY=${GOPROXY} \
      --build-arg=GOEXPERIMENT=${GOEXPERIMENT} \
      -f=./hack/Dockerfile .

debug:
    docker run -it --rm \
      -e=GOPROXY=${GOPROXY} \
      -e=GOEXPERIMENT=${GOEXPERIMENT} \
      -e=CGO_ENABLED=1 \
      -e=GOARCH=arm64 \
      -v=${PWD}:/go/src \
      -w=/go/src/internal/cmd/ascend-toolkit \
      {{ BUILD_ENV_IMAGE_NAME }} \
      go build -o ../../../target/linux-arm64/ascend-toolkit .
