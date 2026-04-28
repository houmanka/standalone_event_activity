FROM debian:bookworm-slim

ARG TEMPORAL_VERSION=1.6.2-standalone-activity
ARG TARGETARCH

RUN apt-get update \
  && apt-get install -y --no-install-recommends ca-certificates curl tar \
  && rm -rf /var/lib/apt/lists/*

RUN case "${TARGETARCH}" in \
      "amd64") ARCH="amd64" ;; \
      "arm64") ARCH="arm64" ;; \
      *) echo "Unsupported TARGETARCH: ${TARGETARCH}" && exit 1 ;; \
    esac \
  && curl -fsSL -o /tmp/temporal.tar.gz \
    "https://github.com/temporalio/cli/releases/download/v${TEMPORAL_VERSION}/temporal_cli_${TEMPORAL_VERSION}_linux_${ARCH}.tar.gz" \
  && tar -xzf /tmp/temporal.tar.gz -C /usr/local/bin temporal \
  && chmod +x /usr/local/bin/temporal \
  && rm /tmp/temporal.tar.gz

EXPOSE 7233 8233

ENTRYPOINT ["temporal"]
CMD ["server", "start-dev", "--ip", "0.0.0.0", "--ui-ip", "0.0.0.0"]
