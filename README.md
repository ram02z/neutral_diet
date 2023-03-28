# Neutral Diet

## Getting started

### Prerequisites

- Go
- GNU Make
- Docker Compose
- [buf](https://github.com/bufbuild/buf)
- [sqlc](https://github.com/kyleconroy/sqlc)

### Environment

There are two files included to track the environment variables:

- [Backend .env.example](./.env.example)
- [Frontend .env.example](./ui/.env.example)

The application also includes `.env.development` files for both the backend and frontend.
Not all required environment variables are included in `.env.development`, i.e. Firebase environment variables.

The application follows the [dotenv convention](https://github.com/bkeepers/dotenv#what-other-env-files-can-i-use) for managing multiple environments.
The order is summarised below:
```bash
.env                # loaded in all cases
.env.local          # loaded in all cases, ignored by git
.env.[mode]         # only loaded in specified mode
.env.[mode].local   # only loaded in specified mode, ignored by git
```
The mode is set using the `APP_ENV` variable on the backend and `NODE_ENV` variable on the frontend.

Once the environment is setup, run the following commands to setup and run the application.

```bash
# Start postgres server (and application server)
# Use .env.docker unless you want to run the applcation server another way
docker-compose --env-file .env.[mode] up

# Create the database schema
make migrate up

# Start application server
go run cmd/app/main.go
# OR for live reloading
air
```

### Database

Before starting to use the application, you need to pre-populate the database.
To populate the database with food items, follow these [steps](./data/README.md). 

The database relies on materialized views for querying the aggregate food items.
Ensure you run the following to setup the materialized views:

```sql
REFRESH MATERIALIZED VIEW aggregate_food_item;
REFRESH MATERIALIZED VIEW regional_aggregate_food_item;
```

### Services

The application server exposes various services:

- `FoodService`
    - Used to create food items and list respective information
    - No authentication
    - [RPCs](https://buf.build/ram02/neutral-diet/docs/main:neutral_diet.food.v1#neutral_diet.food.v1.FoodService)
- `JobService`
    - Used to run scheduled jobs
    - OIDC authentication (Google Cloud Scheduler)
    - [RPCs](https://buf.build/ram02/neutral-diet/docs/main:neutral_diet.job.v1#neutral_diet.job.v1.JobService)
- `UserService`
    - Used by web application's private routes
    - JWT authentication (Firebase Authentication)
    - [RPCs](https://buf.build/ram02/neutral-diet/docs/main:neutral_diet.user.v1#neutral_diet.user.v1.UserService)

## License

[MIT](LICENSE)

## Credits

- [api-go-template](https://github.com/kevinmichaelchen/api-go-template) for the Go backend boilerplate architecture
- [react-pwa](https://github.com/suren-atoyan/react-pwa) for React frontend boilerplate architecture
