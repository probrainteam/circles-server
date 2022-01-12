#!/bin/bash

cmd=$1
where=$2
opt=$3
case "$cmd" in
	-up) # web server
	case "$2" in
		express)
			cd express-server
			npm run $3 # check express-server/package.json
			;;
		gin)
			cd gin-server 
			go run main.go $3
			;;
		*)
		echo "'$2' is unknown server"
		;;
	esac
	;;
	-reset) # DB reset
		if [ "$3" == "" ] ; then
			if [ "$2" == "dev" ] ; then	# rm dev db volumn
				docker compose down dev_db redis -v
				rm -rf dev/data/
			elif [ "$2" == "" ]; then # rm real db volumn
				docker compose down db redis -v
				rm -rf db/data/
			fi
		else
				echo "'$2' and '$3'"

		fi
	;;
	-init) # DB, redis compose up
		if [ "$2" == "dev" ] ; then # dev, mock dummy db 
			docker compose up -d dev_db redis 
		elif [ "$2" == "" ] ; then  # real db
			docker compose up -d db redis 
		else
			echo "'$2' is unknown option"
		fi
	;;
	-down) # DB, redis compose down
		if [ "$2" == "" ] ; then
			docker compose down
		else
			echo "-down : No option"
		fi
	;;
	*)
	echo "'$cmd' is unknown command + '$2' and '$3'"
	;;
esac