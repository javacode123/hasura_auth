package service

import (
	"encoding/json"
	"fmt"
	"github.com/graph-gophers/graphql-go"
	"github.com/hasura_auth/config"
	"github.com/hasura_auth/model"
	"github.com/spf13/cast"
	"testing"
)

func TestUserInfo(t *testing.T) {
	config.GlobalConfig.HasuraHost = "http://127.0.0.1:3000"
	userInfo, err := UserInfo("name", "pwd")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(userInfo)
	}
}

func TestMarshalRsp(t *testing.T) {
	body := []byte(`{
  "data": {
    "user": {
      "id": "8888",
      "name": "tester@tester.com",
      "password": "********",
      "roles": ["admin"]
    }
  }
}`)
	graphqlRsp := &graphql.Response{}
	if err := json.Unmarshal(body, graphqlRsp); err != nil {
		t.Error(err)
	}
	if len(graphqlRsp.Errors) > 0 {
		t.Error(graphqlRsp.Errors)
	}

	userInfo := &model.HasuraUserRsp{}
	fmt.Println(string(graphqlRsp.Data))
	if err := json.Unmarshal(graphqlRsp.Data, userInfo); err != nil {
		t.Error(err)
	}
	res := model.User{
		UserId: cast.ToInt64(userInfo.User.Id),
		Name:   userInfo.User.Name,
		Roles:  userInfo.User.Roles,
	}
	t.Logf("%+v", res)
}
