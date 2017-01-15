CREATE TABLE expenses (
  id SERIAL PRIMARY KEY,
  price_cents INTEGER NOT NULL DEFAULT 0 CHECK (price_cents >= 0),
  date TIMESTAMP WITH TIME ZONE NOT NULL,
  tag VARCHAR(50) NOT NULL,
  notes TEXT
);

CREATE INDEX tag_index ON expenses (tag);