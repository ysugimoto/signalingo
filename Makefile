.PHONY: static

static:
	./bin/go-bindata --pkg=static -o ./static/asset.go ./etc ./public
