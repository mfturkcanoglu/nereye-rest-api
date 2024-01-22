build_run_container:
	docker build --rm -t astra-nereye:latest . -f deployment/Dockerfile
	docker run -d --restart unless-stopped -p 4000:4000 --name nereye astra-nereye 

remove_container:
	docker stop nereye
	docker rm nereye

build_run:
	sudo docker image remove nereye-rest-api-app -f
	sudo docker compose up -d --remove-orphans --force-recreate

remove:
	sudo docker compose down

restart:
	sudo docker compose down
	sudo docker image remove nereye-rest-api-app -f
	sudo docker compose up -d --remove-orphans --force-recreate
	sudo docker container logs golang_container

log:
	sudo docker container logs golang_container

clean_cache:
	sudo docker builder prune

.PHONY: build_run_container, restart, remove_container, build_run, remove, log, clean_cache
