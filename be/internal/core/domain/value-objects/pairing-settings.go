package vo

type PairingSettings struct {
	notifyMe	 		bool
	lowLevelThreshold	int64
}

func NewPairSettings(enableNotification bool,
	lowLevelVal int64) PairingSettings {
	return PairingSettings{
		notifyMe: enableNotification,
		lowLevelThreshold: lowLevelVal,
	}
}

func (s *PairingSettings) NotifyMe() bool {
	return s.notifyMe
}

func (s *PairingSettings) LowLevelThreshold() int64 {
	return s.lowLevelThreshold
}
