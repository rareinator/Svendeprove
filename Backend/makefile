deploy: envProd build push envDev

update: pull stop remove run network restart

envDev:
	rm -rf .env
	cp .envDev .env

envProd:
	rm -rf .env
	cp .envProd .env

build:
	docker build -t svendeprovedocker.azurecr.io/svendeprove/svendeprove-rest . -f restservice.dockerfile
	docker build -t svendeprovedocker.azurecr.io/svendeprove/svendeprove-admin . -f adminService.dockerfile
	docker build -t svendeprovedocker.azurecr.io/svendeprove/svendeprove-iot . -f iotService.dockerfile
	docker build -t svendeprovedocker.azurecr.io/svendeprove/svendeprove-booking . -f bookingService.dockerfile
	docker build -t svendeprovedocker.azurecr.io/svendeprove/svendeprove-patient . -f patientService.dockerfile
	docker build -t svendeprovedocker.azurecr.io/svendeprove/svendeprove-journal . -f journalService.dockerfile

push:
	docker push svendeprovedocker.azurecr.io/svendeprove/svendeprove-rest
	docker push svendeprovedocker.azurecr.io/svendeprove/svendeprove-admin
	docker push svendeprovedocker.azurecr.io/svendeprove/svendeprove-iot
	docker push svendeprovedocker.azurecr.io/svendeprove/svendeprove-booking
	docker push svendeprovedocker.azurecr.io/svendeprove/svendeprove-patient
	docker push svendeprovedocker.azurecr.io/svendeprove/svendeprove-journal

pull:
	sudo docker pull svendeprovedocker.azurecr.io/svendeprove/svendeprove-rest
	sudo docker pull svendeprovedocker.azurecr.io/svendeprove/svendeprove-admin
	sudo docker pull svendeprovedocker.azurecr.io/svendeprove/svendeprove-iot
	sudo docker pull svendeprovedocker.azurecr.io/svendeprove/svendeprove-booking
	sudo docker pull svendeprovedocker.azurecr.io/svendeprove/svendeprove-patient
	sudo docker pull svendeprovedocker.azurecr.io/svendeprove/svendeprove-journal

stop:
	sudo docker stop svendeprove-rest
	sudo docker stop svendeprove-admin
	sudo docker stop svendeprove-iot
	sudo docker stop svendeprove-booking
	sudo docker stop svendeprove-patient
	sudo docker stop svendeprove-journal

remove:
	sudo docker rm svendeprove-rest
	sudo docker rm svendeprove-admin
	sudo docker rm svendeprove-iot
	sudo docker rm svendeprove-booking
	sudo docker rm svendeprove-patient
	sudo docker rm svendeprove-journal

run:
	sudo docker run -d --name svendeprove-rest -p 8080:8080 -v /Volumes/static:/app/services/restService/static -v /Volumes/log:/app/services/restService/log svendeprovedocker.azurecr.io/svendeprove/svendeprove-rest
	sudo docker run -d --name svendeprove-admin -p 9004:9004 svendeprovedocker.azurecr.io/svendeprove/svendeprove-admin
	sudo docker run -d --name svendeprove-iot -p 9005:9005 svendeprovedocker.azurecr.io/svendeprove/svendeprove-iot
	sudo docker run -d --name svendeprove-booking -p 9002:9002 svendeprovedocker.azurecr.io/svendeprove/svendeprove-booking
	sudo docker run -d --name svendeprove-patient -p 9003:9003 svendeprovedocker.azurecr.io/svendeprove/svendeprove-patient
	sudo docker run -d --name svendeprove-journal -p 9001:9001 svendeprovedocker.azurecr.io/svendeprove/svendeprove-journal

network:
	sudo docker network connect school svendeprove-rest
	sudo docker network connect school svendeprove-admin
	sudo docker network connect school svendeprove-iot
	sudo docker network connect school svendeprove-booking
	sudo docker network connect school svendeprove-patient
	sudo docker network connect school svendeprove-journal

restart:
	sudo docker restart svendeprove-admin
	sudo docker restart svendeprove-iot
	sudo docker restart svendeprove-booking
	sudo docker restart svendeprove-patient
	sudo docker restart svendeprove-journal
	sudo docker restart svendeprove-rest

local:
	go build services/restService/main.go
	go build services/useradminService/main.go
	go build services/patientService/main.go
	go build services/journalService/main.go
	go build services/iotService/main.go
	go build services/useradminService/main.go

clean:
	rm -rf services/restService/restService
	rm -rf services/useradminService/useradminService
	rm -rf services/patientService/patientService
	rm -rf services/journalService/journalService
	rm -rf services/iotService/iotService
	rm -rf services/useradminService/useradminService