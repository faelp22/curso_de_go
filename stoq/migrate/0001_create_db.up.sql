
CREATE TABLE tb_produto (
    id         INTEGER         PRIMARY KEY AUTOINCREMENT
                               UNIQUE,
    name       VARCHAR (120)   NOT NULL,
    code       VARCHAR (20),
    price      DECIMAL (10, 2) NOT NULL,
    created_at DATETIME        NOT NULL
                               DEFAULT (datetime('now', 'localtime')),
    updated_at DATETIME        NOT NULL
                               DEFAULT (datetime('now', 'localtime')) 
);

INSERT INTO tb_produto (name, code, price) VALUES ("Notebook 13", "XPS", 14500.00);
INSERT INTO tb_produto (name, code, price) VALUES ("Notebook 15", "DELL", 2500.00);
INSERT INTO tb_produto (name, code, price) VALUES ("Notebook 15", "DELL", 2500.00);
INSERT INTO tb_produto (name, code, price) VALUES ("Notebook 15", "Lenovo", 3500.00);
INSERT INTO tb_produto (name, code, price) VALUES ("Notebook 15", "Lenovo", 4500.00);
insert into tb_produto (name, code, price) values ("Notebook 2", "hp", 3500.00);

-- update tb_produto SET price = 4500 where id == 1;
