package search

import (
	"context"

	"github.com/gwkline/full-stack-skeleton/backend/types"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func (s *Service) Apples(ctx context.Context, input *types.ConnectionInput) (*types.AppleConnection, error) {
	nrTx := newrelic.FromContext(ctx)
	defer nrTx.StartSegment("Apples").End()

	sortBy, direction, query, filter, limit, decodedCursor, err := s.prepareInputs(input)
	if err != nil {
		return nil, err
	}

	tx := s.Repository.DB
	if query != "" {
		tx = tx.Where("variety ILIKE ?",
			"%"+query+"%")
	}

	apples, pageInfo, err := s.Repository.Apple.Connection(tx, "apples", filter, sortBy, direction, limit, decodedCursor)
	if err != nil {
		return nil, err
	}

	edges := s.prepareAppleEdges(apples, direction, decodedCursor, limit)

	bcc := types.AppleConnection{
		Edges:    edges,
		PageInfo: pageInfo,
	}

	return &bcc, nil
}

func (s *Service) prepareAppleEdges(apples []*types.Apple, direction types.AscOrDesc, decodedCursor int, limit int) (edges []*types.AppleEdge) {
	for _, v := range apples {
		cursor := encodeCursor(v.ID)
		edges = append(edges, &types.AppleEdge{
			Cursor: cursor,
			Node:   v,
		})
	}

	return edges
}
