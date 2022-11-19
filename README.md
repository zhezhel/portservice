# Port Service

The current main purpose of this service is to import bulk data into the database.

## DEVELOPMENT

### HOW TO RUN (docker-compose)

1. Pre-requirements: you should have installed `docker compose` and `git`
2. Clone the repository by executing the next command: `git clone git@github.com:zhezhel/portservice.git`
3. Change directory to portservice: `cd portservice`
4. Run `docker compose up  --build` to start service

### HOW TO RUN (locally)

1. Pre-requirements: you should have installed `git`, `go`, `postgresql`
2. Clone repository by executing next command: `git clone git@github.com:zhezhel/portservice.git`
3. Change directory to portservice: `cd portservice`
4. Build executable by running `go build -o ./portservice ./cmd/portservice/main.go`
5. Set database connection string to environment variable `PG_DSN`.  
   Example for *nix-like OS: `export PG_DSN='postgres://admin:admin123@postgres:5432/postgres?sslmode=disable'`
6. Set location where file with ports placed to environment variable `FILE_SOURCE`.  
    Example for *nix-like OS: `export FILE_SOURCE='/Users/Andrii_Zhezhel/Projects/portservice/data/ports.json'`  
    Recommended to use absolute path to avoid issues related to execution of binary from different folders.  
7. Run `./portservice` to start service
