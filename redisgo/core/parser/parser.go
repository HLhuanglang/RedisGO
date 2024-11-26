package parser

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
)

/*
+OK\r\n
-Error message\r\n
:1000\r\n
$5\r\nhello\r\n
*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n
*/

func ParserStream(ctx context.Context, conn net.Conn) {
	rd := bufio.NewReader(conn)
	for {
		data, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		fmt.Printf("data: %+v", data)
	}
}
