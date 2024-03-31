### Architecture ###

Backend part of time-control consists of 3 parts (repositories):
1) time-control (write service, core)

    https://github.com/Le0nar/time-control

2) time-control-auth 

    https://github.com/Le0nar/time-control-auth

3) time-control-read

    https://github.com/Le0nar/time-control-read


### Start ###
1) For run app create a ".env" file with content like in the ".example.env" file.

2) Run docker postgresql:

    - Pull docker image of postgres docker pull postgres

    - Run container docker run --name=tc-write -e POSTGRES_PASSWORD='qwerty' -p 5435:5432 -d postgres

    - Install extension for uuid. Run in db cli: CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

3) Create 2 elements in "activity_event_type" table.
    - INSERT INTO activity_event_type (event_type) values ('check_activity');
    - INSERT INTO activity_event_type (event_type) values ('confirm_activity');

4) Enter "go run cmd/main.go" to command line

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
