package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestGoRoutines(t *testing.T) {
	t.Run("GetTimeSinceStart", GetTimeSinceStartTest)
	t.Run("Greet", GreetTest)
	t.Run("SlowGreet", SlowGreetTest)
}

func GetTimeSinceStartTest(t *testing.T) {
	start = time.Now()
	expected := "- time(0s)"
	result := GetTimeSinceStart()
	fmt.Println("result", result)
	fmt.Println("expected", expected)
	if result == expected {
		t.Errorf("GetTimeSinceStart() = %s; want %s", result, expected)
	}
}

func GreetTest(t *testing.T) {
	done := make(chan bool)
	go Greet("Nice to meet you!", done)
	<-done
}

func SlowGreetTest(t *testing.T) {
	done := make(chan bool)
	go SlowGreet("Nice to meet you!", 1, done)
	<-done
}
