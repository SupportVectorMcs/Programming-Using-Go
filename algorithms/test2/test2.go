package main

import (
    "fmt"
)

type Contact struct {
 name string
 id int
}

func setContactInfo(c *Contact) {
 c.name = "Holly Golightly"
 c.id = 101
}

func main() {
 var c Contact = Contact{name:"Dave", id:33}
 setContactInfo(&c)
 fmt.Println(c)
}
