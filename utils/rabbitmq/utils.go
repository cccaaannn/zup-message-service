package rabbitmq

import "fmt"

func GetQueueNameForUser(userId uint64) string {
	return fmt.Sprintf("user_message_queue_%d", userId)
}
