#########################################################################
# File Name: startup_test.sh
# Author: Lei Tian
# mail: taylor840326@gmail.com
# Created Time: 2018-02-06 17:20
#########################################################################
#!/bin/bash

# Users
go test --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestFindAllUsers
go test --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestAddOneUser
go test --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneUserFastForwardEnable
go test --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneUserFastForwardDisable
go test --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneUserMaxConnections
go test --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneUserActive
go test --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneUserDisactive
go test --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneUserUseSslEnable
go test --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneUserUseSslDisable
go test --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneUserShcemaLockedEnable
go test --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneUserShcemaLockedDisable
go test --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneUserTransactionPersistentEnable
go test --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneUserTransactionPersistentDisable
go test --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestDeleteOneUser

# Servers