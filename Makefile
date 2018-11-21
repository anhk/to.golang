

all: to

-include .deps

to:
	cd src && go build -o ../$@

dep deps:
	echo "to: \\" > .deps
	find src -name '*.go' | awk '{print $$0 " \\"}' >> .deps

clean:
	rm -fr to
