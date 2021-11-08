build: src/main.go
	mkdir -p bin
	go build -o bin/ydbase src/main.go

debug: src/main.go
	mkdir -p bin
	go build -o bin/debug -gcflags "-N -l" src/main.go

clean:
	rm bin/*
