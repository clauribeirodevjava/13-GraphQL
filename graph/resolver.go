package graph

import "github.com/clauribeirodevjava/13-GraphQL.git/graph/db/categoryBO"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryDB *categoryBO.Category
	CourseDB   *categoryBO.Course
}
