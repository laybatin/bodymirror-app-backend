package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/laybatin/bodymirror-app-backend/internal/config"
	"net/http"
)

//ServerVersion 은 서버의 버전을 표시합니다.
type ServerVersion struct {
	Version string `json:"version"`
}

func VersionHandler(c *gin.Context) {
	version := ServerVersion{Version: config.Version}
	c.JSON(http.StatusOK, version)
}
