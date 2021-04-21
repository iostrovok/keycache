package keycache

import (
	"fmt"
	"testing"

	. "github.com/iostrovok/check"
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
