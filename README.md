# GO Starter

![](https://github.com/ganiulis/gilded-rose-go/actions/workflows/test.yml/badge.svg)

- Build image:

```shell
make build
```

- Run tests:

```shell
make test
```

- Run a command inside the container

```shell
make cmd args="go ..."
```

- Run the app with fixtures:

```shell
make fixtures [days=...; default: 2]
```

- Run tests with coverage output in HTML:

```shell
make coverage
```
