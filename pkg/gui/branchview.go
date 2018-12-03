package gui

import (
	"fmt"

	"github.com/isacikgoz/gitbatch/pkg/git"
	"github.com/jroimartin/gocui"
)

// updates the branchview for given entity
func (gui *Gui) updateBranch(g *gocui.Gui, entity *git.RepoEntity) error {
	var err error
	out, err := g.View("branch")
	if err != nil {
		return err
	}
	out.Clear()

	currentindex := 0
	totalbranches := len(entity.Branches)
	for i, b := range entity.Branches {
		if b.Name == entity.Branch.Name {
			currentindex = i
			fmt.Fprintln(out, selectionIndicator+b.Name)
			continue
		}
		fmt.Fprintln(out, tab+b.Name)
	}
	if err = gui.smartAnchorRelativeToLine(out, currentindex, totalbranches); err != nil {
		return err
	}
	return nil
}

// iteration handler for the branchview
func (gui *Gui) nextBranch(g *gocui.Gui, v *gocui.View) error {
	var err error
	entity := gui.getSelectedRepository()
	if err = entity.Checkout(entity.NextBranch()); err != nil {
		if err = gui.openErrorView(g, err.Error(),
			"You should manually resolve this issue"); err != nil {
			return err
		}
		return nil
	}
	if err = gui.updateBranch(g, entity); err != nil {
		return err
	}
	if err = gui.updateCommits(g, entity); err != nil {
		return err
	}
	if err = gui.updateRemoteBranches(g, entity); err != nil {
		return err
	}
	if err = gui.refreshMain(g); err != nil {
		return err
	}
	return nil
}
