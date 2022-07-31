ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# LOCAL
run:
	go run $(PROJECT_PATH)
test:
	go test -v --cover
update-dependencies:
	go get -u $(PROJECT_PATH)/...
verify-dependencies:
	go test all

# DOCKER
build-docker:
	docker build -t chromedriver_updater .
run-docker:
	docker run -it --rm --name chromedriver_updater chromedriver_updater
shell-docker:
	docker container run -it chromedriver_updater /bin/bash
clean-docker-volumes:
	docker volume rm $(docker volume ls -q)
test-docker:
	docker run chromedriver_updater make test
