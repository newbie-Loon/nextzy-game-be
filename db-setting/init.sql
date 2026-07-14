CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,

    guest BOOLEAN NOT NULL DEFAULT FALSE,
    user_id VARCHAR(255) NOT NULL DEFAULT gen_random_uuid(),
    point INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX idx_users_user_id ON users(user_id);

ALTER TABLE users ALTER COLUMN user_id SET DEFAULT gen_random_uuid();


CREATE TABLE point_histories (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,

    user_id VARCHAR(255) NOT NULL,
    point INTEGER NOT NULL,
    earned_date TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE INDEX idx_point_histories_user_id
ON point_histories(user_id);
CREATE INDEX idx_point_histories_earned_date
ON point_histories(earned_date);