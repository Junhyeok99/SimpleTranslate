all:
	go get cloud.google.com/go/translate
	go get golang.org/x/text/language

build: main.go
	go build -o trans.a main.go

clean: trans.a
	rm trans.a