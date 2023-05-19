# License Management

This repository containing the source code for a centralized license management software for many products

## Clone project

```
    $ git clone https://github.com/isophtalic/License.git
    $ cd backend
```

## Run DB in docker-compose

```
    $ docker-compse up -d
```

## Usage/Examples

- Install Golang
    Follow url : https://go.dev/dl/

```
    $ go get -v
```

- Run with debug mode

```
    $ go run cmd/main.go info
    $ go run cmd/main.go migrate
    $ go run cmd/main.go runserver
```

- Build release mode

```
    $ go build -o backend cmd/main.go
```

- Build code by docker

```
    $ docker build -t license_backend -f build/Dockerfile .
    $ docker run -it license_backend bash -c "./backend info" # test build
```
