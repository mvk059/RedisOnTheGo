package replication

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/settings"
	"io"
	"net"
)

func CreateConnection(settings settings.ServerSettings) io.ReadWriter {
	ip := fmt.Sprintf("%s:%d", settings.MasterHost, settings.MasterPort)
	dial, err := net.Dial("tcp4", ip)
	if err != nil {
		fmt.Println("\033[31mCouldn't connect to : \033[0m", err.Error())
		panic("Couldn't connect to " + ip)
	}
	return dial
}
