#!/bin/sh
# export MONGO_URI="mongodb://localhost:27016/shorten-url?rm.failover=1000ms:5x1&rm.monitorRefreshMS=100&rm.nbChannelsPerNode=1"
export DB_URL="mongodb+srv://root:root@cluster0.qfx1p.mongodb.net/short-url?retryWrites=true&w=majority"
export DB_Name="url_database"
export Url_Info_Collection_Name="url_info"
export URL_HOST="http://localhost:8080/"
export PORT="8080"