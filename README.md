# Email - Currency app

#### Email - Currency app for daily getting emails with a currency exchange rate from Monobank.



## Table of contents
* [Architecture](#Architecture)
* [Setup](#Setup)
    * [Technologies Used](#Technologies-Used)
    * [How to Run](#How-to-Run)
    * [Room for improvement](#Room-for-improvement)
        * [Existing Features](#Existing-features)
        * [TODO List](#TODO-List)


## Architecture
`/cmd` package for storing main.go that launch all dependencies.

`/configs` gets values from the environment (.env file).

`/internal` stores `/database` with start migrations.

`/pkg` includes:

*   `/handler` receives HTTP requests and passes on service.

*   `/service` processes handlers and implements business logic.

*   `/repository` works with a database.

*   Files for storing a main entities' structures.

`Docker` files:
*   `Dockerfile` build container & launch all dependencies.

*  `docker-compose.yml` runs application and database in the separate containers to avoid possible data loss if an app crash.


`Makefile` consists of main commands.

`example.env` is an example of .env that stores main secret values.


## Setup
Environmental variables are located in .env file.
  

### Technologies Used
* Go 1.22
* PostgreSQL
* Docker

### How to Run
1. Rename 'example.env' to '.env'.
2. Run `make build && make run`.
3. Run `make migrate_up`.
to apply migrations.



### Room for improvement
#### Existing Features
* Receiving the current dollar (USD) to hryvnia (UAH) exchange rate from Monobank API.
* Receiving an email from a JSON response.
* Following SOLID principles creating scalable infrastructure using interfaces.
* Database Migrations.
* Docker and microservice architecture.


#### TODO List
1. Dividing `/rate` realization in handler into service where business logic should be.
2. CRUD operations for email table.
3. Email validation.
4. [Mailjet]('https://dev.mailjet.com/email/guides/getting-started/') for email sending.
5. Adding CI/CD approach using GitHub Actions.
