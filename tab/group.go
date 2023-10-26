package tab

import "github.com/gdamore/tcell/v2"

var rootGroup Group

// this will be used later to handle split views
type Group struct {
	s         tcell.Screen
	activeTab *Tab
	w         int
	h         int
}

func NewRootGroup(s tcell.Screen, w, h int) *Group {
	// todo rewrite maybe
	g := NewGroup(s, w, h)
	rootGroup = *g
	return &rootGroup
}

func Root() *Group {
	return &rootGroup
}

func NewGroup(s tcell.Screen, w, h int) *Group {
	return &Group{
		s: s,
		w: w,
		h: h,
	}
}

func (g *Group) NewTab() {
	g.activeTab = newTab(g.s, g.w, g.h-1, 0, 0, g)
}

func (g *Group) ActiveTab() *Tab {
	return g.activeTab
}
