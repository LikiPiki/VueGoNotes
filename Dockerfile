FROM postgres:latest
EXPOSE 5432

COPY ./server/tables.sql /docker-entrypoint-initdb.d/
VOLUME  /var/lib/postgresql/data


