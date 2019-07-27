# fabric-samples第一个网络的建立

### 1.启动

```
./byfn.sh -m down  关闭
./byfn.sh -m generate   启动，产生创世节点
```

### 2.运行网络

#### 手动生成`artifacts`

首先让我们运行该`cryptogen`工具。我们的二进制文件位于`bin` 目录中，因此我们需要提供工具所在位置的相对路径。

```
../bin/cryptogen generate --config=./crypto-config.yaml
```

您应该在终端中看到以下内容：

```
org1.example.com
org2.example.com
```

证书和密钥（即MSP材料）将输出到目录 - `crypto-config`- 目录的根`first-network`目录。

接下来，我们需要告诉`configtxgen`工具在哪里查找`configtx.yaml`它需要摄取的 文件。我们将在目前的工作目录中告诉它：

```
export FABRIC_CFG_PATH=$PWD
```

然后，我们将调用该`configtxgen`工具来创建orderer genesis块：

```
../bin/configtxgen -profile TwoOrgsOrdererGenesis -channelID byfn-sys-channel -outputBlock ./channel-artifacts/genesis.block
```

#### 创建通道配置事务

接下来，我们需要创建通道事务工件。请务必替换`$CHANNEL_NAME`或设置`CHANNEL_NAME`为可在整个说明中使用的环境变量：

```
# The channel.tx artifact contains the definitions for our sample channel
export CHANNEL_NAME=mychannel  && ../bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME
```

接下来，我们将在我们构建的通道上为Org1定义锚点对等体。同样，请确保替换`$CHANNEL_NAME`或设置以下命令的环境变量。终端输出将模仿通道事务工件的输出：

```
../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP
```

现在，我们将在同一个通道上为Org2定义锚点对等体：

```
../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org2MSP
```

### 3. 操作网络

手动操作的话，需要修改一下脚步，打开`docker-compose-cli.yaml`文件：

```
gedit docker-compose-cli.yaml
```

注释掉

```
command: /bin/bash
```



首先让我们开始我们的网络：

```
CHANNEL_NAME=$CHANNEL_NAME TIMEOUT=600 docker-compose -f docker-compose-cli.yaml up -d
```

我们将使用以下命令进入CLI容器：`docker exec`

```
docker exec -it cli bash
```

如果成功，您应该看到以下内容：

```
root@0d78bb69300d:/opt/gopath/src/github.com/hyperledger/fabric/peer#
```

##### 创建通道

```
export CHANNEL_NAME=mychannel

peer channel create -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```

##### 加入通道

加入`peer0.org1.example.com`通道：

```
# By default, this joins ``peer0.org1.example.com`` only
peer channel join -b mychannel.block
```



##### 安装链码（智能合约）

```
peer chaincode install -n mycc -v 1.0 -p github.com/chaincode/chaincode_example02/go/

这个链码其实是云上的，不是本地的链码
```

##### 实例化链码

```
peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc -v 1.0 -c '{"Args":["init","a", "100", "b","200"]}' -P "AND ('Org1MSP.peer','Org2MSP.peer')"

#7050是排序服务的默认接口
#'{"Args":["init","a", "100", "b","200"]}' -P "AND ('Org1MSP.peer','Org2MSP.peer')"
#参数设置，比如，a节点，给了100币，b节点，给了200币
#mycc 就是我们的智能合约对象的实例
#
```

##### 查询

```
peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["query","a"]}'
```

##### 转账

```
peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc --peerAddresses peer0.org1.example.com:9051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses peer0.org2.example.com:9051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"Args":["invoke","a","b","10"]}'
```



### 4.快速版

当然，有快速部署

```
./byfn.sh up
```

不出意外的话，他会自动帮你部署，并且加入通道，测试链码，实例化

```
docker exec -it cli bash
peer chaincode query -C mychannel -n mycc -c '{"Args":["query","a"]}'
```

即可查询到