# Calculogic
Victoria Tilley's FYP project for 2024-2025. Contains all the necessary code and scripts to run the project.

# What is Calculogic
A full-stack application built on React and Go in order to help students solve their logical problems, with a particular focus on classical natural deduction and sequent calculus.

# Prerequisites
If you only wish to view the project:
- Docker
If you want to make hot-reloading changes to the project:
- Bun v1.2.4
- Golang v1.23.5

# How to run
Clone the repository, open a terminal at the root directory and run:
```bash
docker compose up --build
```

There is also a requirement to have some environment variables set for the server in an .env file to avoid leaking secrets. These variables are:
```yaml
DB_HOST=<name>
DB_PORT=3306
DB_USER=<user>
DB_PASSWORD=<password>
DB_NAME=<name>
```

There is also a requirement for there to be a `db.env` file in the root directory next to the `docker-compose.yml` file which contains database configuration. These need to be the following:
```yaml
MYSQL_DATABASE=<DB_NAME>
MYSQL_USER=<DB_USER>
MYSQL_PASSWORD=<DB_PASSWORD>
MYSQL_ROOT_PASSWORD=<root password>
```


To run the individual services separately outside of docker containers, perform the following:

## Client
Navigate into `/client` and run the following command:
```bash
bun run dev
```
Which should open a Vite server on `localhost:5173` and should be navigatable on your browser.

## Server
Navigate into `/server` and run the following command:
```bash
go run server.go
```
Which should open the Echo server on `localhost:1323`.