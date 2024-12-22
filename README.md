# Blockchain_DApp
同济大学2024年秋区块链导论期末课程项目



# 启动方式

## 1. 克隆本项目到本地(需要linux环境)

```bash
git clone https://github.com/WWWJ66/Blockchain_DApp.git
```



## 2. 启动区块链部分

```bash
cd Blockchain_DApp/blockchain/network

# 仅首次运行需要
./install-fabric.sh -f 2.5.6 d 

./start.sh
```



## 3. 启动后端部分

```bash
cd Blockchain_DApp/application/backend

go run main.go
```



## 4. 启动前端

```bash
cd Blockchain_DApp/application/frontend

# 仅首次运行需要
npm install

npm run serve
```

