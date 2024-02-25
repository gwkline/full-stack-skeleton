package repo

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/gwkline/full-stack-skeleton/backend/types"
	"gorm.io/gorm"
)

type Repository struct {
	User  types.IGenericRepo[types.User]
	Apple types.IGenericRepo[types.Apple]
	DB    *gorm.DB
}

type Generepo[T types.Modeler] struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	db = db.Session(&gorm.Session{NewDB: true})
	return &Repository{
		User:  &Generepo[types.User]{db: db},
		Apple: &Generepo[types.Apple]{db: db},
		DB:    db,
	}
}

func (g *Generepo[T]) Create(entity *T) (*T, error) {
	result := g.db.Create(entity)
	if result.Error != nil {
		return nil, result.Error
	}
	return entity, nil
}

func (g *Generepo[T]) Update(entity *T) (*T, error) {
	result := g.db.Save(entity)
	if result.Error != nil {
		return nil, result.Error
	}
	return entity, nil
}

func (g *Generepo[T]) Archive(entity *T) error {
	result := g.db.Delete(entity)
	return result.Error
}

func (g *Generepo[T]) Delete(entity *T) error {
	result := g.db.Unscoped().Delete(entity)
	return result.Error
}

// uses = to return an array of *T
func (g *Generepo[T]) ListBy(filters []types.Filter, preloads ...string) ([]*T, error) {
	var entities []*T
	query := g.db

	if len(preloads) > 0 {
		for _, preload := range preloads {
			query = query.Preload(preload)
		}
	}

	for _, filter := range filters {
		if filter.Operator == types.OperatorIncludes {
			query = query.Where(fmt.Sprintf("%s @> ARRAY[?]", filter.Key), filter.Value)
			continue
		}
		if filter.Operator != types.OperatorUnset {
			query = query.Where(fmt.Sprintf("%s %s ?", filter.Key, filter.Operator), filter.Value)
			continue
		}

		switch v := filter.Value.(type) {
		case uint, *uint, int, *int, string, *string, float64, *float64, float32, *float32:
			query = query.Where(filter.Key+" = ?", v)
		case []uint, []string:
			query = query.Where(filter.Key+" IN (?)", v)
		default:
			sentry.CaptureException(fmt.Errorf("unsupported type for ListBy: %T", v))
			fmt.Printf("unsupported type for ListBy: %T\n", v)
			query = query.Where(filter.Key+" = ?", v)
		}
	}

	result := query.Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return entities, nil
}

// uses = to return a single *T
func (g *Generepo[T]) FindBy(filters []types.Filter, preloads ...string) (*T, error) {
	entity := new(T)
	query := g.db

	if len(preloads) > 0 {
		for _, preload := range preloads {
			query = query.Preload(preload)
		}
	}

	for _, filter := range filters {
		if filter.Operator == types.OperatorIncludes {
			query = query.Where(fmt.Sprintf("%s @> ARRAY[?]", filter.Key), filter.Value)
			continue
		}
		if filter.Operator != types.OperatorUnset {
			query = query.Where(fmt.Sprintf("%s %s ?", filter.Key, filter.Operator), filter.Value)
			continue
		}

		switch v := filter.Value.(type) {
		case uint, *uint, int, *int, string, *string:
			query = query.Where(filter.Key+" = ?", v)
		default:
			fmt.Println("unsupported type for FindBy")
			fmt.Printf("Type of v: %T\n", v)
			query = query.Where(filter.Key+" = ?", v)
		}
	}

	result := query.First(&entity)
	if result.Error != nil {
		return nil, result.Error
	}
	return entity, nil
}

// uses ILIKE to return an array of *T
func (g *Generepo[T]) FuzzyFindBy(key string, value any) ([]*T, error) {
	var entities []*T
	result := g.db.Where(key+" LIKE ?", "%"+fmt.Sprint(value)+"%").Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return entities, nil
}
