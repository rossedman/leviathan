# Leviathan

A simple test app.

## /health

Returns 200 OK 

```shell
❯ http :8080/health
HTTP/1.1 200 OK
Content-Length: 3
Content-Type: text/plain; charset=utf-8
Date: Wed, 23 Jun 2021 15:16:19 GMT

OK
```

## /headers

Prints the headers used back to the requester

```shell
❯ http :8080/headers
HTTP/1.1 200 OK
Content-Length: 91
Content-Type: text/plain; charset=utf-8
Date: Wed, 23 Jun 2021 15:16:23 GMT

User-Agent: HTTPie/2.4.0
Accept-Encoding: gzip, deflate
Accept: */*
Connection: keep-alive
```