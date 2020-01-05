DROP TABLE IF EXISTS inverted_indexes;
CREATE TABLE pages (
  id int NOT NULL AUTO_INCREMENT,
  word_id int,
  page_id int,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  PRIMARY KEY (id)
);
