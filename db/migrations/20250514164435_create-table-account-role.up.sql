CREATE TABLE IF NOT EXISTS account_role (
    account_role_id BIGSERIAL PRIMARY KEY,
    account_id int NOT NULL,
    role_id int NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT fk_account
        FOREIGN KEY(account_id)
            REFERENCES account(account_id) 
);