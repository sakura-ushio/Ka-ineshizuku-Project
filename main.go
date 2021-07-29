package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"

	// 注：以下插件均可通过前面加 // 注释，注释后停用并不加载插件
	// 下列插件可与 wdvxdr1123/ZeroBot v1.1.2 以上配合单独使用
	// 词库类
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/atri" // ATRI词库
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/chat" // 基础词库

	// 实用类
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/github"  // 搜索GitHub仓库
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/manager" // 群管
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/runcode" // 在线运行代码

	// 娱乐类
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/ai_false" // 服务器监控
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/minecraft"
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/music"   // 点歌
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/shindan" // 测定

	// b站相关
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/bilibili" // 查询b站用户信息
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/diana"    // 嘉心糖发病

	// 二次元图片
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/image_finder" // 关键字搜图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/lolicon"      // lolicon 随机图片
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/rand_image"   // 随机图片与点评
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/saucenao"     // 以图搜图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/setutime"     // 来份涩图

	// 迫害等
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/gag"   // 禁言套餐
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/snare" // 随机陷害

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/jieba" // 分词
	// 以下为内置依赖，勿动
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
	"github.com/wdvxdr1123/ZeroBot/message"

	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

var content = []string{
	"* OneBot + ZeroBot + Golang ",
	"* Version 1.0.4 - 2021-07-14 14:09:58.581489207 +0800 CST",
	"* Copyright © 2020 - 2021  Kanri, DawnNights, Fumiama, Suika",
	"* Project: https://github.com/FloatTech/ZeroBot-Plugin",
}

func init() {
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[zero][%time%][%lvl%]: %msg% \n",
	})
	log.SetLevel(log.DebugLevel)
	go http.ListenAndServe("0.0.0.0:6060", nil)
}

func main() {
	fmt.Print(
		"====================[ZeroBot-Plugin]====================",
		"\n", strings.Join(content, "\n"), "\n",
		"========================================================",
	) // 启动打印
	zero.Run(zero.Config{
		NickName:      []string{"椛椛", "ATRI", "atri", "亚托莉", "アトリ"},
		CommandPrefix: "/",

		// SuperUsers 某些功能需要主人权限，可通过以下两种方式修改
		// []string{}：通过代码写死的方式添加主人账号
		// os.Args[1:]：通过命令行参数的方式添加主人账号
		SuperUsers: append([]string{"38263547"}, os.Args[1:]...),

		Driver: []zero.Driver{
			&driver.WSClient{
				// OneBot 正向WS 默认使用 6700 端口
				Url:         "ws://127.0.0.1:6700",
				AccessToken: "",
			},
		},
	})

	// 帮助
	zero.OnFullMatchGroup([]string{"help", "/help", ".help", "菜单", "帮助"}, zero.OnlyToMe).SetBlock(true).SetPriority(999).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(
				message.Text(strings.Join(content, "\n")),
			)
		})
	select {}
}
