package ctrl_auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type AuthCtrl struct {
}

func NewAuthCtrl() *AuthCtrl {
	return &AuthCtrl{}
}

func (ctrl AuthCtrl) SignIn(ctx *gin.Context) {
	fmt.Println("Access SignIn api")
	ctx.JSON(200, "access api success")
}
