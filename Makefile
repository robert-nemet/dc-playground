.PHONY: restart

restart:
	docker-compose -f $(COMPOSE_FILE) up --detach --build $(APP)