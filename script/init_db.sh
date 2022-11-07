#!/bin/bash

MYSQL_PWD=password
export MYSQL_PWD

function createDatabase() {
    mysql -u root <<EOS
DROP DATABASE IF EXISTS \`database\`;
CREATE DATABASE IF NOT EXISTS \`database\` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
EOS
}

function executeDDL() {
    mysql -u root database < /var/ddl/master.ddl
    mysql -u root database < /var/ddl/transaction.ddl
}

function executeDML() {
    mysql -u root database < /var/dml/master.dml
}

createDatabase
executeDDL
executeDML
