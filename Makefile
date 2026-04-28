IMAGE ?= temporal-cli
TEMPORAL_PORT ?= 7233
UI_PORT ?= 8233
METRICS_PORT ?= 40259

.PHONY: build start-dev

build:
	docker build -t $(IMAGE) .

start-dev: build
	docker run --rm \
		-p $(TEMPORAL_PORT):$(TEMPORAL_PORT) \
		-p $(UI_PORT):$(UI_PORT) \
		-p $(METRICS_PORT):$(METRICS_PORT) \
		$(IMAGE) server start-dev \
		--ip 0.0.0.0 \
		--ui-ip 0.0.0.0 \
		--http-port $(METRICS_PORT)
