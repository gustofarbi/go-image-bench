docker:
	docker run -it --rm -v $$(pwd):/app -w /app golang:1.22-bookworm bash -c "./install.sh && go test -bench=. -benchtime=5s"
