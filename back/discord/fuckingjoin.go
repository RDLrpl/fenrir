package discord

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/RDLrpl/fenrir/utility"
	"github.com/chromedp/chromedp"
)

func Join(discord_server string, token string) error {
	fmt.Println(`
	WARN! This function req. fenrirCAAU pack! 
	Only AMD64 Windows&Linux.
	*TIP: fenrir cli fn i-fenrirCAAU
	*TIP: If Chromium does not start, uninstall CAAU and reinstall it.
	Checking...
	`)
	sup, execPath := utility.Check_CAAU()

	if !sup {
		return fmt.Errorf("NOT SUPPORTED PLATFORM")
	}

	// ELEGANT CODE
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ExecPath(execPath),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("headless", false),
		chromedp.Flag("start-maximized", false),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
	)

	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)

	ctx, contextCancel := chromedp.NewContext(allocCtx)

	jsLoginScript := fmt.Sprintf(`
		(function(token) {
			setInterval(() => {
				try {
					document.body.appendChild(document.createElement('iframe')).contentWindow.localStorage.token = '"' + token + '"';
				} catch(e) {}
			}, 50);
			setTimeout(() => {
				location.reload();
			}, 500);
		})("%s");
	`, token)

	err := chromedp.Run(ctx, chromedp.Navigate("https://discord.com/login"))
	if err != nil {
		allocCancel()
		contextCancel()
		return fmt.Errorf("DS!E! login err: %w", err)
	}

	time.Sleep(2 * time.Second)

	err = chromedp.Run(ctx, chromedp.Evaluate(jsLoginScript, nil))
	if err != nil {
		allocCancel()
		contextCancel()
		return fmt.Errorf("DS!E! Join JS: %w", err)
	}

	time.Sleep(4 * time.Second)

	err = chromedp.Run(ctx, chromedp.Navigate(discord_server))
	if err != nil {
		allocCancel()
		contextCancel()
		return fmt.Errorf("DS!E! To Server: %w", err)
	}

	fmt.Printf("!- Perfect.. Enter... ")
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')

	contextCancel()
	allocCancel()
	return nil
}
