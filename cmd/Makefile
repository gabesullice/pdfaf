SOURCE=main.go
TARGET=snap.o

$(TARGET): $(SOURCE)
	go build -o $@

build: snap.o

rebuild: clean build

clean:
	rm -f $(TARGET)

.PHONY: clean rebuild
