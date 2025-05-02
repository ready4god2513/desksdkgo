# Teamwork Desk Go SDK

A Go SDK for interacting with the Teamwork Desk API. This SDK provides a simple and intuitive way to interact with Teamwork Desk's REST API.

## Features

- Support for all major Teamwork Desk API endpoints
- Simple and intuitive client interface
- Built-in logging and debugging support
- Command-line interface for quick operations
- Environment variable support for configuration

## Installation

```bash
go get github.com/ready4god2513/desksdkgo
```

## Usage

### Basic Client Usage

```go
package main

import (
    "context"
    "github.com/ready4god2513/desksdkgo/client"
    "github.com/sirupsen/logrus"
)

func main() {
    // Create a new client
    c := client.NewClient(
        "https://yourcompany.teamwork.com/desk/api/v2",
        client.WithAPIKey("your-api-key"), // Get this in your profile settings
        client.WithLogLevel(logrus.DebugLevel),
    )

    // Use the client
    ctx := context.Background()
    
    // List tickets
    tickets, err := c.Tickets.List(ctx, nil)
    if err != nil {
        panic(err)
    }
    
    // Get a specific ticket
    ticket, err := c.Tickets.Get(ctx, 123)
    if err != nil {
        panic(err)
    }
}
```

### Available Resources

The SDK supports the following resources:

- **Tickets**: Manage support tickets
- **Customers**: Manage customer information
- **Companies**: Manage company information
- **Users**: Manage user accounts
- **Tags**: Manage ticket tags
- **Ticket Statuses**: Manage ticket statuses
- **Ticket Types**: Manage ticket types
- **Ticket Priorities**: Manage ticket priorities
- **Help Doc Sites**: Manage help documentation sites
- **Help Doc Articles**: Manage help documentation articles
- **Business Hours**: Manage business hours
- **SLAs**: Manage service level agreements

Each resource supports the following operations:
- `Get`: Retrieve a single resource by ID
- `List`: Retrieve a list of resources with optional filters
- `Create`: Create a new resource
- `Update`: Update an existing resource

### Command Line Interface

The SDK includes a command-line interface for quick operations:

```bash
# List tickets
./desksdkgo --api-key YOUR_API_KEY --resource tickets --action list

# Get a specific ticket
./desksdkgo --api-key YOUR_API_KEY --resource tickets --action get --id 123

# Create a new ticket
./desksdkgo --api-key YOUR_API_KEY --resource tickets --action create --data '{"subject": "New Ticket", "description": "Ticket description"}'

# Update a ticket
./desksdkgo --api-key YOUR_API_KEY --resource tickets --action update --id 123 --data '{"status": "resolved"}'
```

### Configuration

The CLI supports the following configuration options:

- `--api-key`: Teamwork Desk API key (required)
- `--base-url`: Teamwork Desk API base URL (default: https://mycompany.teamwork.com/desk/api/v2)
- `--resource`: Resource to interact with (default: tickets)
- `--action`: Action to perform (get, list, create, update) (default: list)
- `--id`: Resource ID for get/update actions
- `--debug`: Enable debug logging
- `--data`: JSON data to merge with default values for create/update actions

All configuration options can also be set via environment variables:
- `DESK_API_KEY`
- `DESK_BASE_URL`
- `DESK_RESOURCE`
- `DESK_ACTION`

### Filtering

The SDK includes a filter builder for creating complex queries:

```go
filter := client.NewFilter().
    Eq("status", "open").
    And(
        client.NewFilter().Gt("created_at", "2023-01-01"),
        client.NewFilter().Lt("created_at", "2023-12-31"),
    )

tickets, err := c.Tickets.List(ctx, filter.Build())
```

Available filter operators:
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

## License

This project is licensed under the MIT License - see the LICENSE file for details. 