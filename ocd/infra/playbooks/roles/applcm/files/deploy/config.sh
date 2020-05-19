#!/bin/bash
# Command to give privileges to mysql root user
# The Grant is necessary for Applcm-Broker to create applcmdb
mysql -u root -ppassword << EOF
SELECT host,user,Grant_priv,Super_priv FROM mysql.user;
UPDATE mysql.user SET Grant_priv='Y', Super_priv='Y' WHERE User='root';
FLUSH PRIVILEGES;
GRANT ALL ON *.* TO 'root'@'%'; 
EOF
