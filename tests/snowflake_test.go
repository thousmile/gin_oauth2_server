package tests

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestSnowflake(t *testing.T) {
	worker, _ := snowflake.NewNode(1)
	for i := 0; i < 5; i++ {
		s := worker.Generate().String()
		fmt.Printf("len : %d  : %s \n", len(s), s)
	}
}

func TestRand(t *testing.T) {
	n := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000000)
	fmt.Println(strconv.Itoa(n))
}
