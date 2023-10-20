package main

//Categories sample
//"C++", "C#", "Java", "JavaScript", "HTML", "CSS", "PHP", "Go", "Rust", "Node"})

//       __categories__
//      |  category    |
//      |  TEXT        |

// Create categories table
func crerateCategoriesTable() error {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS categories(category TEXT NOT NULL UNIQUE)")
	if err != nil {
		return err
	}
	defer statement.Close()
	statement.Exec()
	return nil
}

func getCategories() ([]string, error) {
	categories := []string{}
	rows, err := db.Query("SELECT * FROM categories")
	if err != nil {
		return categories, err
	}
	defer rows.Close()
	var category string
	for rows.Next() {
		err = rows.Scan(&category)
		if err != nil {
			return categories, err
		}
		categories = append(categories, category)
	}
	err = rows.Err()
	if err != nil {
		return categories, err
	}
	return categories, nil
}

func insertCategories(categories []string) error {
	statement, err := db.Prepare("INSERT INTO categories (category) VALUES(?)")
	if err != nil {
		return err
	}
	defer statement.Close()
	for _, category := range categories {
		_, err = statement.Exec(category)
		if err != nil {
			return err
		}
	}
	return nil
}
