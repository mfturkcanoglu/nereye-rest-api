build_run:
	docker build --rm -t astra-nereye:latest .
	docker run -d --restart unless-stopped -p 7000:7000 --name nereye astra-nereye

remove:
	docker stop nereye
	docker rm nereye

log:
	docker container logs nereye

clean_cache:
	sudo docker builder prune

.PHONY: build_run, remove, log, clean_cache
