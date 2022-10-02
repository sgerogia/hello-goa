package logic_test

import (
	"github.com/sgerogia/hello-goa/logic"
	"log"
	"os"
	"testing"
)

// --- Global variable for child tests ---
var testingCtx struct {
	mathsvc *logic.MathSrvc

}

func TestMain(m *testing.M) {

	// global setup
	testingCtx.mathsvc = newMathSvc()

	// Run test suites
	exitVal := m.Run()

	// clean up here

	// ...and exit test suite
	os.Exit(exitVal)
}

// Wrap normal initializer, explicitly cast to service struct to expose internal methods
func newMathSvc() *logic.MathSrvc {

	return logic.NewMath(
		log.New(os.Stderr, "[math-test] ", log.Ltime),
	).(*logic.MathSrvc)
}

