CREATE SCHEMA IF NOT EXISTS core;

CREATE TABLE IF NOT EXISTS core.user (
    id             UUID       NOT NULL  DEFAULT uuid_generate_v4() PRIMARY KEY
    , telegram_id  BIGINT     NOT NULL
    , is_admin     BOOLEAN    NOT NULL  DEFAULT false
    , created_at   TIMESTAMP  NOT NULL  DEFAULT CURRENT_TIMESTAMP
    , updated_at   TIMESTAMP  
    , deleted_at   TIMESTAMP  
);

CREATE UNIQUE INDEX idx_core_user_telegram_id ON core.user (telegram_id);


CREATE TABLE IF NOT EXISTS core.link (
    id             UUID       NOT NULL  DEFAULT uuid_generate_v4() PRIMARY KEY
    , user_id      UUID       NOT NULL  REFERENCES core.user(id)
    , chat_id      BIGINT     NOT NULL
    , link         TEXT       NOT NULL
    , created_at   TIMESTAMP  NOT NULL  DEFAULT CURRENT_TIMESTAMP
    , updated_at   TIMESTAMP  
    , deleted_at   TIMESTAMP  
)
