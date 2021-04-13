CREATE TABLE public.users (
    id         serial       NOT NULL,
    name       varchar(40)  NULL,
    email      varchar(254) NOT NULL,
    CONSTRAINT users_email_key UNIQUE (email),
    CONSTRAINT users_pkey PRIMARY KEY (id)
);