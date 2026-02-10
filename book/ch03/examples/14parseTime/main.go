package main

 

import (

    "fmt"

    "os"

    "path/filepath"

    "time"

)

Вторая часть содержит следующий код:

func main() {

    var myTime string

    if len(os.Args) != 2 {

        fmt.Printf("usage: %s string\n", filepath.Base(os.Args[0]))

        os.Exit(1)

    }

 

    myTime = os.Args[1]

Последняя часть parseTime.go, в которой и происходит все волшебство, выглядит так:

    d, err := time.Parse("15:04", myTime)

    if err == nil {

        fmt.Println("Full:", d)

        fmt.Println("Time:", d.Hour(), d.Minute())

    } else {

        fmt.Println(err)

    }

}

