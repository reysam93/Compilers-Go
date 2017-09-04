package fxparser

import (
	"strings"
	"testing"
)

func TestFunctions(t *testing.T) {
	input := ""
	p := ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Error: %s\n", err)
	}

	input = `func line(int x, int y){
					//last number in loop is the step
		iter (x := 0; x, 1){	//declares it, scope is the loop
			line(2, 3, y, 5);
		}
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Error: %s\n", err)
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
		circle(4, 5, 2, 0x11000011);
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Error: %s\n", err)
	}
}

func TestWrongFunctions(t *testing.T) {
	input := "line(){}"
	p := ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoEof)
	}

	input = "func line(){} ffawsa"
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoEof)
	}

	input = "func (){} ffawsa"
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoId)
	}

	input = "func line){} ffawsa"
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoLeftPar)
	}

	input = "func line({} ffawsa"
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoRightPar)
	}

	input = "func line() ffawsa"
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoLeftBra)
	}

	input = "func line(){"
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoRightBra)
	}
}

func TestMultipleArgs(t *testing.T) {
	input := "func line(int fuu, int ro, int fjskalf, int fsjapfj){}"
	p := ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Error: %s\n", err)
	}
}

func TestOneArg(t *testing.T) {
	input := `func line(int hola)
	{
	}
	`
	p := ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Error: %s\n", err)
	}
}

func TestNoneArg(t *testing.T) {
	input := "func line(){}"
	p := ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Error: %s\n", err)
	}
}

func TestWrongArgs(t *testing.T) {
	input := "func line(int fuu,){}"
	p := ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoParam)
	}

	input = "func line(int fuu, int far, int){}"
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoId)
	}

	input = "func line(int fuu, int int far, int){}"
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoId)
	}

	input = "func line(int fuu, int roou  int jdsa){}"
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoRightPar)
	}
}

func TestCallFunc(t *testing.T) {
	input := `func line(int hola, int jaja, int vabien){line(2, 3, jaja, 5);}`
	p := ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Error: %s\n", err)
	}

	input = `func line(int hola, int jaja, int vabien){
		line();
		line(2);
		line(2, 3);
		line(2, 3, hola);
		line(2, 3, hola, 5);
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Error: %s\n", err)
	}
}

func TestWrongSentences(t *testing.T) {
	input := `func line(int hola, int vabien){
	line(hola)    }`
	p := ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoScol)
	}

	input = `func line(int hola, int vabien){
	line ident (fuu);    }`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoIdSent)
	}

	input = `func line(int hola, int vabien){
	line );    }`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoIdSent)
	}

	input = `func line(int hola, int vabien){
	line (hola;    }`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoRightPar)
	}
}

func TestParameters(t *testing.T) {
	input := `func rectangle(){}
	func line(int hola, int vabien){
	rectangle(); 
	}`
	p := ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Found error:%s\n", err)
	}

	input = `func move(){}
	func line(int hola, int vabien){
	move(435);
    }`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Found error:%s\n", err)
	}

	input = `func move(){}
	func line(int hola, int vabien){
	move(0xFFF);
    }`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Found error:%s\n", err)
	}

	input = `func move(){}
	func line(int hola, int vabien){
	move(vabien);
    }`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Found error:%s\n", err)
	}

	input = `func move(){}
	func line(int hola, int vabien){
	move(hola, 433, 0x5F, hola);   
    }`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Found error:%s\n", err)
	}
}

func TestWrongParameters(t *testing.T) {
	input := `func line(int hola, int vabien){
	line(vabien,   ); 
	}`
	p := ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoExpr)
	}

	input = `func line(int hola, int vabien){
	line(hola  435);
    }`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoRightPar)
	}

	input = `func line(int hola, int vabien){
	line(,);
    }`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoRightPar)
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

	p := ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Error: %s\n", err)
	}

	input = `func line(int y, int i, int f){
		y = y;
		i =  * (3 -f)**2;
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoExpr)
	}

	input = `func line(int y, int i, int f){
		y = 43
		i =  y * (3 -f)**2;
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoScol)
	}
}

func TestIf(t *testing.T) {
	input := `func line(int hola, int vabien){
			if (True) {	}
		}`
	p := ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Error: %s\n", err)
	}
	input = `func line(int hola, int vabien){
			if (True) {	} else {}
		}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Error: %s\n", err)
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
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Error: %s\n", err)
	}
}

func TestWrongIf(t *testing.T) {
	input := `func line(int hola, int vabien){
			if () {	}
		}`
	p := ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoId)
	}
	input = `func line(int hola, int vabien){
			if 43*(487+2**2) {	}
		}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoLeftPar)
	}
	input = `func line(int hola, int vabien){
			if (43*(487+2**2) {	 int fus; }
		}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoRightPar)
	}
	input = `func line(int hola, int vabien){
			if (43*(487+2**2))  int fus;	}
		}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoLeftBra)
	}
	input = `func line(int hola, int vabien){
			if (43*(487+2**2))  { int fus;	
		}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoRightBra)
	}
	input = `func line(int hola, int vabien){
			if (43*(487+2**2))  { int fus;
			}else
		}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoLeftBra)
	}
	input = `func line(int hola, int vabien){
			if (43*(487+2**2))  { int fus;
			}else {
		}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoRightBra)
	}
}

func TestVarDeclar(t *testing.T) {
	input := `func line(int hola, int vabien){
		int fus;
		bool roo;
		Coord dah;
		}`

	p := ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Error: %s\n", err)
	}

	input = `func line(int hola, int vabien){
		int fus;
		Coord ;
		}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoId)
	}
	input = `func line(int hola, int vabien){
		int fus;
		bool roo
		Coord dah;
		}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoScol)
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
	p := ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Found error:%s\n", err)
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
	p := ParserFromReader("Test_Parser", strings.NewReader(input))
	//p.symbEnvs.Debug = true
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoLeftPar)
	}

	input = `func line(int hola, int i){
		iter (i := 0; 3, 1{
			line(i, i, 3, 0xff);
		}
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoRightPar)
	}

	input = `func line(int hola, int i){
		iter (i := 0 3, 1){
			line(i, i, 3, 0xff);
		}
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoScol)
	}

	input = `func line(int hola, int i){
		iter (i := 0; 3 1{
			line(i, i, 3, 0xff);
		}
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoComa)
	}

	input = `func line(int hola, int i){
		iter (i := 0; , 1{
			line(i, i, 3, 0xff);
		}
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoExpr)
	}

	input = `func line(int fus, int i){
		iter (i := 0; fus, ){
			rect(i, i, 3, 0xff);
		}
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoExpr)
	}
}

func TestWrongInitializartion(t *testing.T) {
	input := `func main(){
		iter ( := 0; 3, 1){
			rect(i, i, 3, 0xff);
		}
	}`
	p := ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoId)
	}

	input = `func line(int fus, int i){
		iter (i  0; fus, 435){
			rect(i, i, 3, 0xff);
		}
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoLoopEq)
	}

	input = `func line(int fus, int i){
		iter (i := ; fus, 435){
			rect(i, i, 3, 0xff);
		}
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoExpr)
	}
}

func TestCombinedExpr (t *testing.T) {
	input := `func main(){
		main(3245 - 2 *324, 324|4- 2, 3*(3+4), 3245 - 2 *324- 2** -5, 1000%(3*5)-(3- 4*2) );
	}`
	p := ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Found error:%s\n", err)
	}
	input = `func main(int ras){
		main(3245 -ras *324- 2** (-5*34);
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoVal)
	}
	input = `func main(){
		main(3245 - 32**-(2*3));
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoVal)
	}
}

func TestlIneqExpr (t *testing.T) {
	input := `func main(){
		rect(True < False, fuuu <= fuuu, False>False, 3324>=0001);
	}`
	p := ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Found error:%s\n", err)
	}
	input = `func main(){
		rect(das, 435<);
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoVal)
	}
	input = `func main(){
		rect(das, 435<=);
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoVal)
	}
	input = `func main(){
		rect(das, 435>);
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoVal)
	}
	input = `func main(){
		rect(das, 435>=);
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoVal)
	}
}

// NOTE: for '-' to be a token, it must be separate of the nnumber
// '2 -3' is not an operation, are two different numbers, the second being negative
func TestLevel3Exp (t *testing.T) {
	input := `
	func main(int fuuu){
		rect(True + False, fuuu  - fuuu, False^False, 3324|0001, 3+4- fuuu);
	}`
	p := ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Found error:%s\n", err)
	}
	input = `
	func main(int das){
		rect(das, 435+);
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoVal)
	}
	input = `
	func main(int das){
		rect(das, 435-);
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoVal)
	}
	input = `func main(int das){
		rect(das, 435^-);
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoVal)
	}
	input = `func main(int das){
		rect(das, 435|);
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoVal)
	}
}

func TestLevel2Exp (t *testing.T) {
	input := `
	func main(int das){
		rect(das*345, 0x324/das, 43%2, (2/43)*das , False & True);
	}`
	p := ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Found error:%s\n", err)
	}
	input = `func main(){
		circle(1, 2/, 2**32);
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoVal)
	}
	input = `func main(){
		circle(3, 2%, fuu**32);
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoVal)
	}
	input = `func main(){
		circle(True, 2&, fuu**32);
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoVal)
	}
}
			
func TestPowExpr (t *testing.T) {
	input := `func main(){
		main(2**432, 4**32, !32**4444, !324**4**2**-3);
	}`
	p := ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Found error:%s\n", err)
	}

	input = `func main(){
		main(111, 2***32, 324**32);
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoVal)
	}

}

func TestNotExpr (t *testing.T) {
	input := `func main(){
		main(!4, !!(3), !!!!!!0x34, 478);
	}`
	p := ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err != nil {
		t.Errorf("Found error:%s\n", err)
	}
	input = `func line(int hola, int i){
			line(!i, !, 3);
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoVal)
	}
	input = `func line(int hola, int i){
			line(!i, !(0x3245, hola);
	}`
	p = ParserFromReader("Test_Parser", strings.NewReader(input))
	if _, err := p.Parse(); err == nil {
		t.Errorf("Expected error:\n", ErrNoRightPar)
	}
}