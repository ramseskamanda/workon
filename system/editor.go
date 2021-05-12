package system

import "os/exec"

type Editor string

const (
	VSCODE Editor = "vscode"
	NEOVIM Editor = "neovim"
	VIM    Editor = "vim"
	EMACS  Editor = "emacs"
)

func (editor Editor) Open(full_path string) error {
	switch editor {
	case VSCODE:
		cmd := exec.Command("code", full_path)
		if err := cmd.Run(); err != nil {
			return err
		}
	case NEOVIM:
		break
	case VIM:
		break
	case EMACS:
		break
	}

	return nil
}
