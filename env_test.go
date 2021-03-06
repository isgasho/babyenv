package babyenv

import (
	"os"
	"strconv"
	"testing"
)

func TestParse(t *testing.T) {
	type config struct {
		A bool   `env:"A"`
		B string `env:"B"`
		C int    `env:"C"`
		D []byte `env:"D"`
		E int64  `env:"E"`
	}

	a := true
	b := "xxx"
	c := 16
	d := []byte("yyy")
	var e int64 = 64

	os.Setenv("A", strconv.FormatBool(a))
	os.Setenv("B", b)
	os.Setenv("C", strconv.FormatInt(int64(c), 10))
	os.Setenv("D", string(d))
	os.Setenv("E", strconv.FormatInt(e, 10))

	var cfg config
	if err := Parse(&cfg); err != nil {
		t.Errorf("error while parsing: %v", err)
		return
	}

	if !cfg.A {
		t.Errorf("failed parsing bool; expected %#v, got %#v", a, cfg.A)
	}
	if cfg.B != b {
		t.Errorf("failed parsing string; expected %#v, got %#v", b, cfg.B)
	}
	if cfg.C != c {
		t.Errorf("failed parsing int; expected %#v, got %#v", c, cfg.C)
	}
	if cfg.D == nil {
		t.Errorf("failed parsing byte[]; expected %#v, got nil", d)
	} else if string(cfg.D) != string(d) {
		t.Errorf("failed parsing []byte; expected %#v, got %#v", d, cfg.D)
	}
	if cfg.E != e {
		t.Errorf("failed parsing int64; expected %#v, got %#v", c, cfg.E)
	}
}

func TestParseWithDefaults(t *testing.T) {
	type config struct {
		A bool   `env:"A" default:"true"`
		B string `env:"B" default:"xxx"`
		C int    `env:"C" default:"16"`
		D []byte `env:"D" default:"yyy"`
		E int64  `env:"E" default:"64"`
	}

	a := true
	b := "xxx"
	c := 16
	d := []byte("yyy")
	var e int64 = 64

	os.Unsetenv("A")
	os.Unsetenv("B")
	os.Unsetenv("C")
	os.Unsetenv("D")
	os.Unsetenv("E")

	var cfg config
	if err := Parse(&cfg); err != nil {
		t.Errorf("error while parsing: %v", err)
		return
	}

	if cfg.A != a {
		t.Errorf("failed parsing bool; expected %#v, got %#v", a, cfg.A)
	}
	if cfg.B != b {
		t.Errorf("failed parsing string; expected %#v, got %#v", b, cfg.B)
	}
	if cfg.C != c {
		t.Errorf("failed parsing int; expected %#v, got %#v", c, cfg.C)
	}
	if cfg.D == nil {
		t.Errorf("failed parsing byte[]; expected %#v, got nil", d)
	} else if string(cfg.D) != string(d) {
		t.Errorf("failed parsing []byte; expected %#v, got %#v", d, cfg.D)
	}
	if cfg.E != e {
		t.Errorf("failed parsing int64; expected %#v, got %#v", e, cfg.E)
	}
}

func TestParsePointers(t *testing.T) {
	type config struct {
		A *bool   `env:"A"`
		B *string `env:"B"`
		C *int    `env:"C"`
		D *[]byte `env:"D"`
		E *int64  `env:"E"`
	}

	a := true
	b := "xxx"
	c := 16
	d := []byte("yyy")
	var e int64 = 64

	os.Setenv("A", strconv.FormatBool(a))
	os.Setenv("B", b)
	os.Setenv("C", strconv.FormatInt(int64(c), 10))
	os.Setenv("D", string(d))
	os.Setenv("E", strconv.FormatInt(e, 10))

	var cfg config
	if err := Parse(&cfg); err != nil {
		t.Errorf("error while parsing: %v", err)
		return
	}

	if cfg.A == nil {
		t.Errorf("failed parsing *bool; expected %#v, got nil", a)
	} else if *cfg.A != a {
		t.Errorf("failed parsing *bool; expected %#v, got %#v", a, *cfg.A)
	}

	if cfg.B == nil {
		t.Errorf("failed parsing *string; expected %#v, got nil", b)
	} else if *cfg.B != b {
		t.Errorf("failed parsing *string; expected %#v, got %#v", b, *cfg.B)
	}

	if cfg.C == nil {
		t.Errorf("failed parsing *int; expected %#v, got nil", c)
	} else if *cfg.C != c {
		t.Errorf("failed parsing *int; expected %#v, got %#v", c, *cfg.C)
	}

	if cfg.D == nil {
		t.Errorf("failed parsing *[]byte; expected %#v, got nil", d)
	} else if string(*cfg.D) != string(d) {
		t.Errorf("failed parsing *[]byte; expected %#v, got %#v", d, *cfg.D)
	}

	if cfg.E == nil {
		t.Errorf("failed parsing *int64; expected %#v, got nil", e)
	} else if *cfg.E != e {
		t.Errorf("failed parsing *int64; expected %#v, got %#v", e, *cfg.E)
	}
}

func TestParsePointersWithDefaults(t *testing.T) {
	type config struct {
		A *bool   `env:"A" default:"true"`
		B *string `env:"B" default:"xxx"`
		C *int    `env:"C" default:"16"`
		D *[]byte `env:"D" default:"yyy"`
	}

	a := true
	b := "xxx"
	c := 16
	d := []byte("yyy")

	os.Unsetenv("A")
	os.Unsetenv("B")
	os.Unsetenv("C")
	os.Unsetenv("D")

	var cfg config
	if err := Parse(&cfg); err != nil {
		t.Errorf("error while parsing: %v", err)
		return
	}

	if cfg.A == nil {
		t.Errorf("failed parsing *bool; expected %#v, got nil", a)
	} else if *cfg.A != a {
		t.Errorf("failed parsing *bool; expected %#v, got %#v", a, *cfg.A)
	}

	if cfg.B == nil {
		t.Errorf("failed parsing *string; expected %#v, got nil", b)
	} else if *cfg.B != b {
		t.Errorf("failed parsing *string; expected %#v, got %#v", b, *cfg.B)
	}

	if cfg.C == nil {
		t.Errorf("failed parsing *int; expected %#v, got nil", c)
	} else if *cfg.C != c {
		t.Errorf("failed parsing *int; expected %#v, got %#v", c, *cfg.C)
	}

	if cfg.D == nil {
		t.Errorf("failed parsing *[]byte; expected %#v, got nil", d)
	} else if string(*cfg.D) != string(d) {
		t.Errorf("failed parsing *[]byte; expected %#v, got %#v", d, *cfg.D)
	}
}

func TestRequiredFlag(t *testing.T) {
	type config struct {
		A bool `env:"A,required"`
	}

	os.Unsetenv("A")

	var cfg config
	if err := Parse(&cfg); err == nil {
		t.Errorf("expected an error because of an unfulfilled 'require' flag")
		return
	}
}
