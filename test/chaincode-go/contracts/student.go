package student

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

// 合约结构体
type SmartContract struct {
	contractapi.Contract
}

// 定义学生结构体
type Student struct {
	ID	string	`json:"id"`
	Name	string	`json:"name"`
	GPA	int	`json:"gpa"`
}

/*
 * 功能: 初始化一些学生信息
 * 参数: 
 * 	- ctx: 链玛上下文对象
 * 返回值: 无
 */
func (s *SmartContract) InitStudent (ctx contractapi.TransactionContextInterface) error {
	// 随机生成一些数据
	students := []Student {
		{ID:"S001",Name:"Guava",GPA:5},
		{ID:"S002",Name:"精彩",GPA:1},
	}
	
	// 存入
	for _, student := range students {
		studentJSON, err := json.Marshal(student)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(student.ID, studentJSON)
		if(err != nil) {
			return err
		}
	}
	
	return nil
}

/*
 * 功能: 创建一条新的学生信息
 * 参数: 
 * 	- ctx: 链玛上下文对象
 * 	- id: 学生ID
 * 	- name: 学生姓名
 * 	- gpa: 学生寄点
 * 返回值: 保存状态?
 */
func (s *SmartContract) CreateStudent (ctx contractapi.TransactionContextInterface, id string, name string, gpa int) error {
	// 判断是否已存在
	exists, err := s.StudentExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the student %s already exists", id)
	}
	
	// 创建新实例
	student := Student {
		ID:	id,
		Name:	name,
		GPA:	gpa,
	}
	// 转为json
	studentJSON, err := json.Marshal(student)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, studentJSON)
}

/*
 * 功能: 更新学生(绩点)
 * 参数: 
 * 	- ctx: 链玛上下文对象
 * 	- id: 学生ID
 * 	- newGPA: 新学生寄点
 * 返回值: 保存状态?
 */
func (s *SmartContract) UpdateStudent (ctx contractapi.TransactionContextInterface,id string,newGPA int) error {
	// 读取学生信息
	student, err := s.ReadStudent(ctx, id)

	// 创建
	student.GPA = newGPA
	studentJSON, err := json.Marshal(student)
	if err != nil {
		return err
	}
	// 写入
	return ctx.GetStub().PutState(id, studentJSON)
}

/*
 * 功能: 根据ID查询学生信息
 * 参数: 
 * 	- ctx: 链玛上下文对象
 * 	- id: 学生ID
 * 返回值: 学生信息
 */
func (s *SmartContract) ReadStudent (ctx contractapi.TransactionContextInterface, id string) (*Student,error) {
	// 根据id查询
	studentJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if studentJSON == nil {
		return nil, fmt.Errorf("the student %s does not exist", id)
	}
	
	// 存储信息
	var student Student
	err = json.Unmarshal(studentJSON, &student)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

/*
 * 功能: 查询所有学生信息
 * 参数: 
 * 	- ctx: 链玛上下文对象
 * 返回值: 学生信息列表
 */
func (s *SmartContract) GetAllStudent (ctx contractapi.TransactionContextInterface)([]*Student, error) {
	// 查询所有信息
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	
	// 读取并保存所有信息
	var students []*Student
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var student Student
		err = json.Unmarshal(queryResponse.Value, &student)
		if err != nil {
			return nil, err
		}
		students = append(students, &student)
	}

	return students, nil
}

/*
 * 功能: 判断id是否存在
 * 参数: 
 * 	- ctx: 链玛上下文对象
 * 	- id: 学生ID
 * 返回值: id已存在返回真，否则反之
 */
func (s *SmartContract) StudentExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	// 根据id读取信息
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

