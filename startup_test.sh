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
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestFindAllQr
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestAddOneQr
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrActive
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrDisactive
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrApply
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrNotapply
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrSchemaname
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrFlagIN
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrClientAddr
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrProxyAddr
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrProxyPort
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrDigest
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrMatchDigest
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrMatchPattern
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrNegateMatchPatternEnable
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrNegateMatchPatternDisable
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrFlagOut
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrReplacePattern
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrDestHostgroup
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrCacheTTL
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrReconnectEnable
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrReconnectDisable
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrTimeOut
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrRetries
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrDelay
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrMirrorFlagOut
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrMirrorHostgroup
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrErrorMsg
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrLogEnable
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneQrLogDisable
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestDeleteOneQr

#scheduler
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestFindAllSchedulers
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestAddOneSchedulers
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneSchedulersActiveEnable
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneSchedulersActiveDisable
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneSchedulersIntervalMs
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneSchedulersArg1
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneSchedulersArg2
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneSchedulersArg3
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneSchedulersArg4
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestUpdateOneSchedulersArg5
go test -v --args -addr 172.18.10.136 -port 13306 -user admin -pass admin -test.run TestDeleteOneSchedulers
