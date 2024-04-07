### Start ###
1) For run app create a ".env" file with content like in the ".example.env" file.

2) Run docker postgresql:

    - Pull docker image of postgres docker pull postgres

    - Run container docker run --name=tc_auth -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d postgres

2.1) Aplpy migrations (for first run)
    - migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

3) Enter "go run cmd/main.go" to command line

### File structure ###

Directories:
1) "cmd" starts app;
2) "config" contains "public" config variables;
3) "schema" contains migration's files;
4) "internal" contains inner modules;
5) "internal/handler" handles and validates request;
6) "internal/service" makes business logic of app;
7) "internal/repository" writes and reads data from database;
8) "internal/util" contains functuins, which used in two or more modules;
9) "internal/modules/:module" contains implementation of module's handler, service and repository.

Files:
1) ".env" contains private values of config variables.
