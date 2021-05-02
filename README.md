# backend-tech-challenge-time
Backend for PentoHQ/tech-challenge-time

## Description
This repository contains the backend for web app for creating, querying, updating and deleting timing sessions.
It has been build using Go, gqlgen (GraphQL), go-pg (Postgres) and go-chi.
By default, it is accesible at
```
http://localhost:8080/query
```
I have also left the GraphQL playground initialized at:
```
http://localhost:8080/
```
for quick testing of the backend.

## Notes
* Due to trying to finish this app in one day, no unit testing was done (integration testing is available, refer to [tech-challenge-time](https://github.com/OscarClemente/tech-challenge-time/tree/main)).
* There is a sleep in main so this go app loads after postgres docker has loaded, this is an ugly and dirty solution, ideally scripts should be used between docker containers to perform a status check.
* Some values are hardcoded, such as DB table and password. They could be easily "softcoded" but I decided to focus on other things during the one day alloted time.
* This backend connects to the DB and if it detects there are no timing sessions stored in the DB it will insert a few timing sessions in the DB. This way it will be possible to test the date filter, as new timers are created with a today date this makes it easier to test this feature.
* The GraphQL API does currently not support subscriptions, I thought about adding them but I realized they were not needed for the frontend with how it was implemented.
* I also wanted to add CI, but of course with no unit tests and having to spend only a day on this, I decided against it.

## Execute
Refer to [tech-challenge-time](https://github.com/OscarClemente/tech-challenge-time/tree/main) to execute the full app.

## Execute only the backend

You can simply run the following command if you have Go installed, this software however requires a DB to be running:
```sh
go run main.go
```
