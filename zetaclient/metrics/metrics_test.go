package metrics

import (
	"fmt"
	. "gopkg.in/check.v1"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type MetricsSuite struct {
	m *Metrics
}

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&MetricsSuite{})

func (ms *MetricsSuite) SetUpSuite(c *C) {
	m, err := NewMetrics()
	c.Assert(err, IsNil)
	m.Start()
	ms.m = m
}

func (ms *MetricsSuite) TestMetrics(c *C) {
	ms.m.RegisterCounter("cnt1", "help to cnt1")
	Counters["cnt1"].Inc()
	time.Sleep(1 * time.Second)
	res, err := http.Get("http://127.0.0.1:8886/metrics")
	c.Assert(err, IsNil)
	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(out))
}
