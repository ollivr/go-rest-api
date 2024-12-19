package tools

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

type ResponsiveImageData struct {
	x72             string
	x192            string
	x256            string
	x384            string
	x512            string
	srcset          string
	src             string
	alt             string
	collectionName  string
	conversionsDisk string
	createdAt       string
	disk            string
	fileName        string
	id              string
	mimeType        string
	modelId         string
	modelType       string
	name            string
	orderColumn     int
	originalUrl     string
	previewUrl      string
	size            int
	updatedAt       string
	uuid            string
}
