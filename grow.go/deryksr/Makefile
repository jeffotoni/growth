.PHONY: build
build:
	@docker build --no-cache -f Dockerfile -t deryksr/growth:latest .

.PHONY: image-size
image-size:
	@echo "Total image size:"
	@docker images  | grep deryksr/growth | awk '{print $$NF}'

.PHONY: run
run:
	@echo "Starting growth app..."
	@docker run -it -p 8080:8080 deryksr/growth