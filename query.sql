--Query yang saya gunakan untuk menghasilkan output yang ditentukan--

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    parent_id INTEGER REFERENCES users(id)
);

INSERT INTO users (name, parent_id) VALUES
('Zaki', 2),
('Ilham', NULL),
('Irwan', 2),
('Arka', 3);

SELECT
    u.id,
    u.name,
    p.name AS parent_name
FROM
    users u
LEFT JOIN
    users p ON u.parent_id = p.id
ORDER BY
    u.id;
