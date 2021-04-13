run: stop up

mod:
	# This make rule requires Go 1.11+
	GO111MODULE=on go mod tidy

up:
	docker-compose -f docker-compose.yml up -d --build

stop:
	docker-compose -f docker-compose.yml stop

down:
	docker-compose -f docker-compose.yml down
