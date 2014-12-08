ckage main
import (
    "fmt"
)

type Contact struct {
    name string
    id int
}

func main() {
    var c Contact = Contact {name:"Dave", id:33}
    var p *Contact = &c

    fmt.Println(c)
    fmt.Println(*p)
    (*p).name = "Holly"
    p.id = 34
    fmt.Println(*p)
}
