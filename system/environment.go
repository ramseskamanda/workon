package system

type Environment string

const (
	TERMINAL        Environment = "Apple_Terminal"
	VSCODE_TERMINAL Environment = "vscode"
	ITERM2          Environment = "iTerm.app"
)
