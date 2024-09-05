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
