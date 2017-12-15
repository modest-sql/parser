UPDATE Rooms
SET Description = Description + '-ED'
WHERE Room LIKE '%2%'
;

SELECT
    RoomId,
    Description
FROM Rooms
WHERE 1 = 1
    AND Description LIKE '%ED';