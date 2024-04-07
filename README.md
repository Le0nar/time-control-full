Backend

For first starting:
1. Clone and start postgresql containers for tc-auth, tc-read, tc-write
    1.1 docker run --name=tc_auth -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d postgres
    1.2 docker run --name=tc-read -e POSTGRES_PASSWORD='qwerty' -p 5434:5432 -d postgres
    1.3 docker run --name=tc-write -e POSTGRES_PASSWORD='qwerty' -p 5435:5432 -d postgres

2. Apply migrations for databases. Look detailed info in service's README.md

3. Add .env file for some services. Look detailed info in service's README.md

4. Run all services
    2.1 tc-auth
    2.1.1 cd tc-auth
    2.2.2 go run cmd/main.go

    2.2 tc-read
    2.2.1 cd tc-read
    2.2.2 go run cmd/main.go

    2.3 tc-write
    2.3.1 cd tc-write
    2.3.2 go run cmd/main.go

    2.4 tc-face-recognising
    2.4.1 cd tc-face-recognising
    2.4.2 go run main.go



For next starting:
1. Run existing containers
    1.1 Check containers ids: docker ps -a
    1.2 docker start {tc_auth container id}
    1.3 docker start {tc-read container id}
    1.4 docker start {tc-write container id}

2. Run all services like in instruction above 

For additional info look at README.md in each service directory