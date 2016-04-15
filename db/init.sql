CREATE TABLE IF NOT EXISTS browsing (
  id int(11) NOT NULL AUTO_INCREMENT,
  src_ip varchar(255) NOT NULL,
  dst_ip varchar(255) NOT NULL,
  src_port int(11) NOT NULL,
  dst_port int(11) NOT NULL,
  timestamp datetime NOT NULL,
  title text,
  url text,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  browsing_time float,
  download int(11),
  PRIMARY KEY (id),
  KEY index_browsing_on_created_at (created_at),
  KEY index_browsing_on_src_ip (src_ip)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS meta (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  value text,
  PRIMARY KEY (id),
  UNIQUE KEY unique_meta_on_name (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS word (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  count int(11) DEFAULT 0,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  KEY index_word_on_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- Created by DOOR to store dpi result
--
-- CREATE TABLE IF NOT EXISTS "dpi" (
--   id int(20) NOT NULL AUTO_INCREMENT,
--   src_ip varchar(255) NOT NULL,
--   dst_ip varchar(255) NOT NULL,
--   src_mac_addr varchar(255) NOT NULL,
--   dst_mac_addr varchar(255) NOT NULL,
--   src_port int(11) NOT NULL,
--   dst_port int(11) NOT NULL,
--   stream_id int(20) NOT NULL,
--   rule_id int(11) NOT NULL,
--   rule varchar(255) NOT NULL,
--   timestamp datetime NOT NULL
--   data text,
--   PRIMARY KEY (id)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
