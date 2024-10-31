.PHONY: build run clean

SOURCES = main.go character.go

build:
	go build -o DateHack.exe $(SOURCES)

run:
	go run $(SOURCES)

clean:
	go clean
	rm -f DateHack.exe
