package stats

func (m *Module) setupCommands() {
}

// 	// Setup Console Commands
// 	console.AddHandler("/coffee", "How many coffees are you on for the day?", m.consoleCoffee)
// 	console.AddHandler("/save", "Did someone save your ass this stream?", m.consoleSave)
// 	console.AddHandler("/workingon", "Set your currently working on text.", m.consoleWorkingOn)
// 	console.AddAlias("/w", "/workingon")
// }

// func (m *WorkingOnModule) consoleCoffee(input string) {
// 	i, err := strconv.Atoi(input)

// 	if err == nil {
// 		m.CoffeeCount = i
// 	} else {
// 		m.CoffeeCount++
// 	}

// 	if m.config.WorkingOn.Output {
// 		Core.SaveFile([]byte(fmt.Sprintf("%02d", m.CoffeeCount)), m.CoffeePath)
// 	}
// 	Core.Log("WORKING", "WORKING", "Coffee Count: "+strconv.Itoa(m.CoffeeCount))
// }

// func (m *WorkingOnModule) consoleSave(input string) {
// 	i, err := strconv.Atoi(input)

// 	if err == nil {
// 		m.SavesCount = i
// 	} else {
// 		m.SavesCount++
// 	}

// 	if m.config.WorkingOn.Output {
// 		Core.SaveFile([]byte(fmt.Sprintf("%02d", m.SavesCount)), m.SavesPath)
// 	}
// 	Core.Log("WORKING", "WORKING", "Saves Count: "+strconv.Itoa(m.SavesCount))
// }

// func (m *WorkingOnModule) consoleWorkingOn(input string) {
// 	m.Message = input
// 	if m.config.WorkingOn.Output {
// 		Core.SaveFile([]byte(input), m.TextPath)
// 	}
// 	Core.Log("WORKING", "LOG", "Set: "+input)
// }
