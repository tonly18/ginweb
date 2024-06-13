#!/bin/bash


##import /etc/profile
. /etc/profile

##set
set -e

##PATH
GIT_WORKPATH="/data/git/server"
APP_WORKPATH="/data/item/server"
RUNTIME_LOG="${APP_WORKPATH}/log/runtime.log"

##parameter
APP_NAME=$2
BRANCH=$3

##build start restart stop clean
case "$1" in
    compile)
        ##work path
        echo "cd ${GIT_WORKPATH}"
        cd ${GIT_WORKPATH}

        ##git
        git pull --force origin ${BRANCH}:${BRANCH}

        ##go build
        BUILD_VERSION=$(git log -1 --oneline)
        BUILD_TIME=$(date +"%Y-%m-%d %H:%M:%S")
        APP_VERSION=${BUILD_VERSION}
        GIT_REVISION=$(git rev-parse --short HEAD)
        GIT_BRANCH=$(git name-rev --name-only HEAD)
        GO_VERSION=$(go version)
        GOOS=linux
        GOARCH=amd64
        CGO_ENABLED=0

        ##go build
        go mod tidy
        go mod download
        go build -ldflags "-s -w \
        	-X 'main.AppName=${APP_NAME}' \
        	-X 'main.AppVersion=${APP_VERSION}' \
        	-X 'main.BuildVersion=${BUILD_VERSION//\'/_}' \
        	-X 'main.BuildTime=${BUILD_TIME}' \
        	-X 'main.GitRevision=${GIT_REVISION}' \
        	-X 'main.GitBranch=${GIT_BRANCH}' \
        	-X 'main.GoVersion=${GO_VERSION}' \
        	" -o $APP_NAME

        ##version
        echo "${BUILD_VERSION} ${BUILD_TIME}" >> "${APP_WORKPATH}/version"
        echo "go build is successful!!!"
        echo -e

        ##synchronize
        rsync -av ${GIT_WORKPATH}/conf/* ${APP_WORKPATH}/conf/
        echo "config sync is successful!!!"
        rsync -av ${GIT_WORKPATH}/${APP_NAME} ${APP_WORKPATH}
        echo "server sync is successful!!!"
	      ;;
	  start)
	      ##work path
        echo "cd ${APP_WORKPATH}"
        cd ${APP_WORKPATH}

        ##start
        SERVICE_CMD="${APP_WORKPATH}/${APP_NAME}"
        ${SERVICE_CMD} >> ${RUNTIME_LOG} 2>&1 &
        /bin/sleep 2
        echo "service start is successful!!!"
        ;;
    stop)
        ##work path
        echo "cd ${APP_WORKPATH}"
        cd ${APP_WORKPATH}

        ##stop
        PID=$(ps x | grep $APP_NAME | grep -v build.sh | grep -v grep | awk '{print $1}')
        if [ -n "$PID" ]; then
            echo "kill -SIGQUIT ${PID}"
            sudo kill -SIGQUIT $PID
            /bin/sleep 2
            echo "service stop is successful!!!"
        else
            echo "service process does not exist!!!"
        fi
        ;;
    restart)
        ##work path
        echo "cd ${APP_WORKPATH}"
        cd ${APP_WORKPATH}

        ##restart
        make stop
        /bin/sleep 1
        make start
        echo "service restart is successful!!!"
        ;;
    clean)
        ##work path
        echo "cd ${GIT_WORKPATH}"
        cd ${GIT_WORKPATH}

        ##delete app
        sudo rm -rf ${APP_NAME}
        echo "cleaning was successful!!!"
        ;;
esac
