package book

type Repository struct{}

func (r Repository) GetBookByTitle() Book {
	book := Book{
		"Best Book",
		"Danial Zaid",
		"Sports",
	}

	return book
}
