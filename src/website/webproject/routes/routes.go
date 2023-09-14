package routes

import (
	"module/src/website/webproject/controller"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/product", controller.NewProduct)
	http.HandleFunc("/insert", controller.InsertProduct)
	http.HandleFunc("/delete", controller.DeleteProduct)
	http.HandleFunc("/edit", controller.EditProduct)
	http.HandleFunc("/update", controller.UpdateProduct)
}
