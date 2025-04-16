package bootstrap

import (
	"fmt"
	"github.com/zeromicro/go-zero/rest"
	"strings"
)

func PrintRoutes(server *rest.Server) {
	fmt.Println("📌 Registered Routes:")

	routes := server.Routes()
	if len(routes) == 0 {
		fmt.Println("No routes found.")
		return
	}

	// 先算出最大宽度
	maxMethodLen := len("METHOD")
	maxPathLen := len("PATH")

	for _, route := range routes {
		if len(route.Method) > maxMethodLen {
			maxMethodLen = len(route.Method)
		}
		if len(route.Path) > maxPathLen {
			maxPathLen = len(route.Path)
		}
	}

	// 分隔线
	sep := fmt.Sprintf("+-%s-+-%s-+",
		strings.Repeat("-", maxMethodLen),
		strings.Repeat("-", maxPathLen),
	)

	// 打印表头
	fmt.Println(sep)
	fmt.Printf("| %-*s | %-*s |\n", maxMethodLen, "METHOD", maxPathLen, "PATH")
	fmt.Println(sep)

	// 打印每一行
	for _, route := range routes {
		fmt.Printf("| %-*s | %-*s |\n", maxMethodLen, route.Method, maxPathLen, route.Path)
	}
	fmt.Println(sep)
}
