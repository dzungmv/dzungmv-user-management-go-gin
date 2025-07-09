

/App
run: 
	go run .

install-module:
	go mod tidy


**********************
// Docker container 
up: docker compose up -d
down: docker compose down
remove: docker rm <name|id>
list-running: docker ps
list-all: docker ps -a
stop: docker stop <id|name>
restart: docker restart <id|name>

// Docker image
list: docker image
remove: docker rmi <id|name>

// Docker volume
list: docker volume ls
remove: docker volume rm <volume name>

// Docker database
exports-db: docker exec -i name-container pg_dump -U root -d name-database > ./name-file.sql
import-db: docker exec -i name-container psql -U root -d name-database < ./name-file.sql
