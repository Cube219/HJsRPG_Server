package main

import bf "golang.org/x/crypto/blowfish"
import "fmt"
import "reflect"

func main() {
    v := []byte("HiHiHiHi")

    c, err := bf.NewCipher([]byte("Key")) 
    if err != nil {
        fmt.Println(err) 
    }

    fmt.Println(string(v[:8]))  

    var enc [8]byte
    c.Encrypt(enc[0:], v)

    fmt.Println(string(enc[:8])) 
    
    fmt.Println(reflect.TypeOf(enc)) 
    fmt.Println(reflect.TypeOf(enc[:])) 
    

    c2, err := bf.NewCipher([]byte("Key")) 
    if err != nil {
        fmt.Println(err) 
    }

    var dec [8]byte
    c2.Decrypt(dec[0:], enc[0:])

    fmt.Println(string(dec[:8]))  
}