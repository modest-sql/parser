SELECT
    t.Columna, *
FROM SomeTable AS t
INNER JOIN AnotherTable a
    ON a.Id = t.Id;

SELECT
    t.Columna, *
FROM SomeTable AS t
INNER JOIN AnotherTable a
    ON a.Id = t.Id
WHERE a.Id > 5;