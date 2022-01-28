# Development Notes!

## Create a database migration

```
migrate create -ext sql -dir src/server/migrate/migrations -seq <some_migration_name>
```

## Architecture and naming conventions

There are three main layers with different responsibilities, each layer having its own types where necessary.

### Controllers

Controllers are REST endpoint handlers. These are responsible for parsing data into a format consumable to an interactor. Like if an ID for something is in a route's arguments but is meant to be an `int`, the controller must parse it. Anything implementation details about REST endpoints should not leak into the interactors.

Controllers should have their own structs for request/response bodies, with their own json tags. Anything meant to be sent to the browser should copy necessary properties from the entities returned from the interactor and create an instance of their own response struct. An entity struct should never be the controller's response to the browser.

### Interactors

Interactors are the business logic layer between controllers. They handle validation of data and should return `data, err` (or just `err` when no data is needed), where `err` is an exported error instance so specific responses in controllers can be made by checking the kind of error returned. Saving or retrieving data from storage is done by calls to a repository.

Interactors should work with Entity structs.

### Repositories

Repositories are responsible for saving and retrieving data from the database, retrieving structs for the Entity the data represents.
