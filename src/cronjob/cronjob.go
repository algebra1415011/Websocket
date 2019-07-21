package cronjob

import (
	"github.com/robfig/cron"
	"server"
	"fmt"
)

func StartcronJob(server *server.Server,msg string){
	cronJob := cron.New()
	fmt.Println("cron job is running ")
	cronJob.AddFunc("*/5 * * * * *", func() {  server.PingConnectedClient(msg) })	
	cronJob.Start()
}