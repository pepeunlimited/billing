CREATE DATABASE IF NOT EXISTS billing CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE billing;

#  -------------
# |subscriptions|
#  -------------
CREATE TABLE plans (
    id            BIGINT NOT NULL AUTO_INCREMENT,
    title_i18n_id BIGINT NOT NULL,
    price_id      BIGINT UNIQUE NOT NULL,
    start_at      DATETIME(3) NOT NULL,
    end_at        DATETIME(3) NOT NULL,
    length        TINYINT UNSIGNED NOT NULL,
    unit          CHAR(7) NOT NULL, #days, weeks, months, years
    PRIMARY KEY (id)
);
CREATE TABLE subscriptions (
    id                    BIGINT      NOT NULL AUTO_INCREMENT,
    user_id               BIGINT      NOT NULL,
    plan_subscriptions    BIGINT,
    start_at              DATETIME(3) NOT NULL,
    end_at                DATETIME(3) NOT NULL,
    FOREIGN KEY (plan_subscriptions) REFERENCES plans (id),
    PRIMARY KEY (id)
);
#  ------
# |orders|
#  ------
# alias order = payment.. what ever..
CREATE TABLE orders (
    id                        BIGINT      NOT NULL AUTO_INCREMENT,
    created_at                DATETIME(3) NOT NULL,
    user_id                   BIGINT      NOT NULL,
    tax_rate_id               BIGINT      NULL,   # reserved for the future
    shipment_cost_id          BIGINT      NULL,   # reserved for the future
    PRIMARY KEY (id)
);
CREATE TABLE order_items (
    id           BIGINT  NOT NULL AUTO_INCREMENT,
    price_id     BIGINT  NOT NULL,
    orders_items BIGINT,
    quantity     TINYINT UNSIGNED DEFAULT 1,
    FOREIGN KEY (orders_items) REFERENCES orders (id),
    PRIMARY KEY (id)
);
CREATE TABLE order_txs (
    id          BIGINT       NOT NULL AUTO_INCREMENT,
    status      CHAR(8)      NOT NULL,
    orders_txs  BIGINT       NULL,
    created_at  DATETIME(3)  NOT NULL,
    FOREIGN KEY (orders_txs) REFERENCES orders (id),
    PRIMARY KEY (id)
);
#  -------------------
# |payment_instruments|
#  -------------------
CREATE TABLE payment_instruments (
     id           TINYINT     NOT NULL AUTO_INCREMENT,
     type         CHAR(12)    UNIQUE NOT NULL,
     type_i18n_id BIGINT      NULL,               # reserved for the future
     PRIMARY KEY (id)
);
#  --------
# |payments|
#  --------
CREATE TABLE payments (
     id                                                     BIGINT NOT NULL AUTO_INCREMENT,
     orders_payments                                        BIGINT UNIQUE,
     instrument_payments                                    TINYINT,
     FOREIGN KEY (orders_payments)                          REFERENCES orders (id),
     FOREIGN KEY (instrument_payments)                      REFERENCES payment_instruments (id),
     PRIMARY KEY (id)
);