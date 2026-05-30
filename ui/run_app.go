package ui

import (
	"context"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	xwid "fyne.io/x/fyne/widget"
	"github.com/RDLrpl/fenrir/back"
	"github.com/RDLrpl/fenrir/back/discord"
)

func RunApp() {
	App := app.NewWithID("org.Fenrir.app")
	Window := App.NewWindow("Fenrir xwx<3")
	Window.Resize(fyne.NewSize(800, 800))
	Window.SetFixedSize(true)

	DiscordBox := container.NewVBox()
	DiscordTitle := widget.NewLabel("Info")

	DiscordTex := widget.NewLabel("Welcome!")

	Discordscroll := container.NewVScroll(DiscordBox)

	var icongif *xwid.AnimatedGif

	changePage := func(page string) {
		if icongif != nil {
			icongif.Stop()
			icongif = nil
		}

		DiscordBox.Objects = nil
		switch page {
		case "info":
			DiscordTitle.SetText("Info")
			DiscordTex.SetText("Welcome!")
		case "settings":
			DiscordTitle.SetText("Settings")
			DiscordTex.SetText("")
		case "whyphy":
			DiscordTitle.SetText("WhYhpY")
			DiscordTex.SetText("WhYhpY is simple GUI Fenrir-Based utility. Check DOCS")

			u := storage.NewFileURI("./res/whyh.gif")

			whyh := widget.NewLabel("->")

			accs := back.GetDsAccounts()
			accounts := []string{}
			targets := []string{}

			for id, _ := range accs {
				accounts = append(accounts, id)
			}

			trgs := back.GetDSTargets()
			for id, _ := range trgs {
				targets = append(targets, id)
			}

			curacc := ""
			curtarg := ""
			var cancelAsend context.CancelFunc

			img, err := xwid.NewAnimatedGif(u)
			if err == nil {
				img.SetMinSize(fyne.NewSize(67, 67))
				img.Start()

				icongif = img

				msgs := widget.NewEntry()

				Row := container.NewHBox(
					img,

					widget.NewSelect(accounts, func(value string) {
						curacc = value
					}),

					widget.NewSelect(targets, func(value string) {
						curtarg = value
					}),

					widget.NewButton("♻️", func() {
						number, err := strconv.Atoi(msgs.Text)
						if curacc != "" && curtarg != "" && err == nil {
							id_chan := back.GetDSTargets()[curtarg]
							account, _ := back.GetDsAccount(curacc)

							messages, err := discord.GetLastMessages(account.Token, id_chan[0], number, account.Proxy, account.ProxyLogin, account.ProxyPass)
							if err != nil {
								return
							}

							for _, msg := range messages {
								mesg := fmt.Sprintf("%s] %s: %s", msg.Timestamp.Format("15:04:02"), msg.Author.Username, msg.Content)
								msgw := widget.NewLabel(mesg)

								DiscordBox.Objects = append(DiscordBox.Objects, msgw)
							}
						}
					}),
					msgs,

					widget.NewButton("asends", func() {
						ctx, cancel := context.WithCancel(context.Background())
						cancelAsend = cancel

						go func() { discord.Asendds(ctx) }()

					}),

					widget.NewButton("STOPASEND", func() {
						if cancelAsend != nil {
							cancelAsend()
						}

					}),
				)

				DiscordBox.Objects = append(DiscordBox.Objects, whyh, Row)
			}
		case "accs":
			DiscordTitle.SetText("Account Manager")
			DiscordTex.SetText("You can add accounts and remove them!")

			accsListContainer := container.NewVBox()

			refreshList := func() {
				accsListContainer.Objects = nil
				accs := back.GetDsAccounts()
				for id, acc := range accs {
					haveproxy := widget.NewCheck("Have Proxy?", nil)
					if acc.Proxy != "" {
						haveproxy.SetChecked(true)
					} else {
						haveproxy.SetChecked(false)
					}
					haveproxy.Disable()

					Acc := container.NewHBox(
						widget.NewLabel(id),
						widget.NewLabel("|"),
						haveproxy,
						widget.NewButton("✏️", func() {

						}),
						widget.NewButton("❌", func() {
							back.RemoveDsAccount(id)
						}),
					)
					accsListContainer.Add(Acc)
				}
				accsListContainer.Refresh()
			}

			Man := container.NewHBox(
				widget.NewButton("♻️", func() {
					refreshList()
				}),
				widget.NewButton("➕", func() {

				}),
			)

			refreshList()
			DiscordBox.Objects = append(DiscordBox.Objects, Man, accsListContainer)
		case "cli":
			DiscordTitle.SetText("cli")
			DiscordTex.SetText("")
		}

		DiscordBox.Refresh()
	}

	discordSidebar := container.NewVBox(
		widget.NewLabel("Discord         !"),
		widget.NewSeparator(),
		widget.NewButton("INFO", func() { changePage("info") }),
		widget.NewButton("WHYPHY", func() { changePage("whyphy") }),
		widget.NewButton("CLIFI", func() { changePage("cli") }),
		widget.NewButton("ACCS", func() { changePage("accs") }),
		widget.NewButton("⚙️", func() { changePage("settings") }),
	)

	discordHeader := container.NewVBox(
		DiscordTitle,
		widget.NewSeparator(),
		DiscordTex,
	)

	discordContent := container.NewBorder(
		discordHeader,
		nil,
		nil,
		nil,
		Discordscroll,
	)

	discordLayout := container.NewBorder(nil, nil, discordSidebar, nil, discordContent)

	telegramSidebar := container.NewVBox(
		widget.NewLabel("TODO"),
	)

	telegramContent := container.NewVBox(
		widget.NewLabel("TODO"),
		widget.NewSeparator(),
	)

	telegramLayout := container.NewBorder(nil, nil, telegramSidebar, nil, telegramContent)

	tabs := container.NewAppTabs(
		container.NewTabItem("Welcome!", container.NewVBox(
			widget.NewLabel("Hello!"),
			widget.NewSeparator(),
		)),
		container.NewTabItem("Discord", discordLayout),
		container.NewTabItem("Telegram", telegramLayout),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	Window.SetContent(tabs)
	Window.ShowAndRun()
}
