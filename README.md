# bigtable-thrift

ubuntu 22.04

## deploy hbase thrift server
```
sudo apt update
sudo apt install unzip openjdk-8-jdk-headless -y
export JAVA_HOME=$(update-alternatives --list java | tail -1 | sed -E 's/\/bin\/java//')
wget https://archive.apache.org/dist/hbase/1.4.3/hbase-1.4.3-bin.tar.gz
tar -zxvf hbase-1.4.3-bin.tar.gz
cd hbase-1.4.3
rm -rf lib/guava-12.0.1.jar
rm -rf lib/protobuf-java-2.5.0.jar
mkdir -p lib/bigtable
cd lib/bigtable
gsutil cp gs://pwm-lowa/bigtable-hbase-1.x-1.14.1.zip .
unzip bigtable-hbase-1.x-1.14.1.zip
```

## modify conf/hbase-site.xml

```
<configuration>
  <property><name>google.bigtable.project.id</name><value>speedy-victory-336109</value></property>
  <property><name>google.bigtable.instance.id</name><value>bigtable1</value></property>
  <property>
    <name>hbase.client.connection.impl</name>
    <value>com.google.cloud.bigtable.hbase1_x.BigtableConnection</value>
  </property>
</configuration>
```

## modify conf/hbase-env.sh
```
export HBASE_CLASSPATH="lib/bigtable/*"
```

## install thrift
```
sudo apt-get install automake bison flex g++ git libboost-all-dev libevent-dev libssl-dev libtool make pkg-config -y
wget http://archive.apache.org/dist/thrift/0.13.0/thrift-0.13.0.tar.gz
tar -zxvf thrift-0.13.0.tar.gz
cd thrift-0.13.0
./configure
make
sudo make install
```

## test with go applicaion
```
sudo apt install golang -y
cd go-thrif2
thrift -gen go hbase.thrift
sudo cp -r gen-go/hbase /usr/lib/go-1.18/src/
go run main.go
```
