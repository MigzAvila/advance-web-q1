ALTER TABLE entries
ADD COLUMN entries timestamp(0) with time zone NOT NULL DEFAULT NOW();
