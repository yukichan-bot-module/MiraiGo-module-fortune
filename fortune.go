package fortune

import (
	"fmt"
	"hash/fnv"
	"strings"
	"sync"
	"time"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
)

var instance *fortune
var logger = utils.GetModuleLogger("com.aimerneige.fortune")

type fortune struct {
}

func init() {
	instance = &fortune{}
	bot.RegisterModule(instance)
}

func (f *fortune) MiraiGoModule() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "com.aimerneige.fortune",
		Instance: instance,
	}
}

// Init 初始化过程
// 在此处可以进行 Module 的初始化配置
// 如配置读取
func (f *fortune) Init() {
}

// PostInit 第二次初始化
// 再次过程中可以进行跨 Module 的动作
// 如通用数据库等等
func (f *fortune) PostInit() {
}

// Serve 注册服务函数部分
func (f *fortune) Serve(b *bot.Bot) {
	b.OnGroupMessage(func(c *client.QQClient, msg *message.GroupMessage) {
		msgString := msg.ToString()
		runes := []rune(msgString)
		if len(runes) <= 3 {
			return
		}
		if string(runes[:3]) == "求签 " {
			things := string(runes[3:])
			things = strings.TrimSpace(things)
			if things == "" {
				return
			}
			fortuneResult := drawAFortuneStick(things, msg.Sender.Uin)
			replyString := fmt.Sprintf("所求事项\"%s\"的求签结果为: %s", things, fortuneResult)
			replyMessage := message.NewSendingMessage().Append(message.NewText(replyString))
			b.SendGroupMessage(msg.GroupCode, replyMessage)
		}
	})
}

// Start 此函数会新开携程进行调用
// ```go
//
//	go exampleModule.Start()
//
// ```
// 可以利用此部分进行后台操作
// 如 http 服务器等等
func (f *fortune) Start(b *bot.Bot) {
}

// Stop 结束部分
// 一般调用此函数时，程序接收到 os.Interrupt 信号
// 即将退出
// 在此处应该释放相应的资源或者对状态进行保存
func (f *fortune) Stop(b *bot.Bot, wg *sync.WaitGroup) {
	// 别忘了解锁
	defer wg.Done()
}

func drawAFortuneStick(things string, uin int64) string {
	unixTime := uint32(time.Now().Unix() / 10000)
	thingsHash := stringHash(things)
	uinHash := stringHash(fmt.Sprint(uin))
	return getFortuneResult(unixTime + thingsHash + uinHash)
}

func getFortuneResult(hash uint32) string {
	var result string = "大凶"
	switch key := hash % 100; {
	case key < 2:
		result = "上吉" // 2
	case key < 10:
		result = "大吉" // 8
	case key < 38:
		result = "上上" // 28
	case key < 42:
		result = "上中" // 4
	case key < 45:
		result = "上平" // 3
	case key < 46:
		result = "上" // 1
	case key < 49:
		result = "中吉" // 3
	case key < 51:
		result = "中上" // 2
	case key < 57:
		result = "中中" // 6
	case key < 66:
		result = "中平" // 9
	case key < 71:
		result = "中" // 5
	case key < 72:
		result = "平中" // 1
	case key < 73:
		result = "平平" // 1
	case key < 74:
		result = "平" // 1
	case key < 99:
		result = "下" // 25
	case key < 100:
		result = "下下" // 1
	default:
		result = "大凶"
	}
	return result
}

func stringHash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
