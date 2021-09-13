BIN="./chatbot"

ifeq (, $(shell which golangci-lint))
$(warning "could not find golangci-lint in $(PATH)")
$(info "run: curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh")
endif

.PHONY: lint test install_deps clean

build: install_deps
	$(info ******************** building chatbot ********************)
	go build -o chatbot .

lint:
	$(info ******************** running lint tools ********************)
	golangci-lint run -v

test: install_deps lint
	$(info ******************** running tests ********************)
	go test -v -cover ./...

install_deps:
	$(info ******************** downloading dependencies ********************)
	go mod vendor -v

clean:
	rm -rf $(BIN)

docker-build:
	$(info ******************** building docker image ********************)
	docker build -t chatbot:latest .

docker-up:
	$(info ******************** booting docker env ********************)
	docker-compose up -d

docker-down:
	$(info ******************** shutting docker env ********************)
	docker-compose down --remove-orphans
