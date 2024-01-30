package system

import (
	"net/http"

	"sophliteos/global"
	"sophliteos/pkg/handle"

	"github.com/gin-gonic/gin"
)

type VersionApi struct{}

func (b *VersionApi) Version(c *gin.Context) {
	c.JSON(http.StatusOK, handle.Success(global.Version))
}
