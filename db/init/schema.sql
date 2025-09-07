create table role (
  id int not null primary key,
  title varchar(20) not null,
  description text not null
);

create table users (
  id varchar(20) not null primary key, -- es varchar por que en el codigo genero un string aleatorio de 16 caracteres
  name varchar(55) not null,
  mat_summary varchar(55) not null,
  pat_summary varchar(55) not null,
  username varchar(55) not null,
  email varchar(255) not null unique,
  password text not null, -- este es tipo text por que ira hasheada
  phone varchar(20),
  image varchar(255),
  role_id int not null references role(id), -- 0 = admin, 1 = usuario
  created_at timestamp default current_timestamp,
  updated_at timestamp
);

create table address (
  id varchar(20) not null primary key,
  id_user varchar(20) not null references users(id) on delete cascade,
  street varchar(40) not null,
  city varchar(30) not null,
  state varchar(30) not null,
  cp varchar(10) not null,
  main boolean not null, --indica si es la direccion principal del cliente
  added_in timestamp default current_timestamp,
  update_at timestamp
);
create table categories (
  id varchar(20) not null primary key,
  name varchar(30) not null,
  description text not null
);

create table subcategories (
  id varchar(20) not null primary key,
  id_category varchar(20) not null references categories(id) on delete cascade,
  name varchar(30) not null,
  description text not null
);

create table products (
  id varchar(16) not null primary key,
  name varchar(30) not null,
  description text not null,
  stock int not null,
  price decimal(10, 2) not null, -- en pesos mexicanos, el tipo de dato decimal(10, 2), permite calculos matematicos en base de datos y mantiene un formato (00.00) ejemplo 1200 = 1,200.00
  category varchar(20)  references categories(id) on delete set null,
  subcategory varchar(20)  references subcategories(id) on delete set null,
  is_active boolean not null default true, -- true si pus si vea, no no ps no vea; permite desactivar un producto sin perder historial
  created_at timestamp default current_timestamp,
  updated_at timestamp
);

create table product_images (
  id varchar(20) not null primary key,
  id_product varchar(20) not null references products(id) on delete cascade,
  image_url varchar(255) not null
);



create table carts (
  id varchar(20) not null primary key,
  id_user varchar(20) not null references users(id) on delete cascade,
  created_at timestamp default current_timestamp
);

create table cart_items (
  id varchar(20) not null primary key,
  id_cart varchar(20) not null references carts(id) on delete cascade,
  id_product varchar(20) not null references products(id) on delete cascade,
  quantity int not null check (quantity > 0)
);

create table orders (
  id varchar(20) not null primary key,
  id_user varchar(20) not null references users(id) on delete cascade,
  status varchar(40) default 'peding' check (status in ('peding', 'paid', 'shipped', 'delivered', 'canceled')), -- pending, paid, shipped, delivered, canceled
  total decimal(10, 2) not null,
  created_at timestamp default current_timestamp
);

create table order_items (
  id varchar(20) not null primary key,
  id_order varchar(20) not null references orders(id) on delete cascade,
  id_product varchar(20) not null references products(id),
  quantity int not null check (quantity > 0),
  price decimal(10, 2) not null
);

create table payments (
  id varchar(20) not null primary key,
  id_order varchar(20) not null references orders(id) on delete cascade,
  method varchar(50) not null check (method in ('paypal', 'stripe')),
  status VARCHAR(50) not null default 'pending' check (status in ('peding','completed', 'failed')),-- pending, completed, failed
  paid_at timestamp
);