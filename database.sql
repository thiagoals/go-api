create database questionarioHospital;

use questionarioHospital;

create table hospital(
    id int primary key not null auto_increment,
    nome varchar(150) not null,
    endereco varchar(255) not null,
    cidade varchar(150) not null,
    estado varchar(150) not null
)engine=InnoDB;

create table topico(
    id int primary key not null auto_increment,
    nome varchar(255) not null
)engine=InnoDB;

create table questao(
    id int primary key not null auto_increment,
    texto varchar(255) not null
)engine=InnoDB;