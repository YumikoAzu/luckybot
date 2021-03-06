package admin

import (
	"sync"

	"github.com/gorilla/mux"
	"github.com/zhangpanyi/luckybot/app/admin/handlers"
)

var once sync.Once

// 初始路由
func InitRoute(router *mux.Router) {
	once.Do(func() {
		router.HandleFunc("/admin/backup", handlers.Backup)
		router.HandleFunc("/admin/deposit", handlers.Deposit)
		router.HandleFunc("/admin/balance", handlers.GetBalance)
		router.HandleFunc("/admin/broadcast", handlers.Broadcast)
		router.HandleFunc("/admin/getactions", handlers.GetActions)
		router.HandleFunc("/admin/subscribers", handlers.Subscribers)
	})
}
