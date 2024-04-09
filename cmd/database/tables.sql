-- public.categories definition

-- Drop table

-- DROP TABLE public.categories;

CREATE TABLE public.categories (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	name varchar(50) NOT NULL,
	status varchar(20) NOT NULL,
	CONSTRAINT categories_pkey PRIMARY KEY (id),
	CONSTRAINT uk_name UNIQUE (name)
);

-- Permissions

ALTER TABLE public.categories OWNER TO postgres;
GRANT ALL ON TABLE public.categories TO postgres;




-- public.image_model definition

-- Drop table

-- DROP TABLE public.image_model;

CREATE TABLE public.images (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	alt varchar(255) NULL,
	description text NULL,
	filename varchar(255) NOT NULL,
	public_id varchar(255) NULL,
	status varchar(255) NOT NULL,
	title varchar(255) NULL,
	url varchar(255) NOT NULL,
	miniature_url text DEFAULT ''::text NOT NULL,
	CONSTRAINT image_pkey PRIMARY KEY (id),
	CONSTRAINT uk_url UNIQUE (url),
	CONSTRAINT uk_filename UNIQUE (filename)
);

-- Permissions

ALTER TABLE public.images OWNER TO postgres;
GRANT ALL ON TABLE public.images TO postgres;


-- public.roles definition

-- Drop table

-- DROP TABLE public.roles;

CREATE TABLE public.roles (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	name varchar(20) NULL,
	CONSTRAINT roles_pkey PRIMARY KEY (id)
);

-- Permissions

ALTER TABLE public.roles OWNER TO postgres;
GRANT ALL ON TABLE public.roles TO postgres;


-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	created_at timestamptz(6) NULL,
	email varchar(255) NULL,
	modified_at timestamptz(6) NULL,
	password varchar(255) NOT NULL,
	username varchar(255) NOT NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id),	
	CONSTRAINT uk_email UNIQUE (email)
);

-- Permissions

ALTER TABLE public.users OWNER TO postgres;
GRANT ALL ON TABLE public.users TO postgres;


-- public.users_roles definition

-- Drop table

-- DROP TABLE public.users_roles;

CREATE TABLE public.users_roles (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	role_id uuid NOT NULL,
	CONSTRAINT users_roles_pkey PRIMARY KEY (id, role_id)
);

-- Permissions

ALTER TABLE public.users_roles OWNER TO postgres;
GRANT ALL ON TABLE public.users_roles TO postgres;


-- public.tb_users_roles foreign keys

ALTER TABLE public.users_roles ADD CONSTRAINT fk_roles FOREIGN KEY (role_id) REFERENCES public.roles(id);


-- public.post_model definition

-- Drop table

-- DROP TABLE public.post_model;

CREATE TABLE public.posts (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	content varchar(255) NULL,
	created_at timestamptz(6) NOT NULL,
	status int2 NULL,
	title varchar(255) NULL,
	updated_at timestamptz(6) NULL,
	author_id uuid NOT NULL,
	category_id uuid NULL,
	image_id uuid NULL,
	CONSTRAINT post_pkey PRIMARY KEY (id)
);

-- Permissions

ALTER TABLE public.posts OWNER TO postgres;
GRANT ALL ON TABLE public.posts TO postgres;


-- public.posts foreign keys

ALTER TABLE public.posts ADD CONSTRAINT fk_user FOREIGN KEY (author_id) REFERENCES public.users(id);
ALTER TABLE public.posts ADD CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES public.categories(id);
ALTER TABLE public.posts ADD CONSTRAINT fk_image FOREIGN KEY (image_id) REFERENCES public.images(id);

-- public.logging_model definition

-- Drop table

-- DROP TABLE public.logging_model;

CREATE TABLE public.logging (
	id uuid NOT NULL,
	data varchar(255) NULL,
	log varchar(255) NULL,
	message varchar(255) NULL,
	timestamp timestamptz(6) NULL,
	user_id uuid NOT NULL,
	CONSTRAINT logging_pkey PRIMARY KEY (id)
);

-- Permissions

ALTER TABLE public.logging OWNER TO postgres;
GRANT ALL ON TABLE public.logging TO postgres;


ALTER TABLE public.images ADD COLUMN category_id uuid NULL;
ALTER TABLE public.images ADD CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES public.categories(id);
