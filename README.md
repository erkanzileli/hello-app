# hello-app

## Motivation

Just a test application. I am using this when I need an echo-server or an endpoint which always returns 200. Also it logs the incomings requests like the format below.

```
2021/02/14 20:41:04 GET /abc
2021/02/14 20:41:04 GET /xyz?param1=value1&param2=value2
2021/02/14 20:41:04 POST /qwe
```

## Usage

Available on [DockerHub](https://hub.docker.com/r/erkanzileli/hello-app)

Run command below for to start the hello-app

```sh
$ docker run -p 8082:8081 --rm erkanzileli/hello-app
```

Test it

```sh
$ curl localhost:8082/xyz
Hello!
```

Looking the container logs

```
2021/02/14 18:18:49 Running on :8081
2021/02/14 18:18:55 GET /xyz
```

## Contribution

If you think something can be better the please do it without any thought. Of course all contributions are welcome.
