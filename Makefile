run:
	go run main.go 8090
clean:
# 杀死服务进程
	kill $(ps aux | grep '[g]o run main.go' | awk '{print $2}')
	kill $(lsof -i:8090|grep 8090|awk '{print $2}')