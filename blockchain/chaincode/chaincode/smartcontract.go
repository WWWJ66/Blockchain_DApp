package chaincode

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// 定义合约结构体
type SmartContract struct {
	contractapi.Contract
}

// 注册用户
func (s *SmartContract) RegisterUser(ctx contractapi.TransactionContextInterface, userID string, userType string, realInfoHash string) error {
	user := User{
		UserID:       userID,
		UserType:     userType,
		RealInfoHash: realInfoHash,
		FishList:    []*Fish{},
	}
	userAsBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(userID, userAsBytes)
	if err != nil {
		return err
	}
	return nil
}

// 水产品上链，传入用户ID、水产品上链信息
func (s *SmartContract) Uplink(ctx contractapi.TransactionContextInterface, userID string, tracecode string, arg1 string, arg2 string, arg3 string, arg4 string,) (string, error) {
	// 获取用户类型
	userType, err := s.GetUserType(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("failed to get user type: %v", err)
	}

	// 通过溯源码获取水产品的上链信息
	FishAsBytes, err := ctx.GetStub().GetState(tracecode)
	if err != nil {
		return "", fmt.Errorf("failed to read from world state: %v", err)
	}
	// 将水产品的信息转换为Fish结构体
	var fish Fish
	if FishAsBytes != nil {
		err = json.Unmarshal(FishAsBytes, &fish)
		if err != nil {
			return "", fmt.Errorf("failed to unmarshal fish: %v", err)
		}
	}

	//获取时间戳,修正时区
	txtime, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return "", fmt.Errorf("failed to read TxTimestamp: %v", err)
	}
	timeLocation, _ := time.LoadLocation("Asia/Shanghai") // 选择你所在的时区
	time := time.Unix(txtime.Seconds, 0).In(timeLocation).Format("2006-01-02 15:04:05")

	// 获取交易ID
	txid := ctx.GetStub().GetTxID()
	// 给水产品信息中加上溯源码
	fish.Tracecode = tracecode
	// 不同用户类型的上链的参数不一致
	switch userType {
	// 水产品养殖者
	case "水产品养殖者":
		// 将传入的水产品上链信息转换为Fisher_input结构体
		fish.Fisher_input.Fi_fishName = arg1
		fish.Fisher_input.Fi_fisherName = arg2
		fish.Fisher_input.Fi_fishedTime = arg3
		fish.Fisher_input.Fi_fishedArea = arg4
		fish.Fisher_input.Fi_Txid = txid
		fish.Fisher_input.Fi_Timestamp = time
	// 工厂
	case "工厂":
		// 将传入的水产品上链信息转换为Factory_input结构体
		fish.Factory_input.Fac_productName = arg1
		fish.Factory_input.Fac_factoryName = arg2
		fish.Factory_input.Fac_processTime = arg3
		fish.Factory_input.Fac_factoryAddress = arg4
		fish.Factory_input.Fac_Txid = txid
		fish.Factory_input.Fac_Timestamp = time
	// 物流公司
	case "物流公司":
		// 将传入的水产品上链信息转换为Driver_input结构体
		fish.Driver_input.Dr_name = arg1
		fish.Driver_input.Dr_phone = arg2
		fish.Driver_input.Dr_transportTime = arg3
		fish.Driver_input.Dr_id = arg4
		fish.Driver_input.Dr_Txid = txid
		fish.Driver_input.Dr_Timestamp = time
	// 商店
	case "商店":
		// 将传入的水产品上链信息转换为Shop_input结构体
		fish.Shop_input.Sh_shopName = arg1
		fish.Shop_input.Sh_shopAddress = arg2
		fish.Shop_input.Sh_storageTime = arg3
		fish.Shop_input.Sh_sellTime = arg4
		fish.Shop_input.Sh_Txid = txid
		fish.Shop_input.Sh_Timestamp = time

	}

	//将水产品的信息转换为json格式
	fishAsBytes, err := json.Marshal(fish)
	if err != nil {
		return "", fmt.Errorf("failed to marshal fish: %v", err)
	}
	//将水产品的信息存入区块链
	err = ctx.GetStub().PutState(tracecode, fishAsBytes)
	if err != nil {
		return "", fmt.Errorf("failed to put fish: %v", err)
	}
	//将水产品存入用户的水产品列表
	err = s.AddFish(ctx, userID, &fish)
	if err != nil {
		return "", fmt.Errorf("failed to add fish to user: %v", err)

	}

	return txid, nil
}

// 添加水产品到用户的水产品列表
func (s *SmartContract) AddFish(ctx contractapi.TransactionContextInterface, userID string, fish *Fish) error {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return err
	}
	user.FishList = append(user.FishList, fish)
	userAsBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(userID, userAsBytes)
	if err != nil {
		return err
	}
	return nil
}

// 获取用户类型
func (s *SmartContract) GetUserType(ctx contractapi.TransactionContextInterface, userID string) (string, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return "", fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return "", fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return "", err
	}
	return user.UserType, nil
}

// 获取用户信息
func (s *SmartContract) GetUserInfo(ctx contractapi.TransactionContextInterface, userID string) (*User, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return &User{}, fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return &User{}, fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

// 获取水产品的上链信息
func (s *SmartContract) GetFishInfo(ctx contractapi.TransactionContextInterface, tracecode string) (*Fish, error) {
	FishAsBytes, err := ctx.GetStub().GetState(tracecode)
	if err != nil {
		return &Fish{}, fmt.Errorf("failed to read from world state: %v", err)
	}

	// 将返回的结果转换为Fish结构体
	var fish Fish
	if FishAsBytes != nil {
		err = json.Unmarshal(FishAsBytes, &fish)
		if err != nil {
			return &Fish{}, fmt.Errorf("failed to unmarshal fish: %v", err)
		}
		if fish.Tracecode != "" {
			return &fish, nil
		}
	}
	return &Fish{}, fmt.Errorf("the fish %s does not exist", tracecode)
}

// 获取用户的水产品ID列表
func (s *SmartContract) GetFishList(ctx contractapi.TransactionContextInterface, userID string) ([]*Fish, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return nil, fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return nil, err
	}
	return user.FishList, nil
}

// 获取所有的水产品信息
func (s *SmartContract) GetAllFishInfo(ctx contractapi.TransactionContextInterface) ([]Fish, error) {
	fishListAsBytes, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	defer fishListAsBytes.Close()
	var fishes []Fish
	for fishListAsBytes.HasNext() {
		queryResponse, err := fishListAsBytes.Next()
		if err != nil {
			return nil, err
		}
		var fish Fish
		err = json.Unmarshal(queryResponse.Value, &fish)
		if err != nil {
			return nil, err
		}
		// 过滤非水产品的信息
		if fish.Tracecode != "" {
			fishes = append(fishes, fish)
		}
	}
	return fishes, nil
}

// 获取水产品上链历史
func (s *SmartContract) GetFishHistory(ctx contractapi.TransactionContextInterface, tracecode string) ([]HistoryQueryResult, error) {
	log.Printf("GetFishHistory: ID %v", tracecode)

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(tracecode)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var records []HistoryQueryResult
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var fish Fish
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &fish)
			if err != nil {
				return nil, err
			}
		} else {
			fish = Fish{
				Tracecode: tracecode,
			}
		}

		timestamp, err := ptypes.Timestamp(response.Timestamp)
		if err != nil {
			return nil, err
		}
		// 指定目标时区
		targetLocation, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			return nil, err
		}

		// 将时间戳转换到目标时区
		timestamp = timestamp.In(targetLocation)
		// 格式化时间戳为指定格式的字符串
		formattedTime := timestamp.Format("2006-01-02 15:04:05")

		record := HistoryQueryResult{
			TxId:      response.TxId,
			Timestamp: formattedTime,
			Record:    &fish,
			IsDelete:  response.IsDelete,
		}
		records = append(records, record)
	}

	return records, nil
}
