# Finance-Bot

Finya is a Telegram bot that helps manage your finances.


# Installation

Clone the repository
Install Docker and Go
Copy [`.env-example`](.env-example) to [`.env`](.env) and fill in the [environment variables](#environment-variables).

Run docker-compose up to start the Postgres database
> docker compose up

Run make run to start the bot
> make run

## Environment Variables

The project uses the following environment variables, which are defined in the [`.env`](.env) file:

- `TELEGRAM_BOT_KEY`: This is your Telegram bot's key. 
Replace `<BOT-KEY>` with your own bot key.

- `LOG_LEVEL`: This is the logging level for your application. 
It can be one of the following values: "DEBUG", "INFO", "WARNING", "ERROR". The default value is "INFO".

You need to copy these environment variables from the [`.env-example`](.env-example) file to the [`.env`](.env) file and fill in the appropriate values.

