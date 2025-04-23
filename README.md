# Desk SDK for Go

A Go client library for the Teamwork Desk API.

## Features

- Full support for Teamwork Desk API v2
- Type-safe API using Go generics
- Comprehensive test coverage
- Structured logging with logrus
- Easy to use and extend
- Command-line interface for common operations

## Installation

```bash
go get github.com/ready4god2513/desksdkgo
```

## Usage

### Command Line Interface

The SDK includes a command-line interface for common operations. You can use it with flags or environment variables:

```bash
# Using flags
./desksdkgo --api-key "your-api-key" --resource tickets --action list

# Using environment variables
export DESK_API_KEY="your-api-key"
export DESK_RESOURCE="tickets"
export DESK_ACTION="list"
./desksdkgo

# Get a specific ticket
./desksdkgo --api-key "your-api-key" --resource tickets --action get --id 123

# Create a new customer
./desksdkgo --api-key "your-api-key" --resource customers --action create
```

Available flags:
- `--api-key`: Desk API key (required)
- `--base-url`: Desk API base URL (default: "https://[yoursite].teamwork.com/desk/api/v2")
- `--resource`: Resource to interact with (tickets, customers, companies, users)
- `--action`: Action to perform (get, list, create, update)
- `--id`: Resource ID for get/update actions

Environment variables:
- `DESK_API_KEY`: Desk API key
- `DESK_BASE_URL`: Desk API base URL
- `DESK_RESOURCE`: Resource to interact with
- `DESK_ACTION`: Action to perform

### Using as a Go Module

#### Creating a Client

```go
import "github.com/ready4god2513/desksdkgo/client"

// Create a new client with default options
deskClient := client.NewClient("https://[yoursite].teamwork.com/desk/api/v2")

// Or with custom options
deskClient := client.NewClient(
    "https://[yoursite].teamwork.com/desk/api/v2",
    client.WithAPIKey("your-api-key"),
    client.WithHTTPClient(customHTTPClient),
    client.WithLogLevel(logrus.InfoLevel),
)
```

#### Using Services

The SDK provides services for different Teamwork Desk resources:

- `Tickets`: Manage tickets
- `Users`: Manage users
- `Customers`: Manage customers
- `Companies`: Manage companies

Each service provides the following methods:

- `Get(ctx context.Context, id int)`: Get a single resource by ID
- `List(ctx context.Context, params url.Values)`: List resources with optional filters
- `Create(ctx context.Context, resource *T)`: Create a new resource
- `Update(ctx context.Context, id int, resource *T)`: Update an existing resource

#### List Options

When listing resources, you can use the `ListOptions` struct to customize the results:

```go
options := &client.ListOptions{
    Page:    1,           // Page number
    PerPage: 50,          // Items per page
    SortBy:  "created_at", // Field to sort by
    SortDir: "desc",      // Sort direction (asc/desc)
    Embed:   "customer",  // Related resources to include
    Fields:  "id,subject", // Fields to return
    Q:       "status:open", // Search query
}

// Convert options to URL values
params := url.Values{}
params.Set("page", strconv.Itoa(options.Page))
// ... or use the Encode() method
queryString := options.Encode()
```

#### Filtering

The SDK supports MongoDB-style filtering through the `FilterBuilder`. This allows you to create complex filters using a fluent interface:

```go
// Simple equality filter
filter := client.NewFilter().
    Eq("status", "open").
    Build()

// Complex filter with multiple conditions
filter := client.NewFilter().
    Gt("id", 2).
    And(
        client.NewFilter().Eq("status", "open"),
        client.NewFilter().Lt("priority", 3),
    ).
    Build()

// Using OR conditions
filter := client.NewFilter().
    Or(
        client.NewFilter().Eq("status", "open"),
        client.NewFilter().Eq("status", "pending"),
    ).
    Build()

// Using IN operator
filter := client.NewFilter().
    In("priority", 1, 2, 3).
    Build()

// Complex nested conditions
filter := client.NewFilter().
    Or(
        client.NewFilter().Gt("age", 20),
        client.NewFilter().Eq("name", "mike"),
        client.NewFilter().And(
            client.NewFilter().Gt("age", 12),
            client.NewFilter().Lte("age", 18),
        ),
    ).
    Build()
```

Available operators:
- `$eq`: Equal to
- `$ne`: Not equal to
- `$lt`: Less than
- `$lte`: Less than or equal to
- `$gt`: Greater than
- `$gte`: Greater than or equal to
- `$in`: In list
- `$nin`: Not in list
- `$and`: Logical AND
- `$or`: Logical OR

To use the filter in a request:

```go
params := url.Values{}
params.Set("filter", filter)
tickets, err := ticketService.List(ctx, params)
```

#### Error Handling

The SDK provides comprehensive error handling:

- HTTP errors (status codes >= 400) are returned as errors
- JSON encoding/decoding errors are properly handled
- Required parameters are validated
- API key validation is performed

Example error handling:

```go
ticket, err := ticketService.Get(ctx, 123)
if err != nil {
    if strings.Contains(err.Error(), "HTTP error: 404") {
        log.Printf("Ticket not found")
    } else if strings.Contains(err.Error(), "HTTP error: 401") {
        log.Printf("Authentication failed")
    } else {
        log.Printf("Unexpected error: %v", err)
    }
    return
}
```

Example:

```go
// Get the ticket service
ticketService := deskClient.Tickets

// Get a ticket
ticket, err := ticketService.Get(context.Background(), 123)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Got ticket: %+v\n", ticket)

// List tickets with filters
params := url.Values{}
params.Set("status", "open")
tickets, err := ticketService.List(context.Background(), params)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Got %d tickets\n", len(tickets))

// Create a customer
customerService := deskClient.Customers
customer := &models.Customer{
    FirstName: "John",
    LastName:  "Doe",
    Email:     "john.doe@example.com",
}
created, err := customerService.Create(context.Background(), customer)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Created customer: %+v\n", created)
```

#### Models

The SDK includes type-safe models for all resources:

```go
import "github.com/ready4god2513/desksdkgo/models"

// Ticket model
ticket := &models.Ticket{
    Subject: "Test Ticket",
}

// Customer model
customer := &models.Customer{
    FirstName: "John",
    LastName:  "Doe",
    Email:     "john.doe@example.com",
}

// Company model
company := &models.Company{
    Name: "Example Corp",
}

// User model
user := &models.User{
    FirstName: "John",
    LastName:  "Doe",
    Email:     "john.doe@example.com",
}
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details. 