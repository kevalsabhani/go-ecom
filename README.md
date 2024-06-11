# ecom

## Local Setup
1. Create environment file `.env` and set all envs. All the required environment variables are already mentioned in the `.env.example` file.
2. Run command `make startdb` to create db container and create database in it.
3. Run command `make migrate-up` to create database tables.
4. To run the server, command `make run`.
