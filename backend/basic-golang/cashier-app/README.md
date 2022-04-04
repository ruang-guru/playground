# Cashier App

## Goals
- Learn how to read and write to a CSV file (database concept)
- Learn CRUD operations
- Learn how to render data to the frontend
- Learn how to make and call APIs
- Learn code reuse: common `repository` code is reused by different frontend (terminal & API)



## 3 modes

### Terminal Mode

Video (click the image):

[![Terminal](https://img.youtube.com/vi/iZOPT3axoi4/maxresdefault.jpg)](https://youtu.be/iZOPT3axoi4)

### TView Mode

Video (click the image):

[![TView](https://img.youtube.com/vi/e1yX1O8pKYs/maxresdefault.jpg)](https://youtu.be/e1yX1O8pKYs)


### API Mode

Video (click the image):

[![API](https://img.youtube.com/vi/lI-RO36El08/maxresdefault.jpg)](https://youtu.be/lI-RO36El08)

Available APIs:
- `/api/user/login?username=<username>&password=<password>`
- `/api/user/logout?username=<username>`
- `/api/dashboard?cash=<cash>`
- `/api/products`
- `/api/cart/add?product_name=<product_name>`
- `/api/cart/clear`
- `/api/carts`

For simplicity, we only use HTTP GET here.
The API also doesn't support multiple concurrent sessions. We will learn about this later :)