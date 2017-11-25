CREATE TABLE peliculas (
    id_pelicula INTEGER PRIMARY KEY,
    titulo CHAR(100)
);

CREATE TABLE peliculas2 (
    id_pelicula INTEGER FOREIGN KEY,
    titulo CHAR(100) NOT NULL,
    fecha DATETIME,
    isTrue BOOLEAN,
    salary FLOAT DEFAULT 1200.0
);