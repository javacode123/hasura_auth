###项目介绍  
>使用 graphql 接口生成 jwt token, 数据查询调用下游的的 graphql 服务

###项目配置  
>配置下游的 graphql 服务：  
>>config.hasura.host = "服务地址" 
> 
>下游 query schema:  
>>user(name: String!, password: String!): User  
type User {  
&#160;&#160;&#160;&#160;id: ID!  
&#160;&#160;&#160;&#160;name: String  
&#160;&#160;&#160;&#160;password: String  
&#160;&#160;&#160;&#160;roles: []string  
}
 
###项目开发  
>通过命令自动生成对应的 scahma:  
go-bindata -ignore=\.go -pkg=schema -o=schema/bindata.go schema/...
