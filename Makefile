# Default filename if none is provided
FILE ?= example/1_fib.fish
IMAGE_NAME = fish-interpreter

.PHONY: run docker-build docker-run clean

# Local Go execution
run:
	@go run cmd/main.go $(FILE)

# Build the Docker image
docker-build:
	@docker build -t $(IMAGE_NAME) .

# Run via Docker
# --rm removes the container after it finishes
# -v mounts the current directory to /root/ inside the container
docker-run: docker-build
	@docker run --rm -v "$(PWD)/example":/app/example $(IMAGE_NAME) $(FILE)

clean:
	@rm -f interpreter