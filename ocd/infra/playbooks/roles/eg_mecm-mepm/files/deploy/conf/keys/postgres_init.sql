CREATE USER lcmcontroller WITH PASSWORD 'PASSWORD_VALUE' CREATEDB;
CREATE DATABASE lcmcontrollerdb
    WITH 
    OWNER = lcmcontroller
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

CREATE USER k8splugin WITH PASSWORD 'PASSWORD_VALUE' CREATEDB;
CREATE DATABASE k8splugindb
    WITH 
    OWNER = k8splugin
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

