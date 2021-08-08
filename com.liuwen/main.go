package main

import (
	"bookstores2/src/controller"
	"net/http"
)

//主方法初始化，处理静态资源
func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))

	http.Handle("/pages/",
		http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages/"))))

	//处理请求
	//用户相关
	http.HandleFunc("/main", controller.IndexHandler) // 首页  查找所有图书
	http.HandleFunc("/toLogin", controller.ToLogin)
	//登录逻辑：
	//step1：验证用户名和密码（密码加密）
	//step2：验证成功,则创建session,保存到服务端,并设置cookie,保存到客户端;
	//step3：验证失败则重新登录。
	http.HandleFunc("/login", controller.Login)
	//注销逻辑：
	//step1：判断客户端是否有cookie
	//step2：有则根据cookie删除服务端的session，再将cookie销毁，返回登录页或首页
	//step3：无则说明未登录，或登录时间已超时
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/register", controller.Register)
	//注册的时候需要验证用户名称是否重复
	http.HandleFunc("/FindUserByName", controller.FindUserByName)

	//图书相关
	//测试，获取所有书籍（前后端分离）  http://localhost:8080/getBooks
	http.HandleFunc("/getBooks", controller.GetBooks)
	//图书相关
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	http.HandleFunc("/toUpdateBookPage", controller.ToUpdateBookPage)
	http.HandleFunc("/deleteBook", controller.DeleteBookById)
	http.HandleFunc("/updateOraddBook", controller.AddOrUpdateBook)
	http.HandleFunc("/queryPrice", controller.QueryPrice)

	//购物车相关
	http.HandleFunc("/AddBook2Cart", controller.AddBook2Cart)
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)
	http.HandleFunc("/deleteCart", controller.DeleteCart)
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)
	http.HandleFunc("/updateCartItem", controller.UpdateCartItem)

	//订单相关（结账，发货，收货）
	//结账逻辑
	//step1：验证用户是否登录，已登录则根据客户端的cookie到服务端获取用户的session信息
	//step2：根据用户session信息（用户uuid）获取购物车
	//step3：创建订单
	//step4：创建订单项（订单详情，辅助结账功能）
	//step5：更新库存
	//step6：清空购物车项
	//step7：清空购物车
	//step8：返回orderUuid
	http.HandleFunc("/checkout", controller.Checkout)
	http.HandleFunc("/getMyOrder", controller.GetMyOrder)
	http.HandleFunc("/getOrders", controller.GetAllOrder)
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfo)
	http.HandleFunc("/sendOrder", controller.SendOrder)
	http.HandleFunc("/takeOrder", controller.TakeOrder)

	////获取SSL 证书和 RSA 私钥
	//utils.GetTLS("utils/pem/cert.pem","utils/pem/key.pem")

	//设置服务器路径,使用默认多路服务器
	http.ListenAndServe("127.0.0.1:8080", nil)

}
