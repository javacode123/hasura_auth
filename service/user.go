package service

import (
	"encoding/json"
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/graph-gophers/graphql-go"
	"github.com/hasura_auth/config"
	"github.com/hasura_auth/model"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
	"strings"
)

var graphqlQueryPath = "/query"

func UserInfo(Name, Password string) (model.User, error) {
	req, err := http.NewRequest("POST",
		config.GlobalConfig.HasuraHost+graphqlQueryPath,
		strings.NewReader(
			fmt.Sprintf(`{
				"query": "{user (name: \"%s\", pwd: \"%s\") {id, name, password}}"
			}`, Name, Password),
		),
	)
	if err != nil {
		return model.User{}, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return model.User{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.User{}, err
	}

	graphqlRsp := &graphql.Response{}
	if err := json.Unmarshal(body, graphqlRsp); err != nil {
		return model.User{}, err
	}
	if len(graphqlRsp.Errors) > 0 {
		return model.User{}, errors.Newf("query hasura meet err: %+v", graphqlRsp.Errors)
	}

	userInfo := &model.HasuraUserRsp{}
	if err := json.Unmarshal(graphqlRsp.Data, userInfo); err != nil {
		return model.User{}, err
	}

	return model.User{
		UserId: cast.ToInt64(userInfo.User.Id),
		Name:   userInfo.User.Name,
		Roles:  userInfo.User.Roles,
	}, nil
}
