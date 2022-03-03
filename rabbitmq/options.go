package rabbitmq

import (
	"context"
	"github.com/tx7do/kratos-transport/common"
	"time"
)

type durableQueueKey struct{}
type headersKey struct{}
type queueArgumentsKey struct{}
type prefetchCountKey struct{}
type prefetchGlobalKey struct{}
type exchangeKey struct{}
type requeueOnErrorKey struct{}
type deliveryMode struct{}
type priorityKey struct{}
type contentType struct{}
type contentEncoding struct{}
type correlationID struct{}
type replyTo struct{}
type expiration struct{}
type messageID struct{}
type timestamp struct{}
type typeMsg struct{}
type userID struct{}
type appID struct{}
type externalAuth struct{}
type durableExchange struct{}

// DurableQueue creates a durable queue when subscribing.
func DurableQueue() common.SubscribeOption {
	return setSubscribeOption(durableQueueKey{}, true)
}

// DurableExchange is an option to set the Exchange to be durable
func DurableExchange() common.Option {
	return setBrokerOption(durableExchange{}, true)
}

// Headers adds headers used by the headers exchange
func Headers(h map[string]interface{}) common.SubscribeOption {
	return setSubscribeOption(headersKey{}, h)
}

// QueueArguments sets arguments for queue creation
func QueueArguments(h map[string]interface{}) common.SubscribeOption {
	return setSubscribeOption(queueArgumentsKey{}, h)
}

// RequeueOnError calls Nack(muliple:false, requeue:true) on amqp delivery when handler returns error
func RequeueOnError() common.SubscribeOption {
	return setSubscribeOption(requeueOnErrorKey{}, true)
}

// ExchangeName is an option to set the ExchangeName
func ExchangeName(e string) common.Option {
	return setBrokerOption(exchangeKey{}, e)
}

// PrefetchCount ...
func PrefetchCount(c int) common.Option {
	return setBrokerOption(prefetchCountKey{}, c)
}

// PrefetchGlobal creates a durable queue when subscribing.
func PrefetchGlobal() common.Option {
	return setBrokerOption(prefetchGlobalKey{}, true)
}

// DeliveryMode sets a delivery mode for publishing
func DeliveryMode(value uint8) common.PublishOption {
	return setPublishOption(deliveryMode{}, value)
}

// Priority sets a priority level for publishing
func Priority(value uint8) common.PublishOption {
	return setPublishOption(priorityKey{}, value)
}

// ContentType sets a property MIME content type for publishing
func ContentType(value string) common.PublishOption {
	return setPublishOption(contentType{}, value)
}

// ContentEncoding sets a property MIME content encoding for publishing
func ContentEncoding(value string) common.PublishOption {
	return setPublishOption(contentEncoding{}, value)
}

// CorrelationID sets a property correlation ID for publishing
func CorrelationID(value string) common.PublishOption {
	return setPublishOption(correlationID{}, value)
}

// ReplyTo sets a property address to to reply to (ex: RPC) for publishing
func ReplyTo(value string) common.PublishOption {
	return setPublishOption(replyTo{}, value)
}

// Expiration sets a property message expiration spec for publishing
func Expiration(value string) common.PublishOption {
	return setPublishOption(expiration{}, value)
}

// MessageId sets a property message identifier for publishing
func MessageId(value string) common.PublishOption {
	return setPublishOption(messageID{}, value)
}

// Timestamp sets a property message timestamp for publishing
func Timestamp(value time.Time) common.PublishOption {
	return setPublishOption(timestamp{}, value)
}

// TypeMsg sets a property message type name for publishing
func TypeMsg(value string) common.PublishOption {
	return setPublishOption(typeMsg{}, value)
}

// UserID sets a property user id for publishing
func UserID(value string) common.PublishOption {
	return setPublishOption(userID{}, value)
}

// AppID sets a property application id for publishing
func AppID(value string) common.PublishOption {
	return setPublishOption(appID{}, value)
}

func ExternalAuth() common.Option {
	return setBrokerOption(externalAuth{}, ExternalAuthentication{})
}

type subscribeContextKey struct{}

// SubscribeContext set the context for common.SubscribeOption
func SubscribeContext(ctx context.Context) common.SubscribeOption {
	return setSubscribeOption(subscribeContextKey{}, ctx)
}

type ackSuccessKey struct{}

// AckOnSuccess will automatically acknowledge messages when no error is returned
func AckOnSuccess() common.SubscribeOption {
	return setSubscribeOption(ackSuccessKey{}, true)
}
