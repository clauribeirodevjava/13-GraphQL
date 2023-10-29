package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name, description, categoryId string) (*Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO Course(ID, Name, Description, CategoryID) values (?,?,?,?)",
		id, name, description, categoryId)
	if err != nil {
		return nil, err
	}
	return &Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryID:  categoryId,
	}, nil

}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("Select ID,Name,Description,CategoryID from Course")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	Courses := []Course{}
	for rows.Next() {
		var id, name, description, categoryId string
		if err := rows.Scan(&id, &name, &description, &categoryId); err != nil {
			return nil, err
		}
		Courses = append(Courses, Course{ID: id, Name: name, Description: description, CategoryID: categoryId})
	}
	return Courses, nil
}

func (c *Course) FindByCategoryID(categoryID string) ([]Course, error) {
	rows, err := c.db.Query("SELECT ID, Name, Description, CategoryID FROM Course WHERE CategoryID = ?", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []Course{}
	for rows.Next() {
		var id, name, description, categoryID string
		if err := rows.Scan(&id, &name, &description, &categoryID); err != nil {
			return nil, err
		}
		courses = append(courses, Course{ID: id, Name: name, Description: description, CategoryID: categoryID})
	}
	return courses, nil
}

func (c *Category) FindByCourseId(courseID string) (Category, error) {
	var id, name, description string
	err := c.db.QueryRow("SELECT c.ID, c.Name, c.Description FROM Category c JOIN Course co ON c.ID = co.CategoryID WHERE co.ID = ?", courseID).
		Scan(&id, &name, &description)
	if err != nil {
		return Category{}, err
	}

	if err != nil {
		return Category{}, err
	}
	return Category{ID: id, Name: name, Description: description}, nil
}
