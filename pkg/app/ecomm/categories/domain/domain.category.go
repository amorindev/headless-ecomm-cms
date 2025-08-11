package domain

type Category struct {
	ID   interface{} `json:"id" bson:"_id"`
	Name string     `json:"name" bson:"name"`
}

func NewCategory(name string) *Category {
	return &Category{
		Name: name,
	}
}