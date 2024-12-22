package chaincode

/*
定义用户结构体
用户ID
用户类型
实名认证信息哈希,包括用户注册的姓名、身份证号、手机号、注册平台同意协议签名的哈希
水产品列表
*/
type User struct {
	UserID       string   `json:"userID"`
	UserType     string   `json:"userType"`
	RealInfoHash string   `json:"realInfoHash"`
	FishList    []*Fish 	`json:"fishList"`
}

/*
定义水产品结构体
溯源码
水产品养殖者输入
工厂输入
物流公司输入
商店输入
*/
type Fish struct {
	Tracecode         string        `json:"tracecode"`
	Fisher_input      Fisher_input  `json:"fisher_input"`
	Factory_input     Factory_input `json:"factory_input"`
	Driver_input      Driver_input  `json:"driver_input"`
	Shop_input        Shop_input    `json:"shop_input"`
}

type HistoryQueryResult struct {
	Record    *Fish `json:"record"`
	TxId      string `json:"txId"`
	Timestamp string `json:"timestamp"`
	IsDelete  bool   `json:"isDelete"`
}

/*
水产品养殖者输入
水产品名称
水产品养殖者名称
捕捞时间
捕捞地点
*/
type Fisher_input struct {
	Fi_fishName    	string `json:"fi_fishName"`
	Fi_fisherName  	string `json:"fi_fisherName"`
	Fi_fishedTime  	string `json:"fi_fishedTime"`
	Fi_fishedArea  	string `json:"fi_fishedArea"`
	Fi_Txid        	string `json:"fi_txid"`
	Fi_Timestamp   	string `json:"fi_timestamp"`
}

/*
工厂输入
商品名称
工厂名称
加工时间
工厂地址
*/
type Factory_input struct {
	Fac_productName     string `json:"fac_productName"`
	Fac_factoryName     string `json:"fac_factoryName"`
	Fac_processTime     string `json:"fac_processTime"`
	Fac_factoryAddress  string `json:"fac_factoryAddress"`
	Fac_Txid            string `json:"fac_txid"`
	Fac_Timestamp       string `json:"fac_timestamp"`
}

/*
物流公司输入
公司名称
联系电话
运输时间
运输单号
*/
type Driver_input struct {
	Dr_name				string `json:"dr_name"`
	Dr_phone    	 	string `json:"dr_phone"`
	Dr_transportTime   	string `json:"dr_transportTime"`
	Dr_id     			string `json:"dr_id"`
	Dr_Txid      		string `json:"dr_txid"`
	Dr_Timestamp 		string `json:"dr_timestamp"`
}

/*
商店输入
商店名称
商店地址
存入时间
销售时间
*/
type Shop_input struct {
	Sh_shopName		string `json:"sh_shopName"`
	Sh_shopAddress	string `json:"sh_shopAddress"`
	Sh_storageTime	string `json:"sh_storageTime"`
	Sh_sellTime		string `json:"sh_sellTime"`
	Sh_Txid			string `json:"sh_txid"`
	Sh_Timestamp	string `json:"sh_timestamp"`
}
