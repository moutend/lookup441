PRAGMA ENCODING="UTF-8";

CREATE TABLE IF NOT EXISTS period (
  sample_rate INTEGER NOT NULL,
  buffer_size INTEGER NOT NULL,
  frequency  INTEGER NOT NULL,
  period REAL NOT NULL,

  PRIMARY KEY(sample_rate, buffer_size, frequency)
);
