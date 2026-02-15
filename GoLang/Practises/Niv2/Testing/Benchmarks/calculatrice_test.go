package calculatrice

import "testing"

// Benchmark simple
func BenchmarkAdditionner(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Additionner(2, 3)
    }
}

// Benchmark avec setup
func BenchmarkMultiplier(b *testing.B) {
    // Setup (non chronométré)
    a, c := 123, 456
    
    // Reset le timer après setup
    b.ResetTimer()
    
    // Benchmark (chronométré)
    for i := 0; i < b.N; i++ {
        Multiplier(a, c)
    }
}

// Benchmark avec subtests
func BenchmarkOperations(b *testing.B) {
    b.Run("addition", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            Additionner(2, 3)
        }
    })
    
    b.Run("multiplication", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            Multiplier(2, 3)
        }
    })
}