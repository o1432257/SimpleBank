# SimpleBank
Golang RESTful HTTP Json Api Practice

Gin + migrate + sqlc + mock + go-test + jwt + paseto + viper


### Tools

- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

### Setup and Start

- Start postgres container:

    ```bash
    make postgres
    ```

- Create simple_bank database:

    ```bash
    make createdb
    ```

- Run db migration up all versions:

    ```bash
    make migrateup
    ```

- Start Server

    ```bash
    make server
    ```
  