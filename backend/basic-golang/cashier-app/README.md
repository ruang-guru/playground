# Cashier App

## Goals
- Learn how to read and write to a CSV file (database concept)
- Learn CRUD operations
- Learn how to render data to the frontend
- Learn how to make and call APIs
- Learn code reuse: common `repository` code is reused by different frontend (terminal & API)

## 3 modes

### Terminal Mode
<!-- video here -->


### TView Mode
<!-- video here -->


### API Mode
<!-- video here -->
- `/api/user/login?username=<username>&password=<password>`
- `/api/user/logout?username=<username>`
- `/api/dashboard?cash=<cash>`
- `/api/products`
- `/api/cart/add?product_name=<product_name>`
- `/api/cart/clear`
- `/api/carts`

For simplicity, we only use HTTP GET here.
The API also doesn't support multiple concurrent sessions. We will learn about this later :)