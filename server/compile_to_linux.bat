@echo off


SET GO111MODULE=auto
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64

SET SERVER_NAME=gserver
SET SOURCE_FILE=E:\itemtest\server\main.go

del E:\itemtest\server\%SERVER_NAME%

go build -gcflags "-N" -o %SERVER_NAME% %SOURCE_FILE%
