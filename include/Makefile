
all:
	@gcc -I./ -c -Wall -Werror -fpic xf_ui.cpp
	# @gcc -shared -o libxf_ui.so xf_ui.o -Wl,--whole-archive -Wl,--no-whole-archive
	@ar cru libxf_ui.so xf_ui.o
	# @g++ -o xf_ui.o -c xf_ui.cpp -shared -o libxf_ui.so -Wl,--whole-archive -lAlgatorc -Wl,--no-whole-archive
	@mv libxf_ui.so ../libs/x64

clean: *.o
	@rm $^