## Service for evaulating math expressions

Service with one endpoint for secure evaulating math expression.

Examples of correct expressions: `2 + 2`, `(2 + 2) * 2`, `2 / 2`

### Installation

Clone this repository 

```bash
git clone https://github.com/timka-123/yandex_web_calc
```

And navigate to project folder

```bash
cd yandex_web_calc
```

#### Running with Docker

```bash
docker compose up
```
or...
```bash
docker-compose up
```

#### Running without Docker

```bash
go run cmd/main.go
```


That will run application on `localhost:8080`


### Usage

Make a `POST` request to `http://localhost:8080/api/v1/calculate` with `{'expression': 'your expression'}` body.

Example curl:
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "expression"
}'
```

### Tests

To run project run command there:
```bash
go test tests/* -v
```