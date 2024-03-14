# This is a simple RESTful WebSite

---

## Installation:
- sqlx `SQL Database interfase`
- goose `Database migration tool`
- echo `Web framework for GO`
- redis `Redis client`
- driver for postgres `PostgreSQL driver`

and other packages that are in the go.mod file

To download all required packages, run:
```go mod download```

## Getting Started

**Start the Server**: Run the following command to start the server. Note that the server runs in a Docker container:
```docker compose up```

**Configure Redis**: Ensure Redis is properly configured, as the server uses it for caching. If authentication is required, set up Redis login credentials accordingly.

**Set Up PostgreSQL**: Configure PostgreSQL to store user information and todo lists. Create necessary tables and configure database connection settings.

## Usage

The server will start at the following address:
```http://localhost:1323```

## Customization

Feel free to customize HTML pages or add your own content as needed. Simply edit the files located in the frontend/html directory.

All HTML pages are located in the following directory:
```frontend/html```

## Security

* Secure Redis and PostgreSQL connections by setting up authentication and encryption.

* Implement user authentication mechanisms to protect sensitive user data.

## Contributing

Contributions are welcome! If you'd like to contribute to the project, please follow these guidelines:

* Fork the repository and create a new branch for your feature or bug fix.
* Ensure your code follows the project's coding style and conventions.
* Submit a pull request with a clear description of your changes and their purpose.

## Very easy to use and very easy to understand :)
