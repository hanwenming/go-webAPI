package service

import (
	"boxProject/tools"
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"io"
	"net/http"
)

type UserResource struct {
	UuId       string `json:"uu_id"`
	CreateTime string `json:"create_time"`
	Password   string `json:"password"`
	RoleId     string `json:"role_id"`
	RoleName   string `json:"role_name"`
	UpdateTime string `json:"update_time"`
	UserName   string `json:"user_name"`
	LockTime   string `json:"lock_time"`
	LoginCount string `json:"login_count"`
}

func (u UserResource) RegisterTo(container *restful.Container) {
	ws := new(restful.WebService)
	//设置匹配的schema和路径
	//ws.Path("/user").Consumes("*/*").Produces("*/*")

	//设置不同method对应的方法，参数以及参数描述和类型
	//参数:分为路径上的参数,query层面的参数,Header中的参数
	//ws.Route(ws.GET("/{id}").
	//	To(u.result).
	//	Doc("方法描述：获取用户").
	//	Param(ws.PathParameter("id", "参数描述:用户ID").DataType("string")).
	//	Param(ws.QueryParameter("name", "用户名称").DataType("string")).
	//	Param(ws.HeaderParameter("token", "访问令牌").DataType("string")).
	//	Do(returns200, returns500))
	//ws.Route(ws.POST("").To(u.result))
	//ws.Route(ws.PUT("/{id}").To(u.result))
	//ws.Route(ws.DELETE("/{id}").To(u.result))
	ws.Route(ws.GET("/123").To(u.FindUser))
	ws.Route(ws.GET("/456").To(u.result))
	container.Add(ws)
}
func (u UserResource) FindUser(req *restful.Request, resp *restful.Response) {
	sqlStr := "select uu_id, create_time  from t_user"
	rows, err := tools.Db.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
	}
	defer rows.Close()
	var user []UserResource
	for rows.Next() {
		var u UserResource
		err := rows.Scan(&u.UuId, &u.CreateTime)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		user = append(user, u)
	}
	fmt.Println(user)
	resp.WriteAsJson(user)

}

func (UserResource) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "Address doc", //空表示结构本省的描述
		"country":  "Country doc",
		"postcode": "PostCode doc",
	}
}

func (u UserResource) result(request *restful.Request, response *restful.Response) {
	io.WriteString(response.ResponseWriter, "this would be a normal response")
}
func returns200(b *restful.RouteBuilder) {
	b.Returns(http.StatusOK, "OK", "success")
}

func returns500(b *restful.RouteBuilder) {
	b.Returns(http.StatusInternalServerError, "Bummer, something went wrong", nil)
}
