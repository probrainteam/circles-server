load data INFILE '/docker-entrypoint-initdb.d/manager.csv' 
INTO TABLE manager 
FIELDS TERMINATED BY ',' 
ENCLOSED BY '"' 
LINES TERMINATED BY '\n' 
IGNORE 1 ROWS
;
load data INFILE '/docker-entrypoint-initdb.d/probrain.csv' 
INTO TABLE probrain
FIELDS TERMINATED BY ',' 
ENCLOSED BY '"' 
LINES TERMINATED BY '\n' 
IGNORE 1 ROWS
;