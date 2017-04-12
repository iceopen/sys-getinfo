package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
)

func main() {
	gin.SetMode("release")
	r := gin.Default()
	//SWAP大小
	r.GET("/swap/info", func(c *gin.Context) {
		swap, _ := mem.SwapMemory()
		c.JSON(200, swap)
	})
	//虚拟内存
	r.GET("/vmem/info", func(c *gin.Context) {
		vmem, _ := mem.VirtualMemory()
		c.JSON(200, vmem)
	})

	//CPU 信息详细
	r.GET("/cpu/info", func(c *gin.Context) {
		cpuinfo, _ := cpu.Info()
		c.JSON(200, cpuinfo)
	})

	//CPU 运行时情况
	r.GET("/cpu/time", func(c *gin.Context) {
		cpuTime, _ := cpu.Times(true)
		c.JSON(200, cpuTime)
	})
	////硬盘列表 信息详细
	r.GET("/disk/list",func(c *gin.Context) {
		diskPart, _ := disk.Partitions(true)
		c.JSON(200, diskPart)
	})
	//硬盘使用 信息详细
	r.GET("/disk/usage", func(c *gin.Context) {
		diskInfo, _ := disk.Usage("/")
		c.JSON(200, diskInfo)
	})
	//制定目录使用 信息详细
	r.POST("/disk/path", func(c *gin.Context){
		path := c.PostForm("path")
		diskInfo, _ := disk.Usage(path)
		c.JSON(200, diskInfo)
	})
	r.Run(":8848") // listen and serve on 0.0.0.0:8080
}