PWD = $(shell pwd)

.PHONY: gen_xf.go env gen_win_h gen_win_c

build-docker: Dockerfile
	@docker build --rm -t xf .

test:
	@docker run -v$(PWD):/usr/src/myapp -it xf go test -v .

gen_exports: 
	@docker run -v$(PWD):/usr/src/myapp -it xf bash -c $$'objdump -TC ./libs/X64/libmsc.so | awk \'$$6 == "Base" {print $$7}\' > EXPORTS'

gen_xf.go: xf.y gen.py
	@python gen.py -e EXPORTS | cat xf.y - > xf.go
	@go fmt xf.go

gen_win_h:
	@python gen_win.py -e EXPORTS -r | cat xunfei.h.y - close.y > win/xunfei/xunfei.h

gen_win_c:
	@python gen_win.py -e EXPORTS | cat xunfei.cpp.y - > win/xunfei/xunfei.cpp

gen_win_go:
	@python gen_win.py -e EXPORTS -g | cat xf_windows.y - > xf_windows.go
	@go fmt xf_windows.go

gen: gen_xf.go

gen_win: gen_win_h gen_win_c

bash: 
	@docker run -v$(PWD):/usr/src/myapp -it xf bash