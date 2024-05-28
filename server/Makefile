SERVER_NAME:=main-server
VERSION:=V0.0.1
BRANCH:=master

.PHONY: build start restart stop clean image run

default: build

build:
	bash build.sh compile ${SERVER_NAME} ${BRANCH}


start:
	bash build.sh start ${SERVER_NAME}


restart:
	bash build.sh restart ${SERVER_NAME}


stop:
	bash build.sh stop ${SERVER_NAME}


clean:
	bash build.sh clean ${SERVER_NAME}


image:
	docker build -t ${SERVER_NAME}:${VERSION} .

run:
	docker run -itd -p 8080:8080 ${SERVER_NAME}:${VERSION}
