CREATE TABLE public.users (
	id_user uuid NOT NULL DEFAULT gen_random_uuid(),
	full_name text NULL,
	email text NOT NULL,
	pass text NOT NULL,
	"role" varchar(6) NOT NULL DEFAULT 'user'::character varying,
	create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	update_at timestamp NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id_user)
);