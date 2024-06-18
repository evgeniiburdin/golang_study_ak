package main

type TVController interface {
	switchON()
	switchOFF()
	GetStatus()
	GetModel()
}

type SamsungTV struct {
	status bool
	model  string
}

func (s *SamsungTV) switchON() {
	s.status = true
}
func (s *SamsungTV) switchOFF() {
	s.status = false
}
func (s *SamsungTV) GetStatus() bool {
	return s.status
}
func (s *SamsungTV) GetModel() string {
	return s.model
}
func (s *SamsungTV) SamsungHub() string {
	return "SamsungHub"
}

type LgTV struct {
	status bool
	model  string
}

func (l *LgTV) switchON() {
	l.status = true
}
func (l *LgTV) switchOFF() {
	l.status = false
}
func (l *LgTV) GetStatus() bool {
	return l.status
}
func (l *LgTV) GetModel() string {
	return l.model
}
func (l *LgTV) LGHub() string {
	return "LGHub"
}
