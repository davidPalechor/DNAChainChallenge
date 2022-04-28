# DNA Chain Mutant Detection
Human/Mutant detection based on their DNA sequence matrix

## Requirements
- Docker
- docker-compose
- [`.env`](https://drive.google.com/drive/u/0/folders/1SNYnUQka7g4G7Cls_wuxw6GsicKBl1AP) environment file
  - This file should contain these vars:
  ```
  PORT=
  DATABASE_URL=
  POSTGRES_PASSWORD=
  RUN_MODE=
  PORT=
  ```

## Building
- Inside the project root run `docker-compose build`

## Running
- Inside the project root run `docker-compose up`

## Migrations
- Export `DATABASE_URL` found in the `.env` file.
- First, start your service using `docker-compose up`
- In another terminal run `docker-compose exec web bee migrate --conn=$DATABASE_URL --driver=postgres`

## Usage
### Find a mutant

- Make a `POST` request to your `localhost:$PORT/v1/users` with a payload:
```json
{
  "dna": ["CTGCG", "CAGTG", "TTATG", "AGAAG", "CCTCT"]
}
```

- Returns `200 OK` if mutant, `400 Bad request` otherwise

### Get Stats
`[GET] localhost:%PORT/v1/stats`
```json
{
  "count_mutant_dna": 5,
  "count_human_dna": 2,
  "ratio": 2.5
}
```

## Coverage
```shell
go test ./tests -coverpkg=./... 
ok      DNAChainChallenge/tests 0.463s  coverage: 70.3% of statements in ./...
```

### Check full API docs at `localhost:$PORT/swagger`
## Test this APP

- Make a `POST` request to `https://dna-chain-app.herokuapp.com/v1/mutant` with a payload:
```json
{
  "dna": ["CTGCG", "CAGTG", "TTATG", "AGAAG", "CCTCT"]
}
```

- `[GET] https://dna-chain-app.herokuapp.com/v1/stats`
- Return
```json
{
  "count_mutant_dna": 5,
  "count_human_dna": 2,
  "ratio": 2.5
}
```