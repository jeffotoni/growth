all: server

clean:
	@rm -rf *.o
	@rm -rf server

static: 
	gcc --static -Wimplicit-int -o server main.c httpd.c

server: main.o httpd.o
	gcc -o server $^

main.o: main.c httpd.h
	gcc -c -o main.o main.c

httpd.o: httpd.c httpd.h
	gcc -c -o httpd.o httpd.c

