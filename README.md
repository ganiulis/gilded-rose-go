# GO Starter

- Build:

```shell
docker build -t gilded-rose .
```

- Run tests:

```shell
make test
```

- Run go inside container

```shell
make go ...
```

- Run app:

```shell
make run [<number-of-days>; default: 2]
```

- Run tests with coverage output:

```shell
make coverage
```
