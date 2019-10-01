.PHONY: deps
deps:
	go mod tidy

.PHONY: image
image:
	docker build -t sayhelloname:latest .

