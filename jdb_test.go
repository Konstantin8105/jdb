package jdb_test

import (
	"os"
	"testing"

	"github.com/Konstantin8105/jdb"
)

type Animal struct {
	Name       string
	AmountLegs int
}

func Test(t *testing.T) {
	filename := "animals.db"

	if err := os.Remove(filename); err != nil {
		err = nil //ignore
	}

	db, err := jdb.Open[Animal](filename)
	if err != nil {
		t.Fatal(err)
	}

	as := []Animal{
		Animal{Name: "Bird", AmountLegs: 2},
		Animal{Name: "Elephan", AmountLegs: 4},
	}

	for i := range as {
		db.Add(as[i])
	}

	isSame := func(vs []Animal) {
		if len(vs) != 2 {
			t.Fatalf("not valid size: %d", len(vs))
		}
		for i := range vs {
			a := vs[i]
			if a.Name != as[i].Name {
				t.Fatalf("not Animal name: %#v", vs)
			}
			if a.AmountLegs != as[i].AmountLegs {
				t.Fatalf("not Animal legs: %#v", vs)
			}
		}
	}

	isSame(db.Get())

	if err := db.Write(); err != nil {
		t.Fatal(err)
	}

	db2, err := jdb.Open[Animal](filename)
	if err != nil {
		t.Fatal(err)
	}

	isSame(db2.Get())
}
