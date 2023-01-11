package main

func main() {
	handlerList := []struct {
		Name    string
		Handler func()
	}{
		{
			Name:    "init_chassis",
			Handler: InitChassis,
		},
		{
			Name:    "init_apps",
			Handler: InitApps,
		},
	}
	for _, handlerInfo := range handlerList {
		handlerInfo.Handler()
	}

	// run chassis
	RunChassis()

}
