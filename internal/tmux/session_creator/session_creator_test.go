package sessioncreator

import "testing"

func TestSessionCreation(t *testing.T) {
	tmuxSessionCreator := New()
	tmuxSessionCreator.NewProjectLocation = TestDefaultProjectRootDir
	tmuxSessionCreator.TmuxinatorConfigLocation = TestTmuxinatorConfigDir
	if err := tmuxSessionCreator.CreateNewSession("test-project-1"); err != nil {
		t.Error(err)
		t.FailNow()
	}
}
