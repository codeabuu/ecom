## Ecommerce Backend (Part1: Cart Microservice)
This GoLang project implements a robust backend system for an e-commerce application, focusing on the cart functionality as a microservice. It provides essential services for managing orders, users, authentication, and of course, shopping carts.

## Key Features
* **Order Management Service(OMS)**: Handles order creation, process and fulfillment.
* **User Management Service**: Manages user registration, login and account details.
* **Authentication Service**
* **Shopping cart service**:
  * Users can add and remove products from their carts
  * Supports instant purchase functionality for a seamless checkout experience
* **Address Managaement**:
  * Users can add, edit, and delete their shipping and billing addresses.
  * Streamlines the checkout process by providing saved addresses.
* **Product management**:
  * Users can view product details and search for specific items. (Full product management might be implemented in a separate microservice)
  
## How To Run
Make sure that your MongoDB username, password &URI is set in your `.env` file at the root directory of this project.

You can start the project with below commands:

1. **`docker-compose up -d`**

2. **`go run main.go`**

## API Tests
You can use postman for API Endpoints tests. Here are some examples of endpoint tests:

#### Signup Endpoint

Here is an example of a succesful signup request and response:
![Signup Endpoint](docs/images/signup.png)

#### Login Endpoint

An example of a susccefull login:
![Login Endpoint](docs/images/login.png)

#### Add to Cart Endpoint

This images shows adding a product for viewing:
![Add product endpoint](docs/images/addproduct.png)

#### Product view

This images shows an already added product:
![view product Endpoint](docs/images/productview.png)

#### Search product

This image shows the product after searching it:
![search product Endpoint](docs/images/searchprod.png)



Click [here](https://docs.google.com/document/d/1UEeHIvhaaA9IXMhYLDiJtD_8RNBqZmoh3kJ09GH3DrY/edit?usp=sharing) to access all the API Endpoints documentation for this project and how to run them.
