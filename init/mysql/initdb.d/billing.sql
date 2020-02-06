CREATE DATABASE IF NOT EXISTS billing CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE billing;

# CREATE TABLE product (
#      id BIGINT NOT NULL AUTO_INCREMENT,
#      name_i18n BIGINT NOT NULL,
#      short_i18n BIGINT NOT NULL,
#      long_i18n BIGINT NOT NULL,
#      sku VARCHAR(30) UNIQUE NOT NULL,
#      created_at INT UNSIGNED NOT NULL,
#      type VARCHAR(200) NOT NULL,
#      FOREIGN KEY (name_i18n) REFERENCES i18n (id),
#      FOREIGN KEY (short_i18n) REFERENCES i18n (id),
#      FOREIGN KEY (long_i18n) REFERENCES i18n (id),
#      PRIMARY KEY (id)
# );
# CREATE TABLE product_price (
#    id BIGINT NOT NULL AUTO_INCREMENT,
#    currency CHAR(3) NOT NULL,
#    start_at INT UNSIGNED NOT NULL,
#    end_at INT UNSIGNED NOT NULL,
#    price DECIMAL(13,4) DEFAULT 0.0000,
#    discount DECIMAL(13,4) DEFAULT 0.0000,
#    product_id BIGINT NOT NULL,
#    FOREIGN KEY (product_id) REFERENCES product (id),
#    FOREIGN KEY (currency) REFERENCES currency (code),
#    PRIMARY KEY (id)
# );

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
    user_id               BIGINT NOT NULL,
    plan_subscriptions    BIGINT,
    start_at              DATETIME(3) NOT NULL,
    end_at                DATETIME(3) NOT NULL,
    status                CHAR(8)     NOT NULL, # paid, canceled, failed, refunded, pending
    FOREIGN KEY (plan_subscriptions) REFERENCES plans (id),
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
    id          BIGINT  NOT NULL AUTO_INCREMENT,
    price_id    BIGINT  NOT NULL,
    orders_items BIGINT,
    quantity    TINYINT UNSIGNED DEFAULT 1,
    FOREIGN KEY (orders_items) REFERENCES orders (id),
    PRIMARY KEY (id)
);
CREATE TABLE order_txs (
    id          BIGINT      NOT NULL AUTO_INCREMENT,
    status      CHAR(10)    NOT NULL,
    orders_txs  BIGINT      NULL,
    created_at  DATETIME(3) NOT NULL,
    FOREIGN KEY (orders_txs) REFERENCES orders (id),
    PRIMARY KEY (id)
);
#  -------
# |payment|
#  -------
CREATE TABLE payments (
     id                                                   BIGINT  NOT NULL AUTO_INCREMENT,
     orders_orders                                        BIGINT UNIQUE,
     method_payment_instruments                           BIGINT,
     FOREIGN KEY (orders_orders)                          REFERENCES orders (id),
     FOREIGN KEY (method_payment_instruments)             REFERENCES payment_instruments (id),
     PRIMARY KEY (id)
);