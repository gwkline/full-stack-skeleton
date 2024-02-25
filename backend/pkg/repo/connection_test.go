package repo

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gwkline/full-stack-skeleton/backend/types"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestConnection(t *testing.T) {
	tests := []struct {
		name          string
		tableName     string
		filters       []types.FilterInput
		sortBy        string
		direction     types.AscOrDesc
		limit         int
		after         int
		mockSetup     func(sqlmock.Sqlmock)
		expectedError string
		expectedCount int
	}{
		{
			name:      "successful query without filters",
			tableName: "users",
			sortBy:    "id",
			direction: types.AscOrDescAsc,
			limit:     10,
			after:     0,
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(1, "John Doe").
					AddRow(2, "Jane Doe")
				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT count(*) FROM "users" WHERE "users"."deleted_at" IS NULL`)).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))
				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL AND users.id > $1 ORDER BY id ASC LIMIT $2`)).
					WillReturnRows(rows)

			},
			expectedCount: 2,
		},
		{
			name:      "query with error on execution",
			tableName: "users",
			sortBy:    "id",
			direction: types.AscOrDescAsc,
			limit:     10,
			after:     0,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT count(*) FROM "users" WHERE "users"."deleted_at" IS NULL`)).
					WillReturnError(gorm.ErrRecordNotFound)
			},
			expectedError: "failed to count: record not found",
		},
		{
			name:      "successful query with filters and descending sort",
			tableName: "users",
			filters: []types.FilterInput{
				{
					Field: "email",
					Value: []string{"test@test.com"},
				},
			},
			sortBy:    "id",
			direction: types.AscOrDescDesc,
			limit:     5,
			after:     0,
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "email"}).
					AddRow(3, "alice@Wonderland.com").
					AddRow(4, "bob@Builder.com")
				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT count(*) FROM "users" WHERE email IN ($1) AND "users"."deleted_at" IS NULL`)).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))
				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT * FROM "users" WHERE email IN ($1) AND "users"."deleted_at" IS NULL AND users.id < $2 ORDER BY id DESC LIMIT $3`)).
					WillReturnRows(rows)
			},
			expectedCount: 2,
		},
		{
			name:      "successful query with after parameter set",
			tableName: "users",
			sortBy:    "id",
			direction: types.AscOrDescAsc,
			limit:     5,
			after:     10, // Setting the after value to simulate pagination
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT users.id as sortByValue, users.id as id FROM "users" WHERE users.id = $1`,
				)).WillReturnRows(sqlmock.NewRows([]string{"sortByValue", "id"}).AddRow(5, 5))

				// Expect a count query first
				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT count(*) FROM "users" WHERE "users"."deleted_at" IS NULL`)).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(5))
				// Expect the main query with conditions applied for pagination
				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL AND users.id > $1 ORDER BY id ASC LIMIT $2`)).
					WithArgs(10, 6). // Note: limit is 5, but we fetch one extra row to determine if there's a next page
					WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
						AddRow(11, "User After 10").
						AddRow(12, "Another User"))
			},
			expectedCount: 5,
		},
		{
			name:      "successful query with sorting by apples.variety",
			tableName: "users",
			sortBy:    "apples.variety",
			direction: types.AscOrDescAsc,
			limit:     5,
			after:     10,
			mockSetup: func(mock sqlmock.Sqlmock) {

				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT apples.variety as sortByValue, users.id as id FROM "users" JOIN apples ON apples.id = users.apple_id WHERE users.id = $1`,
				)).WillReturnRows(sqlmock.NewRows([]string{"sortByValue", "id"}).AddRow(5, 5))

				rows := sqlmock.NewRows([]string{"id", "variety", "apple_id"}).
					AddRow(1, "User One", 1).
					AddRow(2, "User Two", 2)
				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT count(*) FROM "users" LEFT JOIN apples ON apples.id = users.apple_id WHERE "users"."deleted_at" IS NULL`)).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))
				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT "users"."id","users"."created_at","users"."updated_at","users"."deleted_at","users"."email","users"."password_hash","users"."role" FROM "users" LEFT JOIN apples ON apples.id = users.apple_id WHERE "users"."deleted_at" IS NULL AND users.id > $1 ORDER BY apples.variety ASC, users.id ASC LIMIT $2`)).
					WillReturnRows(rows)
			},
			expectedCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()

			db, mock := newMockDB()
			tt.mockSetup(mock)

			result, pageInfo, err := db.User.Connection(db.DB, tt.tableName, tt.filters, tt.sortBy, tt.direction, tt.limit, tt.after)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.expectedCount, pageInfo.Count)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
