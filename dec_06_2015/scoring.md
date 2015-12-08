# How we score
Listed below are the test cases, captured as HTTP request & responses using ngrep.

We run the same request through your service, and check for the response.
The final score is caculated form that.

To make it transparent to participants, we will publish the result of running
our test driver on your service, together with an ngrep dump of all HTTP requests
cominng / going from it.

# Changes

There has been some changes to the scoring table:

- TTL test case (5pts) was not feasible to implement (because the round-trip time
between our client tester & your service will make the returned TTL time incorrect).
Thus, we will not give the 5pts of this test case for anyone
- Some commands are broken down into more tests cases (like invalid key type, invalid value). However, the total score of that command remain unchanged

****

### FLUSHDB (5pts)
`FLUSHDB`

#### FLUSHDB clear all key (5pts)
```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"SET s a"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 23.
Accept-Encoding: gzip.
.
{"command":"RPUSH k b"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 14.
Content-Type: text/plain; charset=utf-8.
.
{"response":1}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 22.
Accept-Encoding: gzip.
.
{"command":"SADD z c"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 19.
Accept-Encoding: gzip.
.
{"command":"GET s"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 20.
Content-Type: text/plain; charset=utf-8.
.
{"response":"EKTYP"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 19.
Accept-Encoding: gzip.
.
{"command":"GET k"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 20.
Content-Type: text/plain; charset=utf-8.
.
{"response":"EKTYP"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 19.
Accept-Encoding: gzip.
.
{"command":"GET z"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 20.
Content-Type: text/plain; charset=utf-8.
.
{"response":"EKTYP"}
```

### EXPIRE (5pts)
`EXPIRE key timeToLive`

#### EXPIRE key after 30 seconds (5pts)
```
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 30.
Accept-Encoding: gzip.
.
{"command":"SET whoareyou me"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 33.
Accept-Encoding: gzip.
.
{"command":"EXPIRE whoareyou 30"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 27.
Accept-Encoding: gzip.
.
{"command":"GET whoareyou"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"me"}
## // wait 5 seconds
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 27.
Accept-Encoding: gzip.
.
{"command":"GET whoareyou"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:58 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"me"}
## // wait 30seconds
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 27.
Accept-Encoding: gzip.
.
{"command":"GET whoareyou"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:53:33 GMT.
Content-Length: 20.
Content-Type: text/plain; charset=utf-8.
.
{"response":"EKTYP"}
```

### DELETE (5pts)
`DEL key`

#### DELETE a key (5pts)
```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:53:33 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 25.
Accept-Encoding: gzip.
.
{"command":"SET foo bar"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:53:33 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"DEL foo"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:53:33 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"GET foo"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:53:33 GMT.
Content-Length: 20.
Content-Type: text/plain; charset=utf-8.
.
{"response":"EKTYP"}
```

## String

### SET (5pts)
`SET key value`

#### basic SET (5pts)
```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}           
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 31.
Accept-Encoding: gzip. 
.
{"command":"SET name grokking"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
```

### GET (5pts)
`GET key`

#### basic GET (4pts)
```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 31.
Accept-Encoding: gzip.
.
{"command":"SET name grokking"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 34.
Accept-Encoding: gzip.
.
{"command":"SET name engineering"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 22.
Accept-Encoding: gzip.
.
{"command":"GET name"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 26.
Content-Type: text/plain; charset=utf-8.
.
{"response":"engineering"}
```

##### GET non-existing key (1pts)
```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 26.
Accept-Encoding: gzip.
.
{"command":"GET notfound"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 20.
Content-Type: text/plain; charset=utf-8.
.
{"response":"EKTYP"}
```

## List

### RPUSH (4pts)
`RPUSH key value [value ...]`

#### RPUSH to a list (3pts)
```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
#
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 35.
Accept-Encoding: gzip.
.
{"command":"RPUSH programmers huy"}
#
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 14.
Content-Type: text/plain; charset=utf-8.
.
{"response":1}
#
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 52.
Accept-Encoding: gzip.
.
{"command":"RPUSH programmers loc khiem minh khanh"}
#
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 14.
Content-Type: text/plain; charset=utf-8.
.
{"response":5}
```

#### RPUSH to a non-list key (1pts)
```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 32.
Accept-Encoding: gzip.
.
{"command":"SET programmers me"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 35.
Accept-Encoding: gzip.
.
{"command":"RPUSH programmers huy"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 20.
Content-Type: text/plain; charset=utf-8.
.
{"response":"EKTYP"}
```

### LPOP (4pts)
`LPOP key`

#### LPOP a list (3pts)
```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 52.
Accept-Encoding: gzip.
.
{"command":"RPUSH programmers loc khiem minh khanh"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 14.
Content-Type: text/plain; charset=utf-8.
.
{"response":4}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 30.
Accept-Encoding: gzip.
.
{"command":"LPOP programmers"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 18.
Content-Type: text/plain; charset=utf-8.
.
{"response":"loc"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 36.
Accept-Encoding: gzip.
.
{"command":"LRANGE programmers 0 2"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 37.
Content-Type: text/plain; charset=utf-8.
.
{"response":["khiem","minh","khanh"]}
```

#### LPOP non-existing key (1pts)
```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 27.
Accept-Encoding: gzip.
.
{"command":"LPOP notfound"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 19.
Content-Type: text/plain; charset=utf-8.
.
{"response":"EKTYP"}
```

### RPOP (4pts)
`RPOP key`

#### RPOP a list (3pts)
```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 52.
Accept-Encoding: gzip.
.
{"command":"RPUSH programmers loc khiem minh khanh"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 14.
Content-Type: text/plain; charset=utf-8.
.
{"response":4}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 30.
Accept-Encoding: gzip.
.
{"command":"RPOP programmers"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 20.
Content-Type: text/plain; charset=utf-8.
.
{"response":"khanh"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 36.
Accept-Encoding: gzip.
.
{"command":"LRANGE programmers 0 2"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 35.
Content-Type: text/plain; charset=utf-8.
.
{"response":["loc","khiem","minh"]}
```

#### RPOP non-existing key (1pts)
```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
#
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
#
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 27.
Accept-Encoding: gzip.
.
{"command":"RPOP notfound"}
#
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 19.
Content-Type: text/plain; charset=utf-8.
.
{"response":"EKTYP"}
```

### LRANGE (4pts)
`LRANGE key start stop`

#### LRANGE a list (2pts)
```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 59.
Accept-Encoding: gzip.
.
{"command":"RPUSH languages ruby python c cpp golang java"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 14.
Content-Type: text/plain; charset=utf-8.
.
{"response":6}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 34.
Accept-Encoding: gzip.
.
{"command":"LRANGE languages 0 1"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 30.
Content-Type: text/plain; charset=utf-8.
.
{"response":["ruby","python"]}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 34.
Accept-Encoding: gzip.
.
{"command":"LRANGE languages 2 4"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 33.
Content-Type: text/plain; charset=utf-8.
.
{"response":["c","cpp","golang"]}
```

#### LRANGE non-existing list (1pts)
```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 33.
Accept-Encoding: gzip.
.
{"command":"LRANGE notfound 0 1"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 19.
Content-Type: text/plain; charset=utf-8.
.
{"response":"EKTYP"}
```

#### LRANGE with invalid range (1pts)
```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 34.
Accept-Encoding: gzip.
.
{"command":"RPUSH languages ruby"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 14.
Content-Type: text/plain; charset=utf-8.
.
{"response":1}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 34.
Accept-Encoding: gzip.
.
{"command":"LRANGE languages 2 4"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 19.
Content-Type: text/plain; charset=utf-8.
.
{"response":"EINV"}
```

### LLEN (4pts)
`LLEN key`

#### LLEN a list (3pts)
```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 52.
Accept-Encoding: gzip.
.
{"command":"RPUSH programmers loc khiem minh khanh"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 14.
Content-Type: text/plain; charset=utf-8.
.
{"response":4}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 30.
Accept-Encoding: gzip.
.
{"command":"LLEN programmers"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 14.
Content-Type: text/plain; charset=utf-8.
.
{"response":4}
```

#### LLEN non-existing key (1pts)
```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 27.
Accept-Encoding: gzip.
.
{"command":"LLEN notfound"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 19.
Content-Type: text/plain; charset=utf-8.
.
{"response":"EKTYP"}
```

## Set

### SADD (4pts)
`SADD key value [value ...]`

#### SADD creates a set (3pts)
```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 46.
Accept-Encoding: gzip.
.
{"command":"SADD asean vietnam lao campuchea"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
```

#### SADD wrong key type (1pts)
```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 29.
Accept-Encoding: gzip.
.
{"command":"SET asean asean"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 46.
Accept-Encoding: gzip.
.
{"command":"SADD asean vietnam lao campuchea"}
#
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 20.
Content-Type: text/plain; charset=utf-8.
.
{"response":"EKTYP"}
```

### SREM (4pts)
`SREM key value [value ...]`

#### SREM removes members from a set (4pts)
```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 34.
Accept-Encoding: gzip.
.
{"command":"SADD myset a b c d e"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 28.
Accept-Encoding: gzip.
.
{"command":"SREM myset a b"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
```

### SCARD (4pts)
`SCARD key`

#### SCARD queries a set size (3pts)
```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 32.
Accept-Encoding: gzip.
.
{"command":"SADD myset a b c d"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 25.
Accept-Encoding: gzip.
.
{"command":"SCARD myset"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 14.
Content-Type: text/plain; charset=utf-8.
.
{"response":4}
```

#### SCARD wrong key type (1pts)
```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
#
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
#
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 34.
Accept-Encoding: gzip.
.
{"command":"RPUSH mylist a b c d"}
#
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 14.
Content-Type: text/plain; charset=utf-8.
.
{"response":4}
#
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 26.
Accept-Encoding: gzip.
.
{"command":"SCARD mylist"}
#
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 20.
Content-Type: text/plain; charset=utf-8.
.
{"response":"EKTYP"}
```

### SINTER (4pts)
`SINTER key [key ...]`

#### SInter intersects sets (4pts)
```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 35.
Accept-Encoding: gzip.
.
{"command":"SADD myseta a b c d e"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 31.
Accept-Encoding: gzip.
.
{"command":"SADD mysetb a b c"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 27.
Accept-Encoding: gzip.
.
{"command":"SADD mysetc a"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 41.
Accept-Encoding: gzip.
.
{"command":"SINTER myseta mysetb mysetc"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 18.
Content-Type: text/plain; charset=utf-8.
.
{"response":["a"]}
```

### SMEMBERS (4pts)
`SMEMBERS key`

#### SMEMBERS queries a set's members (2pts)
```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 34.
Accept-Encoding: gzip.
.
{"command":"SADD myset a b c d e"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 28.
Accept-Encoding: gzip.
.
{"command":"SMEMBERS myset"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 34.
Content-Type: text/plain; charset=utf-8.
.
{"response":["b","c","d","e","a"]}
```

### SREM and then SMEMBERS to query result (2pts)
```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 34.
Accept-Encoding: gzip.
.
{"command":"SADD myset a b c d e"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 28.
Accept-Encoding: gzip.
.
{"command":"SREM myset a b"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 28.
Accept-Encoding: gzip.
.
{"command":"SMEMBERS myset"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:52:53 GMT.
Content-Length: 26.
Content-Type: text/plain; charset=utf-8.
.
{"response":["c","d","e"]}
```

## Save and restore

### SAVE & RESTORE (30pts)
`SAVE`
`RESTORE`

```
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:53:33 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 31.
Accept-Encoding: gzip.
.
{"command":"SET name grokking"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:53:33 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 47.
Accept-Encoding: gzip.
.
{"command":"RPUSH colors green blue red white"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:53:33 GMT.
Content-Length: 14.
Content-Type: text/plain; charset=utf-8.
.
{"response":4}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 54.
Accept-Encoding: gzip.
.
{"command":"SADD countries vietnam singapore vietnam"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:53:33 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 18.
Accept-Encoding: gzip.
.
{"command":"SAVE"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:53:33 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"FLUSHDB"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:53:33 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 21.
Accept-Encoding: gzip.
.
{"command":"RESTORE"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:53:33 GMT.
Content-Length: 17.
Content-Type: text/plain; charset=utf-8.
.
{"response":"OK"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 22.
Accept-Encoding: gzip.
.
{"command":"GET name"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:53:33 GMT.
Content-Length: 23.
Content-Type: text/plain; charset=utf-8.
.
{"response":"grokking"}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 31.
Accept-Encoding: gzip.
.
{"command":"LRANGE colors 0 3"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:53:33 GMT.
Content-Length: 43.
Content-Type: text/plain; charset=utf-8.
.
{"response":["green","blue","red","white"]}
##
T 127.0.0.1:34024 -> 127.0.0.1:8080 [AP]
POST /ledis HTTP/1.1.
Host: 127.0.0.1:8080.
User-Agent: Go 1.1 package http.
Content-Length: 32.
Accept-Encoding: gzip.
.
{"command":"SMEMBERS countries"}
##
T 127.0.0.1:8080 -> 127.0.0.1:34024 [AP]
HTTP/1.1 200 OK.
Date: Tue, 08 Dec 2015 01:53:33 GMT.
Content-Length: 36.
Content-Type: text/plain; charset=utf-8.
.
{"response":["vietnam","singapore"]}
```
