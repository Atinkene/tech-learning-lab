package main

import (
    "fmt"
    "strings"
)

func main() {
    texte := "Dakar, capitale du Sénégal"
    
    // Contient
    fmt.Println(strings.Contains(texte, "Dakar"))        // true
    fmt.Println(strings.ContainsAny(texte, "xyz"))       // false
    
    // Compte les occurrences
    fmt.Println(strings.Count("banana", "a"))            // 3
    
    // Commence/Termine par
    fmt.Println(strings.HasPrefix(texte, "Dakar"))       // true
    fmt.Println(strings.HasSuffix(texte, "Sénégal"))     // true
    
    // Index (position)
    fmt.Println(strings.Index(texte, "capitale"))        // 7
    fmt.Println(strings.LastIndex("banana", "a"))        // 5
    
    // Join (assembler)
    mots := []string{"Go", "est", "génial"}
    phrase := strings.Join(mots, " ")
    fmt.Println(phrase)  // "Go est génial"
    
    // Split (découper)
    parties := strings.Split("a,b,c,d", ",")
    fmt.Println(parties)  // [a b c d]
    
    // Replace (remplacer)
    nouveau := strings.Replace(texte, "Dakar", "Thiès", 1)
    fmt.Println(nouveau)
    
    // ReplaceAll
    fmt.Println(strings.ReplaceAll("banana", "a", "o"))  // bonono
    
    // ToLower / ToUpper
    fmt.Println(strings.ToLower("DAKAR"))                // dakar
    fmt.Println(strings.ToUpper("sénégal"))              // SÉNÉGAL
    
    // Trim (enlever espaces/caractères)
    fmt.Println(strings.TrimSpace("  hello  "))          // "hello"
    fmt.Println(strings.Trim("!!!hello!!!", "!"))        // "hello"
    
    // Repeat
    fmt.Println(strings.Repeat("Go", 3))                 // GoGoGo
    
    // Fields (split sur espaces)
    mots2 := strings.Fields("  un  deux   trois  ")
    fmt.Println(mots2)  // [un deux trois]
    
    // Builder (construction efficace de strings)
    var builder strings.Builder
    for i := 0; i < 5; i++ {
        builder.WriteString(fmt.Sprintf("Ligne %d\n", i))
    }
    fmt.Println(builder.String())
}