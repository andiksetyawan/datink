BEGIN;

-- users Table
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       name VARCHAR NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       birthdate DATE,
                       gender VARCHAR(10),
                       pictures VARCHAR[],
                       last_login TIMESTAMP,
                       verified BOOLEAN,
                       description TEXT,
                       location VARCHAR(255),
                       is_premium BOOLEAN NOT NULL DEFAULT FALSE,
                       created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
                       updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
                       deleted_at timestamp with time zone
);

-- swipes Table
CREATE TABLE swipes (
                        id SERIAL PRIMARY KEY,
                        user_id INT REFERENCES users(id),
                        friend_id INT REFERENCES users(id),
                        swipe_direction VARCHAR, -- 0 for pass, 1 for like
                        created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
                        updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
                        deleted_at timestamp with time zone
);

CREATE TABLE swipe_counts (
                              id SERIAL PRIMARY KEY,
                              user_id INT REFERENCES users(id),
                              swipe_date DATE not null,-- 0 for pass, 1 for like
                              total INT not null,
                              created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
                              updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
                              deleted_at timestamp with time zone
);

CREATE UNIQUE INDEX  swipe_counts_user_id_date_unique ON swipe_counts (user_id, swipe_date);

-- premium_packages Table
CREATE TABLE packages (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(255) NOT NULL,
                          description TEXT,
                          price DECIMAL(10, 2) NOT NULL,
                          created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
                          updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
                          deleted_at timestamp with time zone
);

-- user_premiums Table
CREATE TABLE user_packages (
                               id SERIAL PRIMARY KEY,
                               user_id INT REFERENCES users(id),
                               package_id INT REFERENCES packages(id),
                               start_date timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
                               end_date timestamp with time zone,
                               total_amount DECIMAL(10, 2),

                               created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
                               updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
                               deleted_at timestamp with time zone
);


INSERT INTO packages (name, description, price)
values  ('premium1', 'free', 10000.00);

COMMIT;
