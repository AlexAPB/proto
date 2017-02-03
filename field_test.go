package proto

import "testing"

func TestField(t *testing.T) {
	proto := `repeated foo.bar lots =1 [option1=a, option2=b, option3="happy"];`
	p := newParserOn(proto)
	f := newNormalField()
	err := f.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := f.Repeated, true; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Type, "foo.bar"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Name, "lots"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(f.Options), 3; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[0].Name, "option1"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[0].Constant.Source, "a"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[1].Name, "option2"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[1].Constant.Source, "b"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[2].Constant.Source, "happy"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestFieldSimple(t *testing.T) {
	proto := `string optional_string_piece = 24 [ctype=STRING_PIECE];`
	p := newParserOn(proto)
	f := newNormalField()
	err := f.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := f.Type, "string"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Name, "optional_string_piece"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Sequence, 24; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(f.Options), 1; got != want {
		t.Fatalf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[0].Name, "ctype"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Options[0].Constant.Source, "STRING_PIECE"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestFieldSyntaxErrors(t *testing.T) {
	for i, each := range []string{
		`repeatet foo.bar lots = 1;`,
		`string lots === 1;`,
	} {
		f := newNormalField()
		if f.parse(newParserOn(each)) == nil {
			t.Errorf("uncaught syntax error in test case %d, %#v", i, f)
		}
	}
}

func TestMapField(t *testing.T) {
	proto := ` <string, Project> projects = 3;`
	p := newParserOn(proto)
	f := newMapField()
	err := f.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := f.KeyType, "string"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Type, "Project"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Name, "projects"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := f.Sequence, 3; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestOptionalWithOption(t *testing.T) {
	proto := `optional int32 default_int32    = 61 [default =  41    ];`
	p := newParserOn(proto)
	f := newNormalField()
	err := f.parse(p)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := f.Sequence, 61; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	o := f.Options[0]
	if got, want := o.Name, "default"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := o.Constant.Source, "41"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
