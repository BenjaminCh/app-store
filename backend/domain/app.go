package domain

// App object.
// Implements IApp interface.
type App struct {
	name     string
	image    string
	link     string
	category string
	rank     float64
}

// NewApp allows to create an App Object.
// It returns the newly created App Object.
func NewApp(name string, image string, link string, category string, rank float64) App {
	return App{
		name:     name,
		image:    image,
		link:     link,
		category: category,
		rank:     rank,
	}
}
