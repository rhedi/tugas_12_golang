package main

import "fmt"
import "regexp"

func main()  {
  var makanankesukaan string
  fmt.Println("Masukan buah kesukaan anda :")
  fmt.Scanln(&makanankesukaan)

  var input, err = regexp.Compile(`[A-z]+`)

  if err !=nil{
    fmt.Println(err.Error())
    }
    var hasil = input.FindAllString(makanankesukaan, -1)
    fmt.Println(hasil)
    var cocok = input.MatchString(makanankesukaan)
    fmt.Println(cocok)
    var hitung = input.FindStringIndex(makanankesukaan)
    fmt.Println(hitung)
    var ganti = input.ReplaceAllString(makanankesukaan, " Semangka")
    fmt.Println(ganti)

}
