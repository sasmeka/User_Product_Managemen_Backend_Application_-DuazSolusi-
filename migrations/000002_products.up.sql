CREATE TABLE public.products (
	id_product uuid NOT NULL DEFAULT gen_random_uuid(),
	id_user uuid NOT NULL,
	name_product varchar(255) NOT NULL,
	description text NULL,
	price int NOT NULL DEFAULT 0,
	stock int NOT NULL DEFAULT 0,
	create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	update_at timestamp NULL,
	CONSTRAINT products_pkey PRIMARY KEY (id_product),
	CONSTRAINT products_id_user_fkey FOREIGN KEY (id_user) REFERENCES public.users(id_user)
);