package generator

type SnowflakeManager interface {
	GenerateID() int64
}
