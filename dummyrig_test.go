package goHamlib

import (
	"reflect"
	"testing"
)

func getDummyRig() (*Rig, error) {
	rig := Rig{}
	// rig.SetDebugLevel(RIG_DEBUG_VERBOSE)
	rig.SetDebugLevel(RIG_DEBUG_ERR)
	if err := rig.Init(1); err != nil {
		return nil, err
	}
	if err := rig.Open(); err != nil {
		return nil, err
	}

	return &rig, nil
}

//Initialize rig with invalid data
// func TestInitializeRigWithInvalidData(t *testing.T) {

// 	rig := Rig{}
// 	rig.SetDebugLevel(RIG_DEBUG_NONE)

// 	//rig model must be > 0
// 	if err := rig.Init(0); fmt.Sprint(err) != "init_rig: invalid rig model" {
// 		t.Fatal("error must be raised on invalid rig model")
// 	}

// 	var p Port_t
// 	p.RigPortType = 1
// 	p.Portname = "/dev/myport" // invalid
// 	p.Baudrate = 38400
// 	p.Databits = 8
// 	p.Stopbits = 1
// 	p.Parity = N
// 	p.Handshake = NO_HANDSHAKE

// 	rig.SetPort(p)

// 	if err := rig.Init(1); fmt.Sprint(err) != "init_rig: invalid parameter" {
// 		t.Fatal("error must be raised when port_t struct is supplied with invalid data")
// 	}

// 	rig.Close()
// 	rig.Cleanup()
// }

//Test if RigCaps are populated correctly
func TestDummyRigCapsPopulation(t *testing.T) {

	rig, err := getDummyRig()
	if err != nil {
		t.Fatal(err)
	}

	defer rig.Close()
	defer rig.Cleanup()

	preamps := []int{10}
	if !reflect.DeepEqual(preamps, rig.Caps.Preamps) {
		t.Fatal("Caps.Preamps contains unexpected data; expected:", preamps, "got:", rig.Caps.Preamps)
	}

	atts := []int{10, 20, 30}
	if !reflect.DeepEqual(atts, rig.Caps.Attenuators) {
		t.Fatal("Caps.Attenuators contains unexpected data; expected:", atts, "got:", rig.Caps.Attenuators)
	}

	maxRit := 9990
	if rig.Caps.MaxRit != maxRit {
		t.Fatal("Caps.MaxRit contains unexpected data; expected:", maxRit, "got:", rig.Caps.MaxRit)
	}

	maxXit := 9990
	if rig.Caps.MaxXit != maxXit {
		t.Fatal("Caps.MaxXit contains unexpected data; expected:", maxXit, "got:", rig.Caps.MaxXit)
	}

	maxIfShift := 10000
	if rig.Caps.MaxIfShift != maxIfShift {
		t.Fatal("Caps.MaxIfShift contains unexpected data; expected:", maxIfShift, "got:", rig.Caps.MaxIfShift)
	}

	vfos := []string{"VFOA", "VFOB", "VFO_MEM"}
	if !reflect.DeepEqual(vfos, rig.Caps.Vfos) {
		t.Fatal("Caps.Vfos contains unexpected data; expected:", vfos, "got:", rig.Caps.Vfos)
	}

	// ops := []string{"BAND_UP", "MCL", "RIGHT", "TOGGLE", "TUNE", "UP"}
	// if !reflect.DeepEqual(ops, rig.Caps.Operations) {
	// 	t.Fatal("Caps.Operations contains unexpected data; expected:", ops, "got:", rig.Caps.Operations)
	// }

	modes := []string{"AM", "CW", "CWR", "FM", "LSB", "RTTY", "RTTYR", "USB", "WFM"}
	if !reflect.DeepEqual(modes, rig.Caps.Modes) {
		t.Fatal("Caps.Modes contains unexpected data; expected:", modes, "got:", rig.Caps.Modes)
	}

	getfs := []string{"ABM", "AFC", "AIP", "ANF", "APF", "ARO", "BC", "COMP", "FAGC", "FBKIN", "LOCK", "MBC", "MN", "MON",
		"MUTE", "NB", "NR", "RESUME", "REV", "RF", "RIT", "SATMODE", "SBKIN", "SCOPE", "SQL", "TBURST", "TONE", "TSQL",
		"TUNER", "VOX", "VSC", "XIT"}
	if !reflect.DeepEqual(getfs, rig.Caps.GetFunctions) {
		t.Fatal("Caps.GetFunctions contains unexpected data; expected:", getfs, "got:", rig.Caps.GetFunctions)
	}

	setfs := []string{"ABM", "AFC", "AIP", "ANF", "APF", "ARO", "BC", "COMP", "FAGC", "FBKIN", "LOCK", "MBC", "MN", "MON",
		"MUTE", "NB", "NR", "RESUME", "REV", "RF", "RIT", "SATMODE", "SBKIN", "SCOPE", "SQL", "TBURST", "TONE", "TSQL",
		"TUNER", "VOX", "VSC", "XIT"}
	if !reflect.DeepEqual(setfs, rig.Caps.SetFunctions) {
		t.Fatal("Caps.SetFunctions contains unexpected data; expected:", setfs, "got:", rig.Caps.SetFunctions)
	}

	getlvs := Values{
		{"AF", 0, 0, 0},
		{"AGC", 0, 0, 0},
		{"ALC", 0, 0, 0},
		{"ANTIVOX", 0, 0, 0},
		{"APF", 0, 0, 0},
		{"ATT", 0, 0, 0},
		{"BALANCE", 0, 0, 0},
		{"BKINDL", 0, 0, 0},
		{"BKIN_DLYMS", 0, 0, 0},
		{"COMP", 0, 0, 0},
		{"CWPITCH", 10, 0, 0},
		{"IF", 0, 0, 0},
		{"KEYSPD", 0, 0, 0},
		{"METER", 0, 0, 0},
		{"MICGAIN", 0, 0, 0},
		{"NOTCHF", 0, 0, 0},
		{"NR", 0, 0, 0},
		{"PBT_IN", 0, 0, 0},
		{"PBT_OUT", 0, 0, 0},
		{"PREAMP", 0, 0, 0},
		{"RAWSTR", 0, 0, 0},
		{"RF", 0, 0, 0},
		{"RFPOWER", 0, 0, 0},
		{"SLOPE_HIGH", 0, 0, 0},
		{"SLOPE_LOW", 0, 0, 0},
		{"SQL", 0, 0, 0},
		{"STRENGTH", 0, 0, 0},
		{"SWR", 0, 0, 0},
		{"VOX", 0, 0, 0},
		{"VOXGAIN", 0, 0, 0},
	}
	if !reflect.DeepEqual(getlvs, rig.Caps.GetLevels) {
		t.Fatal("Caps.GetLevels contains unexpected data; expected:", getlvs, "got:", rig.Caps.GetLevels)
	}

	setlvs := Values{
		{"AF", 0, 0, 0},
		{"AGC", 0, 0, 0},
		{"ANTIVOX", 0, 0, 0},
		{"APF", 0, 0, 0},
		{"ATT", 0, 0, 0},
		{"BALANCE", 0, 0, 0},
		{"BKINDL", 0, 0, 0},
		{"BKIN_DLYMS", 0, 0, 0},
		{"COMP", 0, 0, 0},
		{"CWPITCH", 10, 0, 0},
		{"IF", 0, 0, 0},
		{"KEYSPD", 0, 0, 0},
		{"METER", 0, 0, 0},
		{"MICGAIN", 0, 0, 0},
		{"NOTCHF", 0, 0, 0},
		{"NR", 0, 0, 0},
		{"PBT_IN", 0, 0, 0},
		{"PBT_OUT", 0, 0, 0},
		{"PREAMP", 0, 0, 0},
		{"RF", 0, 0, 0},
		{"RFPOWER", 0, 0, 0},
		{"SLOPE_HIGH", 0, 0, 0},
		{"SLOPE_LOW", 0, 0, 0},
		{"SQL", 0, 0, 0},
		{"VOX", 0, 0, 0},
		{"VOXGAIN", 0, 0, 0},
	}
	if !reflect.DeepEqual(setlvs, rig.Caps.SetLevels) {
		t.Fatal("Caps.SetLevels contains unexpected data; expected:", setlvs, "got:", rig.Caps.SetLevels)
	}

	getparms := Values{
		{"ANN", 0, 0, 0},
		{"APO", 0, 0, 0},
		{"BACKLIGHT", 0, 0, 0},
		{"BAT", 0, 0, 0},
		{"BEEP", 0, 0, 0},
		{"KEYLIGHT", 0, 0, 0},
		{"TIME", 0, 0, 0},
	}
	if !reflect.DeepEqual(getparms, rig.Caps.GetParameters) {
		t.Fatal("Caps.GetParameters contains unexpected data; expected:", getparms, "got:", rig.Caps.GetParameters)
	}

	setparms := Values{
		{"ANN", 0, 0, 0},
		{"APO", 0, 0, 0},
		{"BACKLIGHT", 0, 0, 0},
		{"BEEP", 0, 0, 0},
		{"KEYLIGHT", 0, 0, 0},
		{"TIME", 0, 0, 0},
	}
	if !reflect.DeepEqual(setparms, rig.Caps.SetParameters) {
		t.Fatal("Caps.SetParameters contains unexpected data; expected:", setparms, "got:", rig.Caps.SetParameters)
	}

	// tvfos := []int{}
	// if !reflect.DeepEqual(tvfos, rig.Caps.TargetableVfos) {
	// 	t.Fatal("Caps.TargetableVfos contains unexpected data; expected:", tvfos, "got:", rig.Caps.TargetableVfos)
	// }
	if len(rig.Caps.TargetableVfos) != 0 {
		t.Fatal("Caps.TargetableVfos was originally intended to be empty for DummyRig; got:", rig.Caps.TargetableVfos)
	}

	filters := map[string][]int{
		"WFM":  []int{230000},
		"RTTY": []int{2400},
		"USB":  []int{2400},
		"CW":   []int{2400, 500},
		"LSB":  []int{2400},
		"AM":   []int{8000, 2400},
		"FM":   []int{15000, 8000},
	}
	if !reflect.DeepEqual(filters, rig.Caps.Filters) {
		t.Fatal("Caps.Filters contains unexpected data; expected:", filters, "got:", rig.Caps.Filters)
	}

}

//Tests Frequency Get & Sep of the dummyRig
func TestDummyRigSetGetFreq(t *testing.T) {

	rig, err := getDummyRig()
	if err != nil {
		t.Fatal(err)
	}

	defer rig.Close()
	defer rig.Cleanup()

	freq, err := rig.GetFreq(RIG_VFO_A)
	if err != nil {
		t.Fatal(err)
	}

	if freq != 145000000 {
		t.Fatal("frequency of Dummyrig should be 145.000.000 Hz")
	}

	var testFreq float64
	testFreq = 7005000

	// Test Set & Get Frequency on all available VFOs
	for _, vfo := range rig.Caps.Vfos {
		vfoValue := VfoValue[vfo]
		if err := rig.SetFreq(vfoValue, testFreq); err != nil {
			t.Fatal(err)
		}

		freq, err = rig.GetFreq(vfoValue)
		if err != nil {
			t.Fatal(err)
		}

		if freq != testFreq {
			t.Fatalf("Could not set/get Frequency for %s", vfo)
		}
	}
}

//Tests vfo set & go of the dummyRig
func TestDummyRigSetGetVfo(t *testing.T) {

	rig, err := getDummyRig()
	if err != nil {
		t.Fatal(err)
	}

	defer rig.Close()
	defer rig.Cleanup()

	// Test Set & Get Frequency on all available VFOs
	for _, vfo := range rig.Caps.Vfos {
		vfoValue := VfoValue[vfo]
		if err := rig.SetVfo(vfoValue); err != nil {
			t.Fatal(err)
		}

		gvfo, err := rig.GetVfo()
		if err != nil {
			t.Fatal(err)
		}

		if gvfo != vfoValue {
			t.Fatalf("Could not set/get Vfo: %s", vfo)
		}
	}
}

func TestDummyRigModeAndFilters(t *testing.T) {

	rig, err := getDummyRig()
	if err != nil {
		t.Fatal(err)
	}

	defer rig.Close()
	defer rig.Cleanup()

	// for _, mode := range rig.Caps.Modes {
	// 	fmt.Println("Mode:", mode)
	// 	fmt.Println("Filter:", rig.Caps.Filters[mode])
	// }

	// interate over all VFOs
	for _, vfo := range rig.Caps.Vfos {
		vfoValue := VfoValue[vfo]

		// iterate over all modes
		for _, mode := range rig.Caps.Modes {
			modeValue := ModeValue[mode]

			filterAmount := len(rig.Caps.Filters[mode])
			if filterAmount > 0 {
				// iterate over all available filters
				for _, filter := range rig.Caps.Filters[mode] {

					if err := rig.SetMode(vfoValue, modeValue, filter); err != nil {
						t.Fatal(err)
					}
					m, pb, err := rig.GetMode(vfoValue)
					if err != nil {
						t.Fatal(err)
					}
					if m != modeValue {
						t.Fatalf("got mode %s which is inconsistent with set mode %s on vfo %s\n", ModeName[m], mode, vfo)
					}
					if pb != filter {
						t.Fatalf("got filter %dHz which is inconsistent with set filter %dHz for mode %s on vfo %s\n", pb, filter, mode, vfo)
					}
				}
			} else {
				// Not sure if this is the desired behaviour
				// How should a real rig respond to a filter it does not have?
				filter := 500 // Hz
				rig.SetMode(vfoValue, modeValue, filter)
				m, pb, err := rig.GetMode(vfoValue)
				if err != nil {
					t.Fatal(err)
				}
				if m != modeValue {
					t.Fatalf("got mode %s which is inconsistent with set mode %s on vfo %s\n", ModeName[m], mode, vfo)
				}
				if pb != filter {
					t.Fatalf("got filter %dHz which is inconsistent with set filter %dHz for mode %s on vfo %s\n", pb, filter, mode, vfo)
				}
			}
		}
	}

	modeValue := RIG_MODE_CW
	modeName := ModeName[modeValue]

	// Get Narrow Filter for CW
	filter, err := rig.GetPbNarrow(modeValue)
	if err != nil {
		t.Fatal("Could not determin Narrow Passband filter for mode:", modeName)
	}
	if filter != 500 {
		t.Fatalf("Expected for Narrow Passband filter in %s 500Hz; got: %d", modeName, filter)
	}

	// Get Normal Filter for CW
	filter, err = rig.GetPbNormal(modeValue)
	if err != nil {
		t.Fatal("Could not determin Normal Passband filter for mode:", modeName)
	}
	if filter != 2400 {
		t.Fatalf("Expected for Normal Passband filter in %s 2400Hz; got: %d", modeName, filter)
	}

	// Get Wide Filter for CW
	filter, err = rig.GetPbNarrow(modeValue)
	if err != nil {
		t.Fatal("Could not determin Narrow Passband filter for mode:", modeName)
	}
	if filter != 500 {
		t.Fatalf("Expected for Wide Passband filter in %s 500Hz; got: %d", modeName, filter)
	}
}

func TestDummyRigPTT(t *testing.T) {

	rig, err := getDummyRig()
	if err != nil {
		t.Fatal(err)
	}

	defer rig.Close()
	defer rig.Cleanup()

	// Test Set & Get PTT on all available VFOs
	for _, vfo := range rig.Caps.Vfos {
		vfoValue := VfoValue[vfo]

		// iterate over all PTT possibilities
		for pttValue, pttName := range PttName {

			if err := rig.SetPtt(vfoValue, pttValue); err != nil {
				t.Fatal(err)
			}

			ptt, err := rig.GetPtt(vfoValue)
			if err != nil {
				t.Fatal(err)
			}

			if pttValue == RIG_PTT_OFF {
				if ptt != RIG_PTT_OFF {
					t.Fatalf("inconsisted values! Set: %s, Should be: %s", pttName, PttName[RIG_PTT_OFF])
				}
			} else {
				// Dummy Rig returns just "ON" also for "ON_DATA" and "ON_MIC"
				if ptt != RIG_PTT_ON {
					t.Fatalf("inconsisted values! Set: %s, Got: %s", pttName, PttName[RIG_PTT_ON])
				}
			}
		}
	}
}

func TestDummyRigRit(t *testing.T) {

	rig, err := getDummyRig()
	if err != nil {
		t.Fatal(err)
	}

	defer rig.Close()
	defer rig.Cleanup()

	// Test Set & Get RIT on all available VFOs
	for _, vfo := range rig.Caps.Vfos {
		vfoValue := VfoValue[vfo]

		maxRit := rig.Caps.MaxRit

		ritTestValues := []int{-maxRit, -5330, -1, 0, 1, 100, maxRit}
		for _, ritTV := range ritTestValues {
			if err := rig.SetRit(vfoValue, ritTV); err != nil {
				t.Fatal(err)
			}

			rit, err := rig.GetRit(vfoValue)
			if err != nil {
				t.Fatal(err)
			}

			if rit != ritTV {
				t.Fatalf("rit value set (%dHz) does not correspond with read rit value (%dHz)", ritTV, rit)
			}
		}

		// // When rit > maxRit, then maxRit shall be set
		// ritOutOfRangeNegative := -maxRit - 1
		// if err := rig.SetRit(vfoValue, ritOutOfRangeNegative); err != nil {
		// 	t.Fatal(err)
		// }

		// rit, err := rig.GetRit(vfoValue)
		// if err != nil {
		// 	t.Fatal(err)
		// }

		// if rit != -maxRit {
		// 	t.Fatalf("When the set RIT > maxRit, then maxRit should be set; got: %dHz, should be: %dHz", -rit, -maxRit)
		// }
	}
}

func TestDummyRigXit(t *testing.T) {

	rig, err := getDummyRig()
	if err != nil {
		t.Fatal(err)
	}

	defer rig.Close()
	defer rig.Cleanup()

	// Test Set & Get XIT on all available VFOs
	for _, vfo := range rig.Caps.Vfos {
		vfoValue := VfoValue[vfo]

		maxXit := rig.Caps.MaxXit

		xitTestValues := []int{-maxXit, -5330, -1, 0, 1, 100, maxXit}
		for _, xitTV := range xitTestValues {

			if err := rig.SetXit(vfoValue, xitTV); err != nil {
				t.Fatal(err)
			}

			xit, err := rig.GetXit(vfoValue)
			if err != nil {
				t.Fatal(err)
			}

			if xit != xitTV {
				t.Fatalf("xit value set (%dHz) does not correspond with read xit value (%dHz)", xitTV, xit)
			}
		}
	}
}

func TestDummyRigSplitOperations(t *testing.T) {

	rig, err := getDummyRig()
	if err != nil {
		t.Fatal(err)
	}

	defer rig.Close()
	defer rig.Cleanup()

	// TBD Write test case for split operations
}

func TestDummyRigInfo(t *testing.T) {

	rig, err := getDummyRig()
	if err != nil {
		t.Fatal(err)
	}

	defer rig.Close()
	defer rig.Cleanup()

	info, err := rig.GetInfo()

	info_expected := "Nothing much (dummy)"
	if info != info_expected {
		t.Fatalf("info string does not match! got: %s; but expected: %s", info, info_expected)
	}
}

func TestDummyRigPowerStatus(t *testing.T) {

	rig, err := getDummyRig()
	if err != nil {
		t.Fatal(err)
	}

	defer rig.Close()
	defer rig.Cleanup()

	for pValue, pName := range RigPowerName {
		if err := rig.SetPowerStat(pValue); err != nil {
			t.Fatal(err)
		}

		pwrStatus, err := rig.GetPowerStat()
		if err != nil {
			t.Fatal(err)
		}

		if pwrStatus != pValue {
			t.Fatalf("inconsistent values! Set %s; Got: %s", pName, RigPowerName[pwrStatus])
		}
	}
}

func TestDummyRigAntennas(t *testing.T) {

	rig, err := getDummyRig()
	if err != nil {
		t.Fatal(err)
	}
	defer rig.Close()
	defer rig.Cleanup()

	// Test Set & Get Antenna on all available VFOs
	for _, vfo := range rig.Caps.Vfos {
		vfoValue := VfoValue[vfo]

		// iterate over all Antennas
		for aValue, aName := range AntName {

			if err := rig.SetAnt(vfoValue, aValue); err != nil {
				t.Fatal(err)
			}

			ant, err := rig.GetAnt(vfoValue)
			if err != nil {
				t.Fatal(err)
			}

			if ant != aValue {
				t.Fatalf("inconsisted values! Set: %s, Should be: %s", aName, AntName[ant])
			}
		}
	}

	// if err := rig.SetAnt(RIG_VFO_A, 1<<5); err == nil {
	// 	t.Fatal("An invalid Antenna (e.g. 32) should have thrown an error")
	// }
}

// missing function rig_get_resolution
func TestDummyRigTuningSteps(t *testing.T) {

	rig, err := getDummyRig()
	if err != nil {
		t.Fatal(err)
	}
	defer rig.Close()
	defer rig.Cleanup()

	cwRes, err := rig.GetRigResolution(RIG_MODE_CW)
	if err != nil {
		t.Fatal(err)
	}
	if cwRes != 1 {
		t.Fatalf("inconsistent values of rig resolution! Got: %d, should be: %d", cwRes, 1)
	}

	cwTs := rig.Caps.TuningSteps["CW"]

	// Test Set & Get Tuning Steps on all available VFOs
	for _, vfo := range rig.Caps.Vfos {
		vfoValue := VfoValue[vfo]

		// iterate over all Tuning Steps
		for _, ts := range cwTs {

			freq, err := rig.GetFreq(vfoValue)
			if err != nil {
				t.Fatal(err)
			}

			err = rig.SetTs(vfoValue, ts)
			if err != nil {
				t.Fatal(err)
			}

			tsRead, err := rig.GetTs(vfoValue)
			if err != nil {
				t.Fatal(err)
			}

			if tsRead != ts {
				t.Fatalf("inconsistent tuning steps! Got: %d, should be: %d", tsRead, ts)
			}

			// only execute if "UP" is supported
			for _, val := range rig.Caps.Operations {
				if val == "RIG_OP_UP" {

					err = rig.VfoOp(vfoValue, RIG_OP_UP)
					if err != nil {
						t.Fatal(err)
					}

					freqIncremented, err := rig.GetFreq(vfoValue)
					if err != nil {
						t.Fatal(err)
					}

					if freqIncremented != freq+float64(ts) {
						t.Fatalf("inconsistent frequencies; Tuning step: %dHz! Got: %.f, should be: %.f", ts, freqIncremented, freq+float64(ts))
					}
				}
			}
		}
	}
}

func TestDummyRigLevels(t *testing.T) {

	rig, err := getDummyRig()
	if err != nil {
		t.Fatal(err)
	}
	defer rig.Close()
	defer rig.Cleanup()

	// Test Has, Set, Get Level & GetLevelGran on all available VFOs
	// for _, vfo := range rig.Caps.Vfos {
	//	vfoValue := VfoValue[vfo]

	// }
}