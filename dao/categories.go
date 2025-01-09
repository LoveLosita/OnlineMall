package dao

func CheckIfCategoryExists(categoryID int) (bool, error) {
	query := "SELECT id FROM categories WHERE id=?"
	rows, err := Db.Query(query, categoryID)
	if err != nil {
		return false, err
	}
	if rows.Next() { //如果有这个分类
		return true, nil
	} else { //如果没有这个分类
		return false, nil
	}
}

func AddCategory(name string, description string) error {
	query := "INSERT INTO categories(name,description) VALUES(?,?)"
	_, err := Db.Exec(query, name, description)
	if err != nil {
		return err
	}
	return nil
}

func CheckIfCategoryNameExists(name string) (bool, error) {
	query := "SELECT id FROM categories WHERE name=?"
	rows, err := Db.Query(query, name)
	if err != nil {
		return false, err
	}
	if rows.Next() { //如果有这个分类
		return true, nil
	} else { //如果没有这个分类
		return false, nil
	}
}
