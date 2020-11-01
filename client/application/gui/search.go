package gui

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"image/color"
	"nmap/client/application/service/network"
	"sync"
	"time"
)

//查询相关

func search(window fyne.Window) fyne.CanvasObject {
	return widget.NewTabContainer(
		widget.NewTabItemWithIcon("网  络", theme.SearchReplaceIcon(), queryNetwork(window)),
		//widget.NewTabItemWithIcon("网络", theme.SearchReplaceIcon(), networkInfo(window)),

	)
}

var adaptors []*network.Adaptor

// 网络 -> 局域网信息查询 [1.网卡信息查询 2.本机ip查询]
func networkLAN(window fyne.Window) *widget.AccordionItem {
	// 创建IP信息查询按钮
	ipBtn := widget.NewButtonWithIcon(" IP ", theme.InfoIcon(), func() {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			apts, err := network.Adaptors()
			if err != nil {
				fyne.CurrentApp().SendNotification(&fyne.Notification{
					Title:   "本机局域网IP查询[错误]",
					Content: err.Error(),
				})
				dialog.ShowError(err, window)
			}
			fmt.Println("ap", apts)
			adaptors = apts
		}()
		progress("本机局域网IP查询", "正在查询中......", window, wg)
		wg.Wait()
	})
	// 创建网卡信息查询按钮
	adaptorBtn := widget.NewButtonWithIcon("网卡", theme.InfoIcon(), func() {
		var wg sync.WaitGroup
		progress("本机可用网卡/网络适配器查询", "正在查询中......", window, wg)
		wg.Wait()
	})
	// 创建按钮组: 布局容器对象
	child := fyne.NewContainerWithLayout(
		// 设置为垂直布局
		layout.NewVBoxLayout(),
		// 按钮
		ipBtn,
		adaptorBtn,
	)
	// 创建局域网查询按钮导航项
	return &widget.AccordionItem{
		Title: "局域网",
		Detail: fyne.NewContainerWithLayout(
			// 设置布局为靠右
			layout.NewBorderLayout(nil, nil, nil, child),
			child,
		),
		Open: false,
	}
}

// 网络 -> 互联网信息查询 [1.公网IP信息查询 2.网络带宽查询]
func networkWAN(window fyne.Window) *widget.AccordionItem {
	// 创建互联网IP信息查询按钮
	ipBtn := widget.NewButtonWithIcon(" IP ", theme.InfoIcon(), func() {
		var wg sync.WaitGroup
		progress("互联网IP查询", "正在查询中......", window, wg)
		wg.Wait()
	})
	// 创建互联网带宽信息查询按钮
	bandwidthBtn := widget.NewButtonWithIcon("带宽", theme.InfoIcon(), func() {
		var wg sync.WaitGroup
		progress("互联网带宽查询", "正在查询中......", window, wg)
		wg.Wait()
	})
	// 创建按钮组: 布局容器对象
	child := fyne.NewContainerWithLayout(
		// 设置为垂直布局
		layout.NewVBoxLayout(),
		// 按钮
		ipBtn,
		bandwidthBtn,
	)
	// 创建互联网查询按钮导航项
	return &widget.AccordionItem{
		Title: "互联网",
		Detail: fyne.NewContainerWithLayout(
			// 设置布局为靠右
			layout.NewBorderLayout(nil, nil, nil, child),
			child,
		),
		Open: false,
	}
}

// 网络 -> 局域网信息查询结果 [1.网卡信息查询结果 2.本机ip查询结果]
func networkLANResult() *widget.Group {
	if adaptors == nil {
		return nil
	}
	groups := widget.NewGroup("局域网")
	for _, adaptor := range adaptors {
		form := widget.NewForm(
			&widget.FormItem{
				Text:   "网卡序号:",
				Widget: widget.NewLabel(fmt.Sprintf("%v", adaptor.Index)),
			},
			&widget.FormItem{
				Text:   "网卡名称",
				Widget: widget.NewLabel(fmt.Sprintf("%v", adaptor.Name)),
			},
			&widget.FormItem{
				Text:   "最大传输单元:",
				Widget: widget.NewLabel(fmt.Sprintf("%v", adaptor.Mtu)),
			},
			&widget.FormItem{
				Text:   "MAC地址:",
				Widget: widget.NewLabel(fmt.Sprintf("%v", adaptor.Mac)),
			},
			&widget.FormItem{
				Text:   "状态标识:",
				Widget: widget.NewLabel(fmt.Sprintf("%v", adaptor.Flags)),
			},
			&widget.FormItem{
				Text:   "IPv4:",
				Widget: widget.NewLabel(fmt.Sprintf("%v", adaptor.IPv4)),
			},
			&widget.FormItem{
				Text:   "IPv4子网掩码个数:",
				Widget: widget.NewLabel(fmt.Sprintf("%v", adaptor.IPv4MaskCount)),
			},
			&widget.FormItem{
				Text:   "IPv6:",
				Widget: widget.NewLabel(fmt.Sprintf("%v", adaptor.IPv6)),
			},
			&widget.FormItem{
				Text:   "IPv6子网掩码个数:",
				Widget: widget.NewLabel(fmt.Sprintf("%v", adaptor.IPv6MaskCount)),
			},
		)
		groups.Append(form)
	}
	return groups
}

//网络 -> 互联网信息查询结果 [1.公网IP信息查询结果 2.网络带宽查询结果]
func networkWANResult() *widget.Group {
	//// 核心功能还未编写，先用假数据填充
	//ip := widget.NewLabel("114.114.114.114")
	//position := widget.NewLabel("中国江苏省南京市")
	//isp := widget.NewLabel("南京信风网络科技有限公司GreatbitDNS服务器")
	//return widget.NewGroup("互联网",
	//	widget.NewForm(
	//		&widget.FormItem{Text: "IP地址:", Widget: ip},
	//		&widget.FormItem{Text: "所属地:", Widget: position},
	//		&widget.FormItem{Text: "供应商:", Widget: isp},
	//
	//	),
	//)
	return nil
}

// 网络信息查询
/*func queryNetwork(window fyne.Window) fyne.CanvasObject {
	left := widget.NewAccordionContainer(
		networkLAN(window),
		networkWAN(window),
	)
	right := widget.NewVBox()
	if networkLANResult() != nil {
		right.Append(networkLANResult())
	}
	if networkWANResult() != nil {
		right.Append(networkWANResult())
	}
	if networkLANResult() == nil && networkWANResult() == nil {
		right.Append(widget.NewLabelWithStyle("请点击左侧的选项按钮进行查询操作", fyne.TextAlignCenter, fyne.TextStyle{
			Bold:      true,
			Italic:    true,
			Monospace: false,
		}))
	}
	top := fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(),
		left,
		right,
	)
	bottom := widget.NewLabelWithStyle("底部版权信息", fyne.TextAlignCenter, fyne.TextStyle{
		Bold:      false,
		Italic:    true,
		Monospace: false,
	})

	return fyne.NewContainerWithLayout(
		layout.NewBorderLayout(top, bottom, nil, nil),
		top,
		bottom,
	)
}
*/

// 进度条
func progress(title string, message string, window fyne.Window, wg sync.WaitGroup) {
	prog := dialog.NewProgress(title, message, window)
	wg.Add(1)
	go func() {
		wg.Done()
		num := 0.0
		for num < 1.0 {
			time.Sleep(50 * time.Millisecond)
			prog.SetValue(num)
			num += 0.01
		}
		prog.SetValue(1)
		prog.Hide()
	}()
	prog.Show()
}

func queryNetwork2(window fyne.Window) fyne.CanvasObject {

	return widget.NewVBox(
		//layout.NewSpacer(),
		fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithRows(2),
			//layout.NewSpacer(),
			widget.NewHBox(
				widget.NewGroup("局域网",
					widget.NewButtonWithIcon("IP", theme.InfoIcon(), func() {
						fmt.Println("我是")
					}),
					widget.NewButtonWithIcon("网卡", theme.InfoIcon(), func() {
						fmt.Println("您好")
					}),
				),
				widget.NewScrollContainer(
					canvas.NewText("你好", color.RGBA{
						R: 60,
						G: 200,
						B: 0,
						A: 0,
					}),
				),
			),
			widget.NewHBox(
				widget.NewGroup("互联网",
					widget.NewButtonWithIcon("IP", theme.InfoIcon(), func() {
						fmt.Println("www")
					}),
					widget.NewButtonWithIcon("带宽", theme.InfoIcon(), func() {
						fmt.Println("您好")
					}),
				),
				widget.NewLabel("显示的内容"),
			),

		),
		layout.NewSpacer(),
		widget.NewGroup("版权信息",
			widget.NewLabelWithStyle("底部版权信息", fyne.TextAlignCenter, fyne.TextStyle{
				Bold:      false,
				Italic:    true,
				Monospace: false,
			}),
		),


	)

}

// 网络查询
func queryNetwork(window fyne.Window) fyne.CanvasObject {
	//fontColor := color.RGBA{
	//	R: 60,
	//	G: 200,
	//	B: 0,
	//	A: 0,
	//}

	return widget.NewVBox(
		// 网络查询相关布局
		fyne.NewContainerWithLayout(
			// 划分成两列
			layout.NewAdaptiveGridLayout(2),
			// 第1列局域网查询按钮
			widget.NewGroup("局域网",
				widget.NewButtonWithIcon("IP", theme.InfoIcon(), func() {
					fmt.Println("局域网 IP 查询按钮按下")
				}),
				widget.NewButtonWithIcon("网卡", theme.InfoIcon(), func() {
					fmt.Println("局域网 网卡 查询按钮按下")
				}),
			),
			// 第2列互联网查询按钮
			widget.NewGroup("互联网",
				widget.NewButtonWithIcon("IP", theme.InfoIcon(), func() {
					//fmt.Println("互联网 IP 查询按钮按下")
					show("互联网 IP 查询按钮按下")
				}),
				widget.NewButtonWithIcon("带宽", theme.InfoIcon(), func() {
					fmt.Println("互联网 带宽 查询按钮按下")
				}),
				show("1"),
			),
		),
		//layout.NewSpacer(),
		// 第二行显示查询结果
		show("1"),


		layout.NewSpacer(),
		// 版权信息
		widget.NewGroup("版权信息",
			widget.NewLabelWithStyle("底部版权信息", fyne.TextAlignCenter, fyne.TextStyle{
				Bold:      false,
				Italic:    true,
				Monospace: false,
			}),
		),

	)
}
func show(result string) fyne.CanvasObject {
	content := widget.NewMultiLineEntry()
	content.Wrapping = fyne.TextWrapWord
	content.SetText(result)
	fmt.Println(result,"///////////////")
	content.Refresh()
	return content
	//canvas.NewText(fmt.Sprintf("%v : %v", "互联网 IP", "114.114.114.114"), fontColor)
}
//func queryNetwork(window fyne.Window) fyne.CanvasObject {
//	return widget.NewVBox(
//		// 网络查询相关布局
//		fyne.NewContainerWithLayout(
//			// 划分成两行
//			layout.NewGridLayoutWithRows(2),
//			// 第1行分成2列
//			lanLayout(window),
//			// 第2行分成2列
//			wanLayout(window),
//		),
//
//
//		layout.NewSpacer(),
//		// 版权信息
//		widget.NewGroup("版权信息",
//			widget.NewLabelWithStyle("底部版权信息", fyne.TextAlignCenter, fyne.TextStyle{
//				Bold:      false,
//				Italic:    true,
//				Monospace: false,
//			}),
//		),
//
//	)
//}

// 查询 -> 网络 -> 局域网查询布局
//func lanLayout(window fyne.Window) fyne.CanvasObject {
//	fontColor := color.RGBA{
//		R: 60,
//		G: 200,
//		B: 0,
//		A: 0,
//	}
//	return fyne.NewContainerWithLayout(
//		layout.NewAdaptiveGridLayout(2),
//		// 第1列显示按钮
//		query(fontColor),
//		// 第2列显示查询结果
//		showAdaptor("局域网", nil),
//	)
//}

// 查询 -> 网络 -> 局域网查询按钮
//func query(fontColor color.RGBA) fyne.CanvasObject {
//	query := widget.NewGroup("局域网查询",
//		widget.NewButtonWithIcon("IP", theme.InfoIcon(), func() {
//			fmt.Println("局域网  IP信息查询")
//		}),
//		widget.NewButtonWithIcon("网卡信息", theme.InfoIcon(), func() {
//			fmt.Println("局域网  网卡信息查询")
//			adaptors, err := network.Adaptors()
//			if err != nil {
//				log.Fatal(err)
//			}
//			showAdaptor("局域网-网卡信息", adaptors)
//
//		}),
//	)
//	return widget.NewScrollContainer(query)
//
//	/*	return widget.NewGroup("局域网",
//		widget.NewButtonWithIcon("IP", theme.InfoIcon(), func() {
//			fmt.Println("局域网  IP信息查询")
//			showQueryResult("局域网", map[string]string{"IP":"114.114.114.114"}, fontColor)
//		}),
//		widget.NewButtonWithIcon("网卡信息", theme.InfoIcon(), func() {
//			fmt.Println("局域网  网卡信息查询")
//		}),
//	)*/
//}

//var (
//	index         = widget.NewLabel("")
//	mtu           = widget.NewLabel("")
//	name          = widget.NewLabel("")
//	mac           = widget.NewLabel("")
//	flags         = widget.NewLabel("")
//	ipv4          = widget.NewLabel("")
//	ipv4MaskCount = widget.NewLabel("")
//	ipv6          = widget.NewLabel("")
//	ipv6MaskCount = widget.NewLabel("")
//)

// 查询 -> 网络 -> 网卡 -> 查询结果
//func showAdaptor(title string, adaptors []*network.Adaptor) fyne.CanvasObject {
//	if len(adaptors) == 0 {
//		content := widget.NewGroup(title)
//		return widget.NewScrollContainer(content)
//	}
//	screens := make([]fyne.CanvasObject, 0)
//	for _, adaptor := range adaptors {
//		screen := widget.NewForm(
//			&widget.FormItem{Text: "网卡序号:", Widget: index},
//			&widget.FormItem{Text: "最大传输单元:", Widget: mtu},
//			&widget.FormItem{Text: "网卡名称:", Widget: name},
//			&widget.FormItem{Text: "MAC地址", Widget: mac},
//			&widget.FormItem{Text: "状态标识:", Widget: flags},
//			&widget.FormItem{Text: "IPv4:", Widget: ipv4},
//			&widget.FormItem{Text: "IPv4子网掩码个数:", Widget: ipv4MaskCount},
//			&widget.FormItem{Text: "IPv6:", Widget: ipv6},
//			&widget.FormItem{Text: "IPv6子网掩码个数:", Widget: ipv6MaskCount},
//		)
//
//		index.SetText(fmt.Sprintf("%v", adaptor.Index))
//		mtu.SetText(fmt.Sprintf("%v", adaptor.Mtu))
//		name.SetText(fmt.Sprintf("%v", adaptor.Name))
//		mac.SetText(fmt.Sprintf("%v", adaptor.Mac))
//		flags.SetText(fmt.Sprintf("%v", adaptor.Flags))
//		ipv4.SetText(fmt.Sprintf("%v", adaptor.IPv4))
//		ipv4MaskCount.SetText(fmt.Sprintf("%v", adaptor.IPv4))
//		ipv6.SetText(fmt.Sprintf("%v", adaptor.IPv6))
//		ipv6MaskCount.SetText(fmt.Sprintf("%v", adaptor.IPv6MaskCount))
//
//		screens = append(screens, screen)
//	}
//	fmt.Println(adaptors[0].Name,screens)
//	content := widget.NewGroup(title, screens...)
//	return widget.NewScrollContainer(content)
//}

// 查询 -> 网络 -> 互联网查询布局
//func wanLayout(window fyne.Window) fyne.CanvasObject {
//	fontColor := color.RGBA{
//		R: 60,
//		G: 200,
//		B: 0,
//		A: 0,
//	}
//	return fyne.NewContainerWithLayout(
//		layout.NewAdaptiveGridLayout(2),
//		// 第1列显示按钮
//		widget.NewGroup("互联网",
//			widget.NewButtonWithIcon("IP", theme.InfoIcon(), func() {
//				fmt.Println("互联网  IP信息查询")
//			}),
//			widget.NewButtonWithIcon("带宽信息", theme.InfoIcon(), func() {
//				fmt.Println("互联网  带宽信息查询")
//			}),
//		),
//		// 第2列显示查询结果
//		widget.NewGroup("互联网",
//			widget.NewScrollContainer(
//				canvas.NewText(fmt.Sprintf("%v : %v", "互联网 IP", "114.114.114.114"), fontColor),
//			),
//			widget.NewScrollContainer(
//				canvas.NewText(fmt.Sprintf("%v : %v", "互联网 IP", "114.114.114.114"), fontColor),
//			),
//			widget.NewScrollContainer(
//				canvas.NewText(fmt.Sprintf("%v : %v", "互联网 IP", "114.114.114.114"), fontColor),
//			),
//			widget.NewScrollContainer(
//				canvas.NewText(fmt.Sprintf("%v : %v", "互联网 IP", "114.114.114.114"), fontColor),
//			),
//		),
//	)
//}
