package category

import "time"

type Category struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
	IsActive  bool
}

func NewCategory(name string, createdAt time.Time, updatedAt time.Time) Category {
	return Category{
		ID:        0,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Name:      name,
		IsActive:  true,
	}
}
