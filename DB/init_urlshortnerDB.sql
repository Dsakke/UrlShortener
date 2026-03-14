CREATE DATABASE UrlShortener;
CREATE TABLE urls 
(
    id SERIAL PRIMARY KEY,
    urlKey VARCHAR(7), 
    url VARCHAR(255)
);

CREATE INDEX UrlKeyIndex ON urls (urlKey);