CREATE TABLE public."user" (
	id varchar(100) NOT NULL,
	email varchar(100) NOT NULL,
	phone_number varchar(20) NOT NULL,
	"name" varchar(250) NOT NULL,
	"password" varchar(150) NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	created_by varchar(150) NOT NULL,
	updated_at timestamptz NULL,
	updated_by varchar(150) DEFAULT NULL::character varying NULL,
	deleted_at timestamptz NULL,
	deleted_by varchar(150) DEFAULT NULL::character varying NULL,
	CONSTRAINT user_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_user_email ON public."user" USING btree (email);
CREATE INDEX idx_user_id ON public."user" USING btree (id);
CREATE INDEX idx_user_name ON public."user" USING btree (name);
CREATE INDEX idx_user_password ON public."user" USING btree (password);
CREATE INDEX idx_user_phone_number ON public."user" USING btree (phone_number);

INSERT INTO public."user" (id, email, phone_number, "name", "password", created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('40026f7a-6b7d-48e2-a63d-79ff3a0a8c5d', 'ahmadmuharik@gmail.com', '087772488816', 'muharik', '$2a$10$2RLJwYzS2wCGTkgr145miOvXe52nI0JRE.8uOgqEmM8ulmmRufWle', '2024-04-25 00:01:42.893', 'muharik', NULL, NULL, NULL, NULL);


CREATE TABLE public.shop (
	id varchar(100) NOT NULL,
	email varchar(100) NOT NULL,
	phone_number varchar(20) NOT NULL,
	"name" varchar(250) NOT NULL,
	address varchar(250) NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	created_by varchar(150) NOT NULL,
	updated_at timestamptz NULL,
	updated_by varchar(150) DEFAULT NULL::character varying NULL,
	deleted_at timestamptz NULL,
	deleted_by varchar(150) DEFAULT NULL::character varying NULL,
	CONSTRAINT shop_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_shop_address ON public.shop USING btree (address);
CREATE INDEX idx_shop_email ON public.shop USING btree (email);
CREATE INDEX idx_shop_id ON public.shop USING btree (id);
CREATE INDEX idx_shop_name ON public.shop USING btree (name);
CREATE INDEX idx_shop_phone_number ON public.shop USING btree (phone_number);

INSERT INTO public.shop (id, email, phone_number, "name", address, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('97097f86-e32b-4c44-92bb-048a762bb49d', 'ahmadmuharik@gmail.com', '087772488816', 'Muharik', 'Cibinong, Bogor', '2024-04-25 17:19:05.601', 'Muharik', NULL, NULL, NULL, NULL);
INSERT INTO public.shop (id, email, phone_number, "name", address, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('d1fcaf8e-266e-4c66-b042-c375c4c8d3c3', 'rizki@gmail.com', '0882009630636', 'Rizki', 'Cikaret, Bogor', '2024-04-25 17:19:50.341', 'Rizki', NULL, NULL, NULL, NULL);
INSERT INTO public.shop (id, email, phone_number, "name", address, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('1f2f9523-9df2-465b-a041-38813cae7a38', 'albar@gmail.com', '0882009630612', 'Albar', 'Margonda, Depok', '2024-04-25 17:21:04.287', 'Albar', NULL, NULL, NULL, NULL);
INSERT INTO public.shop (id, email, phone_number, "name", address, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('10996d96-2b0b-493d-9d7a-511ae9f7ff32', 'faisal@gmail.com', '0882009630614', 'Faisal', 'Cikini, Jakarta', '2024-04-25 17:21:40.134', 'Faisal', NULL, NULL, NULL, NULL);
INSERT INTO public.shop (id, email, phone_number, "name", address, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('047e2c7b-5ec5-4541-93a8-f1b1c0d33ea7', 'zaenal@gmail.com', '0882009630699', 'Zaneal Telur', 'Parigi, Bintaro', '2024-04-25 17:22:20.994', 'Zaneal Telur', NULL, NULL, NULL, NULL);
INSERT INTO public.shop (id, email, phone_number, "name", address, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('b8fb07fc-dbcc-4239-a5a5-73cdcaa28f08', 'tala@gmail.com', '0882009630688', 'Tala Anderware', 'Parigi, Bintaro', '2024-04-25 17:23:03.095', 'Tala Anderware', NULL, NULL, NULL, NULL);


CREATE TABLE public.product (
	id varchar(100) NOT NULL,
	"name" varchar(250) NOT NULL,
	price numeric NULL,
	stock numeric NULL,
	shop_id varchar(100) NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	created_by varchar(150) NOT NULL,
	updated_at timestamptz NULL,
	updated_by varchar(150) DEFAULT NULL::character varying NULL,
	deleted_at timestamptz NULL,
	deleted_by varchar(150) DEFAULT NULL::character varying NULL,
	CONSTRAINT product_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_product_id ON public.product USING btree (id);
CREATE INDEX idx_product_name ON public.product USING btree (name);
CREATE INDEX idx_product_price ON public.product USING btree (price);
CREATE INDEX idx_product_shop_id ON public.product USING btree (shop_id);
CREATE INDEX idx_product_stock ON public.product USING btree (stock);

INSERT INTO public.product (id, "name", price, stock, shop_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('675cdba4-3094-4762-9d0f-7883f3e55215', 'Chitato', 10000, 108, '97097f86-e32b-4c44-92bb-048a762bb49d', '2024-04-25 00:42:31.160', 'system', NULL, NULL, NULL, NULL);
INSERT INTO public.product (id, "name", price, stock, shop_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('49dae434-e2f3-484d-aa34-ee09a1eb29fa', 'Pocari Sweat 500ml', 5000, 191, 'd1fcaf8e-266e-4c66-b042-c375c4c8d3c3', '2024-04-25 00:43:00.247', 'system', NULL, NULL, NULL, NULL);
INSERT INTO public.product (id, "name", price, stock, shop_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('28a648aa-8873-4143-b83d-b8231b4492eb', 'Chiki ball rasa', 8000, 1000, 'd1fcaf8e-266e-4c66-b042-c375c4c8d3c3', '2024-04-25 00:45:17.399', 'system', NULL, NULL, NULL, NULL);
INSERT INTO public.product (id, "name", price, stock, shop_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('f6a5910c-7b0f-4192-beef-eae33d6bb3f9', 'Pulpen Standar', 1500, 100, 'd1fcaf8e-266e-4c66-b042-c375c4c8d3c3', '2024-04-25 00:45:49.471', 'system', NULL, NULL, NULL, NULL);
INSERT INTO public.product (id, "name", price, stock, shop_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('99a26277-4260-49cb-a758-4ff90f8982a8', 'Susu SGM', 50000, 2, '1f2f9523-9df2-465b-a041-38813cae7a38', '2024-04-25 00:46:27.879', 'system', NULL, NULL, NULL, NULL);
INSERT INTO public.product (id, "name", price, stock, shop_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('85dffee5-1c11-4d21-950f-e6043c63c052', 'Nextar', 2000, 50, '1f2f9523-9df2-465b-a041-38813cae7a38', '2024-04-25 00:47:03.649', 'system', NULL, NULL, NULL, NULL);
INSERT INTO public.product (id, "name", price, stock, shop_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('36dedde5-df77-411a-a06a-052a618c7b7e', 'Milo', 20000, 10, '10996d96-2b0b-493d-9d7a-511ae9f7ff32', '2024-04-25 00:47:24.351', 'system', NULL, NULL, NULL, NULL);
INSERT INTO public.product (id, "name", price, stock, shop_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('869b7d7a-6c6f-4004-8913-b36373b80a96', 'Sirop', 20000, 5, '047e2c7b-5ec5-4541-93a8-f1b1c0d33ea7', '2024-04-25 00:47:45.819', 'system', NULL, NULL, NULL, NULL);
INSERT INTO public.product (id, "name", price, stock, shop_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('69def395-cca9-4aa0-832d-ca9137c782d4', 'Pempes', 50000, 5, 'b8fb07fc-dbcc-4239-a5a5-73cdcaa28f08', '2024-04-25 00:48:05.602', 'system', NULL, NULL, NULL, NULL);
INSERT INTO public.product (id, "name", price, stock, shop_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('c2e8d4c4-b7be-472a-96fa-3a4169bc338d', 'Maxim Neostone', 215000, 2, '97097f86-e32b-4c44-92bb-048a762bb49d', '2024-04-25 00:42:02.683', 'system', NULL, NULL, NULL, NULL);
INSERT INTO public.product (id, "name", price, stock, shop_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('2c23969c-73d9-4c97-9f13-3d3860f030e6', 'Lifebuoy Body Foam', 25000, 5, '97097f86-e32b-4c44-92bb-048a762bb49d', '2024-04-25 00:41:39.792', 'system', NULL, NULL, NULL, NULL);


CREATE TABLE public.warehouse (
	code varchar(100) NOT NULL,
	shop_id varchar(100) NOT NULL,
	"name" varchar(250) NOT NULL,
	address varchar(250) NOT NULL,
	active bool DEFAULT true NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	created_by varchar(150) NOT NULL,
	updated_at timestamptz NULL,
	updated_by varchar(150) DEFAULT NULL::character varying NULL,
	deleted_at timestamptz NULL,
	deleted_by varchar(150) DEFAULT NULL::character varying NULL,
	CONSTRAINT warehouse_pkey PRIMARY KEY (code)
);
CREATE INDEX idx_warehouse_address ON public.warehouse USING btree (address);
CREATE INDEX idx_warehouse_code ON public.warehouse USING btree (code);
CREATE INDEX idx_warehouse_name ON public.warehouse USING btree (name);
CREATE INDEX idx_warehouse_shop_id ON public.warehouse USING btree (shop_id);

INSERT INTO public.warehouse (code, shop_id, "name", address, active, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('WHS1714055856754', '97097f86-e32b-4c44-92bb-048a762bb49d', 'WMS 2', 'Cikaret, Bogor', true, '2024-04-25 21:37:36.754', '', NULL, NULL, NULL, NULL);
INSERT INTO public.warehouse (code, shop_id, "name", address, active, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('WHS1714055574385', '97097f86-e32b-4c44-92bb-048a762bb49d', 'WMS 1', 'Cikaret, Bogor', true, '2024-04-25 21:32:54.386', '', NULL, NULL, NULL, NULL);


CREATE TABLE public."order" (
	invoice varchar(100) NOT NULL,
	user_id varchar(100) NOT NULL,
	amount numeric NULL,
	shop_id varchar(100) NOT NULL,
	payment bool DEFAULT false NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	created_by varchar(150) NOT NULL,
	updated_at timestamptz NULL,
	updated_by varchar(150) DEFAULT NULL::character varying NULL,
	deleted_at timestamptz NULL,
	deleted_by varchar(150) DEFAULT NULL::character varying NULL,
	CONSTRAINT order_pkey PRIMARY KEY (invoice)
);
CREATE INDEX idx_order_amount ON public."order" USING btree (amount);
CREATE INDEX idx_order_invoice ON public."order" USING btree (invoice);
CREATE INDEX idx_order_payment ON public."order" USING btree (payment);
CREATE INDEX idx_order_shop_id ON public."order" USING btree (shop_id);
CREATE INDEX idx_order_user_id ON public."order" USING btree (user_id);

INSERT INTO public."order" (invoice, user_id, amount, shop_id, payment, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('INV1714044665742', '40026f7a-6b7d-48e2-a63d-79ff3a0a8c5d', 555000, '97097f86-e32b-4c44-92bb-048a762bb49d', true, '2024-04-26 01:16:05.765', 'muharik', NULL, NULL, NULL, NULL);


CREATE TABLE public.order_detail (
	invoice varchar(100) NOT NULL,
	product_id varchar(100) NOT NULL,
	qty numeric NULL,
	price numeric NULL,
	amount numeric NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	created_by varchar(150) NOT NULL,
	updated_at timestamptz NULL,
	updated_by varchar(150) DEFAULT NULL::character varying NULL,
	deleted_at timestamptz NULL,
	deleted_by varchar(150) DEFAULT NULL::character varying NULL
);
CREATE INDEX idx_order_detail_amount ON public.order_detail USING btree (amount);
CREATE INDEX idx_order_detail_invoice ON public.order_detail USING btree (invoice);
CREATE INDEX idx_order_detail_price ON public.order_detail USING btree (price);
CREATE INDEX idx_order_detail_product_id ON public.order_detail USING btree (product_id);
CREATE INDEX idx_order_detail_qty ON public.order_detail USING btree (qty);

INSERT INTO public.order_detail (invoice, product_id, qty, price, amount, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('INV1714044665742', 'c2e8d4c4-b7be-472a-96fa-3a4169bc338d', 2, 215000, 430000, '2024-04-25 18:31:05.765', '', '2024-04-26 09:05:00.048', NULL, NULL, NULL);
INSERT INTO public.order_detail (invoice, product_id, qty, price, amount, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('INV1714044665742', '2c23969c-73d9-4c97-9f13-3d3860f030e6', 5, 25000, 125000, '2024-04-25 18:31:05.765', '', '2024-04-26 09:05:00.058', NULL, NULL, NULL);


CREATE TABLE public.warehouse_stock (
	code varchar(100) NOT NULL,
	invoice varchar(100) NOT NULL,
	product_id varchar(100) NOT NULL,
	qty numeric NULL,
	price numeric NULL,
	amount numeric NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	created_by varchar(150) NOT NULL,
	updated_at timestamptz NULL,
	updated_by varchar(150) DEFAULT NULL::character varying NULL,
	deleted_at timestamptz NULL,
	deleted_by varchar(150) DEFAULT NULL::character varying NULL
);
CREATE INDEX idx_warehouse_stock_amount ON public.warehouse_stock USING btree (amount);
CREATE INDEX idx_warehouse_stock_code ON public.warehouse_stock USING btree (code);
CREATE INDEX idx_warehouse_stock_invoice ON public.warehouse_stock USING btree (invoice);
CREATE INDEX idx_warehouse_stock_price ON public.warehouse_stock USING btree (price);
CREATE INDEX idx_warehouse_stock_product_id ON public.warehouse_stock USING btree (product_id);
CREATE INDEX idx_warehouse_stock_qty ON public.warehouse_stock USING btree (qty);

INSERT INTO public.warehouse_stock (code, invoice, product_id, qty, price, amount, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('WHS1714055574385', 'INV1714044665742', '2c23969c-73d9-4c97-9f13-3d3860f030e6', 5, 25000, 125000, '2024-04-26 00:23:51.657', 'system', NULL, NULL, NULL, NULL);
INSERT INTO public.warehouse_stock (code, invoice, product_id, qty, price, amount, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by) VALUES('WHS1714055574385', 'INV1714044665742', 'c2e8d4c4-b7be-472a-96fa-3a4169bc338d', 2, 215000, 430000, '2024-04-26 00:23:51.657', 'system', NULL, NULL, NULL, NULL);