package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"runtime"
	"sync"
	"time"
)

func webParse() {
	runtime.GOMAXPROCS(KernelCount)
	urls := &[]string{
		"https://learning.postman.com/",
		"https://github.com/postmanlabs/",
		"https://github.com/postmanlabs/",
		"https://postman.zoom.us/webinar/register/",
		"https://cheerio.js.org/",
		"https://cheerio.js.org/",
		"https://duckduckgo.com/",
		"https://stackoverflow.com/questions/16566530/what-is-the-simplest-ews-call-to-test-connectivity",
		"https://ria.ru/",
		"https://ru.wikipedia.org/wiki/%D0%A0%D0%98%D0%90_%D0%9D%D0%BE%D0%B2%D0%BE%D1%81%D1%82%D0%B8",
		"https://twitter.com/rianru",
		"https://vk.com/ria",
		"https://ok.ru/ria",
		"https://www.youtube.com/user/rianovosti",
		"https://flipboard.com/@rianovosti",
		"https://zen.yandex.ru/ria",
		"https://twitter.com/riabreakingnews",
		"https://invite.viber.com/?g2=AQAOuCQJzow8L0hQtsB3j3zzc7YMKEHijsslO3ZsCFmFcZGSed0OTOa0ruAXXZn6&lang=ru",
		"https://tgclick.com/rian_ru",
		"https://yandex.ru/maps/org/rossiya_segodnya/1061985604/?ll=37.590466%2C55.737481&z=14",
		"https://www.google.com/maps/place/%D0%A0%D0%98%D0%90+%D0%9D%D0%BE%D0%B2%D0%BE%D1%81%D1%82%D0%B8/@55.7374019,37.5881028,17z/data=!3m1!4b1!4m5!3m4!1s0x46b54ba5ecf15a7f:0x2b9d63d783e6e919!8m2!3d55.7374019!4d37.5902915",
		"https://news.google.com/publications/CAAqBwgKMNu67AEwn6QS?hl=ru&gl=RU&ceid=RU:ru",
		"https://ria.ru/tourism/",
		"https://ria.ru/20231125/finlyandiya-1911801815.html",
		"https://ria.ru/20231125/vzryvy-1911775107.html",
		"https://ria.ru/20231125/izrail-1911796957.html",
		"https://ria.ru/20231125/spetsoperatsiya-1911770224.html",
		"https://ria.ru/20231125/orden-1911778006.html",
		"https://ria.ru/20231125/obstrely-1911784615.html",
		"https://ria.ru/20231125/ssha-1911768367.html",
		"https://ria.ru/20231125/zakon-1911774930.html",
		"https://ria.ru/20231125/makedoniya-1911765591.html",
		"https://ria.ru/20231125/vzryvy-1911780271.html",
		"https://ria.ru/20231125/es-1911765763.html",
		"https://ria.ru/20231125/gazeta_zeit-1911736604.html",
		"https://ria.ru/20231125/kulturnaya_tusovka-1911716092.html",
		"https://ria.ru/20231124/izrail-1911540672.html",
		"https://ria.ru/20231125/evropa-1911671835.html",
		"https://ria.ru/20231125/izrail-1911792108.html",
		"https://ria.ru/20231125/rukovodstvo-1911788742.html",
		"https://ria.ru/20231125/kriminal-1911591731.html",
		"https://ria.ru/20231124/flamingo-1911644350.html",
		"https://ria.ru/20231124/leopard-1911590893.html",
		"https://ria.ru/20231125/ubiystvo-1911707073.html",
		"https://ria.ru/20231125/plan-1911796452.html",
		"https://ria.ru/20231125/dtp-1911790636.html",
		"https://ria.ru/20231125/amaterasu-1911670091.html",
		"https://ria.ru/20231125/tekhnologiya-1911768006.html?rcmd_alg=slotter",
		"https://ria.ru/20231125/most-1911791757.html?rcmd_alg=slotter",
		"https://ria.ru/20231125/izrail-1911789851.html?rcmd_alg=slotter",
		"https://ria.ru/20231125/pobedy-1911766263.html?rcmd_alg=slotter",
		"https://ria.ru/20231125/samolety-1911790734.html?rcmd_alg=slotter",
		"https://ria.ru/20231125/prava-1911781360.html",
		"https://ria.ru/20231125/lantset-1911801180.html",
		"https://ria.ru/20231125/uekhavshie-1911781965.html",
		"https://ria.ru/20231125/blokada-1911772239.html",
		"https://ria.ru/20231125/napoleon-1911567476.html",
		"https://ria.ru/20231123/vystavka-1911265861.html",
		"https://nfw.ria.ru/flv/file.aspx?ID=80430057&type=mp3",
		"https://itunes.apple.com/ru/podcast/%D0%BA%D0%B0%D0%BA-%D1%8D%D1%82%D0%BE-%D0%BF%D0%BE-%D1%80%D1%83%D1%81%D1%81%D0%BA%D0%B8/id1438645662?mt=2",
		"https://www.google.com/podcasts?feed=aHR0cHM6Ly9yc3Muc2ltcGxlY2FzdC5jb20vcG9kY2FzdHMvODUxMS9yc3M%3D",
		"https://tunein.com/podcasts/Education-Podcasts/How-is-it-in-Russian-p473133/",
		"https://castbox.fm/channel/%D0%9A%D0%B0%D0%BA-%D1%8D%D1%82%D0%BE-%D0%BF%D0%BE-%D1%80%D1%83%D1%81%D1%81%D0%BA%D0%B8-id1445548",
		"https://redirect.appmetrica.yandex.com/serve/458217887027478676",
		"https://ria.ru/export/itunes/rss2/kak_jeto_po_russki.xml",
		"https://nfw.ria.ru/flv/file.aspx?ID=80430057&type=mp3",
		"https://ria.ru/20231119/lomonosov-1909636608.html",
		"https://ria.ru/20231124/rubl-1911403066.html",
		"https://ria.ru/20231124/kriminal-1911369602.html",
		"https://ria.ru/20231124/meoty-1911400537.html",
		"https://ria.ru/20231124/tretyakovka-1911357338.html",
		"https://ria.ru/20231124/niderlandy-1911453791.html",
		"https://ria.ru/20231124/russkiy_mir-1911455298.html",
		"https://ria.ru/20231123/sanktsii-1911195495.html",
		"https://rsport.ria.ru/20231125/khokkey-1911775325.html",
		"https://rsport.ria.ru/20231125/bendi-1911777879.html",
		"https://rsport.ria.ru/20231125/figuristy-1911776959.html",
		"https://rsport.ria.ru/20231125/khokkey-1911764835.html",
		"https://rsport.ria.ru/20231125/rozherio_seni-1911754123.html",
		"https://rsport.ria.ru/20231125/magomedkerimov-1911772034.html",
		"https://rsport.ria.ru/20231125/nepomnyaschiy-1911763846.html",
		"https://rsport.ria.ru/20231125/nba-1911762756.html",
		"https://rsport.ria.ru/20231125/monako-1911761777.html",
		"https://vk.com/ria",
		"https://www.odnoklassniki.ru/ria",
		"https://zen.yandex.ru/ria?invite=1",
		"https://rutube.ru/channel/23469114/",
		"https://invite.viber.com/?g2=AQAOuCQJzow8L0hQtsB3j3zzc7YMKEHijsslO3ZsCFmFcZGSed0OTOa0ruAXXZn6",
		"https://tgclick.com/rian_ru",
		"https://twitter.com/rianru",
		"https://twitter.com/riabreakingnews",
		"https://www.tiktok.com/@ria_novosti",
		"https://ria.ru/",
		"https://ria.ru/politics/",
		"https://ria.ru/society/",
		"https://ria.ru/economy/",
		"https://ria.ru/world/",
		"https://ria.ru/incidents/",
		"https://rsport.ria.ru/",
		"https://ria.ru/science/",
	}

	start := time.Now()
	res := getPages(urls)
	totalTime := time.Since(start)
	fmt.Printf("Последовательный результат вычислений 0: %v\n", res[0])
	fmt.Printf("Последовательный результат вычислений 99: %v\n", res[99])
	fmt.Printf("Последовательное время  вычислений: %.3v (c)\n", totalTime.Seconds())

	start = time.Now()
	resParallel := getPagesParallel(urls, GoroutineCnt)
	totalTime = time.Since(start)
	fmt.Printf("Параллельный результат вычислений 0: %v\n", resParallel[0])
	fmt.Printf("Параллельный результат вычислений 99: %v\n", resParallel[99])
	fmt.Printf("Параллельное время  вычислений: %.3v (с)\n", totalTime.Seconds())
}

func getPage(url *string) string {
	resp, err := http.Get(*url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(html)
}

func getPages(urls *[]string) []string {
	pages := make([]string, len(*urls))
	for i := 0; i < len(*urls); i++ {
		html := getPage(&((*urls)[i]))
		pages = append(pages, html)
	}
	return pages
}

func getPagesChannel(urls *[]string, ch chan []string) {
	ch <- getPages(urls)
}

func getPagesParallel(arr *[]string, goroutineCnt int) []string {
	chs := makeStringChannelArray(goroutineCnt)
	partSize := len(*arr) / goroutineCnt
	for i := 0; i < goroutineCnt; i++ {
		chunkArr := make([]string, partSize)
		copy(chunkArr, (*arr)[partSize*i:partSize*(i+1)])
		getPagesChannel(&chunkArr, chs[i])
	}
	chResult := make(chan []string)
	go getPagesFromChannels(chResult, chs)
	return <-chResult
}

func makeStringChannelArray(channelCount int) []chan []string {
	chs := make([]chan []string, channelCount)
	for i := 0; i < len(chs); i++ {
		chs[i] = make(chan []string, 1)
	}
	return chs
}

func getPagesFromChannels(result chan []string, chs []chan []string) {
	var res []string
	cases := make([]reflect.SelectCase, len(chs))
	for i, ch := range chs {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
	}
	remaining := len(chs)
	for remaining > 0 {
		chosen, value, ok := reflect.Select(cases)
		if !ok {
			// Если канал закрыт, убираем case из select-a
			cases[chosen].Chan = reflect.ValueOf(nil)
			remaining -= 1
			continue
		}
		m := sync.Mutex{}
		m.Lock()
		res = append(res, value.Interface().([]string)...)
		m.Unlock()
		close(chs[chosen])
	}
	result <- res
}
