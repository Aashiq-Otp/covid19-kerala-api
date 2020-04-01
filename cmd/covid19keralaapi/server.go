package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	. "github.com/yedhink/covid19-kerala-api/internal/scraper"
)

type Scraper interface {
	GetMainPage(chan [2]string)
	GetLatestPDF(chan string)
}

type Website struct {
	baseURL          string
	mainPageURL      string
	bulletinPageURL  string
	latestPDFFileURL string
}

func (w *Website) GetLatestPDF() (string){
	links := Scrape(w.baseURL, w.bulletinPageURL,"div","entry-content")
	w.latestPDFFileURL = links.Attrs()["href"]
	c <- w.baseURL + w.latestPDFFileURL
}

func (w *Website) GetMainPage(c chan [2]string) {
	links := Scrape(w.baseURL, w.mainPageURL,"h3","entry-title")
	latestFileName := "data/" + strings.ReplaceAll(links.Text(), "/", "-") + ".pdf"
	bulletinPage := links.Attrs()["href"]
	c <- [2]string{latestFileName, bulletinPage}
}

	return w.baseURL + w.latestPDFFileURL
	var website = &Website{
		baseURL: "http://dhs.kerala.gov.in",
func (w *Website) GetMainPage() ([2]string){
			`%e0%b4%ac%e0%b5%81%e0%b4%b3%e0%b5%8d%e0%b4%b3%e0%b4%b1` +
			`%e0%b5%8d%e0%b4%b1%e0%b4%bf%e0%b4%a8%e0%b5%8d%e2%80%8d/`,
	}
	return [2]string{latestFileName, bulletinPage}

	/*
		Glob ignores file system errors such as I/O errors reading directories.
		The only possible returned error is ErrBadPattern, when pattern is malformed.
	*/
	files, err := filepath.Glob("data/*.pdf")
	if err != nil || len(files) == 0{
		fmt.Println(err)
		os.Exit(1)
	}

	chan1 := make(chan [2]string)
	chan2 := make(chan string)
	for _, f := range files {
		go sc.GetMainPage(chan1)
		x := <-chan1
		website.bulletinPageURL = x[1]
		if f == x[0] {
			fmt.Println("The pdf file is already latest")
		} else {
			fmt.Printf("You need latest pdf file : %s(local) != %s(remote)\n", f, x[0])
			go sc.GetLatestPDF(chan2)
			fmt.Printf("lastest file : %s\n",<-chan2)
		}
		website.bulletinPageURL = res[1]
			fmt.Printf("lastest file : %s\n",sc.GetLatestPDF())
