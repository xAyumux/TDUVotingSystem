CREATE TABLE IF NOT EXISTS polls (id INT PRIMARY KEY, title TEXT, description TEXT);
CREATE TABLE IF NOT EXISTS votes(
    userid CHAR(32) PRIMARY KEY,
    id INT,
    result INT
);