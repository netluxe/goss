package resource

import "github.com/aelsabbahy/goss/system"

type Process struct {
	Executable string `json:"-"`
	Running    bool   `json:"running"`
}

func (p *Process) ID() string      { return p.Executable }
func (p *Process) SetID(id string) { p.Executable = id }

func (p *Process) Validate(sys *system.System) []TestResult {
	sysProcess := sys.NewProcess(p.Executable, sys)

	var results []TestResult

	results = append(results, ValidateValue(p, "running", p.Running, sysProcess.Running))

	return results
}

func NewProcess(sysProcess system.Process, ignoreList []string) *Process {
	executable := sysProcess.Executable()
	running, _ := sysProcess.Running()
	return &Process{
		Executable: executable,
		Running:    running.(bool),
	}
}
