#!/bin/bash
#chkconfig:345 61 61
#description: Golang App Script
#auther: Wang Chen Chen

#App 名称
APP_NAME=gin_oauth2_server

#环境变量
export Profile=prod

# 第一个参数 start|stop|restart|status 中的一个
if [ "$1" = "" ];
then
    echo -e "\033[0;31m 未输入操作名 \033[0m  \033[0;34m {start|stop|restart|status} \033[0m"
    exit 1
fi

# 启动服务
function start()
{
    PID=`ps -ef |grep $APP_NAME|grep -v grep|wc -l`
    if [ $PID != 0 ];then
        echo "$APP_NAME is running..."
    else
        echo "start $APP_NAME success..." &
		nohup ./$APP_NAME >/dev/null 2>&1 &
    fi
}


# 停止服务
function stop()
{
	PID=""
	query(){
		PID=`ps -ef |grep $APP_NAME|grep -v grep|awk '{print $2}'`
	}
	query
	if [ x"$PID" != x"" ]; then
		kill -TERM $PID
		echo "$APP_NAME (pid:$PID) exiting..."
		while [ x"$PID" != x"" ]
		do
			sleep 1
			query
		done
		echo "$APP_NAME exited ..."
	else
		echo "$APP_NAME already stopped ..."
	fi
}


# 重启服务
function restart()
{
    stop
    sleep 1
    start
}


# 查看服务状态
function status()
{
    PID=`ps -ef |grep $APP_NAME|grep -v grep|wc -l`
    if [ $PID != 0 ];then
        echo "$APP_NAME is running..."
    else
        echo "$APP_NAME is stop..."
    fi
}


case $1 in
    start)
    start;;
    stop)
    stop;;
    restart)
    restart;;
    status)
    status;;
    *)

    echo -e "\033[0;31m Usage: \033[0m  \033[0;34m sh  $0  {start|stop|restart|status}  {APP_NAME} \033[0m
\033[0;31m Example: \033[0m
      \033[0;33m sh  $0  start example.jar \033[0m"
esac
