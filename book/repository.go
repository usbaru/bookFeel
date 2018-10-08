package book

type Repository struct{}

func (r Repository) GetBookByTitle(title string) Book {
	book := Book{
		"Best Book",
		"Danial Zaid",
		"Sports",
	}

	return book
}
