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
docker compose up
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