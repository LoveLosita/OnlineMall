# [中文版本(Simplified Chinese version)](readme_zh_cn.md)

# 1. Project Features (Current Version: `1.5.0Beta`)

This project is similar to an e-commerce website and supports the following features:

- [x]  User registration and login
- [x]  Encrypted password storage
- [x]  Adding product categories
- [x]  Modifying user information
- [x]  Searching for products
- [x]  Viewing product details
- [x]  Viewing products under a specific category
- [x]  Commenting on products (add, delete, and view comments)
- [x]  Adding products to the shopping cart
- [x]  Retrieving all products in the shopping cart
- [x]  Searching for products in the shopping cart
- [x]  Basic order placement (returns an `order` object to the frontend)
- [ ]  Advanced order placement (affects product inventory and handles concurrency)
- [ ]  Ultimate order placement (simulating WeChat/Alipay payment callbacks based on advanced ordering)
- [x]  Nested comments
- [x]  Anonymous comments
- [x]  Browsing history tracking
- [x]  Merchant and administrator roles (can freely add or remove products)
- [ ]  Chat with AI assistant (using the ChatAI API)
- [ ]  Adding CAPTCHA for login (cost is a concern, may explore implementing 2FA)
- [ ]  Implementing caching (Redis cache)
- [x]  Designing a heat algorithm to display frequently viewed product categories on the homepage
- [ ]  Enabling user-customer service chat (may implement merchant, admin, and user clients using Python if time allows)
- [ ]  Deploying to a personal server for public access
- [ ]  Enhancing security (XSS, SQL injection, CSRF, etc.; SQL injection is already prevented by using placeholders in SQL queries)
- [ ]  Any other features you’d like to add (procrastination mode: ON)

# 2. Project Structure

## 2.1. File Structure Diagram

```go
OnlineMall
│
├── api                  // Stores all API endpoint definitions
│   ├── auth.go           // User authentication-related APIs
│   ├── cart.go           // Shopping cart-related APIs
│   ├── categories.go     // Product category-related APIs
│   ├── order.go          // Order management-related APIs
│   ├── product.go        // Product management-related APIs
│   ├── review.go         // Product review-related APIs
│   └── user.go           // User management-related APIs
│
├── auth                 // Handles authentication functionality
│   ├── check_permission.go  // Checks user permissions
│   └── jwt_generator.go     // Generates and parses JWT tokens
│
├── cmd                  // Contains application startup and initialization logic
│   └── start.go           // Starts the application and initializes services
│
├── dao                  // Database access layer
│   ├── db_connection.go  // Handles database connection
│   ├── cart.go           // Shopping cart data operations
│   ├── categories.go     // Product category data operations
│   ├── order.go          // Order data operations
│   ├── product.go        // Product data operations
│   ├── review.go         // Product review data operations
│   └── user.go           // User data operations
│
├── middleware           // Stores middleware, such as token validation
│   └── token_handler.go    // Handles token validation logic
│
├── model                // Stores data models
│   ├── auth.go           // User authentication model
│   ├── cart.go           // Shopping cart model
│   ├── order.go          // Order model
│   ├── products.go       // Product model
│   ├── review.go         // Product review model
│   ├── user.go           // User model
│   └── map_to_slice.go   // Key-value mapping model
│
├── respond              // Handles API response formatting
│   └── responses.go       // Defines a unified response format
│
├── routers              // Stores route configurations
│   └── router.go          // Configures all API routes
│
├── service              // Business logic layer
│   ├── auth.go           // User authentication-related business logic
│   ├── cart.go           // Shopping cart-related business logic
│   ├── categories.go     // Product category-related business logic
│   ├── order.go          // Order-related business logic
│   ├── product.go        // Product-related business logic
│   ├── review.go         // Product review-related business logic
│   └── user.go           // User-related business logic
│
├── utils                   // Utility functions
│   ├── if_in.go             // Utility functions for checking element existence in sequences
│   ├── pwd_encryption.go    // Password encryption utilities
│   ├── list_in_rank_out.go  // Takes a slice of numbers and returns their ranking
│   └── map_to_int_slice.go  // Converts a map into a slice of custom key-value structures
│ 
├── go.mod               // Go Modules configuration file
├── main.go              // Project entry point, responsible for starting the application
└── readme.md            // Project README file (this document)
```

## 2.2. Directory Details

1. **api**: Contains API definitions for different functionalities such as user authentication, shopping cart management, product categories, product management, order processing, and reviews. Each file corresponds to a specific module.
2. **auth**: Handles user authentication logic, primarily focusing on JWT generation and validation to ensure user identity legitimacy.
3. **cmd**: Contains the application's startup logic. `start.go` serves as the main entry point, responsible for loading configurations and launching the service.
4. **dao**: Data access layer, responsible for database interactions, including CRUD operations. For example, `cart.go` handles shopping cart-related database operations, while `order.go` manages order-related database operations.
5. **middleware**: Middleware layer, typically used for request pre-processing, such as token validation.
6. **model**: Defines data models that map to database tables. Each file represents a specific module's data structure, such as `cart.go` for the shopping cart model and `products.go` for the product model.
7. **respond**: Handles unified API response formatting, ensuring consistency in API return structures for easier frontend processing.
8. **routers**: Defines API routes and their corresponding handlers, managing request paths and logic.
9. **service**: Business logic layer that encapsulates core functionalities. For example, `cart.go` in the service layer handles shopping cart operations.
10. **utils**: Stores utility functions such as element checking and password encryption.
11. **go.mod**: Go Modules configuration file that specifies project dependencies and versions.
12. **main.go**: The main entry point of the project, responsible for initializing configurations and starting the application.
13. **readme.md**: The project's README file, providing basic information, usage instructions, and technology stack details.
14. **readme_zh_cn.md**: Simplified Chinese version of the` README` file (which is also the original version).

# 3. Definition of Status Codes

| Status Code | HTTP Status Code | Description                         | Reason                                                       | Solution                                                     |
| ----------- | ---------------- | ----------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| 20000       | 200              | Success                             | -                                                            | -                                                            |
| 40001       | 401              | Incorrect Username                  | The username provided during login is incorrect and does not match any user in the database. | Provide the correct username.                                |
| 40002       | 401              | Incorrect Password                  | The password provided during login is incorrect and does not match the existing records in the database. | Provide the correct password.                                |
| 40003       | 400              | Invalid Username                    | The username provided during registration already exists in the database. | Provide a unique username.                                   |
| 40004       | 400              | Missing Parameters                  | The number of parameters provided is less than required.     | Provide the required parameters.                             |
| 40005       | 400              | Incorrect Parameter Type            | The parameter type provided is incorrect, preventing it from binding to the structure. | Provide the correct parameter type.                          |
| 40006       | 400              | Parameter Too Long                  | A parameter provided exceeds the allowed length.             | Shorten the parameter length.                                |
| 40007       | 400              | Incorrect Username or Password      | The provided username or password is incorrect.              | Provide the correct username or password.                    |
| 40008       | 400              | Invalid Gender                      | The provided gender is not one of ("male", "female", "other"). | Provide one of the valid options.                            |
| 40009       | 401              | Missing Token                       | The JWT key is missing in the request header.                | Include the JWT key in the header.                           |
| 40010       | 401              | Invalid JWT Signature Method        | The JWT key format is incorrect.                             | Provide a valid JWT token.                                   |
| 40011       | 401              | Invalid Token                       | The JWT token is invalid.                                    | Provide a valid JWT token.                                   |
| 40012       | 401              | Invalid Claims                      | The claims in the JWT token are invalid.                     | Provide a valid JWT token.                                   |
| 40013       | 400              | Invalid User ID                     | The user ID provided for querying user information does not exist. | Provide a valid user ID.                                     |
| 40014       | 401              | Insufficient Permissions            | The user is neither an administrator nor a store owner.      | Have an administrator or store owner perform this action.    |
| 40015       | 404              | Category Not Found                  | The category ID provided when adding a product does not exist. | Provide a valid category ID.                                 |
| 40016       | 400              | Category Name Already Exists        | A duplicate category name was provided when adding a category. | Provide a unique category name.                              |
| 40017       | 404              | Product Not Found                   | Attempted to find a product by ID, but the product does not exist. | Provide a valid product ID.                                  |
| 40018       | 404              | Product Not Found                   | No product was found when searching by keyword.              | Provide an existing product keyword.                         |
| 40019       | 404              | Product List is Empty               | The product list is empty when displaying all products or a specific category. | If displaying all products, add products first. If filtering by category, provide a valid category ID. |
| 40020       | 401              | Invalid Refresh Token               | The provided refresh token is invalid when trying to refresh the access token. | Provide a valid refresh token. If expired, log in again.     |
| 40021       | 400              | Product Already in Cart             | The same product with the same quantity is already in the cart when trying to add it again. | If updating quantity, provide a different quantity; otherwise, choose another product. |
| 40022       | 400              | Quantity Too Large                  | The quantity provided when placing an order or adding a product to the cart exceeds 999. | Provide a quantity less than 999.                            |
| 50001       | 500              | Order Not Found                     | Internal error: order and product records do not match when verifying a purchase. | -                                                            |
| 40024       | 400              | User Has Not Purchased This Product | The user attempted to review a product they have not purchased. | Purchase the product before leaving a review.                |
| 40025       | 400              | User Already Reviewed               | The user attempted to review a product they have already reviewed. | Do not submit duplicate reviews.                             |
| 40026       | 400              | Rating Out of Range                 | The user provided a rating outside the 1-5 range.            | Provide a rating within the valid range.                     |
| 40027       | 400              | Review Too Long                     | The user's review exceeds the allowed character limit.       | Shorten the review text.                                     |
| 40028       | 400              | Parent Comment Not Found            | The parent comment ID provided when replying to a comment does not exist. | Provide a valid parent comment ID.                           |
| 40029       | 404              | Shopping Cart is Empty              | The user requested to display all items in the cart, but the cart is empty. | Add items to the cart before displaying.                     |
| 40030       | 404              | No Product Reviews                  | The user attempted to view reviews for a product, but no reviews exist. | Submit a review before retrieving reviews.                   |
| 40031       | 404              | No Matching Reviews                 | No reviews matching the search keyword were found when searching for product reviews. | Use a different keyword.                                     |
| 40032       | 404              | Review Not Found                    | The review ID provided for query or deletion does not exist. | Provide a valid review ID.                                   |

# 4. General Error Response Examples

Some errors returned by the project are generic, so I only saved their examples in the `apifox` API documentation during the early stages of the project. Later, I did not keep examples. Here, I list these error examples, and you can refer to the table above for their reasons:

## 4.1. Not Logged In

```json
{
    "status": "40009",
    "info": "missing token"
}
```

## 4.2. Invalid JWT Token Signing Method

```json
{
    "status": "40010",
    "info": "invalid signing method"
}
```

## 4.3. Invalid Token / Token Expired

```json
{
    "status": "40011",
    "info": "invalid token"
}
```

## 4.4. Invalid Token Claims

```json
{
    "status": "40012",
    "info": "invalid claims"
}
```

## 4.5. Insufficient Permissions

```json
{
    "status": "40014",
    "info": "unauthorized"
}
```

## 4.6. Missing Parameters

```json
{
    "status": "40004",
    "info": "missing param"
}
```

## 4.7. Incorrect Parameter Type

```json
{
    "status": "40005",
    "info": "wrong param type"
}
```

## 4.8. Parameter Too Long

```json
{
    "status": "40006",
    "info": "param too long"
}
```

# 5. Starting the Project

Ensure that you have the latest version of the Go environment installed locally and that the project has been fully downloaded to your local machine.

First, in the terminal within the main folder of the `OnlineMall` project, run:

```sh
go mod tidy
```

This will download and organize dependencies.

Then, start the project by running:

```sh
go run main.go
```