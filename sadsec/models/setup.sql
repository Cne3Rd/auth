CREATE TABlE IF NOT EXISTS users(
    id SERIAL,
    username VARCHAR(255) NOT NULL UNIQUE PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    iv  VARCHAR(255) NOT NULL,
    is_active  boolean
    verification_token  VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
);