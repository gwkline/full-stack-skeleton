package resolver

import (
	"fmt"
	"time"

	"github.com/gwkline/full-stack-skeleton/backend/graph/loaders"
	"github.com/gwkline/full-stack-skeleton/backend/pkg/repo"
	"github.com/gwkline/full-stack-skeleton/backend/types"
)

func NewLoaders(db *repo.Repository) *loaders.Loaders {
	apple := loaders.NewAppleByIDLoader(loaders.AppleByIDLoaderConfig{
		Fetch: func(ids []uint) ([]*types.Apple, []error) {
			apples, err := db.Apple.ListBy([]types.Filter{{Key: "id", Value: ids}})
			if err != nil {
				return nil, []error{err}
			}

			appleMap := make(map[uint]*types.Apple, len(apples))
			for _, apple := range apples {
				appleMap[apple.ID] = apple
			}

			orderedApples := make([]*types.Apple, len(ids))
			for i, id := range ids {
				apple, ok := appleMap[id]
				if !ok {
					return nil, []error{fmt.Errorf("no apple with ID %d", id)}
				}
				orderedApples[i] = apple
			}

			return orderedApples, nil
		},
		MaxBatch: 100,
		Wait:     1 * time.Millisecond,
	})

	return &loaders.Loaders{
		AppleByID: apple,
	}
}
