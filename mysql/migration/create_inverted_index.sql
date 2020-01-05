DROP TABLE IF EXISTS inverted_index;
CREATE TABLE inverted_index (
  id int NOT NULL AUTO_INCREMENT,
  word_id int,
  page_id int,
  counts int,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  PRIMARY KEY (id)
);
