create table if not exists `categories` (
  id int primary key AUTO_INCREMENT,
  name varchar(255) not null
);


create table if not exists `courses` (
  id int primary key AUTO_INCREMENT,
  name varchar(255) not null,
  category_id integer not null,
  foreign key (category_id) references categories(id)
);