# DotEnv in Go

Read .env file and load environment variables. This libaray allows to list required env variables and validate from all loaded environment variables. 

Example for valid .env file names:
```text
.env
local.env
.env.local
```

## Installation

Library:
```shell
go get github.com/soumayg9673/dotenv
```

## Usage

Add configuration file .env file in the root of your project or know the relative path for the file:
```shell
SERVER_ADDR=8080
SERVER_ENV=development
```

In your Go application, load the .env file
```go
package main

import(
    "log"
    
    "github.com/soumayg9673/soumayg9673"
)

func main() {
    dotenv.AddRqdKey("SERVER_ADD", true)
    dotenv.AddRqdKey("SERVER_ENV", true)

    if err := dotenv.LoadEnvFile(".env"); err != nil {
        log.Println(err)
        return
    }

    fmt.Println(dotenv.GetInt("SERVER_ADD", ":8000"))
    fmt.Println(dotenv.GetString("SERVER_ENV", "local"))

    if err := dotenv.ValidateRqdEnv(); err != nil {
		log.Println(err)
        return
	}
}
```

### Set custom eror messages for required env variables

We have exposed API to customize the error messages for each scenarios:

#### 1. Key not found
```go
func main() {
    // Usage of constant dotenv.ERR_NO_KEY
    dotenv.SetErrorMsg("%s key not found", dotenv.ERR_NO_KEY)
}
```

#### 3. Key-value not found
```go
func main() {
    // Usage of constant dotenv.ERR_NO_KEY_VALUE
    dotenv.SetErrorMsg("%s key-value pair not found", dotenv.ERR_NO_KEY_VALUE)
}
```

#### 3. Value not found
```go
func main() {
    // Usage of constant dotenv.ERR_NO_VALUE
    dotenv.SetErrorMsg("value for key %s not found", dotenv.ERR_NO_VALUE)
}
```

## Guide for using .env file for secrets

There are multiple ways to use .env files for secrets and following are the best practices being followed by developers.

#### 1. Use .env file for all env variables secrets
```shell
In .env file:
SERVER_ADDR=8080
SERVER_ENV=prod
PSQL_DB_NAME=postgres
PSQL_DB_HOST=localhost
PSQL_DB_PORT=5432
PSQL_DB_USER=postgres
PSQL_DB_PWD=admin
```

#### 2. Use .env file for high level configuration and .env.env file for secrets
```shell
In .env file:
SERVER_ADDR=8080
SERVER_ENV=prod

In .env.prod file:
PSQL_DB_NAME=postgres
PSQL_DB_HOST=localhost
PSQL_DB_PORT=5432
PSQL_DB_USER=postgres
PSQL_DB_PWD=admin
```

## Contribution

Contributions are highly recommended. 

Raise an issue for the following:
```text
- Queries
- Contributing to library
```
