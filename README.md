# tax-go

This project is designed to determine the tax implications of stock trades based on Australian capital gain laws.

Please note that this project is currently in development, and efforts are underway to address some limitations:
- Only compatible with files formatted according to Interactive Brokers (IBKR).
- Only utilizes a FIFO algorithm for determining which buy allotments to sell.
- Currently, the tax calculation is done in USD.
- No testing

Models in the repository layer contain comments with details about the structs and the fields.

AUDUSD data is fetched from https://au.investing.com/currencies/aud-usd-historical-data

### Usage

- Ensure your local MongoDB server is running and the port is correctly defined in `config.yaml`.
- Place your IBKR-formatted trade files in the `csvpath` defined in `config.yaml`

Running is a two step process
- run `make parseTrades` to parse the trades into the db
- run  `make processTrades` to process the trades and calculate tax info

A grpc server serves a simple API that returns the transactions for all stocks.
- run `make run` to start the server