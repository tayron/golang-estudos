# Migrate using SQL-Migrate

https://github.com/rubenv/sql-migrate

Execute ```env MYSQL_HOST="servidor_mysql_local" MYSQL_PORT=3306 MYSQL_USER="root" MYSQL_PASSWORD="yakTLS&70c52" DATABASE_NAME="golang_estudos" sql-migrate up -env="development" -dryrun``` to view SQL to execute and execute ```env MYSQL_HOST="servidor_mysql_local" MYSQL_PORT=3306 MYSQL_USER="root" MYSQL_PASSWORD="yakTLS&70c52" DATABASE_NAME="golang_estudos" sql-migrate up -env="development"``` to execute migration at database.
