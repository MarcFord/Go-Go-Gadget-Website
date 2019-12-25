build_web:
	/bin/bash scripts/build_web_app.sh
start_srv:
	web/web -port 3000 -server 127.0.0.1