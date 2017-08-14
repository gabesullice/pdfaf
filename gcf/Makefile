SOURCE=main.go
TARGET=function.so

$(TARGET): $(SOURCE)
	go build -buildmode=plugin -o $@ $<

build: $(TARGET)

rebuild: clean build

clean:
	rm -f $(TARGET)

.PHONY: clean rebuild
