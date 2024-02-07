docker:
	docker run -it --rm -v $$(pwd):/app -w /app golang:1.22-bookworm bash
