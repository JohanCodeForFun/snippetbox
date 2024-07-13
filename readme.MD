# Database Setup

## Prerequisites
- MySQL installed
- Flyway installed (optional but recommended)

## Initial Setup
1. Clone the repository:
   ```sh
   git clone https://github.com/JohanCodeForFun/snippetbox
   cd snippetbox
   ```

2. Create the database:
   ```sh
   mysql -u root -p -e "CREATE DATABASE mydb;"
   ```

3. Run the initialization script:
   ```sh
   mysql -u myuser -p mydb < db/init.sql
   ```

4. (Optional) Load seed data:
   ```sh
   mysql -u myuser -p mydb < db/seeds/seed_data.sql
   ```

## Running Migrations
To apply new migrations, simply add new SQL files to the `db/migrations` directory and run the Flyway migrate command.

## Starting web server
go run ./cmd/web

## Running tests
go test -v ./cmd/web

## Contributing
When adding or modifying migrations, ensure they are numbered sequentially and tested locally before committing.