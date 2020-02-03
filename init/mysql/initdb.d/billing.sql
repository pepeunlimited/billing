CREATE DATABASE IF NOT EXISTS billing CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE billing;

#  -------------
# |subscriptions|
#  -------------
CREATE TABLE plans (
    id            BIGINT NOT NULL AUTO_INCREMENT,
    title_i18n_id BIGINT NOT NULL,
    price_id      BIGINT NOT NULL,
    product_id    BIGINT NOT NULL,
    start_at      DATETIME(3) NOT NULL,
    end_at        DATETIME(3) NOT NULL,
    length        TINYINT UNSIGNED DEFAULT 0,
    unit          CHAR(5) NULL, #day,week,month,year, if null => unlimited
    PRIMARY KEY (id)
);
CREATE TABLE subscriptions (
    id          BIGINT      NOT NULL AUTO_INCREMENT,
    user_id     BIGINT      NOT NULL,
    plan_id     BIGINT,
    start_at    DATETIME(3) NOT NULL,
    end_at      DATETIME(3) NOT NULL,
    FOREIGN KEY (plan_id) REFERENCES plans (id),
    PRIMARY KEY (id)
);
#  -------------------
# |payment_instruments|
#  -------------------
CREATE TABLE payment_instruments (
     id      BIGINT      NOT NULL AUTO_INCREMENT,
     type    CHAR(12)    UNIQUE NOT NULL,
     PRIMARY KEY (id)
);
#  ------
# |orders|
#  ------
# alis orders == payments.. what ever..
CREATE TABLE orders (
    id                     BIGINT      NOT NULL AUTO_INCREMENT,
    created_at             DATETIME(3) NOT NULL,
    user_id                BIGINT      NOT NULL,
    tax_rate_id            BIGINT NULL,   # reserved for the future
    shipment_cost_id       BIGINT NULL,   # reserved for the future
    payment_instruments_id BIGINT,
    FOREIGN KEY (payment_instruments_id) REFERENCES payment_instruments (id),
    PRIMARY KEY (id)
);
CREATE TABLE order_items (
    id          BIGINT  NOT NULL AUTO_INCREMENT,
    price_id    BIGINT  NOT NULL,
    order_id    BIGINT,
    quantity    TINYINT UNSIGNED DEFAULT 1,
    FOREIGN KEY (order_id) REFERENCES orders (id),
    PRIMARY KEY (id)
);
CREATE TABLE order_txs (
    id          BIGINT      NOT NULL AUTO_INCREMENT,
    state       CHAR(10)    NOT NULL,
    order_id    BIGINT      NULL,
    created_at  DATETIME(3) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders (id),
    PRIMARY KEY (id)
);