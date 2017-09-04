package fxsymbols_test

import (
	"testing"
	"strings"
	"fxparser"
)

func TestFuncEnv(t *testing.T) {
	input := `func line(int fuu, int hello,  int world){
			world = 32;
		}
		func function(int param1, int hello){
			int fuu;
			param1 = 543;
			line(342+fuu, param1*(2+hello));
		}`
	p := fxparser.ParserFromReader("Test_Env", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Error: %s\n", err)
	}

	input = `func line(int fuu, int hello,  int world){
			world = 32;
		}
		func line(int param1, int hello){
			int fuu;
			param1 = 543;
		}`
	p = fxparser.ParserFromReader("Test_Env", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error")
	}

	input = `func line(int fuu, int hello,  int world){
			world = 32;
		}
		func line2(int param1, int hello){
			fuu = 0x66;
			param1 = 543;
		}`
	p = fxparser.ParserFromReader("Test_Env", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error")
	}

	input = `func line(int fuu, int hello,  int world){
			world = 32;
		}
		func line2(int param1, int hello){
			fuu = 0x66;
			param1 = 543;
			not_a_function();
		}`
	p = fxparser.ParserFromReader("Test_Env", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error")
	}

	input = `func line(int fuu, int hello,  int world){
			world = 32;
		}
		func function(int param1){
			int fuu;
			param1 = 543;
			line(342+fuu, param1*(2+hello));
		}`
	p = fxparser.ParserFromReader("Test_Env", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error")
	}
}

func TestLoopEnv(t *testing.T) {
	input := `func line(int fuu, int hello,  int world){
			int out_loop;
			int i;
			iter (i := fuu; hello, world) {
				int inside_loop;
				inside_loop = 43;
				out_loop = 22;
			}
			int inside_loop;
		}`
	p := fxparser.ParserFromReader("Test_Env", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Error: %s\n", err)
	}

	input = `func line(int fuu, int hello,  int world){
			int out_loop;
			int i;
			iter (i := fuu; hello, world) {
				int inside_loop;
				inside_loop = 43;
				out_loop = 22;
			}
			inside_loop = 33;
		}`
	p = fxparser.ParserFromReader("Test_Env", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error")
	}

	input = `func line(int fuu, int hello,  int world){
			int out_loop;
			int i;
			iter (i := fuu; hello, world) {
				int inside_loop;
				int out_loop:
			}
			int inside_loop;
		}`
	p = fxparser.ParserFromReader("Test_Env", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error")
	}

	input = `func line(int fuu, int hello,  int world){
			int out_loop;
			iter (i := fuu; hello, world) {
				int inside_loop;
				int out_loop:
			}
			int inside_loop;
		}`
	p = fxparser.ParserFromReader("Test_Env", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error")
	}
}

func TestIfEnv(t *testing.T) {
	input := `func line(int fuu, int hello,  int world){
			int out_if;
			if (!False) {
				out_if = 43;
				int in_if;
				in_if = 22;
			}else{
				out_if = 43;
				int in_if;
				in_if = 22;
			}
			int in_if;
		}`
	p := fxparser.ParserFromReader("Test_Env", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Error: %s\n", err)
	}

	input = `func line(int fuu, int hello,  int world){
			int out_if;
			if (!False) {
				int out_if;
			}else{
				out_if = 43;
				int in_if;
				in_if = 22;
			}
			int in_if;
		}`
	p = fxparser.ParserFromReader("Test_Env", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error")
	}

	input = `func line(int fuu, int hello,  int world){
			int out_if;
			if (!False) {
				int in_if;
			}else{
				in_if = 34;
			}
			int in_if;
		}`
	p = fxparser.ParserFromReader("Test_Env", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error")
	}

	input = `func line(int fuu, int hello,  int world){
			int out_if;
			if (!False) {
				int in_if;
			}else{
			}
			in_if = 34;
		}`
	p = fxparser.ParserFromReader("Test_Env", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error")
	}

	input = `func line(int fuu, int hello,  int world){
			int out_if;
			if (!False) {
				int in_if;
			}else{
				int in_if;
				in_if = 34;
			}
			in_if = 34;
		}`
	p = fxparser.ParserFromReader("Test_Env", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error")
	}

	input = `func line(int fuu, int hello,  int world){
			int out_if;
			if (!False) {
			}else{
				int out_if;
			}
			in_if = 34;
		}`
	p = fxparser.ParserFromReader("Test_Env", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error")
	}
}

func TestBuiltIn(t *testing.T) {
	input := `func line(int fuu, int hello,  int world){
			circle();
			rect();
		}`
	p := fxparser.ParserFromReader("Test_Env", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Error: %s\n", err)
	}

	input = `func circle(int fuu, int hello,  int world){
			int out_if;
		}`
		p = fxparser.ParserFromReader("Test_Env", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error")
	}

	input = `func rect(int fuu, int hello,  int world){
			int out_if;
		}`
		p = fxparser.ParserFromReader("Test_Env", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error")
	}
}