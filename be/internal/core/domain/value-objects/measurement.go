package vo

type Measurement struct {
	temperature  float32
	soilMoisture float32
	humidity     float32
	pressure     float32
	altitude     float32
}

func NewMeasurement(
	t float32,
	sm float32,
	h float32,
	p float32,
	a float32) Measurement {
	return Measurement{
		temperature:  t,
		soilMoisture: sm,
		humidity:     h,
		pressure:     p,
		altitude:     a,
	}
}

func (m *Measurement) Temperature() float32 {
	return m.temperature
}

func (m *Measurement) SoilMoisture() float32 {
	return m.soilMoisture
}

func (m *Measurement) Humidity() float32 {
	return m.humidity
}

func (m *Measurement) Pressure() float32 {
	return m.pressure
}

func (m *Measurement) Altitude() float32 {
	return m.altitude
}
