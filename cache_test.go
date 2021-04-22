package keycache

import (
	"fmt"
	. "github.com/iostrovok/check"
	"sync"
	"testing"
)

type testSuite struct{}

var _ = Suite(&testSuite{})

func TestService(t *testing.T) { TestingT(t) }

func (s *testSuite) TestSyntax(c *C) {
	c.Assert(1, NotNil)
}

func (s *testSuite) TestEncode1(c *C) {

	id := 12345432
	data := &Data{
		MyLongData1: "asdasdadasdad-1",
		MyLongData2: "asdasdadasdad-2",
	}
	md5 := GetMD5(data)

	t := &MyTestItem{
		id:   id,
		Data: data,
		md5:  md5,
	}

	kc := New()
	c.Assert(kc.Set(t), IsNil)

	t2 := &MyTestItem{
		id:  id,
		md5: md5,
	}

	c.Assert(kc.Get(t2), IsNil)
	c.Assert(t2.Find, Equals, true)
	c.Assert(t2.Data, DeepEquals, data)
}

func (s *testSuite) TestEncodeMulti(c *C) {

	data := &Data{
		MyLongData1: "asdasdadasdad-1",
		MyLongData2: "asdasdadasdad-2",
	}
	md5 := GetMD5(data)
	kc := New()
	for i := 1; i < 1000; i++ {
		data.MyLongData1 = fmt.Sprintf("data-%d", i)
		t := &MyTestItem{
			id:   i,
			Data: data,
			md5:  md5,
		}
		c.Assert(kc.Set(t), IsNil)
	}

	for i := 1; i < 1000; i++ {
		t2 := &MyTestItem{
			id:  i,
			md5: md5,
		}

		c.Assert(kc.Get(t2), IsNil)
		c.Assert(t2.Find, Equals, true)
		c.Assert(t2.Data.MyLongData1, DeepEquals, fmt.Sprintf("data-%d", i))
	}
}

func (s *testSuite) TestSize(c *C) {

	maxSize := 1000

	// cache with max size
	kc := New(maxSize)

	wg := sync.WaitGroup{}
	count := 2
	for count > 0 {
		wg.Add(1)
		count--
		go func() {
			defer wg.Done()
			for i := 1; i < 100000; i++ {
				data := &Data{
					MyLongData1: fmt.Sprintf("data-%d", i),
				}
				t := &MyTestItem{
					id:   i,
					Data: data,
					md5:  GetMD5(data),
				}

				c.Assert(kc.Set(t), IsNil)
			}
		}()
	}

	wg.Wait()

	found := 0
	for i := 1; i < 100000; i++ {
		data := &Data{
			MyLongData1: fmt.Sprintf("data-%d", i),
		}
		t2 := &MyTestItem{
			id:  i,
			md5: GetMD5(data),
		}

		c.Assert(kc.Get(t2), IsNil)
		if t2.Find {
			found++
			c.Assert(t2.Data.MyLongData1, Equals, fmt.Sprintf("data-%d", i))
		}
	}

	c.Assert(found > maxSize/2 && found <= maxSize, Equals, true)
}
