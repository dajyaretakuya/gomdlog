package route

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dajyaretakuya/gomdlog/action"
)

func dispatchHandler(w http.ResponseWriter, r *http.Request) {
	routeMap := make(map[string]action.Root)
	routeMap["/"] = action.Home{w, r}

	fullPath := r.URL.Path
	fullPathLength := strings.Count(fullPath, "")
	if strings.Compare(string(fullPath[fullPathLength-2]), "/") == 0 {
		fullPath = fullPath[0 : fullPathLength-2]
	}
	lastSlashPosition := strings.LastIndex(fullPath, "/")
	firstLevelPath := ""
	if lastSlashPosition <= 0 {
		firstLevelPath = fullPath[0:]
	} else {
		firstLevelPath = fullPath[0:lastSlashPosition]
	}
	//handler := routeMap[path]
	fmt.Fprintf(w, "fullPath is: %s, level path: %s, index of / is: %d", fullPath, firstLevelPath, strings.LastIndex(fullPath, "/"))
}

// Dispatch 分发请求
func Dispatch() {
	http.HandleFunc("/", dispatchHandler)
}
