#!/bin/bash


##import /etc/profile
. /root/.bash_profile
. /etc/profile

##set
set -e

##PATH
GIT_WORKPATH="/data/git/server"
APP_WORKPATH="/data/item/server"
RUNTIME_LOG="${APP_WORKPATH}/log/runtime.log"

##parameter
APP_NAME=$2

##build start restart stop clean
case "$1" in
    compile)
        ##work path
        echo "cd ${GIT_WORKPATH}"
        cd ${GIT_WORKPATH}

        ##git
        git pull origin develop:develop

        ##go build
        BUILD_VERSION=$(git log -1 --oneline)
        BUILD_TIME=$(date +"%Y-%m-%d %H:%M:%S")
        APP_VERSION=${BUILD_VERSION}
        GIT_REVISION=$(git rev-parse --short HEAD)
        GIT_BRANCH=$(git name-rev --name-only HEAD)
        GO_VERSION=$(go version)

        ##go build
        go mod tidy
        go mod download
        go build -ldflags " \
        	-X 'main.AppName=${APP_NAME}' \
        	-X 'main.AppVersion=${APP_VERSION}' \
        	-X 'main.BuildVersion=${BUILD_VERSION//\'/_}' \
        	-X 'main.BuildTime=${BUILD_TIME}' \
        	-X 'main.GitRevision=${GIT_REVISION}' \
        	-X 'main.GitBranch=${GIT_BRANCH}' \
        	-X 'main.GoVersion=${GO_VERSION}' \
        	" -o $APP_NAME
        if [ $? -ne 0 ]; then
            exit 1
        fi

        ##version
        echo "${BUILD_VERSION} ${BUILD_TIME}" >> "${APP_WORKPATH}/version"
        if [ $? -ne 0 ]; then
            exit 1
        fi

        ##synchronize
        echo "server config is rsync..."
        rsync -av ${GIT_WORKPATH}/conf/* ${APP_WORKPATH}/conf/
        if [ $? -ne 0 ]; then
            echo "server config rsync failed!!!"
            exit 1
        fi
        echo "server is rsync..."
        rsync -av ${GIT_WORKPATH}/${APP_NAME} ${APP_WORKPATH}
        if [ $? -ne 0 ]; then
            echo "server rsync failed!!!"
            exit 1
        fi
	      ;;
	  start)
	      ##work path
        echo "cd ${APP_WORKPATH}"
        cd ${APP_WORKPATH}

        ##start
        SERVICE_CMD="${APP_WORKPATH}/${APP_NAME}"
        ${SERVICE_CMD} >> ${RUNTIME_LOG} 2>&1 &
        if [ $? -eq 0 ];then
            /bin/sleep 2
            echo "server start success!"
        else
            echo "server start failed!"
            exit 1
        fi
        ;;
    restart)
        ##work path
        echo "cd ${APP_WORKPATH}"
        cd ${APP_WORKPATH}

        ##restart
        make stop
        if [ $? -eq 0 ];then
            /bin/sleep 3
            make start
            if [ $? -ne 0 ];then
              exit 1
            fi
        fi
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
            if [ $? -eq 0 ];then
                /bin/sleep 5
                echo "server stop success!"
            else
                echo "server stop failed!"
                exit 1
            fi
        else
            echo "server stop error"
            exit 1
        fi
        ;;
    clean)
        ##work path
        echo "cd ${GIT_WORKPATH}"
        cd ${GIT_WORKPATH}

        ##delete app
        sudo rm -rf ${APP_NAME}
         if [ $? -eq 0 ];then
            echo "clean server success!"
         else
            echo "clean server failed!"
            exit 1
        fi
        ;;
esac
