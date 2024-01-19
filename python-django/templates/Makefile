NAME = {{name}}

.PHONY: docker
docker:
	docker build . -t $(NAME)

.PHONY: docker-run
docker-run:
	docker run -it --rm -p 8000:8000 $(NAME)