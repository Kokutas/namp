package gui

// 客户端GUI页面展示基本

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/flopp/go-findfont"
	"github.com/go-vgo/robotgo"
	"math"
	"os"
	"runtime"
	"strings"
	"time"
)

// 当前的页面展示
const preferenceCurrentTab = "currentTab"
// 初始化操作：初始化fyne对中文的支持
func init() {
	// 1.获取字体库的文件全路径名
	fontPaths := findfont.List()
	// 2.遍历字体库全路径名
	for _, fontPath := range fontPaths {
		// 3.根据操作系统匹配中文字体库并设置环境变量
		switch strings.ToLower(runtime.GOOS) {
		case "windows":
			switch {
			case strings.Contains(strings.ToLower(fontPath), "simkai"):
				os.Setenv("FYNE_FONT", fontPath)
				break
			case strings.Contains(strings.ToLower(fontPath), "simhei"):
				os.Setenv("FYNE_FONT", fontPath)
				break
			}
		case "darwin":
			switch {
			case strings.Contains(strings.ToLower(fontPath), "stkaiti"):
				os.Setenv("FYNE_FONT", fontPath)
				break
			case strings.Contains(strings.ToLower(fontPath), "stheiti"):
				os.Setenv("FYNE_FONT", fontPath)
				break
			}
		case "linux":
		default:
			switch {
			case strings.Contains(strings.ToLower(fontPath), "kai"):
				os.Setenv("FYNE_FONT", fontPath)
				break
			case strings.Contains(strings.ToLower(fontPath), "hei"):
				os.Setenv("FYNE_FONT", fontPath)
				break
			}
		}
	}
}

func Start() {
	// 字体环境变量设置关闭
	defer os.Unsetenv("FYNE_FONT")

	// 创建一个ID为当前时间的应用对象
	application := app.NewWithID(time.Local.String())
	// 设置应用的icon为fyne中的logo
	application.SetIcon(theme.FyneLogo())
	// 设置默认的主题
	application.Settings().SetTheme(theme.DarkTheme())
	// 根据应用创建一个标题为：NAMP 的GUI窗口对象
	window := application.NewWindow("NMAP")
	// 设置window窗口的尺寸
	// -- 通过robotgo获取屏幕的信息
	sx, sy := robotgo.GetScreenSize()
	// 如果当前设备是移动端
	if fyne.CurrentDevice().IsMobile() {
		// 设置窗口的尺寸为最大
		window.Resize(fyne.NewSize(sx, sy))
	} else {
		// 设置窗口的尺寸为50%，50%
		window.Resize(fyne.NewSize(int(math.Ceil(0.5*float64(sy))), int(math.Ceil(0.5*float64(sy)))))
	}
	// 设置窗口居中
	window.CenterOnScreen()
	// 创建tab容器
	tabs := widget.NewTabContainer(
		// 新建一个欢迎页tab
		widget.NewTabItemWithIcon("欢迎", theme.HomeIcon(), welcome(application)),
		// 新建一个帮助页的tab
		//widget.NewTabItemWithIcon("帮助", theme.HelpIcon(), help(window)),
		// 查询tab
		widget.NewTabItemWithIcon("查询", theme.SearchIcon(), search(window)),

	)
	// 设置容器的位置
	tabs.SetTabLocation(widget.TabLocationLeading)
	// 记录操作的记录，方便后续打开的是上一次的记录
	tabs.SelectTabIndex(application.Preferences().Int(preferenceCurrentTab))
	// 设置窗口对象的内容为tabs
	window.SetContent(tabs)
	// 启动并展示窗口
	window.ShowAndRun()

	// 展示上一次操作的页面
	application.Preferences().SetInt(preferenceCurrentTab, tabs.CurrentTabIndex())

}
