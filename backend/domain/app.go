package domain

// App object.
// Implements IApp interface.
type App struct {
	Name     string  `json:"name"`
	Image    string  `json:"image"`
	Link     string  `json:"link"`
	Category string  `json:"category"`
	Rank     float64 `json:"rank"`
}

// NewApp allows to create an App Object.
// It returns the newly created App Object.
func NewApp(name string, image string, link string, category string, rank float64) App {
	// TODO : Add a bit of logic here
	// Checking if image is a proper URL, same as link
	return App{
		Name:     name,
		Image:    image,
		Link:     link,
		Category: category,
		Rank:     rank,
	}
}
