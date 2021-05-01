# backend-tech-challenge-time
Backend for PentoHQ/tech-challenge-time
Built with Go, gqlgen (GraphQL) and go-pg

## Execute
Refer to [tech-challenge-time](https://github.com/OscarClemente/tech-challenge-time/tree/main) to execute the full app.

## Description
This repository contains a web app for creating, querying, updating and deleting timing sessions.
It has been built using svelte and urql/svelte for connecting to a GraphQL backend.

## Notes
* Due to trying to finish this app in one day, no unit testing was done (integration testing is available, refere to [tech-challenge-time](https://github.com/OscarClemente/tech-challenge-time/tree/main)) and some values are still hardcoded.
* There is a sleep in main so this go app loads after postgres docker has loaded, this is an ugly and dirty solution, ideally scripts should be used between docker containers to perform a status check.

## Execute only the frontend

You can simply run the following command if you have Go installed:
```sh
go run main.go
```
