FROM postgres:16

COPY ./db/*.sql /docker-entrypoint-initdb.d/
