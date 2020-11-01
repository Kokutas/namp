package gui

/*func help(window fyne.Window) fyne.CanvasObject {
	// 返回一个新的tab容器
	return widget.NewTabContainer(
		widget.NewTabItemWithIcon("局域网", theme.HelpIcon(), lan(window)),
	)
}
func lan(window fyne.Window) fyne.CanvasObject {
	return fyne.NewContainerWithLayout(
		// 两列的格子布局
		layout.NewAdaptiveGridLayout(2),
		// 创建左边的
		widget.NewScrollContainer(loadDialogGroup(window)),
		widget.NewScrollContainer(loadWindowGroup()),
	)
}
func loadDialogGroup(win fyne.Window) *widget.Group {
	return widget.NewGroup("查询",
		widget.NewButtonWithIcon("网卡信息",theme.SearchIcon(), func() {
			prog := dialog.NewProgress("[本机网卡信息]", "正在查询...", win)

			go func() {
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
		}),
	)
}

func loadWindowGroup() fyne.Widget {
	windowGroup := widget.NewGroup("Windows",
		widget.NewButton("New window", func() {
			w := fyne.CurrentApp().NewWindow("Hello")
			w.SetContent(widget.NewLabel("Hello World!"))
			w.Show()
		}),
		widget.NewButton("Fixed size window", func() {
			w := fyne.CurrentApp().NewWindow("Fixed")
			w.SetContent(fyne.NewContainerWithLayout(layout.NewCenterLayout(), widget.NewLabel("Hello World!")))

			w.Resize(fyne.NewSize(240, 180))
			w.SetFixedSize(true)
			w.Show()
		}),
		widget.NewButton("Centered window", func() {
			w := fyne.CurrentApp().NewWindow("Central")
			w.SetContent(fyne.NewContainerWithLayout(layout.NewCenterLayout(), widget.NewLabel("Hello World!")))

			w.CenterOnScreen()
			w.Show()
		}))

	drv := fyne.CurrentApp().Driver()
	if drv, ok := drv.(desktop.Driver); ok {
		windowGroup.Append(
			widget.NewButton("Splash Window (only use on start)", func() {
				w := drv.CreateSplashWindow()
				w.SetContent(widget.NewLabelWithStyle("Hello World!\n\nMake a splash!",
					fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))
				w.Show()

				go func() {
					time.Sleep(time.Second * 3)
					w.Close()
				}()
			}))
	}

	otherGroup := widget.NewGroup("Other",
		widget.NewButton("Notification", func() {
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Fyne Demo",
				Content: "Testing notifications...",
			})
		}))

	return widget.NewVBox(windowGroup, otherGroup)
}

func loadImage(f fyne.URIReadCloser) *canvas.Image {
	data, err := ioutil.ReadAll(f)
	if err != nil {
		fyne.LogError("Failed to load image data", err)
		return nil
	}
	res := fyne.NewStaticResource(f.Name(), data)

	return canvas.NewImageFromResource(res)
}

func loadText(f fyne.URIReadCloser) string {
	data, err := ioutil.ReadAll(f)
	if err != nil {
		fyne.LogError("Failed to load text data", err)
		return ""
	}
	if data == nil {
		return ""
	}

	return string(data)
}
*/