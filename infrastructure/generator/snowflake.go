package generator

func (m *Module) GenerateID() int64 {
	return m.snode.Snode.Generate().Int64()
}
