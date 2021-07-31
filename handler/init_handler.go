package handler

import (
	"gin_oauth2_server/constant"
	"gin_oauth2_server/domain"
	"gin_oauth2_server/mapper"
	"gin_oauth2_server/util"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ClientInitHandler 客户端授权模式
func ClientInitDataHandler(c *gin.Context) {
	baseMapper := mapper.BaseMapper
	s1 := util.PasswordEncrypt("7KutwpFgFXv0hcvkBO")
	s2 := util.PasswordEncrypt("nssbTtp5FO6NjZpUwP")
	s3 := util.PasswordEncrypt("VIUvXZmVXmOFh1gYWK")
	var clients = []domain.ClientDetails{
		{
			ClientId:    "7KutwpFgFXv0hcvkBO",
			Secret:      s1,
			Name:        "商户管理后台",
			Logo:        "https://images.xaaef.com/20210719104730.jpg",
			Description: "商户管理后台-描述",
			ClientType:  1,
			GrantTypes:  "[\"we_chat\",\"password\",\"tencent_qq\",\"sms\"]",
			DomainName:  "www.xaaef.com",
			Scope:       "read,write",
			Status:      constant.Normal,
		},
		{
			ClientId:    "nssbTtp5FO6NjZpUwP",
			Secret:      s2,
			Name:        "vue.js总部管理后台",
			Logo:        "https://tl329pszs0mg.oss-cn-shenzhen.aliyuncs.com/avatar/20210728105923.png",
			Description: "vue.js总部管理后台-描述",
			ClientType:  1,
			GrantTypes:  "[\"*\"]",
			DomainName:  "www.baidu.com",
			Scope:       "read,write",
			Status:      constant.Normal,
		},
		{
			ClientId:    "VIUvXZmVXmOFh1gYWK",
			Secret:      s3,
			Name:        "华为IOT园区项目组",
			Logo:        "https://tl329pszs0mg.oss-cn-shenzhen.aliyuncs.com/avatar/20210728110412.jpg",
			Description: "华为IOT园区项目组-描述",
			ClientType:  1,
			GrantTypes:  "[\"client_credentials\"]",
			DomainName:  "www.xaaef.com",
			Scope:       "read,write",
			Status:      constant.Disable,
		},
	}

	baseMapper.Create(clients)

	c.JSON(http.StatusOK, domain.Ok2(clients))
	return
}

// 初始化用户
func UserInitHandler(c *gin.Context) {
	baseMapper := mapper.BaseMapper
	u1 := "admin"
	u2 := "hello"
	// Create a new Node with a Node number of 1
	worker, _ := snowflake.NewNode(1)

	users := []domain.UserInfo{
		{
			UserId:    worker.Generate().String(),
			Avatar:    "https://images.xaaef.com/b9a7abacafd747bbb74cf7cb3de36c1e.png",
			Username:  u1,
			Mobile:    "15071526233",
			Email:     "15071526233@qq.com",
			Nickname:  "管理员",
			Password:  util.PasswordEncrypt(u1),
			Status:    constant.Normal,
			AdminFlag: constant.YES,
		},
		{
			UserId:    worker.Generate().String(),
			Avatar:    "https://images.xaaef.com/63e240e0f1de45c1ad7b818405e659c3.jpg",
			Username:  u2,
			Mobile:    "15071526222",
			Email:     "15071526222@qq.com",
			Nickname:  "测试人员",
			Password:  util.PasswordEncrypt(u2),
			Status:    constant.Disable,
			AdminFlag: constant.NO,
		},
	}
	baseMapper.Create(users)
	c.JSON(http.StatusOK, domain.Ok2(users))
	return
}
