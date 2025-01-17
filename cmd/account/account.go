package main

import (
	"net/http"

	"github.com/Scfy-Code/IM/app/account/router"
	"github.com/Scfy-Code/IM/sys"
)

var (
	app *http.Server
)

func init() {
	app = &http.Server{
		Addr:    ":8088",
		Handler: sys.UniversalHandler,
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../web/static"))))
	http.Handle("/login.scfy", router.NewloginTemplate())
	http.Handle("/login.action", router.NewLoginRouter())
	http.Handle("/regist.scfy", router.NewRegistTemplate())
	http.Handle("/regist.action", router.NewRegistRouter())
}
func main() {
	//sys.UniversalHandler.Handle("", nil)
	app.ListenAndServe()
}
