# Pico HTTP Server in C

This is a very simple HTTP server for Unix, using `fork()`. It's very easy to use.

## How to use

1. Include header `httpd.h`.
2. Write your route method, handling requests.
3. Call `serve_forever("8000")` to start serving on http://127.0.0.1:8000/.

See `main.c`, an interesting example.

To log stuff, use `fprintf(stderr, "message");`

View `httpd.h` for more information.

## Quick start

1. Run `make`.
2. Run `./server` or `./server [port]` (port = 8000 by default).
3. Open http://localhost:8000/ or http://localhost:8000/test in browser to see request headers.

## Testing and benchmarking

I suggest using [Siege](https://github.com/JoeDog/siege) utility for testing and benchmarking the Pico HTTP server.

```sh
> siege -i -f urls.txt
```

