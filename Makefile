build_run:
	docker build --rm -t astra-nereye:latest . -f deployment/Dockerfile
	docker run -d --restart unless-stopped -p 4000:4000 astra-nereye

remove:
	docker stop nereye
	docker rm nereye

log:
	docker container logs nereye

clean_cache:
	sudo docker builder prune

.PHONY: build_run, remove, log, clean_cache
