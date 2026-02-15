package calculatrice

import (
    "os"
    "testing"
    "fmt"
)

// Helper function
func assertEqual(t *testing.T, got, want int) {
    t.Helper()  // Marque comme helper pour stack trace
    if got != want {
        t.Errorf("got %d, want %d", got, want)
    }
}

func TestAvecHelper(t *testing.T) {
    resultat := Additionner(2, 3)
    assertEqual(t, resultat, 5)
}

// Setup/Teardown pour tous les tests
func TestMain(m *testing.M) {
    // Setup global
    fmt.Println("Setup avant tous les tests")
    
    // Exécuter les tests
    code := m.Run()
    
    // Teardown global
    fmt.Println("Cleanup après tous les tests")
    
    os.Exit(code)
}

// Cleanup pour un test spécifique
func TestAvecCleanup(t *testing.T) {
    // Setup
    fichier, _ := os.Create("temp.txt")
    
    // Cleanup (exécuté même si le test panic)
    t.Cleanup(func() {
        fichier.Close()
        os.Remove("temp.txt")
    })
    
    // Test...
}