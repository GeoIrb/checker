package site

import "checker/app"

func Insert(env app.Data, results chan Result) {
	defer env.Completion("InsertData")

	query := app.Load("sql")

	connection, _ := env.Prepare(query["insert"].(string))
	defer connection.Close()

	for r := range results {
		if _, err := connection.Exec(r.Type, r.ID, r.Status, r.Keywords, r.Count); err != nil {
			env.Err("InsertResult error: %s", err.Error())
		}
	}
}

func Select(env app.Data) (data Data) {
	defer env.Completion("SelectData")

	query := app.Load("sql")

	env.Select(&data.Sites, query["select"].(string))

	if app.GetPath() == "system" {
		env.Get(&data.List, query["select_system_word"].(string))
	}

	return
}
