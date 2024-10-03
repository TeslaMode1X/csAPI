package facadePattern

import "fmt"

// Подсистема: DVD Player
type DVDPlayer struct{}

func (d *DVDPlayer) On() {
	fmt.Println("DVD Player is On")
}

func (d *DVDPlayer) Play(movie string) {
	fmt.Printf("Playing movie: %s\n", movie)
}

func (d *DVDPlayer) Off() {
	fmt.Println("DVD Player is Off")
}

// Подсистема: Projector
type Projector struct{}

func (p *Projector) On() {
	fmt.Println("Projector is On")
}

func (p *Projector) Off() {
	fmt.Println("Projector is Off")
}

// Подсистема: Sound System
type SoundSystem struct{}

func (s *SoundSystem) On() {
	fmt.Println("Sound System is On")
}

func (s *SoundSystem) SetVolume(level int) {
	fmt.Printf("Sound volume set to %d\n", level)
}

func (s *SoundSystem) Off() {
	fmt.Println("Sound System is Off")
}

// Подсистема: Lights
type Lights struct{}

func (l *Lights) Dim(level int) {
	fmt.Printf("Lights dimmed to %d%%\n", level)
}

// Facade: Home Theater
type HomeTheaterFacade struct {
	dvd       *DVDPlayer
	projector *Projector
	sound     *SoundSystem
	lights    *Lights
}

func NewHomeTheaterFacade(dvd *DVDPlayer, projector *Projector, sound *SoundSystem, lights *Lights) *HomeTheaterFacade {
	return &HomeTheaterFacade{
		dvd:       dvd,
		projector: projector,
		sound:     sound,
		lights:    lights,
	}
}

func (h *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("Get ready to watch a movie...")
	h.lights.Dim(10)
	h.projector.On()
	h.sound.On()
	h.sound.SetVolume(5)
	h.dvd.On()
	h.dvd.Play(movie)
}

func (h *HomeTheaterFacade) EndMovie() {
	fmt.Println("Shutting movie theater down...")
	h.dvd.Off()
	h.sound.Off()
	h.projector.Off()
}

func Facade() {
	// Создаем экземпляры подсистем
	dvd := &DVDPlayer{}
	projector := &Projector{}
	sound := &SoundSystem{}
	lights := &Lights{}

	// Создаем фасад
	homeTheater := NewHomeTheaterFacade(dvd, projector, sound, lights)

	// Используем фасад для просмотра фильма
	homeTheater.WatchMovie("Inception")
	fmt.Println()
	homeTheater.EndMovie()
}
