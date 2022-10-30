### Mini Project Golang : Movie API
REST API application to manage notes with rating system

### How to Run
1. Clone this repository.
2. Copy the `.env` file.
   ```shell
   cp .env.example .env
   ```
3. Fill the values inside the `.env` file for the database configurations.
4. Create a new database called `movie-api`.
   ```sql
   CREATE DATABASE "movie-api";
   ```
5. Run the application with this command.
   ```shell
   go run main.go
   ```