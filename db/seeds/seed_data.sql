-- Add some dummy records (which we'll use in the next couple of chapters).
INSERT INTO snippets (title, content, created, expires) VALUES (
'An old silent pond',
'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō', UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);
INSERT INTO snippets (title, content, created, expires) VALUES (
'Over the wintry forest',
'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki', UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);
INSERT INTO snippets (title, content, created, expires) VALUES (
'First autumn morning',
'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo', UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
);

-- SELECT id, title, expires  FROM snippets;
-- +----+------------------------+---------------------+
-- | id | title                  | expires             |
-- +----+------------------------+---------------------+
-- |  1 | An old silent pond     | 2025-05-31 19:13:07 |
-- |  2 | Over the wintry forest | 2025-05-31 19:13:08 |
-- |  3 | First autumn morning   | 2024-06-07 19:13:08 |
-- +----+------------------------+---------------------+
-- 3 rows in set (0.00 sec)
