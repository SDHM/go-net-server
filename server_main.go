package main
import "github.com/SDHM/server"

func main() {
	upgradeServer := server.NewUpGradeServer()
	go upgradeServer.Run(":21231");
}
