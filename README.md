# Golang Bootcamp

This project is the challenge that was requested in order to get certified with the academy-go-q12021 badge. It is a very simple API that implements an approach to clean architecture, consumes a REST API, reads and writes in a CSV and also implements concurrency with go routines.

In this example, we are using Rick & Morty's characters for our data sample.

## Running the project

- Get/update dependencies by running:

``` bash
go mod tidy
```

- Run the project:

``` bash
go run main.go
```

## Basic endpoints

- Get all characters in the CSV file:

``` bash
GET <hostname>:<port>/characters/
```

- Get a single character by its ID:

``` bash
GET <hostname>:<port>/character/{id}
```

- Get a character from an external API by its ID and write it into the CSV file:

``` bash
GET <hostname>:<port>/character/external/{id}
```

- Get all characters concurrently:

``` bash
GET <hostname>:<port>/characters/concurrent/
```
