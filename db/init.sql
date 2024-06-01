-- init.sql

-- Run migration scripts
SOURCE db/migrations/001_create_tables.sql;
SOURCE db/migrations/002_add_indexes.sql;

-- Run seed scripts
SOURCE db/seeds/seed_data.sql;