FROM centos:7

COPY proxysql-master  /usr/local/bin/proxysql-master


CMD /usr/local/bin/proxysql-master  -s 'admin/admin@172.18.7.204:6032/main' -p 3333 -l /tmp/pm.log


