package repo

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/gwkline/full-stack-skeleton/backend/types"
	"gorm.io/gorm"
)

func (g *Generepo[T]) Connection(tx *gorm.DB, tableName string, filters []types.FilterInput, sortBy string, direction types.AscOrDesc, limit int, after int) ([]*T, *types.PageInfo, error) {
	var sortByValue interface{}
	var err error
	if after != 0 && after != 9999999999999999 {
		if strings.Contains(sortBy, ".") {
			sortByValue, err = g.getSortByValueFromCursorJoins(tableName, after, sortBy, direction)
		} else {
			sortByValue, err = g.getSortByValueFromCursor(tableName, after, sortBy, direction)
		}
	}
	if err != nil {
		return nil, nil, fmt.Errorf("failure getting cursor: %w", err)
	}

	tx = addJoins(tx, sortBy, filters, tableName)

	if sortBy != "id" {
		tx = tx.Order(fmt.Sprintf("%s %s, %s.id %s", sortBy, direction, tableName, direction))
	} else {
		tx = tx.Order(fmt.Sprintf("%s %s", sortBy, direction))
	}

	tx, err = doFilter(tx, tableName, filters)
	if err != nil {
		return nil, nil, err
	}

	// Get a count of the total number of records
	unlimited := tx
	var count int64
	if err := unlimited.Model(new(T)).Count(&count).Error; err != nil {
		return nil, nil, fmt.Errorf("failed to count: %w", err)
	}

	// after = 11
	// direction = asc
	if sortBy != "id" && sortByValue != nil {
		tx = tx.Where(fmt.Sprintf("(%s, %s.id) %s= (?, ?)", sortBy, tableName, getOrderOperator(direction)), sortByValue, after)
	} else {
		tx = tx.Where(fmt.Sprintf("%s.id %s ?", tableName, getOrderOperator(direction)), after)
	}
	// WHERE "scripts.id > 11"

	// Apply the limit
	tx = tx.Limit(limit + 1)

	// Fetch the results
	var finalResults []*T
	if err := tx.Model(new(T)).Find(&finalResults).Error; err != nil {
		return nil, nil, fmt.Errorf("failed to find: %w", err)
	}

	hasNextPage := false
	if len(finalResults) > limit {
		finalResults = finalResults[:len(finalResults)-1]
		hasNextPage = true
	}

	startCursor, endCursor := "", ""
	if len(finalResults) > 0 {
		start := *finalResults[0]
		end := *finalResults[len(finalResults)-1]
		startCursor = encodeCursor(start.GetID())
		endCursor = encodeCursor(end.GetID())
	}

	return finalResults, &types.PageInfo{
		Count:       int(count),
		HasNextPage: hasNextPage,
		StartCursor: startCursor,
		EndCursor:   endCursor,
	}, nil
}

func (g *Generepo[T]) getSortByValueFromCursor(tableName string, after int, sortBy string, direction types.AscOrDesc) (interface{}, error) {
	dbP := *g.db
	db := dbP.Session(&gorm.Session{NewDB: true})
	db = db.Table(tableName).Select(fmt.Sprintf("%s.%s as sortByValue, %s.id as id", tableName, sortBy, tableName))

	var result []map[string]interface{}
	err := db.Where(fmt.Sprintf("%s.id = ?", tableName), after).Find(&result).Error
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no results found for after cursor")
	}

	sortByValue := result[0]["sortbyvalue"]

	return sortByValue, nil
}

func (g *Generepo[T]) getSortByValueFromCursorJoins(mainTableName string, after int, sortBy string, direction types.AscOrDesc) (interface{}, error) {
	dbP := *g.db
	db := dbP.Session(&gorm.Session{NewDB: true})
	db = db.Table(mainTableName).Select(fmt.Sprintf("%s as sortByValue, %s.id as id", sortBy, mainTableName))

	// We need to join the related table to get the sort by value
	// e.g. the current table is "scripts" and the sort is "users.name"
	relatedTable := strings.Split(sortBy, ".")[0] // "users"

	// JOIN "users" ON "scripts".user_id = "users".id
	joinClause := fmt.Sprintf("JOIN %s ON %s.id = %s.%s_id", relatedTable, relatedTable, mainTableName, getSingularObjectName(relatedTable))
	db = db.Joins(joinClause)

	var result []map[string]interface{}
	err := db.Where(fmt.Sprintf("%s.id = ?", mainTableName), after).Find(&result).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no results found for after cursor")
	}

	sortByValue := result[0]["sortbyvalue"]

	return sortByValue, nil
}

// getOrderOperator returns the correct operator based on the sorting direction
func getOrderOperator(direction types.AscOrDesc) string {
	if direction == types.AscOrDescDesc {
		return "<"
	}
	return ">"
}

func addJoins(db *gorm.DB, sortBy string, filters []types.FilterInput, mainTableName string) *gorm.DB {
	// A map to keep track of which joins have been added
	joinedTables := make(map[string]bool)

	// Check if sortBy requires a join
	if strings.Contains(sortBy, ".") {
		relatedTable := strings.Split(sortBy, ".")[0]
		if !joinedTables[relatedTable] {
			db = db.Joins(fmt.Sprintf("LEFT JOIN %s ON %s.id = %s.%s_id", relatedTable, relatedTable, mainTableName, getSingularObjectName(relatedTable)))
			joinedTables[relatedTable] = true
		}
	}

	// Check if any filters require a join
	for _, filter := range filters {
		if strings.Contains(filter.Field, ".") {
			relatedTable := strings.Split(filter.Field, ".")[0]
			if !joinedTables[relatedTable] {
				db = db.Joins(fmt.Sprintf("LEFT JOIN %s ON %s.id = %s.%s_id", relatedTable, relatedTable, mainTableName, getSingularObjectName(relatedTable)))
				joinedTables[relatedTable] = true
			}
		}
	}

	return db
}

func doFilter(tx *gorm.DB, tableName string, filter []types.FilterInput) (*gorm.DB, error) {
	repo := NewRepository(tx)

	fieldHandlers := map[string]func(f types.FilterInput) (*gorm.DB, error){
		"users.some_weird_field": func(f types.FilterInput) (*gorm.DB, error) {
			// this is how you can build filters where some data processing is required
			// this is a contrived example, but you can imagine a scenario where you need to do some data processing
			// before you can filter on a field
			users, err := repo.User.FuzzyFindBy("email", f.Value)
			if err != nil {
				return nil, err
			}

			return tx.Where("some_weird_field IN (?)", users), nil
		},
	}

	for _, f := range filter {
		handler, ok := fieldHandlers[f.Field]
		if ok {
			var err error
			tx, err = handler(f)
			if err != nil {
				return nil, err
			}
		} else {
			return tx.Where(fmt.Sprintf("%s IN (?)", f.Field), f.Value), nil
		}
	}

	return tx, nil
}

func getSingularObjectName(tableName string) string {
	if strings.HasSuffix(tableName, "ies") {
		return strings.TrimSuffix(tableName, "ies") + "y"
	} else if strings.HasSuffix(tableName, "s") {
		return strings.TrimSuffix(tableName, "s")
	}
	return tableName
}

func encodeCursor(id uint) string {
	return base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(int(id))))
}
