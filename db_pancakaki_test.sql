CREATE TABLE tbl_admin(
	id SERIAL PRIMARY KEY,
<<<<<<< HEAD
	name VARCHAR(50),
	password VARCHAR(100),
	is_delete bool default false,
	FOREIGN KEY (role_id) REFERENCES tbl_bank(id)
	);

CREATE TABLE tbl_store(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50),
	description VARCHAR(50),
	address TEXT,
	customer_id INT,
	is_delete bool default false
);

CREATE TABLE tbl_employee(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50),
	password VARCHAR(50),
	store_id INT,
	is_delete bool default false
);

CREATE TABLE tbl_customer(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50),
	password VARCHAR(50),
	no_hp BIGINT UNIQUE,
	email VARCHAR(50),
	address TEXT,
	photo VARCHAR(200),
	loyalti INT,
	balance INT,
	bank_id INT,
	account_number BIGINT,
	is_delete bool default false,
	FOREIGN KEY (bank_id) REFERENCES tbl_bank(id)
	);

CREATE TABLE tbl_role(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50),
	is_delete bool default false
	);

CREATE TABLE tbl_bank(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50),
	is_delete bool default false
);

CREATE TABLE tbl_merk(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50),
	store_id INT,
	is_delete bool default false,
	FOREIGN KEY (store_id) REFERENCES tbl_store(id)
	);

CREATE TABLE tbl_courier(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50),
	interval smallint,
	store_id INT,
	is_delete bool default false,
	FOREIGN KEY (store_id) REFERENCES tbl_store(id)
	);
	
CREATE TABLE tbl_discount(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50),
	discount SMALLINT,
	store_id INT,
	is_delete bool default false,
	FOREIGN KEY (store_id) REFERENCES tbl_store(id)
	);
	
CREATE TABLE tbl_product(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50),
	price INT,
	stock INT,
	description TEXT,
	created_at date,
	update_at date,
	discount_id INT,
	merk_id INT,
	store_id INT,
	is_delete bool default false,
	FOREIGN KEY (discount_id) REFERENCES tbl_discount(id),
	FOREIGN KEY (merk_id) REFERENCES tbl_merk(id),
	FOREIGN KEY (store_id) REFERENCES tbl_store(id)
	);
	
CREATE TABLE tbl_product_image(
id SERIAL PRIMARY KEY,
image_url VARCHAR(200),
product_id int,
is_delete bool default false,
FOREIGN KEY (product_id) REFERENCES tbl_product(id)
);

CREATE TABLE review(
id SERIAL PRIMARY KEY,
review TEXT,
product_id INT,
customer_id INT,
is_delete bool default false,
FOREIGN KEY (product_id) REFERENCES tbl_product(id),
FOREIGN KEY (customer_id) REFERENCES tbl_customer(id)
);
=======
	username VARCHAR(100), -- Pancakaki123 ---
	password VARCHAR(100), -- Pancakaki123 ---
	is_delete bool,
	);
	
CREATE TABLE tbl_membership(
	id SERIAL PRIMARY KEY,
	name VARCHAR(100), -- Gold ---
	tax DOUBLE PRECISION -- 5.0 
	price BIGINT --- 4000000
	is_delete bool,
	);

CREATE TABLE tbl_bank(
	id SERIAL PRIMARY KEY, 
	name VARCHAR(100), ----- Mandiri 
	bank_account BIGINT UNIQUE ------ 0233488485934534334
	);

	
CREATE TABLE tbl_merk(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50)); ------- Samsung

CREATE TABLE tbl_owner(
	id SERIAL PRIMARY KEY,
	name VARCHAR(100), -------	Chauzar
	no_hp BIGINT UNIQUE, ------------ 08345343344
	email VARCHAR(50), ------------- cha@gmail.com
	password VARCHAR(100), ------------ rahasia
	membership_id INT,
	FOREIGN KEY (membership_id) REFERENCES tbl_membership (id) -------- 1 ( refer)
	)

CREATE TABLE tbl_store(
	id SERIAL PRIMARY KEY,
	name VARCHAR(100), ------ Toko Semangat 45
	no_hp BIGINT UNIQUE, -------- 089234435345 -
	email VARCHAR(50), ---------- Toko@gmail.com
	address TEXT, ----------- Jln raya Adhmad Yani 
	name_bank VARCHAR(50),	---------------- Mandiri ------------
	name_account varchar(100), ---------------- Toko Bagus ---------
	bank_account BIGINT UNIQUE, ----------- 0892333434534342334
	is_deleted bool
	);
	
CREATE TABLE tbl_discount(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50), ------------- Lebaran Idul Fitri
	discount SMALLINT, ------------- 2 
	store_id INT, ------------ 1
	FOREIGN KEY (store_id) REFERENCES tbl_store (id)); 
	
CREATE TABLE tbl_product(
	id SERIAL PRIMARY KEY,
	name VARCHAR(100), ------------ S7
	price DOUBLE PRECISION, ---------- 20000000
	stock INT, ------------- 100
	description TEXT, ------------ Ini hape terbagus loh
	tax DOUBLE PRECISION, ---------- 5 (Isi Default Berdasarkan Membership)
	shipping_cost DOUBLE PRECISION, ------- 15000
	discount_id int, ----- 1 ( Bisa Kosong )
	merk_id int, --------- 2
	store_id int,---------- 1
	is_deleted bool,
	FOREIGN KEY (merk_id) REFERENCES tbl_merk (id)),
	FOREIGN KEY (store_id) REFERENCES tbl_store (id);

CREATE TABLE tbl_product_image(
id SERIAL PRIMARY KEY,
image_url VARCHAR(200), ------------- link photo 
product_id int);

CREATE TABLE tbl_customer(
	id SERIAL PRIMARY KEY,
	name VARCHAR(100), ------------- Chauzar
	no_hp BIGINT UNIQUE, ------------- 089233435345
	address text, ------------------ Jln Raya Ciburial
	password VARCHAR(100)); ----------------- Rahasiaa
>>>>>>> chauzar


CREATE TABLE tbl_transaction_order(
id SERIAL PRIMARY KEY,
quantity INT, ---------- 3
total int, ----------- harga product * qty
customer_id int,
product_id int,
<<<<<<< HEAD
FOREIGN KEY (customer_id) REFERENCES tbl_customer(id),
FOREIGN KEY (product_id) REFERENCES tbl_product(id)
);

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
=======
detail_order_id int);
>>>>>>> chauzar

CREATE TABLE tbl_transaction_detail_order(
id SERIAL PRIMARY KEY,
buy_date date, ------------ 25-05-30
status VARCHAR(50), ----------- sedang disiapkan (Secara Default Udah Ke Isi Begitu)
total_price BIGINT, --------- jumlah dari total
Photo VARCHAR(100));

