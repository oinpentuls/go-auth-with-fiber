# Golang authentication using Fiber and JWT

## Setup

###
Create database and set the config in `.env` file  
run sql script

```bash
mysql -u <username> -p <database_name> < ./schema.sql
```

###
Get dependecy modules
```bash
go mod tidy
```

## Run

###
```bash
go run .
```

OR

```bash
go build
```

and then run the executable file