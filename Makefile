PWD = $(shell pwd)

.PHONY: gen_xf.go env

build-docker: Dockerfile
	@docker build --rm -t xf .

test:
	@docker run -v$(PWD):/usr/src/myapp -it xf go test -v .

gen_exports: 
	@docker run -v$(PWD):/usr/src/myapp -it xf bash -c $$'objdump -TC ./libs/X64/libmsc.so | awk \'$$6 == "Base" {print $$7}\' > EXPORTS'

gen_xf.go: xf.y gen.py
	@python gen.py -e EXPORTS | cat xf.y - > xf.go
	@go fmt xf.go

gen: gen_xf.go

bash: 
	@docker run -v$(PWD):/usr/src/myapp -it xf bash