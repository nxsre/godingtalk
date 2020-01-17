package godingtalk

import (
	"fmt"
	"github.com/jinzhu/copier"
	jsoniter "github.com/json-iterator/go"
	"math/big"
	"net/url"
	"strconv"
	"strings"
)

type User struct {
	OAPIResponse
	Userid     string
	Name       string
	Mobile     string
	Tel        string
	Remark     string
	Order      int
	IsAdmin    bool
	IsBoss     bool
	IsLeader   bool
	IsSys      bool `json:"is_sys"`
	SysLevel   int  `json:"sys_level"`
	Active     bool
	Department []int
	Position   string
	Email      string
	Avatar     string
	Extattr    interface{}
}

type UserInfo struct {
	OAPIResponse
	UnionId         string           `json:"unionid"`
	OpenId          string           `json:"openid"`
	Userid          string           `json:"userid"`
	Roles           []Role           `json:"roles"`
	IsLeaderInDepts LeaderInDepts    `json:"isLeaderInDepts"`
	IsBoss          bool             `json:"isBoss"`
	HiredDate       int64            `json:"hiredDate"`
	IsSenior        bool             `json:"isSenior"`
	Department      []int            `json:"department"`
	OrderInDepts    map[int]*big.Int `json:"orderInDepts"`
	Mobile          string           `json:"mobile"`
	ErrMsg          string           `json:"errmsg"`
	Active          bool             `json:"active"`
	Avatar          string           `json:"avatar"`
	IsAdmin         bool             `json:"isAdmin"`
	Tags            interface{}      `json:"tags"`
	IsHide          bool             `json:"isHide"`
	JobNumber       string           `json:"jobnumber"`
	Name            string           `json:"name"`
	StateCode       string           `json:"stateCode"`
	Position        string           `json:"position"`
	Email           string           `json:"email"`
	Remark          string           `json:"remark"`
	Tel             string           `json:"tel"`
	WorkPlace       string           `json:"workPlace"`
}

type LeaderInDepts map[int]bool

func (u *UserInfo) UnmarshalJSON(b []byte) error {
	tmpUser := struct {
		OAPIResponse
		UnionId         string      `json:"unionid"`
		OpenId          string      `json:"openid"`
		Userid          string      `json:"userid"`
		Roles           []Role      `json:"roles"`
		IsLeaderInDepts string      `json:"isLeaderInDepts"`
		IsBoss          bool        `json:"isBoss"`
		HiredDate       int64       `json:"hiredDate"`
		IsSenior        bool        `json:"isSenior"`
		Department      []int       `json:"department"`
		OrderInDepts    string      `json:"orderInDepts"`
		Mobile          string      `json:"mobile"`
		ErrMsg          string      `json:"errmsg"`
		Active          bool        `json:"active"`
		Avatar          string      `json:"avatar"`
		IsAdmin         bool        `json:"isAdmin"`
		Tags            interface{} `json:"tags"`
		IsHide          bool        `json:"isHide"`
		JobNumber       string      `json:"jobnumber"`
		Name            string      `json:"name"`
		StateCode       string      `json:"stateCode"`
		Position        string      `json:"position"`
		Email           string      `json:"email"`
		Remark          string      `json:"remark"`
		Tel             string      `json:"tel"`
		WorkPlace       string      `json:"workPlace"`
	}{}
	err := jsoniter.Unmarshal(b, &tmpUser)
	if err != nil {
		return err
	}

	copier.Copy(u, &tmpUser)

	leaderInDepts := strings.TrimFunc(tmpUser.IsLeaderInDepts, func(r rune) bool {
		switch r {
		case rune('{'), rune('}'):
			return true
		}
		return false
	})

	isLeaderInDepts := LeaderInDepts{}
	for _, dept := range strings.Split(leaderInDepts, ",") {
		lead := strings.Split(dept, ":")
		if len(lead) == 2 {
			if deptId, err := strconv.Atoi(lead[0]); err == nil {
				switch lead[1] {
				case "true":
					isLeaderInDepts[deptId] = true
				case "false":
					isLeaderInDepts[deptId] = false
				}
			}
		}
	}

	orderInDepts := strings.TrimFunc(tmpUser.OrderInDepts, func(r rune) bool {
		switch r {
		case rune('{'), rune('}'):
			return true
		}
		return false
	})

	deptOrder := map[int]*big.Int{}
	for _, dept := range strings.Split(orderInDepts, ",") {
		orders := strings.Split(dept, ":")
		if len(orders) == 2 {
			if deptId, err := strconv.Atoi(orders[0]); err == nil {
				i := big.Int{}
				if order, ok := i.SetString(orders[1], 10); ok {
					deptOrder[deptId] = order
				}
			}
		}
	}

	u.OrderInDepts = deptOrder
	u.IsLeaderInDepts = isLeaderInDepts

	return nil
}

type Role struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	GroupName string `json:"groupName"`
	Type      int    `json:"type"`
}

type UserList struct {
	OAPIResponse
	HasMore  bool
	Userlist []User
}

type Department struct {
	OAPIResponse
	Id                    int
	Name                  string
	ParentId              int
	Order                 int
	DeptPerimits          string
	UserPerimits          string
	OuterDept             bool
	OuterPermitDepts      string
	OuterPermitUsers      string
	OrgDeptOwner          string
	DeptManagerUseridList string
}

type DepartmentList struct {
	OAPIResponse
	Departments []Department `json:"department"`
}

// DepartmentList is 获取部门列表
func (c *DingTalkClient) DepartmentList(id int, fetch_child bool) (DepartmentList, error) {
	var data DepartmentList
	params := url.Values{}
	if id == 0 {
		id = 1
	}
	params.Add("id", fmt.Sprintf("%d", id))
	params.Add("fetch_child", fmt.Sprintf("%t", fetch_child))

	err := c.httpRPC("department/list", params, nil, &data)
	return data, err
}

//DepartmentDetail is 获取部门详情
func (c *DingTalkClient) DepartmentDetail(id int) (Department, error) {
	var data Department
	params := url.Values{}
	params.Add("id", fmt.Sprintf("%d", id))
	err := c.httpRPC("department/get", params, nil, &data)
	return data, err
}

//UserList is 获取部门成员
func (c *DingTalkClient) UserList(departmentID, offset, size int) (UserList, error) {
	var data UserList
	if size > 100 {
		return data, fmt.Errorf("size 最大100")
	}

	params := url.Values{}
	params.Add("department_id", fmt.Sprintf("%d", departmentID))
	params.Add("offset", fmt.Sprintf("%d", offset))
	params.Add("size", fmt.Sprintf("%d", size))
	err := c.httpRPC("user/list", params, nil, &data)
	return data, err
}

//CreateChat is 
func (c *DingTalkClient) CreateChat(name string, owner string, useridlist []string) (string, error) {
	var data struct {
		OAPIResponse
		Chatid string
	}
	request := map[string]interface{}{
		"name":       name,
		"owner":      owner,
		"useridlist": useridlist,
	}
	err := c.httpRPC("chat/create", nil, request, &data)
	return data.Chatid, err
}

//UserInfoByID 获取用户详情
func (c *DingTalkClient) UserInfoByID(userid string) (UserInfo, error) {
	var data UserInfo
	params := url.Values{}
	params.Add("userid", userid)
	err := c.httpRPC("user/get", params, nil, &data)
	return data, err
}

//UserInfoByMobile 获取用户详情
func (c *DingTalkClient) UseridByMobile(mobile string) (string, error) {
	var data UserInfo
	params := url.Values{}
	params.Add("mobile", mobile)
	err := c.httpRPC("user/get_by_mobile", params, nil, &data)
	return data.Userid, err
}

//UserInfoByCode 校验免登录码并换取用户身份
func (c *DingTalkClient) UserInfoByCode(code string) (User, error) {
	var data User
	params := url.Values{}
	params.Add("code", code)
	err := c.httpRPC("user/getuserinfo", params, nil, &data)
	return data, err
}

//UseridByUnionId 通过UnionId获取玩家Userid
func (c *DingTalkClient) UseridByUnionId(unionid string) (string, error) {
	var data struct {
		OAPIResponse
		UserID string `json:"userid"`
	}

	params := url.Values{}
	params.Add("unionid", unionid)
	err := c.httpRPC("user/getUseridByUnionid", params, nil, &data)
	if err != nil {
		return "", err
	}

	return data.UserID, err
}
