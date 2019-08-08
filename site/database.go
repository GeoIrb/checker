package site

import "github.com/checker/app"

func Insert(cfg app.Data, results chan Result) {
	defer cfg.Completion("InsertData")

	query := app.Load("sql")

	connection := cfg.Prepare(query["insert"].(string))
	defer connection.Close()

	for r := range results {
		if _, err := connection.Exec(r.Type, r.ID, r.Status, r.Keywords, r.Count); err != nil {
			cfg.Err("InsertResult error: %s", err.Error())
		}
	}
}

func Select(cfg app.Data) (data Data) {
	defer cfg.Completion("SelectData")

	query := app.Load("sql")

	cfg.Select(&data.Sites, query["select"].(string))

	if cfg.Name == "system" {
		cfg.Get(&data.List, query["select_system_word"].(string))
	}

	return
}
