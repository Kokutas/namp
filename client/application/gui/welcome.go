package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/cmd/fyne_demo/data"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"net/url"
)

func welcome(application fyne.App) fyne.CanvasObject  {
	// 创建一个logo的图片资源
	logo:=canvas.NewImageFromResource(data.FyneScene)
	// 如果当前设备是移动端
	if fyne.CurrentDevice().IsMobile() {
		// 设置logo的图片资源的最小尺寸
		logo.SetMinSize(fyne.NewSize(171,125))
	} else {
		// 设置logo的图片资源的最小尺寸
		logo.SetMinSize(fyne.NewSize(171,125))
	}
	// 返回一个垂直方向布局的盒子
	return widget.NewVBox(
		// NewSpacer返回一个可以填充垂直和水平空间的spacer对象，主要用于盒子布局
		layout.NewSpacer(),
		//设置字体加粗、水平居中的label
		widget.NewLabelWithStyle("欢迎使用go开发的nmap工具",fyne.TextAlignCenter,fyne.TextStyle{
			Bold:      true,
			Italic:    true,
			Monospace: false,
		}),
		// 设置logo图像
		widget.NewHBox(layout.NewSpacer(),logo,layout.NewSpacer()),
		// 设置一个水平方向布局排列的盒子
		widget.NewHBox(
			layout.NewSpacer(),
			widget.NewHyperlinkWithStyle("官网",parseUrl("https://www.taobao.com"),fyne.TextAlignCenter,fyne.TextStyle{
				Bold:      false,
				Italic:    true,
				Monospace: false,
			}),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)
}
// url解析函数
func parseUrl(urlString string)*url.URL  {
	link,err:=url.Parse(urlString)
	if err != nil {
		fyne.LogError("can not parse URL:",err)
	}
	return link
}
