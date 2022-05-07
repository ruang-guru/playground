# Cashier App

## Requirements

- There are two users, admin and employee.
- User should able to login to the system
- User should able to logout from the system
- User should able to see the list of product
- User should able to add product to the cart
- User should able to clear the cart
- User should able to see the product list in the cart
- User should able to make a payment
- Admin user should able to see summary of sales in the dashboard
- Admin user should able to see all sales product in all time
- Admin user should able to see all sales product in specific period
- Admin user should able to see sales of specific product in all time
- Admin user should able to see sales of specific product in specific period

## Available APIs

- `POST`: `/api/user/login`
- `POST`: `/api/user/logout`
- `GET`: `/api/products`
- `POST`: `/api/cart/add`
- `GET`: `/api/cart/clear`
- `GET`: `/api/carts`
- `GET`: `/api/dashboard?cash=<cash>`
- `GET`: `/api/admin/sales`
- `GET`: `/api/admin/sales?product_name=<product_name>`
- `GET`: `/api/admin/sales?start_period=<yyyy-mm-dd>&end_period=<yyyy-mm-dd>`
