## Setup

Create Directory

```
$ mkdir ${GOPATH}/src/github.com/sgykfjsm/cockroach-sample-by-go
$ cd $_
$ mkdir bin // the directory for saving cockroach db binary
$ cd $_
```

Download Binary

```
$ wget https://binaries.cockroachdb.com/cockroach-latest.darwin-10.9-amd64.tgz
$ tar zxf ./cockroach-latest.darwin-10.9-amd64.tgz
```

Create Symlink

```
$ ln -sv cockroach-latest.darwin-10.9-amd64 cockroach
'cockroach' -> 'cockroach-latest.darwin-10.9-amd64'
```

Show Version to Confirm

```
$ ./bin/cockroach/cockroach version
Build Tag:    beta-20170126
Build Time:   2017/01/26 16:04:08
Distribution: CCL
Platform:     darwin amd64
Go Version:   go1.7.4
C Compiler:   4.2.1 Compatible Clang 3.8.0 (tags/RELEASE_380/final)
```

## Play cockroachdb with sql command

https://www.cockroachlabs.com/docs/start-a-local-cluster.html

Start your first node

```
# Remove data directory if you run cockroachdb on darwin
$ rm -rf ./cockroach-data ./.store
$ ./bin/cockroach/cockroach start --background
CockroachDB node starting at 2017-01-29 21:10:53.690109761 +0900 JST
build:      CCL beta-20170126 @ 2017/01/26 16:04:08 (go1.7.4)
admin:      http://localhost:8080
sql:        postgresql://root@localhost:26257?sslmode=disable
logs:       cockroach-data/logs
store[0]:   path=cockroach-data
status:     initialized new cluster
clusterID:  418606af-5e9d-4c00-a2a6-75458a5ff5ff
nodeID:     1
```

Add more nodes to the cluster

```
# 2nd
$ ./bin/cockroach/cockroach start --store=.store/node2 --port=26258 --http-port=8081 --join=localhost:26257 --background
# 3rd
$ ./bin/cockroach/cockroach start --store=.store/node3 --port=26259 --http-port=8082 --join=localhost:26257 --background
```

Use the build-in SQL Client

```
$ ./bin/cockroach/cockroach sql
# Welcome to the cockroach SQL interface.
# All statements must be terminated by a semicolon.
# To exit: CTRL + D.
root@:26257> CREATE DATABASE bank;
CREATE DATABASE
root@:26257> CREATE TABLE bank.accounts (id INT PRIMARY KEY, balance DECIMAL);
CREATE TABLE
root@:26257> INSERT INTO bank.accounts VALUES (1, 1000.50);
INSERT 1
root@:26257> SELECT * FROM bank.accounts;
+----+---------+
| id | balance |
+----+---------+
|  1 |  1000.5 |
+----+---------+
(1 row)
root@:26257> \q
```

Stop the cluster

```
$ ./bin/cockroach/cockroach quit
ok
$ ./bin/cockroach/cockroach quit --port=26258
ok
$ ./bin/cockroach/cockroach quit --port=26259
ok
```

## Build a Test Apps By Golang

Create a new user 'maxroach' by `cockroach user` command

```
$ ./bin/cockroach/cockroach user set maxroach
INSERT 1
```

Create Database as `root` user

```
$ ./bin/cockroach/cockroach sql -e 'CREATE DATABASE bank'
CREATE DATABASE
```


Grant privileges to the `maxroach` user.  See also: https://www.cockroachlabs.com/docs/grant.html

```
$ ./bin/cockroach/cockroach sql -e 'GRANT ALL ON DATABASE bank TO maxroach'
GRANT
```

Create a table in the new database

```
$ ./bin/cockroach/cockroach sql --database=bank --user=maxroach -e 'CREATE TABLE accounts (id INT PRIMARY KEY, balance INT)'
```

### Simple Code by postgresql driver

Install postgresql driver. See also https://www.cockroachlabs.com/docs/install-client-drivers.html and https://godoc.org/github.com/lib/pq

```
$ go get github.com/lib/pq
```

Write code. See also [./app1/main.go](./app1/main.go)

### Transactions from a client

Install CockroachDB Go client. See also https://www.cockroachlabs.com/docs/transactions.html#transaction-retries

```
$ mkdir $GOPATH/src/github.com/cockroachdb
$ cd $_
$ git clone https://github.com/cockroachdb/cockroach-go
```

Write code. See also [./app2/main.go](./app2/main.go)
