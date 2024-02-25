package some_service

import (
	"context"
	"fmt"
)

func (s *Service) SomeMethod(ctx context.Context) error {
	fmt.Println("doing some logic!")

	return nil
}
