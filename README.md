# go-crud-rest-api

## Technologies used:
- Go
- Docker

### Prerequisites
#### go
Make sure you have the go installed using this command:
```bash script
go version
```

#### Docker
Make sure you have docker installed using this command:
```bash script
docker --version
```

#### Curl
Make sure you have curl installed using this command:
```bash script
curl --version
```

### Build the code
``` bash
go build .
```

### Run the code
``` bash
go run .
```

#### Running the application locally
#####  Create docker image of app
Creating a docker image should be as simple as
``` bash
docker build -t goapp .
```

##### 🐘 Run the Postgres container
```bash script
docker-compose up -d db
```

##### 🏗️ Build the go app image
```bash script
docker compose build
```

##### 👟 Run the Go Container
```bash script
docker compose up goapp
```

##### 🧪 Test the applications endpoints

Request to get the all the users:
```bash script
curl --location --request GET 'http://localhost:8080/users'
```
Example of a response:
`[
{
"id": 1,
"name": "aaa",
"email": "aaa@mail"
},
{
"id": 2,
"name": "bbb",
"email": "bbb@mail"
}
]`

Request to create a new user
```bash script
curl --location --request POST 'http://localhost:8080/users' \
--header 'Content-Type: application/json' \
--data-raw '{"name": "aaa","email": "aaa@mail"}'
```

Request to get one specific user:
```bash script
curl --location --request GET 'http://localhost:8080/users/2'
```
Example of a response:
`{
"name": "new",
"email": "new@mail"
}`

Request to update a user
```bash script
curl --location --request PUT 'http://localhost:8080/users/2' \
--header 'Content-Type: application/json' \
--data-raw '{"name": "new","email": "new@mail"}'
```

Request to delete a user
```bash script
curl --location --request DELETE 'http://localhost:8080/users/3'
```

## Contact
This project is maintained by [MikAoJk](CODEOWNERS)