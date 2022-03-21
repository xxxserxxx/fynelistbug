package main

import (
	"fmt"
	"path"
	"strings"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("net.ser1.fynelistbug")

	topDir, err := storage.ParseURI(path.Join(a.Storage().RootURI().String(), "fynelistbug"))
	if err != nil {
		panic(err)
	}
	if yes, err := storage.Exists(topDir); err != nil {
		panic(err)
	} else if !yes {
		if err := storage.CreateListable(topDir); err != nil {
			panic(err)
		}
	}

	subDir, err := storage.ParseURI(path.Join(topDir.String(), "subdir"))
	if err != nil {
		panic(err)
	}
	if yes, err := storage.Exists(subDir); err != nil {
		panic(err)
	} else if !yes {
		if err := storage.CreateListable(subDir); err != nil {
			panic(err)
		}
	}
	if ok, err := storage.CanList(subDir); !ok {
		panic(fmt.Errorf("%s is not a listable directory", subDir))
	} else if err != nil {
		panic(err)
	}

	fmt.Printf("subDir => %s\n", subDir)
	fs, err := storage.List(subDir)
	if err != nil {
		panic(err)
	}
	rv := make(map[string]string)
	var nilCount, notNilCount int
	for _, f := range fs {
		if f == nil {
			nilCount++
			continue
		}
		notNilCount++
		name := f.Name()
		ps := strings.SplitN(name, ".", 2)
		if len(ps) > 1 {
			rv[ps[1]] = ps[0]
		}
	}
	ls := fmt.Sprintf("%d nils, %d not-nils for %s", nilCount, notNilCount, subDir.Name())
	w := a.NewWindow("Fyne List Bug")
	l := widget.NewLabel(ls)
	w.SetContent(l)
	w.ShowAndRun()
}
