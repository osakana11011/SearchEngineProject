DROP TABLE IF EXISTS pages;
CREATE TABLE pages (
  id int NOT NULL AUTO_INCREMENT,
  title VARCHAR(255) NOT NULL,
  url VARCHAR(2083) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  PRIMARY KEY (id)
);
