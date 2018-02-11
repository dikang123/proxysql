#########################################################################
# File Name: startup_test.sh
# Author: Lei Tian
# mail: taylor840326@gmail.com
# Created Time: 2018-02-06 17:20
#########################################################################
#!/bin/bash

# Users
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestFindAllUsers
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestAddOneUser
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneUserFastForwardEnable
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneUserFastForwardDisable
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneUserMaxConnections
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneUserActive
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneUserDisactive
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneUserUseSslEnable
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneUserUseSslDisable
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneUserShcemaLockedEnable
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneUserShcemaLockedDisable
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneUserTransactionPersistentEnable
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneUserTransactionPersistentDisable
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestDeleteOneUser

# Servers
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestFindAllServers
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestAddOneServer
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneServerStatusToOnline
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneServerStatusToOfflineSoft
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneServerStatusToOfflineHard
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneServerWeight
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneServerCompressionEnable
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneServerCompressionDisable
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneServerMaxConnection
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneServerMaxReplication
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneServerUseSslEnable
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneServerUseSslDisable
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneServerMaxLatencyMs
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneServerComment
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestDeleteOneServer

#Queryrules