package stepImpl

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/getgauge-contrib/gauge-go/gauge"
	"github.com/getgauge-contrib/gauge-go/testsuit"
)

var _ = gauge.Step(`Run <command>`, func(command string) {
	// Build command
	tokens := strings.Split(command, " ")
	binary, err := exec.LookPath(tokens[0])
	if err != nil {
		testsuit.T.Fail(fmt.Errorf("can not found binary: %s", binary))
	}
	cmd := exec.Command(binary, tokens[1:]...)

	// Run command
	out, err := cmd.CombinedOutput()
	gauge.WriteMessage(string(out))
	if err != nil {
		testsuit.T.Fail(fmt.Errorf("Command exit with %q: %q\n", err, out))
	}
})

var _ = gauge.Step(`Check folder <folder> is exists`, func(folder string) {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		testsuit.T.Fail(fmt.Errorf("folder %s is not exists", folder))
	}
})
