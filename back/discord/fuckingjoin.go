package discord

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/go-vgo/robotgo"
)

func Join(discord_server string, token string) error {
	// bezumie

	// ABSOLUTE ELEGANT CODE
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ExecPath("/usr/bin/thorium-browser"),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("headless", false),
		chromedp.Flag("start-maximized", true),
		chromedp.Flag("user-data-dir", fmt.Sprintf("/tmp/thorium_%d", time.Now().Unix())),

		chromedp.Flag("disable-dev-tools", "false"),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.Flag("excludeSwitches", "enable-automation"),
	)

	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, contextCancel := chromedp.NewContext(allocCtx)

	err := chromedp.Run(ctx, chromedp.Navigate("https://discord.com/login"))
	if err != nil {
		return fmt.Errorf("TG!D! Thorium: %w", err)
	}
	robotgo.MoveClick(500, 500, "left", true)

	robotgo.KeyDown("ctrl")
	robotgo.KeyDown("shift")

	robotgo.KeyTap("j")

	robotgo.KeyUp("shift")
	robotgo.KeyUp("ctrl")
	time.Sleep(4000 * time.Millisecond)

	robotgo.KeyDown("ctrl")
	robotgo.KeyTap("v")
	time.Sleep(100 * time.Millisecond)
	robotgo.KeyUp("ctrl")

	robotgo.TypeStr("allow pasting")

	time.Sleep(100 * time.Millisecond)
	robotgo.KeyTap("enter")

	time.Sleep(100 * time.Millisecond)
	towrite := fmt.Sprintf(`function login(token) {
	setInterval(() => {
		document.body.appendChild(document.createElement('iframe')).contentWindow.localStorage.token = '"' + token + '"';
	}, 50);
	setTimeout(() => {
		location.reload();
	}, 2500);
}
login("%s");`, token)

	robotgo.WriteAll(towrite)

	robotgo.KeyDown("ctrl")
	robotgo.KeyTap("v")
	robotgo.KeyUp("ctrl")

	time.Sleep(100 * time.Millisecond)
	robotgo.KeyTap("enter")

	time.Sleep(4000 * time.Millisecond)
	err = chromedp.Run(ctx, chromedp.Navigate(discord_server))
	if err != nil {
		return fmt.Errorf("TG!D! Thorium: %w", err)
	}

	fmt.Printf("Press Enter to close browser and return... ")
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')

	contextCancel()
	allocCancel()

	return nil
}

// js := fmt.Sprintf(`(function(){const t="%s";setInterval(()=>{const f=document.createElement("iframe");document.body.appendChild(f);f.contentWindow.localStorage.token='"'+t+'"';f.remove()},50);setTimeout(()=>location.reload(),1800)})()`, token)
/*
	towrite := fmt.Sprintf(`function login(token) {
		setInterval(() => {
			document.body.appendChild(document.createElement('iframe')).contentWindow.localStorage.token = '"' + token + '"'
		}, 50);
		setTimeout(() => {
			location.reload();
		}, 2500);
	}
	login("%s");`, token)
*/
