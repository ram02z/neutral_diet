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

The application follows the [dotenv convention](https://github.com/bkeepers/dotenv#what-other-env-files-can-i-use) for managing multiple environments.
The order is summarised below:
```
.env                # loaded in all cases
.env.local          # loaded in all cases, ignored by git
.env.[mode]         # only loaded in specified mode
.env.[mode].local   # only loaded in specified mode, ignored by git
```
The mode is set using the `APP_ENV` variable on the backend and `NODE_ENV` variable on the frontend.

Once the environment is setup, run the following commands to setup and run the application.

```bash
# Start postgres
docker-compose --env-file .env.[mode] up

# Create the database schema
make migrate up

# Start application
go run cmd/app/main.go
# OR for live reloading
air
```

## Database

To populate the database with food items, follow these [steps](./data/README.md). 

## License

[MIT](LICENSE)

## Credits

- [api-go-template](https://github.com/kevinmichaelchen/api-go-template) for the Go backend boilerplate architecture
- [react-pwa](https://github.com/suren-atoyan/react-pwa) for React frontend boilerplate architecture
