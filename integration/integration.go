package integration

import (
	"log"
	"net/http"
	"net/http/pprof"
	"os"

	"github.com/felixge/fgprof"
	"github.com/goccy/go-json"
	"github.com/gorilla/mux"
	"github.com/kaz/pprotein/internal/git"
	"github.com/kaz/pprotein/internal/tail"
)

var (
	httplogPath       = getEnvOrDefault("PPROTEIN_HTTPLOG", "/var/log/nginx/access.log")
	slowlogPath       = getEnvOrDefault("PPROTEIN_SLOWLOG", "/var/log/mysql/mysql-slow.log")
	gitRepositoryPath = getEnvOrDefault("PPROTEIN_GIT_REPOSITORY", ".")
)

func NewDebugHandler(gitDir string) http.Handler {
	r := mux.NewRouter()
	RegisterDebugHandlersWithGitDir(r, gitDir)
	return r
}

// RegisterDebugHandlers は下位互換性のために残します
func RegisterDebugHandlers(r *mux.Router) {
	RegisterDebugHandlersWithGitDir(r, "")
}

func RegisterDebugHandlersWithGitDir(r *mux.Router, gitDir string) {
	// gitDirが指定されている場合は、それを使用。そうでなければデフォルト値を使用
	repoPath := gitDir
	if repoPath == "" {
		repoPath = gitRepositoryPath
	}

	r.Use(gitRepositoryMiddlewareWithPath(repoPath))

	r.Handle("/debug/log/httplog", tail.NewTailHandler(httplogPath))
	r.Handle("/debug/log/slowlog", tail.NewTailHandler(slowlogPath))

	r.Handle("/debug/fgprof", fgprof.Handler())

	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)
	r.HandleFunc("/debug/pprof/{h:.*}", pprof.Index)
}

func gitRepositoryMiddleware(next http.Handler) http.Handler {
	return gitRepositoryMiddlewareWithPath(gitRepositoryPath)(next)
}

func gitRepositoryMiddlewareWithPath(repoPath string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			defer next.ServeHTTP(rw, r)

			headInfo, err := git.GetInfo(repoPath)
			if err != nil {
				log.Printf("failed to get git info: %v", err)
				return
			}

			data, err := json.Marshal(headInfo)
			if err != nil {
				log.Printf("failed to marshal git info: %v", err)
				return
			}

			rw.Header().Set("X-Git-Repository", string(data))
		})
	}
}

func getEnvOrDefault(key string, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
