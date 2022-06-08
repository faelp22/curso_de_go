
CREATE TABLE IF NOT EXISTS tb_produto
(
    id serial NOT NULL,
    name character varying(120) NOT NULL,
    code character varying(20) NOT NULL,
    price double precision NOT NULL,
    created_at time without time zone NOT NULL DEFAULT NOW(),
    updated_at time without time zone NOT NULL DEFAULT NOW(),
    CONSTRAINT tb_produto_pkey PRIMARY KEY (id)
);

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON tb_produto
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();

INSERT INTO tb_produto (name, code, price) VALUES 
    ('Notebook 13 XPS', 'DELL', 14500),
    ('Notebook 15', 'DELL', 2500),
    ('Notebook 14', 'DELL', 2500),
    ('Notebook 15', 'Lenovo', 3500),
    ('Notebook 13', 'Lenovo', 4568),
    ('Tablet', 'SAM', 8450),
    ('Macbook 13 Pro M1', 'Apple', 18500.00),
    ('TV 55', 'SONY', 4500.00),
    ('TV 45', 'SONY', 3500.00),
    ('TV 32', 'SAM', 2500.00),
    ('TV 60', 'LG', 6500.00),
    ('TV 50', 'LG', 4800.00);
