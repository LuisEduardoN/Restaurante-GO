create database restaurante_db;

\connect restaurante_db;


CREATE TABLE IF NOT exists PRATO(
    codigo SERIAL NOT NULL constraint PRATO_PK PRIMARY KEY,
    nome varchar(50) not null,
    ingredientes varchar(250) not null,
    tipo varchar(100) not null,
    preco decimal not null
);


