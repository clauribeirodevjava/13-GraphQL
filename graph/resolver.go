package graph

import (
	"github.com/clauribeirodevjava/13-GraphQL/graph/internal/database"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryField *database.Category
	CourseField   *database.Course
}
