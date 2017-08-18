SOURCE=main.go
TARGET=function.so
CHROME=chrome
INDEXJS=index.js
SHIM=cloud-functions-go-shim
FUNCTION=function
ZIP=$(FUNCTION).zip

$(TARGET): $(SOURCE)
	go build -buildmode=plugin -o $@ $<

build: $(TARGET)

rebuild: clean build

$(CHROME):
	./download-chrome.sh

$(FUNCTION): $(TARGET) $(CHROME) $(INDEXJS) $(SHIM)
	mkdir -p $@
	cp -r $(CHROME)/chrome-linux $@/
	cp $(TARGET) $@/
	cp $(INDEXJS) $@/
	cp $(SHIM) $@/

$(ZIP): $(FUNCTION)
	cd $< && zip -r ../$@ ./*

zip: $(ZIP)

clean:
	rm -rf $(TARGET) $(CHROME) $(FUNCTION) $(ZIP)

.PHONY: clean rebuild zip
