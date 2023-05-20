CREATE TABLE tbl_admin(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50),
	password VARCHAR(100)
	);
	
CREATE TABLE tbl_merk(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50));

CREATE TABLE tbl_packet(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50),
	interval smallint);
	
CREATE TABLE tbl_role(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50));
	
CREATE TABLE tbl_discount(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50),
	discount SMALLINT);
	
CREATE TABLE tbl_customer(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50),
	no_hp BIGINT,
	address TEXT,
	photo VARCHAR(200),
	bank_id INT,
	account_number BIGINT
	)

CREATE TABLE tbl_bank(
id SERIAL PRIMARY KEY,
name VARCHAR(50))

CREATE TABLE tbl_product(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50),
	price INT,
	stock INT,
	description TEXT,
	created_at date,
	update_at date,
	is_delete bool,
	discount_id int,
	merk_id int);
	
CREATE TABLE tbl_product_image(
id SERIAL PRIMARY KEY,
image_url VARCHAR(200),
product_id int);

CREATE TABLE review(
id SERIAL PRIMARY KEY,
review TEXT,
product_id INT,
customer_id INT);


CREATE TABLE tbl_transaction_order(
id SERIAL PRIMARY KEY,
quantity INT,
buy_date date,
status VARCHAR(100),
total int,
customer_id int,
product_id int);

ALTER TABLE tbl_admin
ADD COLUMN role_id INT;

ALTER TABLE tbl_admin
DROP COLUMN role_id;

ALTER TABLE tbl_customer
DROP COLUMN role_id;

ALTER TABLE tbl_customer
DROP COLUMN account_number;

ALTER TABLE tbl_customer
DROP COLUMN bank_id;

DROP TABLE tbl_role;


ALTER TABLE tbl_customer
ADD COLUMN balance BIGINT;


ALTER TABLE tbl_admin
ADD CONSTRAINT fk_admin_role
FOREIGN KEY (role_id)
REFERENCES tbl_role(id);

ALTER TABLE tbl_customer
ADD COLUMN role_id INT;

ALTER TABLE tbl_transaction_order
ADD COLUMN packet_id INT;

ALTER TABLE tbl_bank
ADD COLUMN bank_account BIGINT;

ALTER TABLE tbl_customer
ADD CONSTRAINT fk_customer_role
FOREIGN KEY (role_id)
REFERENCES tbl_role(id);

ALTER TABLE tbl_product
DROP COLUMN review_id;

ALTER TABLE tbl_product
ADD CONSTRAINT fk_product_discount
FOREIGN KEY (discount_id)
REFERENCES tbl_discount(id);

ALTER TABLE tbl_product
ADD CONSTRAINT fk_product_merk
FOREIGN KEY (merk_id)
REFERENCES tbl_merk(id);

ALTER TABLE tbl_product_image
ADD CONSTRAINT fk_product_image_product
FOREIGN KEY (product_id)
REFERENCES tbl_product(id);

ALTER TABLE tbl_review
ADD CONSTRAINT fk_review_product
FOREIGN KEY (product_id)
REFERENCES tbl_product(id);

ALTER TABLE tbl_review
ADD CONSTRAINT fk_review_customer
FOREIGN KEY (customer_id)
REFERENCES tbl_customer(id);

ALTER TABLE tbl_transaction_order
ADD CONSTRAINT fk_transaction_order_customer
FOREIGN KEY (customer_id)
REFERENCES tbl_customer(id);

ALTER TABLE tbl_transaction_order
ADD CONSTRAINT fk_transaction_order_product
FOREIGN KEY (product_id)
REFERENCES tbl_product(id);

ALTER TABLE tbl_transaction_order
ADD CONSTRAINT fk_transaction_order_packet
FOREIGN KEY (packet_id)
REFERENCES tbl_packet(id);