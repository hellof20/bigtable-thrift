FROM openjdk:8
ADD hbase-1.4.3 /hbase
WORKDIR /hbase
CMD ["bin/hbase","thrift2","start"]
