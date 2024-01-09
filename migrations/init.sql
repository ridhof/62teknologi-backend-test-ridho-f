CREATE TYPE price_type as ENUM ('$', '$$', '$$$', '$$$$', '$$$$$');
CREATE TYPE transaction_type as ENUM ('pickup', 'delivery');

CREATE TABLE users (
	id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,

	name VARCHAR(127) NOT NULL,
	image_url VARCHAR(255) NOT NULL,

	created_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (name, image_url)
VALUES
 ('Ghan', 'https://i.imgur.com/2380ghn.jpeg'),
 ('Zelda', 'https://i.imgur.com/zPJO5j9.jpeg');

CREATE TABLE businesses (
	id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,

	alias VARCHAR(127) UNIQUE NOT NULL,
	name VARCHAR(255) NOT NULL,
	image_url VARCHAR(255) NOT NULL,
	transactions transaction_type ARRAY NOT NULL,
	latitude NUMERIC(9,7) NOT NULL,
	longitude NUMERIC(10, 7) NOT NULL,
	price price_type NOT NULL,
	location_address_1 VARCHAR(255) NOT NULL,
	location_address_2 VARCHAR(255),
	location_address_3 VARCHAR(255),
	city VARCHAR(255) NOT NULL,
	zip_code VARCHAR(7) NOT NULL,
	country VARCHAR(3) NOT NULL,
	state VARCHAR(3) NOT NULL,
	display_address VARCHAR(255) ARRAY NOT NULL,
	phone VARCHAR(31) NOT NULL,
	display_phone VARCHAR(31) NOT NULL,

	created_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO businesses 
 (alias, name, image_url, transactions, latitude, longitude, price, location_address_1, city, zip_code, country, state, display_address, phone, display_phone)
VALUES
 ('kismo-house', 'Kismo Coffee House Malang', 'https://i.imgur.com/NLf8tka.png', '{"pickup"}', -7.9297124, 112.6035438, '$$$', 'Jl. Baiduri Sepah No.4, Tlogomas, Kec.Lowokwaru', 'Malang', '65144', 'id', 'ejv', '{"Jl. Baiduri Sepah, Tlogomas", "Kec.Lowokwaru, Malang"}', '081936051915', '(+62) 819-3605-1915'),
 ('common-grounds-fx-sudirman', 'Common Grounds Coffee - FX Sudirman', 'https://i.imgur.com/7CYPcx3.png', '{"pickup", "delivery"}', -6.2247853, 106.8042443, '$$$$', 'FX Lifestyle Center FI.02, Jl. Jenderal Sudirman, Jl. Pintu Satu Senayan', 'Jakarta Pusat', '10270', 'id', 'jkt', '{"FX Lifestyle Centre FI.02", "Gelora, Tanah Abang"}', '08170200131', '(+62) 817-020-0131');

CREATE TABLE categories (
	id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,

	title VARCHAR(127) NOT NULL,
	alias VARCHAR(127) NOT NULL,

	created_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO categories (title, alias)
VALUES
 ('Coffee', 'coffee'),
 ('Coffee House', 'coffee-house'),
 ('Common Grounds', 'common-grounds');

CREATE TABLE business_categories (
	id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,

	business_id BIGINT NOT NULL REFERENCES businesses (id) ON DELETE RESTRICT,
	category_id BIGINT NOT NULL REFERENCES categories (id) ON DELETE RESTRICT,

	created_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO business_categories (business_id, category_id)
VALUES
 (1, 1),
 (1, 2),
 (2, 1),
 (2, 3);

CREATE TABLE reviews (
	id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,

	text TEXT NOT NULL,
	rating INT NOT NULL,
	user_id BIGINT NOT NULL REFERENCES users (id) ON DELETE RESTRICT,
	business_id BIGINT NOT NULL REFERENCES businesses (id) ON DELETE RESTRICT,

	created_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO reviews (text, rating, user_id, business_id)
VALUES
 ('such a cozy, homey, clean, warming cafe with a good interior. good for wfc.', 4, 1, 1),
 ('one of cafe in malang that should be in the list with a homey concept.', 5, 2, 1),
 ('you cannot be uncertain with the coffee, surely delicious, and have promo. the barista are warming and so helpful.', 5, 1, 2),
 ('if you are visiting here, you should go upstair and take a seat at outdoor.', 5, 2, 2);
