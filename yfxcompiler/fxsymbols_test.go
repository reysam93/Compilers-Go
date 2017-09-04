package yfxcompiler

import (
	"strings"
	"testing"
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
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Syms_Env")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}
	input = `func line(int fuu, int hello,  int world){
			world = 32;
		}
		func line(int param1, int hello){
			int fuu;
			param1 = 543;
		}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Syms_Env")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = `func line(int fuu, int hello,  int world){
			world = 32;
		}
		func line2(int param1, int hello){
			fuu = 0x66;
			param1 = 543;
			not_defined_var = 43;
		}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Syms_Env")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
	input = `func line(int fuu, int hello,  int world){
			world = 32;
		}
		func line2(int param1, int hello){
			fuu = 0x66;
			param1 = 543;
			not_a_function();
		}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Syms_Env")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
	input = `func line(int fuu, int hello,  int world){
			world = 32;
		}
		func function(int param1){
			int fuu;
			param1 = 543;
			line(342+fuu, param1*(2+hello));
		}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Syms_Env")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
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
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Syms_Env")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
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
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Syms_Env")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
	input = `func line(int fuu, int hello,  int world){
			int out_loop;
			iter (i := fuu; hello, world) {
				int inside_loop;
				inside_loop = 43;
				out_loop = 22;
			}
			int inside_loop;
		}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Syms_Env")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
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
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Syms_Env")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = `func line(int fuu, int hello,  int world){
			int out_loop;
			iter (i := fuu; hello, world) {
				int inside_loop;
				int out_loop:
			}
			int inside_loop;
		}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Syms_Env")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
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
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Syms_Env")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
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
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Syms_Env")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
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
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Syms_Env")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = `func line(int fuu, int hello,  int world){
			int out_if;
			if (!False) {
				int in_if;
			}else{
			}
			in_if = 34;
		}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Syms_Env")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
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
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Syms_Env")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = `func line(int fuu, int hello,  int world){
			int out_if;
			if (!False) {
			}else{
				int out_if;
			}
			in_if = 34;
		}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Syms_Env")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
}

func TestBuiltIn(t *testing.T) {
	input := `func line(int fuu, int hello,  int world){
			circle();
			rect();
		}`
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Syms_Env")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}

	input = `func circle(int fuu, int hello,  int world){
			int out_if;
		}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Syms_Env")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = `func rect(int fuu, int hello,  int world){
			int out_if;
		}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Syms_Env")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
}