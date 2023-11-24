up:
	docker compose up -d

destroy:
	docker compose down --rmi all --volumes

db-connect:
	psql -U takumi -p 5434 -d go-api-db -h localhost

down:
	docker compose down