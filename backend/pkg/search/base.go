package search

import (
	"encoding/base64"
	"strconv"

	"github.com/gwkline/full-stack-skeleton/backend/pkg/repo"
	"github.com/gwkline/full-stack-skeleton/backend/types"
)

type Service struct {
	Repository *repo.Repository
}

func Init(repo *repo.Repository) *Service {
	return &Service{
		Repository: repo,
	}
}

func (s *Service) prepareInputs(input *types.ConnectionInput) (sortBy string, direction types.AscOrDesc, searchString string, filter []types.FilterInput, limit int, decodedCursor int, err error) {
	sortBy, direction = s.getSortInputs(input.SortBy, input.Direction)
	searchString = s.getSearchString(input.Query)
	filter = s.getFilter(input.Filters)
	limit = s.getLimit(input.Limit)
	decodedCursor, err = s.getDecodedCursor(input.After, direction)
	return
}

func (s *Service) getSortInputs(sortByInput *string, directionInput *types.AscOrDesc) (string, types.AscOrDesc) {
	sortBy := "id"
	if sortByInput != nil && *sortByInput != "" {
		sortBy = *sortByInput
	}

	direction := types.AscOrDescAsc
	if directionInput != nil {
		direction = *directionInput
	}

	return sortBy, direction
}

func (s *Service) getSearchString(query *string) string {
	if query != nil {
		return *query
	}
	return ""
}

func (s *Service) getFilter(filters []*types.FilterInput) []types.FilterInput {
	result := []types.FilterInput{}
	for _, item := range filters {
		result = append(result, *item)
	}
	return result
}

func (s *Service) getLimit(limit *int) int {
	if limit != nil {
		return *limit
	}
	return 10
}

func (s *Service) getDecodedCursor(after *string, direction types.AscOrDesc) (int, error) {
	if after != nil && *after != "" {
		return decodeCursor(*after)
	}

	if (after == nil || *after == "") && direction == types.AscOrDescDesc {
		return 9999999999999999, nil
	}

	return 0, nil
}

func decodeCursor(encodedCursor string) (int, error) {
	// return strconv.Atoi(encodedCursor)
	b, err := base64.StdEncoding.DecodeString(encodedCursor)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(b))
}

func encodeCursor(id uint) string {
	// return strconv.Itoa(int(id))
	return base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(int(id))))
}
