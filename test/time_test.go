package test

import (
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	time1 := time.Now()
	time2 := time1.Add(2 * time.Minute)
	calculate := time2.Sub(time1)
	logrus.Infof("Result calculate: %v", calculate)
	isThreshold := calculate == 2*time.Minute
	logrus.Infof("Is Threshold value: %v", isThreshold)
}
