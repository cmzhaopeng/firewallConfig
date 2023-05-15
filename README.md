
# README

Read IP address list from database, write to firewall IP list group.

## .env example

DSN= "host=127.0.0.1 user=postgres password=yourdb_password dbname=postgres port=5432"
FWIP=xx.xx.xx.xx
FWUSER=admin
FWPASS=yourpassword
TIMEOUT= 5

The Variable TIMEOUT should not be set too short because it should wait for the ending of write the command to the firewall.


## Error information for plink

If you find the log file, permit.log have the connection reset by peer, you should check the ssh connect to firewall address through putty.exe. The cause may be the crypto algorithm is not be supported. When you use putty to connect the firewall successfully, the plink can connect to it too. 


