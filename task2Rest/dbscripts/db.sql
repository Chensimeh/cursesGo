CREATE TABLE web_origins (
    client_id character varying(36) NOT NULL,
    value character varying(255)
);
INSERT INTO web_origins(client_id,value) values(1,1);
SELECT * FROM web_origins;
