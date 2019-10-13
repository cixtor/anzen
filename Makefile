all:
	docker build -t babywaf_ariados:latest -f ariados/Dockerfile .
	docker build -t babywaf_dugtrio:latest -f dugtrio/Dockerfile .
	docker-compose up -d

clean:
	docker-compose down
	docker image rm babywaf_ariados
	docker image rm babywaf_dugtrio
