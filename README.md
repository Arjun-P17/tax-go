# tax-go

This project is designed to determine the tax implications of stock trades based on Australian capital gain laws.

Please note that this project is currently in development, and efforts are underway to address some limitations:
- Only compatible with files formatted according to Interactive Brokers (IBKR).
- Only utilizes a FIFO algorithm for determining which buy allotments to sell.
- Currently, the calculation is done in USD.
- No testing

### Usage
- Ensure your local MongoDB server is running and the port is correctly defined in `config.yaml`.
- Place your IBKR-formatted trade files in the `csvpath` defined in `config.yaml`

- Running is a two step process
- run `make parseTrades` to parse the trades into the db
- run  `make processTrades` to process the trades and calculate tax info