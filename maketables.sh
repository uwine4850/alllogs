#!/bin/bash

sudo docker exec -i mysql mysql --defaults-extra-file=/schema/mysql.cnf -e "SOURCE /schema/profile.sql"
sudo docker exec -i mysql mysql --defaults-extra-file=/schema/mysql.cnf -e "SOURCE /schema/project.sql"