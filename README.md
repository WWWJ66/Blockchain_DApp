# 区块链启动教程



## 1.  运行以下命令拉取*Hyperledger Fabric*镜像

```bash
# 不要忘了后面的参数'd'，其它所需文件我已下载好，所以只需docker拉取镜像即可
$ bash ./install-fabric.sh d
```



## 2. 启动网络

具体教程可参考[将智能合约部署到通道 — hyperledger-fabricdocs master 文档](https://hyperledger-fabric.readthedocs.io/zh-cn/latest/deploy_chaincode.html)

```bash
# 第一次启动网络前线清空网络，防止上次测试后未清空
$ ./network.sh down

# 启动网络并创建通道
$ ./network.sh up createChannel

。。。太多了先不写了
```



