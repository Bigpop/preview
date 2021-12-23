run:
	go run main.go  &
	sh ./script/open_browser.sh http://localhost:8090

clean:
	kill -9 $(ps -ef | grep 'go run main.go' | awk '{print $2}')
