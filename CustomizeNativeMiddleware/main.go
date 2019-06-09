/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-9
* Time: 下午10:47
* */
// 自定义原生中间件
package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

func MiddleAuth(router httprouter.Handle) httprouter.Handle {
	return httprouter.Handle(
		func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
			token := request.Header.Get("token")
			if strings.EqualFold(token,"tokenId") {
				router(writer,request,params)
				return
			}
			writer.WriteHeader(401)
			writer.Write([]byte("401"))
		})
}

func RegisRouter()  *httprouter.Router{
	router := httprouter.New()

	router.GET("/home",MiddleAuth(home))

	
	return router
}

func home(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {

}

func main() {
	router := RegisRouter()

	http.ListenAndServe(":8085",router)
}

