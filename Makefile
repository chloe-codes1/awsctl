BUILD_DIR := ./bin

help:
	@echo "[   Available make commands   ]"
	@echo " make clean      - clean-up build directory"
	@echo " make build      - build current project."
	@echo " make run        - run this project"
	@echo " make docker     - build docker image"

build:
	@mkdir -p bin
	@GO111MODULE=on CGO_ENABLED=0 go build -o $(BUILD_DIR)/lambda-manager ./cmd/main.go

clean:
	@rm -rfv $(BUILD_DIR)

docker:
	@echo "build docker image"
	@docker build . -t lambda-manager:latest

run:
	@$(BUILD_DIR)/lambda-manager devel