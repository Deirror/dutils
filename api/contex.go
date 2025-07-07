package api

// Avoids string keys to prevent collisions.
type CtxKey string

const ReqIDKey CtxKey = "requestID"
