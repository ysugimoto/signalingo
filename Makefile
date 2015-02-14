.PHONY: static

static:
	./bin/go-bindata --pkg=env -o ./env/static.go ./conf
