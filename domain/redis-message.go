package domain

// RedisMessage is the model used to describe messages coming from redis
type RedisMessage struct {
	key    string
	values map[string]interface{}
}

// NewRedisMessage is the constructor for redis messages
func NewRedisMessage(key string, values map[string]interface{}) *RedisMessage {
	return &RedisMessage{
		key:    key,
		values: values,
	}
}

// ID is the public getter for the field id
func (r *RedisMessage) ID() string {
	return r.values["id"].(string)
}

// Key is the public getter for the field key
func (r *RedisMessage) Key() string {
	return r.key
}

// Values is the public getter for the field values
func (r *RedisMessage) Values() map[string]interface{} {
	return r.values
}
