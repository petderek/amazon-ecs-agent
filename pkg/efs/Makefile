.PHONY: all clean
all: efs-debug efs-cli efs-plugin

efs-debug: cmd/efs-debug/*.go
	go build ./cmd/efs-debug

efs-cli: cmd/efs-cli/*.go
	go build ./cmd/efs-cli

efs-plugin: cmd/efs-plugin/*.go
	go build ./cmd/efs-plugin

clean:
	rm -f efs-cli efs-plugin efs-debug
