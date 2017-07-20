FROM centos:7

COPY proxysql-master  /usr/local/bin/proxysql-master


CMD /usr/local/bin/proxysql-master  -s 'admin/admin@127.0.0.1:6032/main' -p 3333 -l /tmp/pm.log

