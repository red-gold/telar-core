package utils

import (
	"strconv"
	"testing"
	"time"
)

// TestMomentToTime Covert moment unix to time.Time
func TestMomentToTime(t *testing.T) {
	currentMoment := UTCNowUnix()
	t.Logf("currentMoment: %d", currentMoment)

	currentTime := UTCUnixToTime(currentMoment)
	t.Logf("currentTime: %v", currentTime)

	var convertedTime time.Time
	var err error

	t.Logf("currentMoment in string: %s", strconv.FormatInt(currentMoment, 10))

	convertedTime, err = MomentToTime(strconv.FormatInt(currentMoment, 10))
	if err != nil {
		t.Errorf("error converting moment unix time string to time.Time : %s", err.Error())
		t.Fail()
	}
	convertedTime.Equal(currentTime)
}

// TestUTCNowUnix Get moment UTC NOW
func TestUTCNowUnix(t *testing.T) {

}

// TestUTCUnixToTime Get UTC unix time in go time
func TestUTCUnixToTime(t *testing.T) {

}

// TestIsTimeExpired
func TestIsTimeExpired(t *testing.T) {
	currentMoment := UTCNowUnix()
	t.Logf("currentMoment: %d", currentMoment)

	currentTime := UTCUnixToTime(currentMoment)
	t.Logf("currentTime: %v", currentTime)

	// Token not expired
	count := 360
	beforeMoment := currentTime.Add(time.Duration(-count) * time.Second)
	t.Logf("beforeMoment: %v", beforeMoment)

	status := IsTimeExpired(TimeUnix(beforeMoment), 3600)
	if status {
		t.Errorf("error expire time. Token is not expired! -> %v  ", beforeMoment)
		t.Fail()
	}

	// Token is expired
	count = 8601
	beforeMoment = currentTime.Add(time.Duration(-count) * time.Second)
	t.Logf("beforeMoment: %v", beforeMoment)

	status = IsTimeExpired(TimeUnix(beforeMoment), 3600)
	if !status {
		t.Errorf("error expire time. Token is expired! -> %v  ", beforeMoment)
		t.Fail()
	}
}
