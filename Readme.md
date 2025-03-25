This is a simple HTTP server, that serves data from a database that contains some
fraction of Ethereum onchain data.

# Running locally (optional)

Start the server:

docker-compose up --build


You should now be able to call the API:

curl http://localhost:8080/metrics/ethereum/transaction_fees_hourly?date=2020-09-07


## Open issue

API calls return with 500, with the server printing an error message containing:
... runtime error: invalid memory address or nil pointer dereference ... pkg/metrics.go:44 ...


While reading the codebase it would be nice if you can:
- explain what you're seeing
- how you're interpreting the codebase
- spot issues and explain how you would fix them
