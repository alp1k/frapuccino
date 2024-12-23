# hot-coffee

## Context

Have you ever wondered how your favorite coffee shop manages to handle a flurry of orders during the morning rush, keep track of inventory so they never run out of your preferred blend, or remember that you like your coffee with an extra shot of espresso?

Behind the scenes, coffee shops rely on sophisticated management systems that coordinate orders, inventory, menu items, and customer preferences in real-time. These systems ensure that baristas can focus on crafting the perfect cup while the technology handles the complexities of order processing, stock management, and data recording.

The `hot-coffee` (coffee shop management system) project is a simplified version of these real-world applications, designed to give you hands-on experience with the core principles behind such operational software. Imagine an application that allows staff to:

- **Manage Orders:** Create, modify, close, and delete customer orders efficiently.
- **Oversee Inventory:** Track ingredient stock levels to prevent shortages and ensure freshness.
- **Update the Menu:** Add new drinks or pastries, adjust prices as needed, and keep the offerings up to date.


### API Endpoints

#### Orders

| Method | Endpoint          | Description                       | Response                   |
|--------|-------------------|-----------------------------------|----------------------------|
| POST   | `/orders`         | Creates a new order.             | 201 Created                |
| GET    | `/orders`         | Retrieves all orders.             | 200 OK                     |
| GET    | `/orders/{id}`    | Retrieves a specific order by ID. | 200 OK     |
| PUT    | `/orders/{id}`    | Updates an existing order.        | 200 OK     |
| DELETE | `/orders/{id}`    | Deletes an order.                 | 204 No Content  |
| POST   | `/orders/{id}/close` | Closes an open order.         | 200 OK   |

#### Menu Items

| Method | Endpoint          | Description                        | Response                   |
|--------|-------------------|------------------------------------|----------------------------|
| POST   | `/menu`           | Adds a new menu item.             | 201 Created                |
| GET    | `/menu`           | Retrieves all menu items.          | 200 OK                     |
| GET    | `/menu/{id}`      | Retrieves a specific menu item.    | 200 OK     |
| PUT    | `/menu/{id}`      | Updates an existing menu item.     | 200 OK     |
| DELETE | `/menu/{id}`      | Deletes a menu item.               | 204 No Content  |

#### Inventory

| Method | Endpoint          | Description                        | Response                   |
|--------|-------------------|------------------------------------|----------------------------|
| POST   | `/inventory`      | Adds a new inventory item.        | 201 Created                |
| GET    | `/inventory`      | Retrieves all inventory items.     | 200 OK                     |
| GET    | `/inventory/{id}`  | Retrieves a specific inventory item. | 200 OK    |
| PUT    | `/inventory/{id}`  | Updates an inventory item.         | 200 OK    |
| DELETE | `/inventory/{id}`  | Deletes an inventory item.         | 204 No Content  |

#### Aggregations

| Method | Endpoint                  | Description                        | Response                   |
|--------|---------------------------|------------------------------------|----------------------------|
| GET    | `/reports/total-sales`    | Retrieves the total sales amount.  | 200 OK                     |
| GET    | `/reports/popular-items`   | Retrieves a list of popular menu items. | 200 OK                  |




  # Request Examples

- **Create/Update Order Request:**
```http 
POST /orders
Content-Type: application/json

{
    "customer_name": "Tyler Derden",
    "items": [
        {
            "product_id": "latte",
            "quantity": 2
        },
        {
            "product_id": "muffin",
            "quantity": 1
        }
    ]
}
```

- **Add\Update menu item Request:**
```http 
POST /orders
Content-Type: application/json

{
        "product_id": "latte",
        "name": "Caffe Latte",
        "description": "Espresso with steamed milk",
        "price": 3.5,
        "ingredients": [
            {
                "ingredient_id": "espresso_shot",
                "quantity": 1
            },
            {
                "ingredient_id": "milk",
                "quantity": 200
            }
        ]
    }
```

- **Add\Update inventory item Request:**
```http 
POST /orders
Content-Type: application/json

{
        "ingredient_id": "espresso_shot",
        "name": "Espresso Shot",
        "quantity": 490,
        "unit": "shots"
    }
```


- **Total Sales Aggregation Response:**
```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "total_sales": 1500.50
}
```

# Error Examples:

```http 
HTTP/1.1 400 Bad Request
Content-Type: application/json

{
  "error": "Invalid product ID in order items."
}
```

```http
HTTP/1.1 400 Bad Request
Content-Type: application/json

{
  "error": "Insufficient inventory for ingredient 'Milk'. Required: 200ml, Available: 150ml."
}

```

## Database stored in given directory (--dir flag)(Default: data)

#### JSON Files:

1. `orders.json`:

```json
[
  {
    "order_id": "order123",
    "customer_name": "Alice Smith",
    "items": [
      {
        "product_id": "latte",
        "quantity": 2
      },
      {
        "product_id": "muffin",
        "quantity": 1
      }
    ],
    "status": "open",
    "created_at": "2023-10-01T09:00:00Z"
  },
  {
    "order_id": "order124",
    "customer_name": "Bob Johnson",
    "items": [
      {
        "product_id": "espresso",
        "quantity": 1
      }
    ],
    "status": "closed",
    "created_at": "2023-10-01T09:30:00Z"
  }
]
```

2. `menu_items.json`:
```json
[
  {
    "product_id": "latte",
    "name": "Caffe Latte",
    "description": "Espresso with steamed milk",
    "price": 3.50,
    "ingredients": [
      {
        "ingredient_id": "espresso_shot",
        "quantity": 1
      },
      {
        "ingredient_id": "milk",
        "quantity": 200
      }
    ]
  },
  {
    "product_id": "muffin",
    "name": "Blueberry Muffin",
    "description": "Freshly baked muffin with blueberries",
    "price": 2.00,
    "ingredients": [
      {
        "ingredient_id": "flour",
        "quantity": 100
      },
      {
        "ingredient_id": "blueberries",
        "quantity": 20
      },
      {
        "ingredient_id": "sugar",
        "quantity": 30
      }
    ]
  },
  {
    "product_id": "espresso",
    "name": "Espresso",
    "description": "Strong and bold coffee",
    "price": 2.50,
    "ingredients": [
      {
        "ingredient_id": "espresso_shot",
        "quantity": 1
      }
    ]
  }
]
```

**Note:** The ingredients field in each menu item lists the ingredients required to prepare that item. The quantity is specified in units appropriate for the ingredient (e.g., grams, milliliters).

3. `inventory.json`:

```json
[
  {
    "ingredient_id": "espresso_shot",
    "name": "Espresso Shot",
    "quantity": 500, // Number of shots
    "unit": "shots"
  },
  {
    "ingredient_id": "milk",
    "name": "Milk",
    "quantity": 5000, // In milliliters
    "unit": "ml"
  },
  {
    "ingredient_id": "flour",
    "name": "Flour",
    "quantity": 10000, // In grams
    "unit": "g"
  },
  {
    "ingredient_id": "blueberries",
    "name": "Blueberries",
    "quantity": 2000,  // In grams
    "unit": "g"
  },
  {
    "ingredient_id": "sugar",
    "name": "Sugar",
    "quantity": 5000, // In grams
    "unit": "g"
  }
]
```




### Usage
s

```sh
$ ./hot-coffee --help
Coffee Shop Management System

Usage:
  hot-coffee [--port <N>] [--dir <S>] 
  hot-coffee --help

Options:
  --help       Show this screen.
  --port N     Port number.
  --dir S      Path to the data directory.
```



---
## Author

This project has been done by:

tkoszhan, malpamys