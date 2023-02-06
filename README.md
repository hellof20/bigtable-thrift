# bigtable-thrift

## Prerequirement
- Ubuntu 22.04 VM
- GCP VM with service account， have permission to access Bigtable
- Bigtable instance

## Deploy Hbase compatible with Bigtable
### Install Hbase
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

### Modify conf/hbase-site.xml

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

### Modify conf/hbase-env.sh
```
export HBASE_CLASSPATH="lib/bigtable/*"
```

## Test with native hbase shell
<img width="1710" alt="image" src="https://user-images.githubusercontent.com/8756642/216880206-4bc1fd5a-34ae-4594-a711-2aa8aa1f2dca.png">

## Test with golang
### Start Hbase Thrift server
```
bin/hbase thrift2 start
```
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

