# App-server

first run go mod tidy to download dependency

To run Seeders pass the -runseeders flag
To run migration pass the -runmigration flag
To run server pass the -startserver flag

- go run main.go -runseeders -runmigration -startserver


How To generate swagger files and documents.

swag init -g main.go --output docs/fibersimple/

after running this command yml files is generated inside docs/fibersimple.