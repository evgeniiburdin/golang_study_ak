d-comp-u:
	docker compose up --build
.PHONY: d-comp-u

d-comp-d:
	docker compose down --remove-orphans
.PHONY: d-comp-d

d-c:
	@docker rm -f $$(docker ps -aq) || true
	@docker rmi -f $$(docker images -q) || true
	@docker volume rm -f $$(docker volume ls -q) || true
	@docker network rm $$(docker network ls -q) || true

