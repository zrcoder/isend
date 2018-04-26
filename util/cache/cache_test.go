package cache
 
import (
	"testing"
	"fmt"
)
 
type person struct {
	age int
	name string
}
 
func init() {
	Example()
}
 
func Example() {
	cache := NewWithCapacity(2)
 
	p1 := person{18, "Lily"}
	p2 := person{20, "Lucy"}
	if ok := cache.Add("p1", p1); !ok {
		fmt.Println("add error")
	}
	if ok := cache.Add("p2", p2); !ok {
		fmt.Println("add error")
	}
 
	p, found := cache.Search("p1")
	if found {
		pp, ok := p.(person)
		if ok {
			fmt.Println("pp ia a person, age:", pp.age, " name:", pp.name )
		} else {
			fmt.Println("pp is not a person~")
		}
	}
}
func TestAll(t *testing.T)  {
	cache := NewWithCapacity(2)
	if cache.capacity != 2 {
		t.Error("wrong capacity")
	}
 
	if ok := cache.Add("kkk", "vvv"); !ok {
		t.Error("add error")
	}
	if ok := cache.Add("kkk", "vvv"); ok {
		t.Error("add error: kkk is already in cache, but added again")
	}
	t.Log(cache.items)
 
	p1 := person{18, "Lily"}
	p2 := person{20, "Lucy"}
	if ok := cache.Add("p1", p1); !ok {
		t.Error("add error")
	}
	if ok := cache.Add("p2", p2); !ok {
		t.Error("add error")
	}
	t.Log(cache.items)
 
	cache.Add("kkk", "vvv")
	t.Log(cache.items)
 
	if ok := cache.Remove("kkk"); !ok {
		t.Error("remove error: kkk is in cache, but cannot remove")
	}
	if ok := cache.Remove("abc"); ok {
		t.Error("remove error: abs isn't in cache, but removed")
	}
	t.Log(cache.items)
 
	if ok := cache.Replace("p1", "s1"); !ok {
		t.Log("p1 isn't in cache, so can't replace")
	}
	if ok := cache.Replace("p2", "s2"); !ok {
		t.Log("p2 isn't in cache so can't replace")
	}
	t.Log(cache.items)
 
	if v, foud := cache.Search("p1"); foud {
		t.Log(v)
	}
	if v, foud := cache.Search("p2"); foud {
		t.Log(v)
	}
	if v, foud := cache.Search("kkk"); foud {
		t.Log(v)
	}
	if v, foud := cache.Search("xxxx"); foud {
		t.Log(v)
	}
}
