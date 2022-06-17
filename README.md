# Backend Test

## How to run

```bash
# clone this repo
git clone https://github.com/sampriskwila/be-test.git

# move to project directory
cd be-test

# run docker build
make up

# run the application
go run main.go

# stop docker container
make down

# clean up docker images
make prune
```

## How to access endpoint

### POST - localhost:8080/api/v1/balance

```json
request:

{
  "wallet_id": "228dfafa-0e46-412f-9339-97e5e18bba1e",
  "amount": 6000
}
```

### GET - localhost:8080/api/v1/balance/{wallet_id}

```json
response:

{
  "data": {
    "wallet_id": "228dfafa-0e46-412f-9339-97e5e18bba1e",
    "amount": 6000,
    "is_threshold": false,
    "updated_at": "2022-06-17T11:51:25.8853061+07:00"
  }
}
```
