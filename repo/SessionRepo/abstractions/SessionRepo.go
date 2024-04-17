package abstractions

type SessionRepo interface {
	Set(key, value any)
	Get(key any) any
	Delete(key any)
	SessionId() any
}
