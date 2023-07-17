package mockfuncoes

type MockCalculadora struct{}

func (m MockCalculadora) Somar(a, b int) int {
	return 10
}
