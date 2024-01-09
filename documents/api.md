# Bizsearch - APIs

## /business

### GET /search

Request:

```json
{
    parameter: {
        latitude: float,
        longitude: float,
        radius: integer,
        limit: integer,
        offset: integer
    }
}
```

Response:

```json
{
    data: [
        {
            id: integer,
            alias: string,
            name: string,
            image_url: string,
            is_closed: boolean,
            url: string,
            review_count: integer,
            categories: [
                {
                    alias: string,
                    title: string
                },
            ],
            rating: float,
            coordinates: {
                longitude: float,
                latitude: float
            },
            transactions: [
                string,
                string
            ],
            price: string,
            location: {
                address1: string,
                address2: string,
                address3: string,
                city: string,
                zip_code: string,
                country: string,
                state: string,
                display_address: [
                    string,
                    string
                ]
            },
            phone: string,
            display_phone: string,
            distance: float
        }
    ],
    total: integer,
    coordinates: {
        longitude: float,
        latitude: float
    }
}
```

1. Query table `businesses`
2. Join with `reviews` table and aggregate review_counts
3. Filter business by `review_counts` more than 0
4. Calculate `distance` by given `latitude` and `longitude` parameter compared to business latitude and longitude data and filter by less than given radius parameter
5. Sort by shortest distance
6. Join with `business_categories` table to get business's categories
7. Join with `business_transaction_types` table to get business's transaction types
8. Count total businesses data based on given parameter
9. Limit and offset by given limit and offset parameter

### POST /

Request:

```json
{
    body: {
        alias: string,
        name: string,
        image_url: string,
        categories: [
            {
                id: bigint
            }
        ],
        coordinates: {
            latitude: float,
            longitude: float
        },
        transactions: [
            string
        ],
        price: string,
        location: {
            address1: string,
            address2: string,
            address3: string,
            city: string,
            zip_code: string,
            country: string,
            state: string,
            display_address: [
                string,
                string
            ]
        },
        phone: string,
        display_phone: string
    }
}
```

Response:

```json
{
    data: {
        id: integer,
        alias: string,
        name: string,
        image_url: string,
        url: string,
        categories: [
            {
                alias: string,
                title: string
            },
        ],
        rating: float,
        coordinates: {
            longitude: float,
            latitude: float
        },
        transactions: [
            string,
            string
        ],
        price: string,
        location: {
            address1: string,
            address2: string,
            address3: string,
            city: string,
            zip_code: string,
            country: string,
            state: string,
            display_address: [
                string,
                string
            ]
        },
        phone: string,
        display_phone: string
    }
}
```

Steps:

1. Validate `categories`
2. Insert into `business_transaction_types` table
3. Insert into `business_categories` table
4. Insert into `businesses` table

### PUT /

Request:

```json
{
    body: {
        id: integer,
        alias: string,
        name: string,
        image_url: string,
        categories: [
            {
                id: bigint
            }
        ],
        coordinates: {
            latitude: float,
            longitude: float
        },
        transactions: [
            string
        ],
        price: string,
        location: {
            address1: string,
            address2: string,
            address3: string,
            city: string,
            zip_code: string,
            country: string,
            state: string,
            display_address: [
                string,
                string
            ]
        },
        phone: string,
        display_phone: string
    }
}
```

Response:

```json
{
    data: {
        id: integer,
        alias: string,
        name: string,
        image_url: string,
        url: string,
        categories: [
            {
                alias: string,
                title: string
            },
        ],
        rating: float,
        coordinates: {
            longitude: float,
            latitude: float
        },
        transactions: [
            string,
            string
        ],
        price: string,
        location: {
            address1: string,
            address2: string,
            address3: string,
            city: string,
            zip_code: string,
            country: string,
            state: string,
            display_address: [
                string,
                string
            ]
        },
        phone: string,
        display_phone: string
    }
}
```

Steps:

1. Validate `categories`
2. Compare `categories` and `transaction-types`, then delete the unused data and add the new data
3. Update `businesses` data with given id

### DEL /

Request:

```json
{
    request: {
        id: integer
    }
}
```

Response:

```json
{
    header: {
        status: 200
    }
}
```

Steps:

1. Find business data with given id from `businesses` table
2. Delete business categories data
3. Delete business transactions data
4. Delete business reviews data
5. Delete business data from `businesses` table
