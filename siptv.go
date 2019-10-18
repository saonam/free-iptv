package main

import (
	"context"
	"errors"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
)

func UpdateSmartTVApp(m3uURL string, macAddress string) error {

	log.Println("update m3uURL of device with mac address: " + macAddress)

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var res string

	const macSel = `//*[@id="mac_file"]`
	const urlSel = `//*[@id="1"]/input`

	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(`https://siptv.app/mylist/`),
		chromedp.WaitVisible(macSel),
		chromedp.SendKeys(macSel, macAddress),
		chromedp.WaitVisible(urlSel),
		chromedp.SendKeys(urlSel, m3uURL),
		chromedp.Click(`//*[@id="url_table"]/tbody/tr[1]/td[6]/input[2]`),
		chromedp.WaitVisible(`//*[@id="boxContent"]`),
		chromedp.Text(`//*[@id="boxContent"]`, &res),
	})
	if err != nil {
		return err
	}

	if `Adding...` != strings.TrimSpace(res) {
		return errors.New(res)
	}

	return nil
}
