# Fromhel Clients Tables Fix

This project is designed to migrate and fix client data in the `fromhelClientsDB` MongoDB database. It reads client data, processes it, and inserts the updated data back into the database.

### Files and Directories

- **main.go**: Entry point of the application. Connects to MongoDB and initiates the migration process.
- **services/**: Contains utility functions used in the migration process.
  - **date.go**: Provides the `ParseDate` function to parse date strings.
  - **trim.go**: Provides the `Trim` function to clean up strings.
  - **uuid.go**: Provides the `GenerateId` function to generate new UUIDs.
- **usecases/**: Contains the core business logic.
  - **migrate.go**: Implements the `Migrate` function to handle the migration process.

## Getting Started

### Prerequisites

- Go 1.22.8 or later
- MongoDB instance

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/FromhelStudio/fromhel-clients-tables-fix.git
    cd fromhel-clients-tables-fix
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Create a `.env` file based on the `.env.example` file:
    ```sh
    cp .env.example .env
    ```

4. Update the `MONGODB_URI` in the `.env` file with your MongoDB connection string.

### Running the Application

To run the application, execute the following command:

```sh
go run [main.go](http://_vscodecontentref_/9)
```