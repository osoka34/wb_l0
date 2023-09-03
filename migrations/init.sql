-- Создание таблицы "order"
CREATE TABLE public.orders (
                          order_uid VARCHAR(255) PRIMARY KEY,
                          track_number VARCHAR(255),
                          entry VARCHAR(255),
                          locale VARCHAR(255),
                          internal_signature VARCHAR(255),
                          customer_id VARCHAR(255),
                          delivery_service VARCHAR(255),
                          shardkey VARCHAR(255),
                          sm_id INT,
                          date_created TIMESTAMP,
                          oof_shard VARCHAR(255)
);

-- Создание таблицы "delivery"
CREATE TABLE public.delivery (
                          id SERIAL PRIMARY KEY,
                          order_uid VARCHAR(255) REFERENCES orders(order_uid),
                          name VARCHAR(255),
                          phone VARCHAR(255),
                          zip VARCHAR(255),
                          city VARCHAR(255),
                          address VARCHAR(255),
                          region VARCHAR(255),
                          email VARCHAR(255)
);

-- Создание таблицы "payment"
CREATE TABLE public.payment (
                         id SERIAL PRIMARY KEY,
                         order_uid VARCHAR(255) REFERENCES orders(order_uid),
                         transaction VARCHAR(255),
                         request_id VARCHAR(255),
                         currency VARCHAR(255),
                         provider VARCHAR(255),
                         amount INT,
                         payment_dt INT,
                         bank VARCHAR(255),
                         delivery_cost INT,
                         goods_total INT,
                         custom_fee INT
);

-- Создание таблицы "item"
CREATE TABLE public.item (
                      id SERIAL PRIMARY KEY,
                      order_uid VARCHAR(255) REFERENCES orders(order_uid),
                      chrt_id INT,
                      track_number VARCHAR(255),
                      price INT,
                      rid VARCHAR(255),
                      name VARCHAR(255),
                      sale INT,
                      size VARCHAR(255),
                      total_price INT,
                      nm_id INT,
                      brand VARCHAR(255),
                      status INT,
                      CHECK (price >= 0 AND total_price >= 0)
);


INSERT INTO public.orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
VALUES
    ('order1', 'track1', 'entry1', 'locale1', 'internal1', 'customer1', 'delivery1', 'shard1', 1, '2023-08-30', 'oof1'),
    ('order2', 'track2', 'entry2', 'locale2', 'internal2', 'customer2', 'delivery2', 'shard2', 2, '2023-08-31', 'oof2');

-- Генерация данных для таблицы "delivery"
INSERT INTO public.delivery (order_uid, name, phone, zip, city, address, region, email)
VALUES
    ('order1', 'name1', 'phone1', 'zip1', 'city1', 'address1', 'region1', 'email1'),
    ('order2', 'name2', 'phone2', 'zip2', 'city2', 'address2', 'region2', 'email2');

-- Генерация данных для таблицы "payment"
INSERT INTO public.payment (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
VALUES
    ('order1', 'trans1', 'request1', 'currency1', 'provider1', 100, 1630387200, 'bank1', 10, 90, 5),
    ('order2', 'trans2', 'request2', 'currency2', 'provider2', 200, 1630473600, 'bank2', 20, 180, 10);

-- Генерация данных для таблицы "item"
INSERT INTO public.item (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
VALUES
    ('order1', 1, 'track1', 50, 'rid1', 'item1', 10, 'size1', 50, 1, 'brand1', 1),
    ('order2', 2, 'track2', 100, 'rid2', 'item2', 20, 'size2', 100, 2, 'brand2', 2);

