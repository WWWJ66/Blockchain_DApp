package controller

import (
	"backend/pkg"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 保存了水产品上链与查询的函数

func Uplink(c *gin.Context) {
	// 与userID不一样，取ID从第二位作为溯源码，即18位，雪花ID的结构如下（转化为10进制共19位）：
	// +--------------------------------------------------------------------------+
	// | 1 Bit Unused | 41 Bit Timestamp |  10 Bit NodeID  |   12 Bit Sequence ID |
	// +--------------------------------------------------------------------------+
	fisher_tracecode := pkg.GenerateID()[1:]
	args := buildArgs(c, fisher_tracecode)
	if args == nil {
		return
	}
	res, err := pkg.ChaincodeInvoke("Uplink", args)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "uplink failed" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":		200,
		"message":	"uplink success",
		"txid":		res,
		"tracecode": 	args[1],
	})
}

// 获取水产品的上链信息
func GetFishInfo(c *gin.Context) {
	res, err := pkg.ChaincodeQuery("GetFishInfo", c.PostForm("tracecode"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "查询失败：" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "query success",
		"data":    res,
	})

}

// 获取用户的水产品ID列表
func GetFishList(c *gin.Context) {
	userID, _ := c.Get("userID")
	res, err := pkg.ChaincodeQuery("GetFishList", userID.(string))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "query failed" + err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "query success",
		"data":    res,
	})
}

// 获取所有的农产品信息
func GetAllFishInfo(c *gin.Context) {
	res, err := pkg.ChaincodeQuery("GetAllFishInfo", "")
	fmt.Print("res", res)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "query failed" + err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "query success",
		"data":    res,
	})
}

// 获取水产品上链历史
func GetFishHistory(c *gin.Context) {
	res, err := pkg.ChaincodeQuery("GetFishHistory", c.PostForm("trececode"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "query failed" + err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "query success",
		"data":    res,
	})
}

func buildArgs(c *gin.Context, fisher_tracecode string) []string {
	var args []string
	userID, _ := c.Get("userID")
	userType, _ := pkg.ChaincodeQuery("GetUserType", userID.(string))
	args = append(args, userID.(string))
	fmt.Print(userID)
	// 水产品养殖者不需要输入溯源码，其他用户需要，通过雪花算法生成ID
	if userType == "水产品养殖者" {
		args = append(args, fisher_tracecode)
	} else {
		// 检查溯源码是否正确
		res, err := pkg.ChaincodeQuery("GetFishInfo", c.PostForm("tracecode"))
		if res == "" || err != nil || len(c.PostForm("tracecode")) != 18 {
			c.JSON(200, gin.H{
				"message": "请检查溯源码是否正确!!",
			})
			return nil
		} else {
			args = append(args, c.PostForm("tracecode"))
		}
	}
	args = append(args, c.PostForm("arg1"))
	args = append(args, c.PostForm("arg2"))
	args = append(args, c.PostForm("arg3"))
	args = append(args, c.PostForm("arg4"))
	return args
}
