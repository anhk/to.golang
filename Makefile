export GOPROXY=https://goproxy.cn,direct
export GO111MODULE=on

OBJ = to

default: $(OBJ)

$(OBJ):
	go mod tidy && go build -gcflags "-N -l" -o ${OBJ} ./

clean:
	rm -fr $(OBJ)

-include .deps
dep:
	echo -n "$(OBJ):" > .deps
	find . -path ./vendor -prune -o -name '*.go' -print | awk '{print $$0 " \\"}' >> .deps
	echo "" >> .deps