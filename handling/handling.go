package handling

import (
	"geoirb/checker/app"
	"geoirb/checker/site"
	"strings"
	"sync"
	"time"
)

func Start(cfg app.Data) {
	wg := &sync.WaitGroup{}

	theards := app.Load("thread")["number"].(int)
	timeout := time.Duration(app.Load("time")["request"].(int)) * time.Second
	name := cfg.Name[strings.LastIndex(cfg.Name, "/")+1:]

	data := site.Select(cfg)
	resChan := make(chan site.Result)

	n := int(len(data.Sites) / theards)
	if int(len(data.Sites)%theards) > 0 {
		n++
	}

	i := 0
	for i < theards {
		i++

		wg.Add(1)
		go func(wg *sync.WaitGroup, sl site.Sites) {
			defer wg.Done()

			for _, s := range sl {

				var result site.Result
				result.Type = name
				result.ID = s.ID

				url := s.URL
				if result.Type == "hash" {
					url = s.UpURL
				}

				html, err := site.HTTPGet(url, timeout)

				if err != nil && err.Error() == "No response" {
					result.Status = 2
					cfg.Err(url)
					resChan <- result
					continue
				}

				switch result.Type {
				case "hash":
					result.Status, _ = hash(html, s.ControlSUM)
				case "system":
					result.Status, result.Keywords, result.Count = system(string(html), data.List)
				case "keywords":
					result.Status, result.Keywords, result.Count = keywords(string(html), s.Type, s.Keywords)
				default:
					cfg.Err("Unknown checker type")
				}

				resChan <- result
			}
		}(wg, data.Sites[:n])

		data.Sites = data.Sites[n:]
		if len(data.Sites) < n {
			n = len(data.Sites)
		}
	}

	go site.Insert(cfg, resChan)

	wg.Wait()
	close(resChan)
}
