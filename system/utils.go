package system

import (
	"github.com/manifoldco/promptui"
)

func AskInitializationQuestions(cfg *Config) error {
	directory_prompt := &promptui.Prompt{
		Label:     "Root Project Directory (full path)",
		Default:   "/",
		IsVimMode: true,
		AllowEdit: true,
	}
	dir, err := directory_prompt.Run()
	if err != nil {
		return err
	}
	editor_prompt := promptui.Select{
		Label:     "Preferred Editor",
		IsVimMode: true,
		Items:     []string{string(VSCODE), string(NEOVIM), string(VIM), string(EMACS)},
	}
	_, editor, err := editor_prompt.Run()
	if err != nil {
		return err
	}
	editor_profile_prompt := promptui.Select{
		Label:     "Preferred Profile (for Terminal-based environments)",
		IsVimMode: true,
		Items:     []string{string(VSCODE), string(NEOVIM), string(VIM), string(EMACS)},
	}
	_, editor_profile, err := editor_profile_prompt.Run()
	if err != nil {
		return err
	}

	cfg = &Config{
		Path:    dir,
		Editor:  editor,
		Profile: editor_profile,
	}

	return nil
}
