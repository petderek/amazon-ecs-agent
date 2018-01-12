# Copyright 2014-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You
# may not use this file except in compliance with the License. A copy of
# the License is located at
#
# 	http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is
# distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF
# ANY KIND, either express or implied. See the License for the specific
# language governing permissions and limitations under the License.

PAUSE_CONTAINER_IMAGE = "amazon/amazon-ecs-pause"
PAUSE_CONTAINER_TAG = "0.1.0"
PAUSE_CONTAINER_TARBALL = "amazon-ecs-pause.tar"

# Variable to determine branch/tag of amazon-ecs-cni-plugins
ECS_CNI_REPOSITORY_REVISION=master

# Variable to override cni repository location
ECS_CNI_REPOSITORY_SRC_DIR=$(PWD)/amazon-ecs-cni-plugins


.PHONY: all gobuild static docker release certs test clean netkitten test-registry run-functional-tests gremlin benchmark-test gogenerate run-integ-tests image-cleanup-test-images pause-container get-cni-sources cni-plugins test-artifacts test-artifacts-in-docker

all: docker

# Dynamic go build; useful in that it does not have -a so it won't recompile
# everything every time
gobuild:
	./scripts/build false

# Basic go build
static:
	./scripts/build

builder-image:
	@docker build -f scripts/dockerfiles/Dockerfile.build -t "amazon/amazon-ecs-agent-build:make" .

out-dir:
	mkdir -p ./out/test-artifacts
	mkdir -p ./out/cni-plugins

# 'build-in-docker' builds the agent within a dockerfile and saves it to the ./out
# directory
build-in-docker: builder-image out-dir
	@docker run --net=none \
	  -e TARGET_OS="${TARGET_OS}" \
	  -e LDFLAGS="-X github.com/aws/amazon-ecs-agent/agent/config.DefaultPauseContainerTag=$(PAUSE_CONTAINER_TAG) \
	  -X github.com/aws/amazon-ecs-agent/agent/config.DefaultPauseContainerImageName=$(PAUSE_CONTAINER_IMAGE)" \
	  -v "$(PWD)/out:/out" \
	  -v "$(PWD):/go/src/github.com/aws/amazon-ecs-agent" \
	  "amazon/amazon-ecs-agent-build:make"

# 'docker' builds the agent dockerfile from the current sourcecode tree, dirty
# or not
docker: certs build-in-docker pause-container-release cni-plugins
	@cd scripts && ./create-amazon-ecs-scratch
	@docker build -f scripts/dockerfiles/Dockerfile.release -t "amazon/amazon-ecs-agent:make" .
	@echo "Built Docker image \"amazon/amazon-ecs-agent:make\""

# 'docker-release' builds the agent from a clean snapshot of the git repo in
# 'RELEASE' mode
docker-release: pause-container-release cni-plugins out-dir
	@docker build -f scripts/dockerfiles/Dockerfile.cleanbuild -t "amazon/amazon-ecs-agent-cleanbuild:make" .
	@docker run --net=none \
	  -e TARGET_OS="${TARGET_OS}" \
	  -v "$(PWD)/out:/out" \
	  -e LDFLAGS="-X github.com/aws/amazon-ecs-agent/agent/config.DefaultPauseContainerTag=$(PAUSE_CONTAINER_TAG) \
	  -X github.com/aws/amazon-ecs-agent/agent/config.DefaultPauseContainerImageName=$(PAUSE_CONTAINER_IMAGE)" \
	  -v "$(PWD):/src/amazon-ecs-agent" \
	  "amazon/amazon-ecs-agent-cleanbuild:make"

# Release packages our agent into a "scratch" based dockerfile
release: certs docker-release
	@./scripts/create-amazon-ecs-scratch
	@docker build -f scripts/dockerfiles/Dockerfile.release -t "amazon/amazon-ecs-agent:latest" .
	@echo "Built Docker image \"amazon/amazon-ecs-agent:latest\""

gogenerate:
	./scripts/gogenerate

# We need to bundle certificates with our scratch-based container
certs: misc/certs/ca-certificates.crt
misc/certs/ca-certificates.crt:
	docker build -t "amazon/amazon-ecs-agent-cert-source:make" misc/certs/
	docker run "amazon/amazon-ecs-agent-cert-source:make" cat /etc/ssl/certs/ca-certificates.crt > misc/certs/ca-certificates.crt

test:
	. ./scripts/shared_env && go test -race -timeout=25s -v -cover $(shell go list ./agent/... | grep -v /vendor/)

test-silent:
	. ./scripts/shared_env && go test -timeout=25s -cover $(shell go list ./agent/... | grep -v /vendor/)

benchmark-test:
	. ./scripts/shared_env && go test -run=XX -bench=. $(shell go list ./agent/... | grep -v /vendor/)

test-artifacts: out-dir
	mkdir -p ./out/test-artifacts
	go test -race -tags integration -o ./out/test-artifacts/unix-engine-tests -c ./agent/engine
	go test -race -tags integration -o ./out/test-artifacts/unix-stats-tests -c ./agent/stats
	go test -race -tags integration -o ./out/test-artifacts/unix-app-tests -c ./agent/app
	go test -tags functional -o ./out/test-artifacts/unix-simple-tests -c ./agent/functional_tests/tests/generated/simpletests_unix/
	go test -tags functional -o ./out/test-artifacts/unix-handwritten-tests -c ./agent/functional_tests/tests/

	# Race detector on windows requires CGO
	GOOS=windows CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go test -race -tags integration -o ./out/test-artifacts/windows-engine-tests.exe -c ./agent/engine
	GOOS=windows CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go test -race -tags integration -o ./out/test-artifacts/windows-stats-tests.exe -c ./agent/stats
	GOOS=windows CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go test -race -tags integration -o ./out/test-artifacts/windows-app-tests.exe -c ./agent/app
	GOOS=windows go test -tags functional -o ./out/test-artifacts/windows-simple-tests.exe -c ./agent/functional_tests/tests/generated/simpletests_windows/
	GOOS=windows go test -tags functional -o ./out/test-artifacts/windows-handwritten-tests.exe -c ./agent/functional_tests/tests/

test-artifacts-in-docker: builder-image out-dir
	@docker run --net=none \
		-v "$(PWD)/out:/out" \
		-v "$(PWD):/go/src/github.com/aws/amazon-ecs-agent" \
		"amazon/amazon-ecs-agent-build:make" \
		make test-artifacts

# Run our 'test' registry needed for integ and functional tests
test-registry: netkitten volumes-test squid awscli image-cleanup-test-images fluentd
	@./scripts/setup-test-registry

test-in-docker:
	docker build -f scripts/dockerfiles/Dockerfile.test -t "amazon/amazon-ecs-agent-test:make" .
	# Privileged needed for docker-in-docker so integ tests pass
	docker run --net=none -v "$(PWD):/go/src/github.com/aws/amazon-ecs-agent" --privileged "amazon/amazon-ecs-agent-test:make"

run-functional-tests: testnnp test-registry
	. ./scripts/shared_env && go test -tags functional -timeout=30m -v ./agent/functional_tests/...

pause-container: out-dir
	@docker build -f scripts/dockerfiles/Dockerfile.buildPause -t "amazon/amazon-ecs-build-pause-bin:make" .
	@docker run --net=none \
		-v "$(PWD)/misc/pause-container:/out" \
		-v "$(PWD)/misc/pause-container/buildPause:/usr/src/buildPause" \
		"amazon/amazon-ecs-build-pause-bin:make"

	$(MAKE) -C misc/pause-container $(MFLAGS)
	@docker rmi -f "amazon/amazon-ecs-build-pause-bin:make"

pause-container-release: pause-container out-dir
	@docker save ${PAUSE_CONTAINER_IMAGE}:${PAUSE_CONTAINER_TAG} > "$(PWD)/out/${PAUSE_CONTAINER_TARBALL}"

get-cni-sources:
	git submodule update --init --checkout

cni-plugins: get-cni-sources out-dir
	@docker build -f scripts/dockerfiles/Dockerfile.buildCNIPlugins -t "amazon/amazon-ecs-build-cniplugins:make" .
	docker run --rm --net=none \
		-e GIT_SHORT_HASH=$(shell cd $(ECS_CNI_REPOSITORY_SRC_DIR) && git rev-parse --short HEAD) \
		-e GIT_PORCELAIN=$(shell cd $(ECS_CNI_REPOSITORY_SRC_DIR) && git status --porcelain 2> /dev/null | wc -l | sed 's/^ *//') \
		-v "$(PWD)/out/cni-plugins:/go/src/github.com/aws/amazon-ecs-cni-plugins/bin/plugins" \
		-v "$(ECS_CNI_REPOSITORY_SRC_DIR):/go/src/github.com/aws/amazon-ecs-cni-plugins" \
		"amazon/amazon-ecs-build-cniplugins:make"
	@echo "Built amazon-ecs-cni-plugins successfully."

run-integ-tests: test-registry gremlin
	. ./scripts/shared_env && go test -race -tags integration -timeout=5m -v ./agent/engine/... ./agent/stats/... ./agent/app/...

codebuild: docker test-artifacts-in-docker
	docker save -o ./out/test-artifacts/agent.tar "amazon/amazon-ecs-agent:make"
	TARGET_OS="windows" ./scripts/build

netkitten:
	$(MAKE) -C misc/netkitten $(MFLAGS)

volumes-test:
	$(MAKE) -C misc/volumes-test $(MFLAGS)

# TODO, replace this with a build on dockerhub or a mechanism for the
# functional tests themselves to build this
.PHONY: squid awscli fluentd
squid:
	$(MAKE) -C misc/squid $(MFLAGS)

gremlin:
	$(MAKE) -C misc/gremlin $(MFLAGS)

awscli:
	$(MAKE) -C misc/awscli $(MFLAGS)

fluentd:
	$(MAKE) -C misc/fluentd $(MFLAGS)

image-cleanup-test-images:
	$(MAKE) -C misc/image-cleanup-test-images $(MFLAGS)

testnnp:
	$(MAKE) -C misc/testnnp $(MFLAGS)

.get-deps-stamp:
	go get golang.org/x/tools/cmd/cover
	go get github.com/golang/mock/mockgen
	go get golang.org/x/tools/cmd/goimports
	touch .get-deps-stamp

get-deps: .get-deps-stamp

clean:
	# ensure docker is running and we can talk to it, abort if not:
	docker ps > /dev/null
	rm -f misc/certs/ca-certificates.crt &> /dev/null
	rm -rf out/*
	$(MAKE) -C $(ECS_CNI_REPOSITORY_SRC_DIR) clean
	-$(MAKE) -C misc/netkitten $(MFLAGS) clean
	-$(MAKE) -C misc/volumes-test $(MFLAGS) clean
	-$(MAKE) -C misc/gremlin $(MFLAGS) clean
	-$(MAKE) -C misc/testnnp $(MFLAGS) clean
	-$(MAKE) -C misc/image-cleanup-test-images $(MFLAGS) clean
	-rm -f .get-deps-stamp

