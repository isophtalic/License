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

## API Reference

- ### User

```http
   /api/v1/users
```

| HTTP Method | URL                | Query                               | JSON                                                    | Description                                         |
| :---------- | :----------------- | :---------------------------------- | :------------------------------------------------------ | :-------------------------------------------------- |
| `GET`       | `/`                | page, per_page, sort                |                                                         |                                                     |
| `GET`       | `/filter`          | `role/status`, page, per_page, sort |                                                         |                                                     |
| `GET`       | `/search`          | `name/email`, page, per_page, sort  |                                                         |                                                     |
| `POST`      | `/change-password` |                                     | `email`, `password`, `new_password`, `confirm_password` | if you need to **_reset_** password with admin role |
| `POST`      | `/change-password` |                                     | `password`, `new_password`, `confirm_password`          | if you need to **_change_** password                |
| `POST`      | `/add`             |                                     | `email`, `password`, `name`, `role`, `status`           |                                                     |
| `PUT`       | `/`                |                                     | `email`, `name`, `role`, `status`                       |                                                     |

_**Required**_ : API Key

- ### Product

```http
   /api/v1/product
```

| HTTP method | URL                 | Query                | JSON                                                          | Form-data                             |
| :---------- | :------------------ | :------------------- | :------------------------------------------------------------ | :------------------------------------ |
| `GET`       | `/`                 | page, per_page, sort |                                                               |                                       |
| `GET`       | `/{id}`             |                      |                                                               |                                       |
| `GET`       | `/search`           | email                |                                                               |                                       |
| `POST`      | `/`                 |                      | `name`, `description`, `company`, `email`, `phone`, `address` |                                       |
| `PATCH`     | `/{id}`             |                      | `productId`, `description`, `company`, `email`                |                                       |
| `PATCH`     | `/status/ {id}`     |                      | status                                                        |                                       |
| `POST`      | `/{id}/key`         | type                 | status                                                        |                                       |
| `POST`      | `/{id}/key/upload/` | type                 |                                                               | **_keys_** : private.pem & public.pem |

_**Required**_ : API Key

- ### Product option

```http
   /api/v1/product-option
```

| HTTP method | URL     | Query | JSON                                                               | Form-data |
| :---------- | :------ | :---- | :----------------------------------------------------------------- | :-------- |
| `POST`      | `/`     |       | `name`, `description`, `productID`, `optionDetail: [{key, value}]` |           |
| `GET`       | `/{id}` |       |                                                                    |           |
| `PATCH`     | `/{id}` |       | `productID`, `optionDetail: [{key, value}]`                        |           |
| `DELETE`    | `/{id}` |       |                                                                    |           |
| `DELETE`    | `/{id}` |       |                                                                    |           |

_**Required**_ : API Key

- ### Option Detail

```http
   /api/v1/option-detail
```

| HTTP method | URL     | Query | JSON | Form-data |
| :---------- | :------ | :---- | :--- | :-------- |
| `DELETE`    | `/{id}` |       |      |           |

_**Required**_ : API Key

## License

[CyRadar](https://cyradar.com/)
