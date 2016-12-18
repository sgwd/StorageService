package data

import (
	"time"
	"fmt"
	"strconv"
)

type Bucket struct {
	Name string
}

func NewBucket(name string) *Bucket {
	return &Bucket{
		Name: name}
}

func (c Bucket) Print() {
	fmt.Println("Bucket name: "+ c.Name)
}
