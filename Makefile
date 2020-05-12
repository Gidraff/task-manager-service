# This is adapted from https://github.com/thockin/go-build-template
BIN := api

REGISTRY ?= gidraff

VERSION := $(shell git describe --tags --always --dirty)


###
### The variables should not need tweaking.
###


SRC_DIRS := cmd pkg # Dirs which hold app source

ALL_PLATFORMS := linux/amd64

# Used internally. Users should pass GOOS and/or GOARCH
OS := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))

###
### Docker image variables
###

BASEIMAGE ?= gcr.io/distroless/static

IMAGE := $(REGISTRY)/$(BIN)
TAG := $(VERSION)__$(OS)_$(ARCH)

BUILD_IMAGE ?= golang:1.13-alpine

TEST_IMAGE ?= gidraff/golang:1.13-alpine-test

# If you want to build all binaries, see the 'all-build' rule.
# If you want to build all containers, see the 'all-containers' rule.

all: build

build-%:
	@$(MAKE) build 								\
		--no-print-directory					\
		GOOS=$(firstword $(subst _, ,$*)) 		\
		GOARCH=$(lastword $(subst _, ,$*))

container-%:
	@$(MAKE) container							\
		--no-print-directory					\
		GOOS=$(firstword $(subst -, ,$*)) 		\
		GOARCH=$(lastword $(subst -, ,$*))		\

push-%:
	@$(MAKE) push								\
		--no-print-directory					\
		GOOS=$(firstword $(subst -, ,$*)) 		\
		GOARCH=$(lastword $(subst -, ,$*))		\

all-build: $(addprefix build-, $(subst /,_, $(ALL_PLATFORMS))) \

all-container: $(addprefix container-, $(subst /,_, $(ALL_PLATFORMS))) \


build: bin/$(OS)_$(ARCH)/$(BIN) \

# Required directories for build/test
BUILD_DIRS := bin/$(OS)_$(ARCH)				\
							.go/bin/$(OS)_$(ARCH) \
							.go/cache

OUTBIN = bin/$(OS)_$(ARCH)/$(BIN)
$(OUTBIN): .go/$(OUTBIN).stamp
	@true

# .PHONY: foo
# foo: ;@echo $@

.PHONY: .go/$(OUTBIN).stamp
.go/$(OUTBIN).stamp: $(BUILD_DIRS)
	@echo "making $(OUTBIN)"
	@docker run 					\
			-i						\
			--rm 					\
			-u $$(id -u):$(id -g) 	\
			-v $$(pwd):/src 		\
			-w /src 				\
			-v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin \
			-v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin/$(OS)_$(ARCH) \
			-v $$(pwd)/.go/cache:/.cache 				\
			--env HTTP_PROXY=$(HTTP_PROXY) 				\
			--env HTTPS_PROXY=$(HTTPS_PROXY) 			\
			$(BUILD_IMAGE) 								\
			/bin/sh -c "								\
				ARCH=$(ARCH)							\
				OS=$(OS) 								\
				VERSION=$(VERSION)						\
				./build/build.sh						\
			"
	@if ! cmp -s .go/$(OUTBIN) $(OUTBIN); then 			\
		mv .go/$(OUTBIN) $(OUTBIN);						\
		date >$@;										\
	fi

# Example make shell CMD="-c 'date > datafile'"
shell: $(BUILD_DIRS)
	@echo "launching a shell in the container build environment"
	@docker run 										\
			-ti 										\
			--rm 										\
			-u $$(id -u):$$(id -g) 						\
			-v $$(pwd):/src 							\
			-w /src 									\
			-v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin 	\
			-v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin/$(OS)_$(ARCH)    \
			-v $$(PWD)/.go/cache:/.cache				\
			--env HTTP_PROXY=$(HTTP_PROXY) 				\
			--env HTTPS_PROXY=$(HTTPS_PROXY) 			\
			$(BUILD_IMAGE) 								\
			/bin/sh $(CMD)

# Used to track state in hidden files
DOTFILE_IMAGE = $(subst /,_,$(IMAGE)-$(TAG))

container: .container-$(DOTFILE_IMAGE) say_container_name
.container-$(DOTFILE_IMAGE): bin/$(OS)_$(ARCH)/$(BIN) in.Dockerfile
	@sed										\
			-e 's|{ARG_BIN}|$(BIN)|g' 			\
			-e 's|{ARG_ARCH}|$(ARCH)|g' 		\
			-e 's|{ARG_OS}|$(OS)|g' 			\
			-e 's|{ARG_FROM}|$(BASEIMAGE)|g' 	\
			in.Dockerfile > .dockerfile-$(OS)_$(ARCH)
	@docker build -t $(IMAGE):$(TAG) -t $(IMAGE):latest -f .dockerfile-$(OS)_$(ARCH) .
	@docker images -q $(IMAGE):$(TAG) > $@

say_container_name:
				@echo "container: $(IMAGE):$(TAG)"

push: .push-$(DOTFILE_IMAGE) say_push_name
.push-$(DOTFILE_IMAGE): .container-$(DOTFILE_IMAGE)
				@docker push $(IMAGE):$(TAG)

push-latest: .push-$(DOTFILE_IMAGE) say_push_name_latest
				@docker push $(IMAGE):latest

say_push_name:
	@echo "pushed: $(IMAGE):$(TAG)"

say_push_name_latest:
	@echo "pushed: $(IMAGE):$(TAG)"

version:
	@echo $(VERSION)


test: $(BUILD_DIRS)
	@docker run 								\
			-i 									\
			--rm 								\
			-u $$(id -u):$$(id -g) 				\
			-v $$(pwd):/src 					\
			-w /src 							\
			-v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin	\
			-v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin/$(OS)_$(ARCH) 	\
			-v $$(pwd)/.go/cache:/.cache 				\
			-v $$(pwd)/config:/config 					\
			-v $$(pwd)/cmd/taskman/test_data:/test_data	\
			--env HTTP_PROXY=$(HTTP_PROXY) 				\
			--env HTTPS_PROXY=$(HTTPS_PROXY) 			\
			--env TASKMAN_API_KEY=DUMMY 				\
			--env TASKMAN_DSN=DUMMY 					\
			$(TEST_IMAGE) 								\
			/bin/sh -c " 								\
					ARCH=$(ARCH) 						\
					OS=$(OS)							\
					VERSION=$(VERSION) 					\
					./build/test.sh $(SRC_DIRS)			\
			"

ci: $(BUILD_DIRS)
	@docker run 										\
			-i 											\
			--rm 										\
			-u $$(id -u):$$(id -g) 						\
			-v $$(pwd):/src 							\
			-w /src 									\
			-v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin 	\
			-v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin/$(OS)_$(ARCH) 	\
			-v $$(pwd)/.go/cache:/.cache				\
			-v $$(pwd)/reports:/reports 				\
			-v $$(pwd)/config:/config					\
			-v $$(pwd)/cmd/taskman/test_data:/test_data \
			-v $$(pwd)/:/coverage 						\
			--env HTTP_PROXY=$(HTTP_PROXY) 				\
			--env HTTPS_PROXY=$(HTTPS_PROXY) 			\
			--env TASKMAN_API_KEY=DUMMY 				\
			--env TASKMAN_DSN=DUMMY 					\
			$(TEST_IMAGE) 								\
			/bin/sh -c "								\
					ARCH=$(ARCH) 						\
					OS=$(OS) 							\
					VERSION=$(VERSION) 					\
					./build/test_ci.sh $(SRC_DIRS) 		\
			"



$(BUILD_DIRS):
	@mkdir -p $@

clean:  container-clean bin-clean

container-clean:
	rm -rf .container-* .dockerfile-* .push-*

bin-clean:
	rm -rf .go bin
