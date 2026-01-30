package calculatrice

import (
    "testing"
)

// Test simple
func TestAdditionner(t *testing.T) {
    resultat := Additionner(2, 3)
    attendu := 5
    
    if resultat != attendu {
        t.Errorf("Additionner(2, 3) = %d; attendu %d", resultat, attendu)
    }
}

// Table-driven test (pattern Go standard)
func TestAdditionnerTable(t *testing.T) {
    tests := []struct {
        name     string
        a        int
        b        int
        attendu  int
    }{
        {"positifs", 2, 3, 5},
        {"négatifs", -2, -3, -5},
        {"mixte", -2, 3, 1},
        {"zéro", 0, 5, 5},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            resultat := Additionner(tt.a, tt.b)
            if resultat != tt.attendu {
                t.Errorf("Additionner(%d, %d) = %d; attendu %d", 
                    tt.a, tt.b, resultat, tt.attendu)
            }
        })
    }
}

// Test avec erreur
func TestDiviser(t *testing.T) {
    // Cas normal
    resultat, err := Diviser(10, 2)
    if err != nil {
        t.Errorf("Erreur inattendue: %v", err)
    }
    if resultat != 5.0 {
        t.Errorf("Diviser(10, 2) = %f; attendu 5.0", resultat)
    }
    
    // Cas d'erreur
    _, err = Diviser(10, 0)
    if err == nil {
        t.Error("Division par zéro devrait retourner une erreur")
    }
}

// Subtests
func TestOperations(t *testing.T) {
    t.Run("addition", func(t *testing.T) {
        if Additionner(1, 1) != 2 {
            t.Error("1 + 1 devrait être 2")
        }
    })
    
    t.Run("soustraction", func(t *testing.T) {
        if Soustraire(5, 3) != 2 {
            t.Error("5 - 3 devrait être 2")
        }
    })
}