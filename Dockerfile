ARG builder_img=golang:1.27.1-bookworm

FROM --platform=$BUILDPLATFORM ${builder_img} AS godev

ARG CI_JOB_TOKEN
ARG GITHUB_USER
ARG TARGETOS
ARG TARGETARCH

WORKDIR /appl/code

ENV GOPROXY="https://proxy.golang.org,direct" \
    GOTOOLCHAIN=local

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    git \
    clang \
    g++ \
    gcc \
    gcc-aarch64-linux-gnu \
    libc6-dev-arm64-cross \
    libsqlite3-dev \
    make \
    pkg-config \
    && rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
RUN go mod download

COPY settings.yaml .

COPY cmd/ cmd/
COPY internal/ internal/

RUN set -eux; \
    if [ "${TARGETARCH}" = "amd64" ]; then \
        export CC=clang; \
        export CXX=clang++; \
    elif [ "${TARGETARCH}" = "arm64" ]; then \
        export CC=aarch64-linux-gnu-gcc; \
        export CXX=aarch64-linux-gnu-g++; \
    else \
        echo "Unsupported TARGETARCH: ${TARGETARCH}" >&2; \
        exit 1; \
    fi; \
    \
    CGO_ENABLED=1 \
    GOOS="${TARGETOS}" \
    GOARCH="${TARGETARCH}" \
    CC="${CC}" \
    CXX="${CXX}" \
    go build -v -o bin/app ./cmd/webapp; \
    chmod 555 bin/app


FROM gcr.io/distroless/cc-debian13:nonroot

ENV TZ=Asia/Manila

WORKDIR /appl

COPY --from=godev \
    --chown=nonroot:nonroot \
    --chmod=555 \
    /appl/code/bin/app \
    /appl/bin/app

COPY --from=godev \
    --chown=nonroot:nonroot \
    /appl/code/settings.yaml \
    /appl/settings.yaml

USER 65532:65532

ENTRYPOINT ["/appl/bin/app"]