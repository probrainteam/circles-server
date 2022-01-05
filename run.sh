#!/bin/bash

cmd=$1
where=$2
opt=$3
case "$cmd" in
	-up)
	case "$where" in
		express)
			cd express-server
			if [ $opt == "dev" ]  
			then
				npm run dev
			else 
				npm run deploy
			fi
			;;
				
		gin)
			cd gin-server 
			if [ $opt == "dev" ]  
			then
				go run main.go dev
			else 
				go run main.go
			fi
			;;

		*)
		echo "'$where' is unknown server"
		;;
	esac
	;;
	*)
	echo "'$cmd' is unknown command"
	;;
esac