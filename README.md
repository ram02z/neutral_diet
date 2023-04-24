# Neutral Diet

## Getting started

### Prerequisites

- [Go](https://go.dev/)
- [GNU Make](https://www.gnu.org/software/make/)
- [Docker Compose](https://docs.docker.com/compose/)
- [buf](https://github.com/bufbuild/buf)
- [sqlc](https://github.com/kyleconroy/sqlc)
- Firebase Project

### Environment

The application includes `.env.development` files with sample working
development environment variables. Not all required environment variables are
included in `.env.development`, i.e. Firebase environment variables.

- `GOOGLE_APPLICATION_CREDENTIALS`: path to Google Cloud service account key (JSON)
    - [Docs](https://cloud.google.com/docs/authentication/application-default-credentials#GAC)


Refer to [.env.example](./.env.example) for the complete list of required
environment variables.

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
# Use .env.docker unless you want to run the application server another way
docker-compose --env-file .env.[mode] up
```

Alternatively, you can run the application directly:
```bash
# Start application server
go run cmd/app/main.go
# OR for live reloading
air
```

Once the database and server are running, you can create the database schema using the migration files:
```bash
make migrate-up
```

### Database

Before starting to use the application, you need to pre-populate the database.
To populate the database with food items, follow these [steps](./data/README.md). 

The database relies on materialized views for querying the aggregate food items used for the search functionality.
Ensure you run the following to setup the materialized views:

```sql
REFRESH MATERIALIZED VIEW aggregate_food_item;
REFRESH MATERIALIZED VIEW regional_aggregate_food_item;
```

### UI

Once the database and server instances are setup and running, you can follow
the UI setup [guide](/ui/README.md).

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
