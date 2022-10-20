#!/bin/bash

function createDatabase() {
    mysql --defaults-extra-file=/etc/mysql/conf.d/my.cnf <<EOS
DROP DATABASE IF EXISTS \`database\`;
CREATE DATABASE IF NOT EXISTS \`database\` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
EOS
}

function executeDDL() {
    mysql --defaults-extra-file=/etc/mysql/conf.d/my.cnf database < /var/ddl/master.ddl
    mysql --defaults-extra-file=/etc/mysql/conf.d/my.cnf database < /var/ddl/transaction.ddl
}

function executeDML() {
    mysql --defaults-extra-file=/etc/mysql/conf.d/my.cnf database < /var/dml/master.dml
}

createDatabase
executeDDL
executeDML
