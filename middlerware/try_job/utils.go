package try_job

import (
	"fmt"
	"net"
)

func JobKeyGenerator(conn net.Conn, content string) string {
	return fmt.Sprintf("%s:%s", conn.RemoteAddr().String(), content)
}
func ContentMaker(content, cod string) string {
	return fmt.Sprintf("%s%s", content, cod)
}
