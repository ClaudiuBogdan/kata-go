package maxflow

import (
	"testing"
)

func TestFindMaxFlow(t *testing.T) {
	tests := []struct {
		name           string
		network        *FlowNetwork
		source         int
		sink           int
		expectedMaxFlow int
	}{
		{
			name:           "Simple network",
			network:        createSimpleNetwork(),
			source:         0,
			sink:           5,
			expectedMaxFlow: 23,
		},
		{
			name:           "Complex network",
			network:        createComplexNetwork(),
			source:         0,
			sink:           7,
			expectedMaxFlow: 28,
		},
		{
			name:           "Single path network",
			network:        createSinglePathNetwork(),
			source:         0,
			sink:           3,
			expectedMaxFlow: 5,
		},
		{
			name:           "Disconnected network",
			network:        createDisconnectedNetwork(),
			source:         0,
			sink:           3,
			expectedMaxFlow: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maxFlow := tt.network.FindMaxFlow(tt.source, tt.sink)
			if maxFlow != tt.expectedMaxFlow {
				t.Errorf("Expected max flow %d, but got %d", tt.expectedMaxFlow, maxFlow)
			}

			if !isValidFlow(tt.network, tt.source, tt.sink) {
				t.Errorf("Invalid flow in the network")
			}
		})
	}
}

func createSimpleNetwork() *FlowNetwork {
	fn := NewFlowNetwork(6)
	fn.AddEdge(0, 1, 16)
	fn.AddEdge(0, 2, 13)
	fn.AddEdge(1, 2, 10)
	fn.AddEdge(1, 3, 12)
	fn.AddEdge(2, 1, 4)
	fn.AddEdge(2, 4, 14)
	fn.AddEdge(3, 2, 9)
	fn.AddEdge(3, 5, 20)
	fn.AddEdge(4, 3, 7)
	fn.AddEdge(4, 5, 4)
	return fn
}

func createComplexNetwork() *FlowNetwork {
	fn := NewFlowNetwork(8)
	fn.AddEdge(0, 1, 10)
	fn.AddEdge(0, 2, 5)
	fn.AddEdge(0, 3, 15)
	fn.AddEdge(1, 2, 4)
	fn.AddEdge(1, 4, 9)
	fn.AddEdge(1, 5, 15)
	fn.AddEdge(2, 3, 4)
	fn.AddEdge(2, 5, 8)
	fn.AddEdge(3, 6, 16)
	fn.AddEdge(4, 5, 15)
	fn.AddEdge(4, 7, 10)
	fn.AddEdge(5, 6, 15)
	fn.AddEdge(5, 7, 10)
	fn.AddEdge(6, 2, 6)
	fn.AddEdge(6, 7, 10)
	return fn
}

func createSinglePathNetwork() *FlowNetwork {
	fn := NewFlowNetwork(4)
	fn.AddEdge(0, 1, 10)
	fn.AddEdge(1, 2, 5)
	fn.AddEdge(2, 3, 15)
	return fn
}

func createDisconnectedNetwork() *FlowNetwork {
	fn := NewFlowNetwork(4)
	fn.AddEdge(0, 1, 10)
	fn.AddEdge(2, 3, 5)
	return fn
}

func isValidFlow(fn *FlowNetwork, source, sink int) bool {
	// Check flow conservation
	for u := 0; u < fn.V; u++ {
		if u == source || u == sink {
			continue
		}
		inFlow := 0
		outFlow := 0
		for v := 0; v < fn.V; v++ {
			inFlow += fn.Flow[v][u]
			outFlow += fn.Flow[u][v]
		}
		if inFlow != outFlow {
			return false
		}
	}

	// Check capacity constraints
	for u := 0; u < fn.V; u++ {
		for v := 0; v < fn.V; v++ {
			if fn.Flow[u][v] < 0 || fn.Flow[u][v] > fn.Graph[u][v] {
				return false
			}
		}
	}

	return true
}
