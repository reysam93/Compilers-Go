package yfxcompiler

import (
	"strings"
	"testing"
)

func TestFunctions(t *testing.T) {
	input := ""
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}



	input = `func line(int x, int y){
					//last number in loop is the step
		iter (x := 0; x, 1){	//declares it, scope is the loop
			circle(2, 3, y, 5);
		}
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}

	input = `//macro definition
	func line(int x, int y, int i){
					//last number in loop is the step
		iter (i := 0; x, 1){	//declares it, scope is the loop
			circle(2, 3, y, 5);
		}
	}

	//macro entry
	func main(){
		int i;
		iter (i := 0; 3, 1){
			circle(i, i, 3, 0xff);
		}
		int j;
		iter (j := 0; 8, 2){	//loops 0 2 4 6 8
			circle(j, j, 8, 0xff);
		}
		rect(4, 5, 2, 0x11000011);
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}
}

func TestWrongFunctions(t *testing.T) {
	input := "line(){}"
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = "func line(){} ffawsa"
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = "func (){}"
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = "func line){}"
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = "func line({}"
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = "func line()"
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = "func line(){"
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
}

func TestMultipleArgs(t *testing.T) {
	input := "func line(int fuu, int ro, int fjskalf, int fsjapfj){}"
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}
}

func TestOneArg(t *testing.T) {
	input := `func line(int hola)
	{
	}
	`
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}
}

func TestNoneArg(t *testing.T) {
	input := "func line(){}"
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}
}

func TestWrongArgs(t *testing.T) {
	input := "func line(int fuu,){}"
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = "func line(int fuu, int far, int){}"
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = "func line(int fuu, int int far, int){}"
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = "func line(int fuu, int roou  int jdsa){}"
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
}

func TestCallFunc(t *testing.T) {
	input := `func line(int hola, int jaja, int vabien){line(2, 3, jaja, 5);}`
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}

	input = `func line(int hola, int jaja, int vabien){
		line();
		line(2);
		line(2, 3);
		line(2, 3, hola);
		line(2, 3, hola, 5);
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}
}

func TestWrongSentences(t *testing.T) {
	input := `func line(int hola, int vabien){
	line(hola)    }`
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = `func line(int hola, int vabien){
	line ident (fuu);    }`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = `func line(int hola, int vabien){
	line );    }`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = `func line(int hola, int vabien){
	line (hola;    }`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
}

func TestParameters(t *testing.T) {
	input := `func rectangle(){}
	func line(int hola, int vabien){
	rectangle(); 
	}`
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}

	input = `func move(){}
	func line(int hola, int vabien){
	move(435);
    }`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}

	input = `func move(){}
	func line(int hola, int vabien){
	move(0xFFF);
    }`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}

	input = `func move(){}
	func line(int hola, int vabien){
	move(vabien);
    }`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}

	input = `func move(){}
	func line(int hola, int vabien){
	move(hola, 433, 0x5F, hola);   
    }`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}
}

func TestWrongParameters(t *testing.T) {
	input := `func line(int hola, int vabien){
	line(vabien,   ); 
	}`
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = `func line(int hola, int vabien){
	line(hola  435);
    }`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = `func line(int hola, int vabien){
	line(,);
    }`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
}

func TestAsign(t *testing.T) {
	input := `func line(int x, int y, bool f){
		int z;
		int i;
		y = 3;
		x = 0xffff;
		z = x;
		i = y * (3 -f)**2;
	}`

	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}

	input = `func line(int y, int i, int f){
		y = y;
		i =  * (3 -f)**2;
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = `func line(int y, int i, int f){
		y = 43
		i =  y * (3 -f)**2;
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
}

func TestIf(t *testing.T) {
	input := `func line(int hola, int vabien){
			if (True) {	}
		}`
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}
	input = `func line(int hola, int vabien){
			if (True) {	} else {}
		}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}
	input = `func line(int fuus, int distance){
			if (True) {
				line(43,5,fuus);
				line(distance);
			} else {
				line(distance);
				int j;
				iter (j := 0; 8, 2){	//loops 0 2 4 6 8
					line(j, j, 8, 0xff);
					line(4, 5, 2, 0x11000011);
				}
			}
		}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}
}

func TestWrongIf(t *testing.T) {
	input := `func line(int hola, int vabien){
			if () {	}
		}`
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
	input = `func line(int hola, int vabien){
			if 43*(487+2**2) {	}
		}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
	input = `func line(int hola, int vabien){
			if (43*(487+2**2) {	 int fus; }
		}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
	input = `func line(int hola, int vabien){
			if (43*(487+2**2))  int fus;	}
		}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
	input = `func line(int hola, int vabien){
			if (43*(487+2**2))  { int fus;	
		}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
	input = `func line(int hola, int vabien){
			if (43*(487+2**2))  { int fus;
			}else
		}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
	input = `func line(int hola, int vabien){
			if (43*(487+2**2))  { int fus;
			}else {
		}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
}

func TestVarDeclar(t *testing.T) {
	input := `func line(int hola, int vabien){
		int fus;
		bool roo;
		Coord dah;
		}`

	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}

	input = `func line(int hola, int vabien){
		int fus;
		Coord ;
		}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
	input = `func line(int hola, int vabien){
		int fus;
		bool roo
		Coord dah;
		}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
}

func TestLoops(t *testing.T) {
	input := `
	func main(){
		int i;
		int j;
		iter (i := 0; 3, 1){
			circle(i, i, 3, 0xff);
		}
		iter (j := 0; 8, 2){	//loops 0 2 4 6 8
			iter (j := 0; 8, 2){	//loops 0 2 4 6 8
				circle(j, j, 8, 0xff);
				circle(4, 5, 2, 0x11000011);
			}
		circle(4, 5, 2, 0x11000011);
		}
	}`
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}
}

func TestWrongLoops(t *testing.T) {
	input := `
	func line(int hola, int vabien){
		int i;
		iter i := 0; 3, 1){
			circle(i, i, 3, 0xff);
		}
	}`
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = `func line(int hola, int i){
		iter (i := 0; 3, 1{
			line(i, i, 3, 0xff);
		}
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = `func line(int hola, int i){
		iter (i := 0 3, 1){
			line(i, i, 3, 0xff);
		}
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = `func line(int hola, int i){
		iter (i := 0; 3 1{
			line(i, i, 3, 0xff);
		}
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = `func line(int hola, int i){
		iter (i := 0; , 1{
			line(i, i, 3, 0xff);
		}
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = `func line(int fus, int i){
		iter (i := 0; fus, ){
			rect(i, i, 3, 0xff);
		}
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
}

func TestWrongInitializartion(t *testing.T) {
	input := `func main(){
		iter ( := 0; 3, 1){
			rect(i, i, 3, 0xff);
		}
	}`
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = `func line(int fus, int i){
		iter (i  0; fus, 435){
			rect(i, i, 3, 0xff);
		}
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}

	input = `func line(int fus, int i){
		iter (i := ; fus, 435){
			rect(i, i, 3, 0xff);
		}
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
}

func TestCombinedExpr (t *testing.T) {
	input := `func main(){
		main(3245 - 2 *324, 324|4- 2, 3*(3+4), 3245 - 2 *324- 2** -5, 1000%(3*5)-(3- 4*2) );
	}`
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}
	input = `func main(int ras){
		main(3245 -ras *324- 2** (-5*34);
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
	input = `func main(){
		main(3245 - 32**-(2*3));
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
}

func TestlIneqExpr (t *testing.T) {
	input := `func main(){
		rect(True < False, fuuu <= fuuu, False>False, 3324>=0001);
	}`
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}
	input = `func main(){
		rect(das, 435<);
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
	input = `func main(){
		rect(das, 435<=);
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
	input = `func main(){
		rect(das, 435>);
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
	input = `func main(){
		rect(das, 435>=);
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
}

func TestLevel3Exp (t *testing.T) {
	input := `
	func main(int fuuu){
		rect(True + False, fuuu  - fuuu, False^False, 3324|0001, 3+4- fuuu);
	}`
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}
	input = `
	func main(int das){
		rect(das, 435+);
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
	input = `
	func main(int das){
		rect(das, 435-);
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
	input = `func main(int das){
		rect(das, 435^-);
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
	input = `func main(int das){
		rect(das, 435|);
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
}

func TestLevel2Exp (t *testing.T) {
	input := `
	func main(int das){
		rect(das*345, 0x324/das, 43%2, (2/43)*das , False & True);
	}`
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}
	input = `func main(){
		circle(1, 2/, 2**32);
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
	input = `func main(){
		circle(3, 2%, fuu**32);
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
	input = `func main(){
		circle(True, 2&, fuu**32);
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
}

func TestPowExpr (t *testing.T) {
	input := `func main(){
		main(2**432, 4**32, !32**4444, !324**4**2**-3);
	}`
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}

	input = `func main(){
		main(111, 2***32, 324**32);
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
}

func TestNotExpr (t *testing.T) {
	input := `func main(){
		main(!4, !!(3), !!!!!!0x34, 478);
	}`
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n", nerrors)
	}
	input = `func line(int hola, int i){
			line(!i, !, 3);
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
	input = `func line(int hola, int i){
			line(!i, !(0x3245, hola);
	}`
	scn = NewScanner(NewText(strings.NewReader(input)), "Test_Parser")
	NewEnvStack()
	FXParse(scn)
	if nerrors == 0 {
		t.Errorf("Error expected\n")
	}
}