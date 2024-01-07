# Bizsearch - Business Search

## Tech Stack Requirement

### Backend

- Laravel / PHP / NodeJS / Java Spring Boot / Python / Golang
- PostgreSQL / MySQL
- Reference API (https://docs.developer.yelp.com/reference/v3_business_search)[https://docs.developer.yelp.com/reference/v3_business_search]

## Specification

1. Create a database that could provide data from API ref
2. Create a POST `/business` endpoint to add new data to database
3. Create a PUT `/business` endpoint to edit a data from database
4. Create a DEL `/business` endpoint to delete a data from database
5. Create a GET `/business/search` to get data from database with at least 5 parameters such as: `term`, `location`, `latitude`, `longitude`, `radius`, `categories`, `locale`, `limit`, `offset`, `sort_by`, `price`, `open_now`, `open_at`, `attributes`
6. Reusable and clean code (optional)

## Yelp Spec Review

- Only business with reviews are returned in `/business/search`

Query Params:

- `location` is string and required if `latitude` or `longitude` is not provided, i.e. "New York City", "NYC", "350 5th Ave, New York, NY 10118"
- `latitude` is a number and required if `location` is not provided. Both `latitude` and `longitude` should exist.
- `longitude` same as `latitude`.
- `term` is a string and can be filled with category or business name such as "food", "restaurant", or "sushitei". if not provided, will go to default searching from small number to popular categories.
- `radius` is an integer and max value is 40,000 meters.
- `categories` is a string separated by comma to provide list of supported categories. i.e. `bars,french`.
- `locale` is a string with format of `{language_code}_{country_code}`.
- `price` is array of integers (in a string format) separated by comma. with 1 = "$", 2 = "$$", and so on. i.e. `1,2,3` will result $, $$, and $$$.
- `open_now` is a boolean, filter businesses that are open now. `open_now` and `open_at` can't be used together.
- `open_at` is an integer, filter businesses that are open by given Unix time in the timezone of the search location.
- `attributes` is array of strings. can be combined by separating it with comma. list of options are: 
    - `hot_and_new`
    - `request_a_quote`
    - `reservation`
    - `waitlist_reservation`
    - `deals`
    - `gender_neutral_restrooms`
    - `open_to_all`
    - `wheelchair_accessible`
    - `liked_by_vegetarians`
    - `outdoor_seating`
    - `parking_garage`
    - `parking_lot`
    - `parking_street`
    - `parking_valet`
    - `parking_validated`
    - `parking_bike`
    - `restaurants_delivery`
    - `restaurants_takeout`
    - `wifi_free`
    - `wifi_paid`
- `limit` is an integer, to return number of result
- `offset` is an integer, to offset the returned results by this amount
- `sort_by` is a string to decide the algorithm modes. default is `best_match`, and the list are `best_match`, `rating`, `review_count`, or `distance`

## Sample Response Data

### Business / Businesses

- `id`, string, i.e. `j1S3NUrkB3BVT49n_e76NQ`
- `alias`, string i.e. `best-bagel-and-coffee-new-york`
- `name`, string
- `image_url`, string
- `is_closed`, boolean
- `url`, string, looks like the yelp page url of this business
- `review_count`, integer
- `categories`, relation one to many with `category`
- `rating`, double/float, max 5
- `coordinates` is consist of `latitude` and `longitude`, both are double/float
- `transactions` is a list of string, i.e. `pickup` and `delivery`
- `price` is a string `$`
- `location` contains of:
    - `address1` a string
    - `address2` a string
    - `address3` a string
    - `city` a string, i.e `New York`
    - `zip_code` is a string of number, i.e. `10001`
    - `country` is country code, i.e. `US`
    - `state` is a state code, i.e. `NY`
    - `display_address` is a list of address, i.e. `225 W 33th St` and `New York, NY 10001`
- `phone` is a string, i.e. `+12125644409`
- `display_phone` is a string, i.e. `(212) 564-4409`
- `distance` is a double/float, `5202.248023208401`

### Category / Categories

- `alias`, string, lowercase, only using `_` as symbol and no whitespace
- `title`, string

### Review

- `id` string
- `url` string, review url
- `text` string, review text
- `rating` integer with max value is 5
- `time_created` datetime with format i.e. `2024-01-01 13:50:17`
- `user` is a relation to user profile

### User

- `id` string
- `name` string, user name
- `profile_url` string, url of user profile
- `image_url` string, url of user image
