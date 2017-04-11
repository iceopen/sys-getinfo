package main

import (
	"github.com/go-martini/martini"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/martini-contrib/render"
	"net/http"
)

func init() {
	martini.Env = martini.Prod
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	//SWAP大小
	m.Get("/swap/info", func(r render.Render) {
		swap, _ := mem.SwapMemory()
		r.JSON(200, swap)
	})
	//虚拟内存
	m.Get("/vmem/info", func(r render.Render) {
		vmem, _ := mem.VirtualMemory()
		r.JSON(200, vmem)
	})
	//CPU 运行时情况
	m.Get("/cpu/time", func(r render.Render){
		cpuTime, _ := cpu.Times(true)
		r.JSON(200, cpuTime)
	})
	//CPU 信息详细
	m.Get("/cpu/info",func(r render.Render){
		cpuinfo, _ := cpu.Info()
		r.JSON(200, cpuinfo)
	})
	////硬盘列表 信息详细
	m.Get("/disk/list", func(r render.Render) {
		diskPart, _ := disk.Partitions(true)
		r.JSON(200, diskPart)
	})
	//硬盘使用 信息详细
	m.Get("/disk/usage", func(r render.Render){
		diskInfo, _ := disk.Usage("/")
		r.JSON(200, diskInfo)
	})
	//制定目录使用 信息详细
	m.Post("/disk/path", func(r render.Render, req *http.Request){
		path := req.FormValue("path")
		diskInfo, _ := disk.Usage(path)
		r.JSON(200, diskInfo)
	})
	//m.Run()
	m.RunOnAddr(":8848")
}