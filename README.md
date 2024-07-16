# Bigtable-Thrift

## Prerequirement
注意: HBase的版本和Thrift的版本
- Ubuntu 22.04 VM
- GCP VM with service account， have permission to access Bigtable
- Bigtable instance

## Deploy HBase compatible with Bigtable
### Install HBase
```
sudo apt update
sudo apt install unzip openjdk-8-jdk-headless maven -y
wget https://archive.apache.org/dist/hbase/2.2.1/hbase-2.2.1-bin.tar.gz
tar -zxvf hbase-2.2.1-bin.tar.gz
```

### Bigtable HBase dependencies
```
cd hbase-2.2.1
mkdir -p lib/bigtable
cd lib/bigtable
```
将pom.xml放到lib/bigtable下
```
mvn dependency:go-offline dependency:copy-dependencies -DoutputDirectory=.
```

### Modify conf/hbase-site.xml

```
<configuration>
  <property>
    <name>google.bigtable.project.id</name>
    <value>speedy-victory-336109</value>
  </property>
  <property>
    <name>google.bigtable.instance.id</name>
    <value>bigtable1</value>
  </property>
  <property>
    <name>hbase.client.connection.impl</name>
    <value>com.google.cloud.bigtable.hbase2_x.BigtableConnection</value>
  </property>
</configuration>
```

### Modify conf/hbase-env.sh
```
export HBASE_CLASSPATH="lib/bigtable/*"
export JAVA_HOME=$(update-alternatives --list java | tail -1 | sed -E 's/\/bin\/java//')
```

## Test with native hbase shell
```
bin/hbase shell
```
![image](https://github.com/user-attachments/assets/947ff6c8-ef46-440b-afa0-e41ad8da7d3c)

![image](https://github.com/user-attachments/assets/b315b682-9376-419b-a148-7e68b483d9c9)


## Start Hbase Thrift server
```
bin/hbase thrift2 start
```

## Test with golang
### Install Thrift
```
sudo apt-get install automake bison flex g++ git libboost-all-dev libevent-dev libssl-dev libtool make pkg-config -y
wget http://archive.apache.org/dist/thrift/0.13.0/thrift-0.13.0.tar.gz
tar -zxvf thrift-0.13.0.tar.gz
cd thrift-0.13.0
./configure
make
sudo make install
```

### prepare golang applicaion
- Install golang
```
sudo apt install golang -y
cd go-thrif2
thrift -gen go hbase.thrift
sudo cp -r gen-go/hbase /usr/lib/go-1.18/src/
```
- Modify main.go
Change 'HOST' to your VM IP address

- Test it
```
go run main.go
```
<img width="836" alt="image" src="https://user-images.githubusercontent.com/8756642/216880488-3e175cca-d9b3-4ecf-9532-12b2f7a96d92.png">

