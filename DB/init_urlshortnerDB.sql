CREATE DATABASE UrlShortener;
CREATE TABLE urls 
(
    id VARCHAR(10) PRIMARY KEY,
    urlKey VARCHAR(7), 
    url VARCHAR(255)
);

CREATE INDEX UrlKeyIndex ON urls (urlKey);