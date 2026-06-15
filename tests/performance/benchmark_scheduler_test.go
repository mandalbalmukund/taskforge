package performance

import "testing"

func BenchmarkScheduler(
	b *testing.B,
) {

	for i := 0; i < b.N; i++ {

		// execute workflow
	}
}