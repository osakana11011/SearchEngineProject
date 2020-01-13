DROP TABLE IF EXISTS inverted_list;
CREATE TABLE inverted_list (
  id INT NOT NULL AUTO_INCREMENT,
  word_id INT NOT NULL ,
  document_id INT NOT NULL,
  offset_list TEXT,
  tf FLOAT,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  PRIMARY KEY (id)
);
ALTER TABLE inverted_list ADD INDEX index_name(word_id);
