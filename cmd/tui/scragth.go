package main

// func layout(g *gocui.Gui) error {
// 	var views = []string{treeView, textView, pathView}
// 	maxX, maxY := g.Size()
// 	for _, view := range views {
// 		x0, y0, x1, y1 := viewPositions[view].getCoordinates(maxX, maxY)
// 		if v, err := g.SetView(view, x0, y0, x1, y1); err != nil {
// 			v.SelFgColor = gocui.ColorBlack
// 			v.SelBgColor = gocui.ColorGreen

// 			v.Title = " " + view + " "
// 			if err != gocui.ErrUnknownView {
// 				return err

// 			}
// 			if v.Name() == treeView {
// 				v.Highlight = true
// 				drawTree(g, v, tree)
// 				// v.Autoscroll = true
// 			}
// 			if v.Name() == textView {
// 				drawJSON(g, v)
// 			}

// 		}
// 	}
// 	if helpWindowToggle {
// 		height := strings.Count(helpMessage, "\n") + 1
// 		width := -1
// 		for _, line := range strings.Split(helpMessage, "\n") {
// 			width = int(math.Max(float64(width), float64(len(line)+2)))
// 		}
// 		if v, err := g.SetView(helpView, maxX/2-width/2, maxY/2-height/2, maxX/2+width/2, maxY/2+height/2); err != nil {
// 			if err != gocui.ErrUnknownView {
// 				return err

// 			}
// 			fmt.Fprintln(v, helpMessage)

// 		}
// 	} else {
// 		g.DeleteView(helpView)
// 	}
// 	_, err := g.SetCurrentView(treeView)
// 	if err != nil {
// 		log.Fatal("failed to set current view: ", err)
// 	}
// 	return nil

// }
