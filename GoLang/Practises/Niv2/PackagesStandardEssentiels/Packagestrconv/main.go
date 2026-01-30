package main

import (
    "fmt"
    "strconv"
)

func main() {
    // String vers Int
    entier, err := strconv.Atoi("42")
    if err != nil {
        fmt.Println("Erreur:", err)
    }
    fmt.Println(entier + 10)  // 52
    
    // Int vers String
    texte := strconv.Itoa(42)
    fmt.Println(texte + "0")  // "420"
    
    // ParseInt (avec base)
    n64, _ := strconv.ParseInt("1010", 2, 64)  // Base 2 (binaire)
    fmt.Println(n64)  // 10
    
    hexVal, _ := strconv.ParseInt("FF", 16, 64)  // Base 16 (hexa)
    fmt.Println(hexVal)  // 255
    
    // ParseFloat
    f, _ := strconv.ParseFloat("3.14159", 64)
    fmt.Println(f)
    
    // ParseBool
    b, _ := strconv.ParseBool("true")
    fmt.Println(b)  // true
    
    // FormatInt
    binaire := strconv.FormatInt(42, 2)
    fmt.Println(binaire)  // "101010"
    
    hexa := strconv.FormatInt(255, 16)
    fmt.Println(hexa)  // "ff"
    
    // FormatFloat
    s := strconv.FormatFloat(3.14159, 'f', 2, 64)
    fmt.Println(s)  // "3.14"
    
    // Quote (échapper string)
    quoted := strconv.Quote("Hello\nWorld")
    fmt.Println(quoted)  // "Hello\nWorld"
}