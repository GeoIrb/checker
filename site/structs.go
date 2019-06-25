package site

type Result struct {
	Type     string
	ID       int
	Status   int
	Keywords string
	Count    int
}

type Site struct {
	ID         int    `db:"id"`
	UpURL      string `db:"up_url"`
	URL        string `db:"control_keywords_page"`
	ControlSUM string `db:"up_url_hash"`
	Keywords   string `db:"control_keywords"`
	Type       string `db:"control_keywords_type"`
}

type Sites []Site

type Data struct {
	List  string `db:"list"`
	Sites Sites
}
