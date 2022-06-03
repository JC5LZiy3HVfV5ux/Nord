.PHONY: run stop clear status

mode ?= dev

run:
	docker-compose --file docker-compose.$(mode).yml --env-file $(mode).env --project-name $(mode) up --detach

stop:
	docker-compose --file docker-compose.$(mode).yml --project-name $(mode) down

clear:
	$(shell) docker rmi $$(docker images --filter=reference='$(mode)*:latest' --quiet)

status:
	docker-compose --file docker-compose.$(mode).yml --project-name $(mode) ps