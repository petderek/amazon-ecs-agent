.PHONY: all clean
all: efs-debug efs-plugin

efs-debug: cmd/efs-debug/*.go *.go
	go build ./cmd/efs-debug

efs-plugin: cmd/efs-plugin/*.go *.go
	go build ./cmd/efs-plugin

clean:
	rm -f efs-plugin efs-debug
