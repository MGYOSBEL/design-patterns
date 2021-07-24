package devices

import "fmt"

type Source int

const (
	CD Source = iota
	DVD
	Radio
)

type Stereo struct {
	source    Source
	poweredOn bool
	voulume   int
}

func NewStereo() Stereo {
	return Stereo{
		poweredOn: false,
		source:    Radio,
		voulume:   0,
	}
}

func (st Stereo) On() {
	st.poweredOn = true
	st.voulume = 2
	fmt.Println("Stereo have been powered on...")
	fmt.Printf("Source is %s \n", getSource(st.source))
}

func (st Stereo) Off() {
	st.poweredOn = false
	fmt.Println("Stereo have been powered off...")
}

func (st Stereo) SetCD() {
	(&st).setSource(CD)
}

func (st Stereo) SetDVD() {
	(&st).setSource(DVD)
}

func (st Stereo) SetVoulme(volume int) {
	st.voulume = volume
	fmt.Printf("Stereo volume set to %d\n", volume)
}

func (st Stereo) SetRadio() {
	(&st).setSource(Radio)
}

func (st *Stereo) setSource(source Source) {
	st.source = source
	fmt.Printf("Source is %s \n", getSource(st.source))

}

func getSource(source Source) string {
	switch source {
	case CD:
		return "CD"
	case DVD:
		return "DVD"
	default:
		return "Radio"
	}
}
