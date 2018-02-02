main: main.go handler/handler.go
	go build -o main
main.zip: main
	zip main.zip main
clean:
	rm -f main main.zip
