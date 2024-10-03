package compositePattern

type Component interface {
	showPrice()
}

// compositePattern
//computer := &composite_pattern.Composite{name: "PC", price: 590000}
//
//// hardware
//motherBoard := &composite_pattern.Composite{name: "Mother Board", price: 50000}
//hdd := &composite_pattern.Leaf{name: "HDD", price: 10000}
//cpu := &composite_pattern.Leaf{name: "CPU", price: 80000}
//ram := &composite_pattern.Leaf{name: "RAM", price: 50000}
//videocard := &composite_pattern.Leaf{name: "Video Card", price: 300000}
//
//motherBoard.addComponent(hdd, cpu, ram, videocard)
//
//// devices
//devices := &composite_pattern.Composite{name: "Devices", price: 100000}
//mouse := &composite_pattern.Leaf{name: "Mouse", price: 20000}
//keyboard := &composite_pattern.Leaf{name: "Keyboard", price: 20000}
//headsets := &composite_pattern.Leaf{name: "Headsets", price: 60000}
//
//devices.addComponent(mouse, keyboard, headsets)
//
//// add everything into computer
//computer.addComponent(motherBoard, devices)
//
//computer.showPrice()
