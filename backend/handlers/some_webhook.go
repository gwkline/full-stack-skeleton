package handlers

import (
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gwkline/full-stack-skeleton/backend/pkg/workers"
	"github.com/hibiken/asynq"
)

func (h *Handler) SomeWebhookHandler(wg *sync.WaitGroup) gin.HandlerFunc {
	return func(c *gin.Context) {
		if wg != nil {
			defer wg.Done()
		}

		c.JSON(200, gin.H{"message": "Webhook received"})

		messageURL := c.PostForm("message-url")
		if messageURL == "" {
			fmt.Println("no stored email found")
			return
		}

		task, err := workers.NewProcessSomeDataTask(messageURL, 0)
		if err != nil {
			fmt.Printf("failed creating new teddy task: %v\n", err)
			return
		}

		_, err = h.Service.Queue.Enqueue(c, task, asynq.Retention(24*time.Hour))
		if err != nil {
			fmt.Printf("failed enqueuing task: %v", err)
			return
		}
	}
}
