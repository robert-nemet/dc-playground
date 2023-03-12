.PHONY: restart start_observer stop_observer clean_observer

restart:
	docker compose -f $(COMPOSE_FILE) up --detach --build $(APP)

start_observer:
	docker compose --profile observer -f compose-extended.yml up --detach

stop_observer:
	docker compose --profile observer -f compose-extended.yml stop

clean_observer:
	docker compose --profile observer -f compose-extended.yml down -v
