.PHONY: static

static:
	./bin/go-bindata --pkg=signaling -o ./static.go ./conf
