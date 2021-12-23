run:
	go run main.go  &
	sh ./script/open_browser.sh http://localhost:8090

clean:
# 杀死服务进程
	kill $(ps aux | grep '[g]o run main.go' | awk '{print $2}')

