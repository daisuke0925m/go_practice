CREATE TABLE tasks (
  id SERIAL NOT NULL,
  title varchar(255) DEFAULT NULL,
  note text DEFAULT NULL,
  completed integer DEFAULT 0,
  created_at TIMESTAMP DEFAULT NULL,
  updated_at TIMESTAMP DEFAULT NULL,
  PRIMARY KEY(id)
);
CREATE INDEX task_id on tasks (id);
