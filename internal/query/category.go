package query

import (
	"strings"
)

type CategoryLabel struct {
	Name  string
	Title string
}

func (q *Query) CategoryLabels(limit, offset int) (results []CategoryLabel) {
	s := q.db.NewScope(nil).DB()

	s = s.Table("categories").
		Select("label_name AS name").
		Joins("JOIN labels l ON categories.category_id = l.id").
		Group("label_name").
		Limit(limit).Offset(offset)

	if err := s.Scan(&results).Error; err != nil {
		log.Errorf("categories: %s", err.Error())
		return results
	}

	for i, l := range results {
		results[i].Title = strings.Title(l.Name)
	}

	return results
}
