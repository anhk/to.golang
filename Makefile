export GOPROXY=https://goproxy.cn,direct
export GO111MODULE=on

OBJ = to
INSTALL_PATH = /usr/local/bin/$(OBJ)

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Darwin)
	CODESIGN = sudo codesign --force --sign - $(INSTALL_PATH)
else
	CODESIGN =
endif

default: $(OBJ)

$(OBJ):
	go build -mod=vendor -gcflags "-N -l" -o $@ ./src

install: $(OBJ)
	cp $(OBJ) $(INSTALL_PATH)
	$(CODESIGN)

codesign:
ifeq ($(UNAME_S),Darwin)
	sudo codesign --force --sign - $(INSTALL_PATH)
endif

vendor:
	go mod vendor

clean:
	rm -fr $(OBJ)

-include .deps

dep:
	echo -n "$(OBJ):" > .deps
	find . -name '*.go' | awk '{print $$0 " \\"}' >> .deps
	echo "" >> .deps
