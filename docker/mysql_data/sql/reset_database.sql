
DROP DATABASE IF EXISTS veterinariadb;
CREATE DATABASE veterinariadb
CHARACTER
SET utf8mb4
COLLATE utf8mb4_general_ci;
use veterinariadb;
-- public.usuarios definition

-- Drop table

-- DROP TABLE public.usuarios;

CREATE TABLE usuarios
(
  id int NOT NULL
  AUTO_INCREMENT,
  email varchar
  (255) NOT NULL,
  password varchar
  (255) NOT NULL,
  role varchar
  (255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  enabled BOOLEAN NOT NULL,
  CONSTRAINT usuarios_pk PRIMARY KEY
  (id)
);
  -- clientes definition

  -- Drop table

  -- DROP TABLE clientes;

  CREATE TABLE clientes
  (
    id int NOT NULL
    AUTO_INCREMENT,
  nombre varchar
    (255) NOT NULL,
  direccion varchar
    (255) NOT NULL,
  telefono varchar
    (255) NULL,
  usuario_id INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL,
  enabled BOOLEAN NULL,
  CONSTRAINT clientes_pk PRIMARY KEY
    (id)
);


    -- clientes foreign keys

    ALTER TABLE clientes ADD CONSTRAINT clientes_fk FOREIGN KEY (usuario_id) REFERENCES usuarios(id) ON DELETE CASCADE;
    -- vendedores definition

    -- Drop table

    -- DROP TABLE vendedores;

    CREATE TABLE vendedores
    (
      id int NOT NULL
      AUTO_INCREMENT,
  nombre varchar
      (255) NOT NULL,
  direccion varchar
      (255) NOT NULL,
  telefono varchar
      (255) NULL,
  usuario_id INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL,
  enabled BOOLEAN,
  CONSTRAINT vendedores_pk PRIMARY KEY
      (id)
);


      -- vendedores foreign keys

      ALTER TABLE vendedores ADD CONSTRAINT vendedores_fk FOREIGN KEY (usuario_id) REFERENCES usuarios(id) ON DELETE CASCADE;
      -- mascotas definition

      -- Drop table

      -- DROP TABLE mascotas;

      CREATE TABLE mascotas
      (
        id int NOT NULL
        AUTO_INCREMENT,
  cliente_id INTEGER NOT NULL,
  nombre varchar
        (255) NOT NULL,
  tipo varchar
        (255) NOT NULL,
  peso FLOAT NOT NULL,
  fecha_nacimiento TIMESTAMP NOT NULL,
  castrado BOOLEAN NOT NULL,
  created_at TIMESTAMP NOT NULL,
  enabled BOOLEAN,
  CONSTRAINT mascotas_pk PRIMARY KEY
        (id)
);


        -- mascotas foreign keys

        ALTER TABLE mascotas ADD CONSTRAINT mascotas_fk FOREIGN KEY (cliente_id) REFERENCES clientes(id) ON DELETE CASCADE;
        -- combos definition

        -- Drop table

        -- DROP TABLE combos;

        CREATE TABLE combos
        (
          id int NOT NULL
          AUTO_INCREMENT,
  peso FLOAT NOT NULL,
  complementos INTEGER NOT NULL,
  mascota_id INTEGER NOT NULL,
  vendedor_id INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL,
  enabled BOOLEAN,
  CONSTRAINT combos_pk PRIMARY KEY(id, mascota_id, vendedor_id)
);


          -- combos foreign keys

          ALTER TABLE combos ADD CONSTRAINT combos_fk FOREIGN KEY (vendedor_id) REFERENCES vendedores(id) ON DELETE CASCADE;
          ALTER TABLE combos ADD CONSTRAINT combos_fk_1 FOREIGN KEY (mascota_id) REFERENCES mascotas(id) ON DELETE CASCADE;

-- CLIENTES --
INSERT INTO usuarios (email,password,`role`,created_at,enabled) values ("cliente_federico@gmail.com","$2a$12$IM7lA8GM8tTd3yAKZvCjWOwXG4MW8p1jsnrIpPGMrocaOiMgZ5STi","CLIENTE",NOW(),true);
INSERT INTO clientes (nombre,direccion,telefono,usuario_id,created_at,enabled) values ("federico","chacabuco 425","454237",1,NOW(),true);
INSERT INTO mascotas (cliente_id,nombre,tipo,peso,fecha_nacimiento,castrado,created_at,enabled) values (1,"chicho","PERRO",15.50,'2020-01-01 10:10:10',true,NOW(),true);
INSERT INTO mascotas (cliente_id,nombre,tipo,peso,fecha_nacimiento,castrado,created_at,enabled) values (1,"chicho","PERRO",15.50,'2020-01-01 10:10:10',true,NOW(),true);
INSERT INTO mascotas (cliente_id,nombre,tipo,peso,fecha_nacimiento,castrado,created_at,enabled) values (1,"camila","GATO",13.78,'2020-01-01 10:10:10',true,NOW(),true);




-- VENDEDORES --
INSERT INTO usuarios (email,password,`role`,created_at,enabled) values ("vendedor_jaun@gmail.com","$2a$12$IM7lA8GM8tTd3yAKZvCjWOwXG4MW8p1jsnrIpPGMrocaOiMgZ5STi","VENDEDOR",NOW(),true);
INSERT INTO vendedores (nombre,direccion,telefono,usuario_id,created_at,enabled) values ("juan","chacabuco 298","12345231",2,NOW(),true);
INSERT INTO combos (peso,complementos,mascota_id,vendedor_id,created_at,enabled) values (15.50,3,1,1,NOW(),true);
INSERT INTO combos (peso,complementos,mascota_id,vendedor_id,created_at,enabled) values (15.50,0,1,1,NOW(),true);

INSERT INTO usuarios (email,password,`role`,created_at,enabled) values ("vendedor_carlos@gmail.com","$2a$12$IM7lA8GM8tTd3yAKZvCjWOwXG4MW8p1jsnrIpPGMrocaOiMgZ5STi","VENDEDOR",NOW(),true);
INSERT INTO vendedores (nombre,direccion,telefono,usuario_id,created_at,enabled) values ("carlos","chacabuco 212","12345231",3,NOW(),true);
INSERT INTO combos (peso,complementos,mascota_id,vendedor_id,created_at,enabled) values (15.50,0,2,2,NOW(),true);
INSERT INTO combos (peso,complementos,mascota_id,vendedor_id,created_at,enabled) values (15.50,0,2,2,NOW(),true);
INSERT INTO combos (peso,complementos,mascota_id,vendedor_id,created_at,enabled) values (15.50,0,2,2,NOW(),true);
INSERT INTO combos (peso,complementos,mascota_id,vendedor_id,created_at,enabled) values (15.50,0,1,2,NOW(),true);
INSERT INTO combos (peso,complementos,mascota_id,vendedor_id,created_at,enabled) values (15.50,0,1,2,NOW(),true);
