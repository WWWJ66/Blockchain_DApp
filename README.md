# 区块链启动教程



## 1.  运行以下命令拉取*Hyperledger Fabric*镜像

```bash
# 不要忘了后面的参数'd'，其它所需文件我已下载好，所以只需docker拉取镜像即可
bash ./install-fabric.sh d
```



## 2. 启动网络并测试test链码

具体教程可参考[将智能合约部署到通道 — hyperledger-fabricdocs master 文档](https://hyperledger-fabric.readthedocs.io/zh-cn/latest/deploy_chaincode.html)

```bash
# 进入network文件夹
cd network

# 第一次启动网络前线清空网络，防止上次测试后未清空
./network.sh down

# 启动网络并创建通道
./network.sh up createChannel

# 将test链码部署到通道上
./network.sh deployCC -ccn test -ccp ../test/chaincode-go -ccl go

# 配置peer环境变量
export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/
peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

# 调用初始化函数
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n test --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"InitStudent","Args":[]}'

# 调用查询函数
peer chaincode query -C mychannel -n test -c '{"Args":["GetAllStudent"]}'

# 调用插入函数
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n test --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"CreateStudent","Args":["S003", "口苗", "3"]}'

# 调用查询函数
peer chaincode query -C mychannel -n test -c '{"Args":["GetAllStudent"]}'

# 调用按id查询函数
peer chaincode query -C mychannel -n test -c '{"Args":["ReadStudent", "S002"]}'

# 完成test，关闭网络清空数据
./network.sh down
```



