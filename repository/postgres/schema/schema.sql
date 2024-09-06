CREATE TABLE tenants (
    id text DEFAULT gen_random_text () PRIMARY KEY,
    name text NOT NULL,
    last_modified_at timestamp with time zone NOT NULL,
    last_modified_by text NOT NULL,
    deleted_at timestamp with time zone
);

CREATE TABLE skills (
    id text DEFAULT gen_random_text () PRIMARY KEY,
    name text NOT NULL,
    tenant_id text NOT NULL REFERENCES tenants (id),
    last_modified_at timestamp with time zone NOT NULL,
    last_modified_by text NOT NULL REFERENCES users (id),
    deleted_at timestamp with time zone
);

CREATE TABLE users (
    id text DEFAULT gen_random_text () PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email_address text NOT NULL,
    tenant_id text NOT NULL REFERENCES tenants (id),
    last_modified_at timestamp with time zone NOT NULL,
    last_modified_by text NOT NULL,
    deleted_at timestamp with time zone
);

CREATE TABLE user_presences (
    user_id text PRIMARY KEY,
    presence text NOT NULL,
    last_modified_at timestamp with time zone NOT NULL,
    last_modified_by text NOT NULL
);

CREATE TABLE user_skills (
    user_id string PRIMARY KEY REFERENCES users (id),
    skill_id string NOT NULL REFERENCES skills (id),
    skill_level integer NOT NULL
);

CREATE TABLE queues (
    id text DEFAULT gen_random_uuid() PRIMARY KEY,
    name text NOT NULL,
    tenant_id text NOT NULL REFERENCES tenants(id),
    last_modified_at timestamp with time zone NOT NULL,
    last_modified_by text NOT NULL REFERENCES users(id),
    deleted_at timestamp with time zone
);

CREATE TABLE customers (
    id text DEFAULT gen_random_uuid() PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    last_modified_at timestamp with time zone NOT NULL,
    last_modified_by text NOT NULL REFERENCES users(id),
    deleted_at timestamp with time zone,
    tenant_id text NOT NULL REFERENCES tenants(id)
);

CREATE TABLE customer_phone_numbers (
    id text DEFAULT gen_random_text() PRIMARY KEY,
    customer_id text NOT NULL REFERENCES customers(id),
    phone_number text NOT NULL,
    type text NOT NULL,
    last_modified_at timestamp with time zone NOT NULL,
    last_modified_by text NOT NULL REFERENCES users(id)
);

CREATE TABLE customer_notes (
    id text DEFAULT gen_random_uuid() PRIMARY KEY,
    note text NOT NULL,
    customer_id text NOT NULL REFERENCES customers(id),
    last_modified_at timestamp with time zone NOT NULL,
    last_modified_by text NOT NULL REFERENCES users(id)
);

CREATE TABLE customer_email_addresses (
    id text DEFAULT gen_random_uuid() PRIMARY KEY,
    email_address text NOT NULL,
    type text NOT NULL,
    customer_id text NOT NULL REFERENCES customers(id),
    last_modified_at timestamp with time zone NOT NULL,
    last_modified_by text NOT NULL REFERENCES users(id)
);

CREATE TABLE customer_addresses (
    id text DEFAULT gen_random_uuid() PRIMARY KEY,
    street_address text NOT NULL,
    state text,
    zip_code text,
    country text,
    customer_id text NOT NULL REFERENCES customers(id),
    last_modified_at timestamp with time zone NOT NULL,
    last_modified_by text NOT NULL REFERENCES users(id)
);

CREATE TABLE interactions (
    id text DEFAULT gen_random_uuid() PRIMARY KEY,
    queue_id text NOT NULL REFERENCES queues(id),
    state text NOT NULL,
    state_modified_at timestamp with time zone NOT NULL,
    user_id text REFERENCES users(id),
    tenant_id text NOT NULL REFERENCES tenants(id),
    created_at timestamp with time zone NOT NULL,
    type text NOT NULL
);

CREATE TABLE interaction_notes (
    id text DEFAULT gen_random_uuid() PRIMARY KEY,
    note text NOT NULL,
    interaction_id text NOT NULL REFERENCES interactions(id),
    last_modified_at timestamp with time zone NOT NULL,
    last_modified_by text NOT NULL REFERENCES users(id)
);
