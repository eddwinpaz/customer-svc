CREATE TABLE public."content" (
	id serial4 NOT NULL,
	content_uuid bpchar(255) NOT NULL,
	"text" int4 NOT NULL,
	customer_id int4 NOT NULL,
	created_on timestamp(0) NOT NULL,
	updated_on timestamp(0) NOT NULL,
	CONSTRAINT content_pkey PRIMARY KEY (id)
);

CREATE TABLE public.customer (
	id serial4 NOT NULL,
	customer_uuid varchar NOT NULL,
	first_name varchar NOT NULL,
	last_name varchar NOT NULL,
	email varchar NOT NULL,
	created_on timestamp(0) NOT NULL,
	last_login timestamp(0) NOT NULL,
	"password" bpchar(255) NOT NULL,
	CONSTRAINT customer_pkey PRIMARY KEY (id)
);