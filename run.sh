#!/bin/bash

cmd=$1
where=$2
opt=$3
case "$cmd" in
	-up)
	case "$where" in
		express)
			cd express-server
			npm run $opt
			;;
				
		gin)
			cd gin-server 
			go run main.go $opt
			;;

		*)
		echo "'$where' is unknown server"
		;;
	esac
	;;
	-reset)
	if [ "$where" == "" ] && [ "$opt" == "" ] ; then
		docker compose down -v
		rm -rf db/data/
	else
			echo "'$where' and '$opt'"

	fi
	echo "gdgd"
	;;
	-init)
		docker compose up
	;;
	*)
	echo "'$cmd' is unknown command"
	;;
esac