ALTER TABLE entries
ADD COLUMN created_at timestamp(0) with time zone NOT NULL DEFAULT NOW();
