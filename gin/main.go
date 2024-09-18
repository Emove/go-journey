package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	engine := gin.Default()
	engine.GET("/download", download)
	_ = engine.Run(":8080")
}

func download(ctx *gin.Context) {

	filePath := "D:\\download\\asm.rar"

	//打开文件
	_, errByOpenFile := os.Open(filePath)
	//非空处理
	if errByOpenFile != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "失败",
			"error":   "资源不存在",
		})
		//ctx.Redirect(http.StatusFound, "/404")
		return
	}
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename=asm.rar")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.File(filePath)
	return
}
