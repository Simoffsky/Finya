# Finance-Bot

Finya is a Telegram bot that helps manage your finances.


## Installation

### Clone the repository

To clone the repository, run the following command in your terminal:

```sh
git clone https://github.com/Simoffsky/Finya.git
```

### Install Docker and Go

1. Install Docker from [official Docker website](https://www.docker.com/products/docker-desktop). 

2. Install Go from [official Go downloads page](https://golang.org/dl/).

### Configure and Run
Copy [`.env-example`](.env-example) to [`.env`](.env) and fill in the [environment variables](#environment-variables).

Run docker-compose up to start the Postgres database
```sh
docker compose up
```
Run make run to start the bot

```sh
make run
```

### Environment Variables

The project uses the following environment variables, which are defined in the [`.env`](.env) file:

- `TELEGRAM_BOT_KEY`: This is your Telegram bot's key. 
Replace `<BOT-KEY>` with your own bot key.

- `LOG_LEVEL`: This is the logging level for your application. 
It can be one of the following values: "DEBUG", "INFO", "WARNING", "ERROR". The default value is "INFO".

You need to copy these environment variables from the [`.env-example`](.env-example) file to the [`.env`](.env) file and fill in the appropriate values.


## Continuous Integration/Continuous Deployment (CI/CD)

This project uses GitHub Actions for CI/CD. The workflow is defined in [`.github/workflows/go.yaml`](.github/workflows/go.yaml). 

The workflow is triggered on every push or pull request to the `master` branch. It includes the following steps:

- Linting with golangci-lint
- Building the project with `go build`
- Running tests with `go test`

## Testing

Unit tests are located in the same directories as the code they test. For example, tests for the code in `pkg/telebot` are located in `pkg/telebot/bot_test.go`.

To run all tests, use the `make test` command:

```sh
make test