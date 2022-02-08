# Setup

You will need to create a project in Firebase, which is used for authentication. Get some admin credentials (Project Settings > Service Accounts > Generate new private key) and save the JSON file they provide in the root of the project in a file called `service-account-file.json`. You will also need the Firebase config for the front end (Project Settings > General, somewhere under Your Apps (you might need to make one) you will see a firebaseConfig object) saved in the root of the project as `public-firebase-config.json`.

Requirements:

- `docker`
- `docker-compose`

Create a `.env` file using `.env.example` as a base, changing the database username and password as you please.

Run `./run.sh prod`

View the site at [http://localhost:5004](http://localhost:5004)

# Development

Requirements:

- `docker`
- `docker-compose`
- `node` 16+

Create a `.env` file and setup Firebase as described above. Run the server with `./run.sh dev`. Any time you make changes to Go code run `docker-compose restart api` in a separate terminal window to restart the server with your changes.

The frontend in development is built and hosted by [Vite](https://vitejs.dev/). Run `npm i` in the project root to install dependencies, then change directories to `src/static` and run `npm run dev` to start the Vite dev server.

View the site at [http://localhost:3000](http://localhost:3000).

# Development Notes!

## Create a database migration

```
migrate create -ext sql -dir src/server/migrate/migrations -seq <some_migration_name>
```

## Architecture and naming conventions

There are three main layers with different responsibilities, each layer having its own types where necessary. The architecture is inspired by Clean Architecture and DDD.

### Controllers

Controllers are REST endpoint handlers. These are responsible for parsing data into a format consumable to an interactor. Like if an ID for something is in a route's arguments but is meant to be an `int`, the controller must parse it. Anything implementation details about REST endpoints should not leak into the interactors.

Controllers should have their own structs for request/response bodies, with their own json tags. Anything meant to be sent to the browser should copy necessary properties from the entities returned from the interactor and create an instance of their own response struct. An entity struct should never be the controller's response to the browser.

### Interactors

Interactors are the business logic layer between controllers. They handle validation of data and should return `data, err` (or just `err` when no data is needed), where `err` is an exported error instance so specific responses in controllers can be made by checking the kind of error returned. Saving or retrieving data from storage is done by calls to a repository.

Interactors should work with Entity structs.

### Repositories

Repositories are responsible for saving and retrieving data from the database, retrieving structs for the Entity the data represents.
