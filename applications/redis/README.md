
### REDIS

#### DOCKER BUILD
```bash
$ docker build --no-cache -t jeffotoni/redis .
[+] Building 5.7s (5/5) FINISHED                                                                                                                                                                                                        docker:desktop-linux
 => [internal] load build definition from Dockerfile                                                                                                                                                                                                    0.0s
 => => transferring dockerfile: 476B                                                                                                                                                                                                                    0.0s
 => [internal] load metadata for docker.io/library/redis:latest                                                                                                                                                                                         4.7s
 => [internal] load .dockerignore
 ...
 View build details: docker-desktop://dashboard/build/desktop-linux/desktop-linux/vpln5blckxph241b5g5ipdtwe]
```

#### DOCKER RUN
```bash
$ docker run --rm -d -p 6379:6379 --name my-redis jeffotoni/redis
56648714d72f9422f5d1ba846b8ad81d515b743a4d0d79df9fdecd3bcedfc69c
```

#### DOCKER EXEC
```bash
$ docker exec -it my-redis redis-cli
127.0.0.1:6379> keys *
(empty array)
127.0.0.1:6379> SET key1 '{"id":2222,"name":"jeff"}'
OK
127.0.0.1:6379> keys *
1) "key1"
127.0.0.1:6379> get key1
"{\"id\":2222,\"name\":\"jeff\"}"
127.0.0.1:6379> del key1
(integer) 1
127.0.0.1:6379> HSET myhash id 39338
(integer) 1
127.0.0.1:6379> HSET myhash name "jeff"
(integer) 1
127.0.0.1:6379> HGET myhash id
"value1"
127.0.0.1:6379> HGET myhash name
"value2"
127.0.0.1:6379> keys *
1) "myhash"
127.0.0.1:6379> FLUSHALL
OK
127.0.0.1:6379> keys *
(empty array)
```