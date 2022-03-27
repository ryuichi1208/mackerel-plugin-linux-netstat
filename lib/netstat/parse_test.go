package netstat

import (
	"io"
	"os"
	"reflect"
	"testing"
)

func Test_source_read(t *testing.T) {
	type fields struct {
		path   string
		parser parser
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &source{
				path:   tt.fields.path,
				parser: tt.fields.parser,
			}
			got, err := s.read()
			if (err != nil) != tt.wantErr {
				t.Errorf("source.read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("source.read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitKeyValues(t *testing.T) {
	type args struct {
		s     string
		kvSep string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := splitKeyValues(tt.args.s, tt.args.kvSep)
			if (err != nil) != tt.wantErr {
				t.Errorf("splitKeyValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("splitKeyValues() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("splitKeyValues() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_parseSocket(t *testing.T) {
	path := "testdata/socket"
	fp, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}
	type args struct {
		file io.Reader
	}
	tests := []struct {
		name string
		args args
		want map[string]float64
	}{
		{
			name: "socket parse",
			args: args{
				file: fp,
			},
			want: map[string]float64{
				"tcp_memory_size": 4096,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSocket(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSocket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseCompact(t *testing.T) {
	fp, err := os.Open("testdata/netstat")
	if err != nil {
		t.Error(err)
	}
	type args struct {
		file io.Reader
	}
	tests := []struct {
		name string
		args args
		want map[string]float64
	}{
		{
			name: "test_parse_compact",
			args: args{
				file: fp,
			},
			want: map[string]float64{
				"IpExtInBcastOctets":              2.648922e+06,
				"IpExtInBcastPkts":                16958,
				"IpExtInCEPkts":                   0,
				"IpExtInCsumErrors":               0,
				"IpExtInECT0Pkts":                 36810,
				"IpExtInECT1Pkts":                 0,
				"IpExtInMcastOctets":              72,
				"IpExtInMcastPkts":                2,
				"IpExtInNoECTPkts":                430727,
				"IpExtInNoRoutes":                 2,
				"IpExtInOctets":                   7.41110769e+08,
				"IpExtInTruncatedPkts":            0,
				"IpExtOutBcastOctets":             0,
				"IpExtOutBcastPkts":               0,
				"IpExtOutMcastOctets":             0,
				"IpExtOutMcastPkts":               0,
				"IpExtOutOctets":                  1.12781653e+08,
				"IpExtReasmOverlaps":              0,
				"TcpExtArpFilter":                 0,
				"TcpExtBusyPollRxPackets":         0,
				"TcpExtDelayedACKLocked":          2,
				"TcpExtDelayedACKLost":            93,
				"TcpExtDelayedACKs":               1562,
				"TcpExtEmbryonicRsts":             0,
				"TcpExtIPReversePathFilter":       0,
				"TcpExtListenDrops":               0,
				"TcpExtListenOverflows":           0,
				"TcpExtLockDroppedIcmps":          0,
				"TcpExtOfoPruned":                 0,
				"TcpExtOutOfWindowIcmps":          17,
				"TcpExtPAWSActive":                0,
				"TcpExtPAWSEstab":                 11,
				"TcpExtPAWSPassive":               0,
				"TcpExtPFMemallocDrop":            0,
				"TcpExtPruneCalled":               0,
				"TcpExtRcvPruned":                 0,
				"TcpExtSyncookiesFailed":          6,
				"TcpExtSyncookiesRecv":            0,
				"TcpExtSyncookiesSent":            0,
				"TcpExtTCPACKSkippedChallenge":    0,
				"TcpExtTCPACKSkippedFinWait2":     0,
				"TcpExtTCPACKSkippedPAWS":         6,
				"TcpExtTCPACKSkippedSeq":          19,
				"TcpExtTCPACKSkippedSynRecv":      0,
				"TcpExtTCPACKSkippedTimeWait":     0,
				"TcpExtTCPAbortFailed":            0,
				"TcpExtTCPAbortOnClose":           0,
				"TcpExtTCPAbortOnData":            55,
				"TcpExtTCPAbortOnLinger":          0,
				"TcpExtTCPAbortOnMemory":          0,
				"TcpExtTCPAbortOnTimeout":         8,
				"TcpExtTCPAutoCorking":            1764,
				"TcpExtTCPBacklogDrop":            0,
				"TcpExtTCPChallengeACK":           0,
				"TcpExtTCPDSACKIgnoredNoUndo":     445,
				"TcpExtTCPDSACKIgnoredOld":        0,
				"TcpExtTCPDSACKOfoRecv":           0,
				"TcpExtTCPDSACKOfoSent":           0,
				"TcpExtTCPDSACKOldSent":           94,
				"TcpExtTCPDSACKRecv":              750,
				"TcpExtTCPDSACKUndo":              0,
				"TcpExtTCPDeferAcceptDrop":        0,
				"TcpExtTCPDirectCopyFromBacklog":  75525,
				"TcpExtTCPDirectCopyFromPrequeue": 4.8327202e+07,
				"TcpExtTCPFACKReorder":            0,
				"TcpExtTCPFastOpenActive":         0,
				"TcpExtTCPFastOpenActiveFail":     0,
				"TcpExtTCPFastOpenCookieReqd":     0,
				"TcpExtTCPFastOpenListenOverflow": 0,
				"TcpExtTCPFastOpenPassive":        0,
				"TcpExtTCPFastOpenPassiveFail":    0,
				"TcpExtTCPFastRetrans":            58,
				"TcpExtTCPForwardRetrans":         2,
				"TcpExtTCPFromZeroWindowAdv":      0,
				"TcpExtTCPFullUndo":               0,
				"TcpExtTCPHPAcks":                 42009,
				"TcpExtTCPHPHits":                 141523,
				"TcpExtTCPHPHitsToUser":           12656,
				"TcpExtTCPHystartDelayCwnd":       0,
				"TcpExtTCPHystartDelayDetect":     0,
				"TcpExtTCPHystartTrainCwnd":       296,
				"TcpExtTCPHystartTrainDetect":     9,
				"TcpExtTCPLossFailures":           0,
				"TcpExtTCPLossProbeRecovery":      431,
				"TcpExtTCPLossProbes":             818,
				"TcpExtTCPLossUndo":               1,
				"TcpExtTCPLostRetransmit":         0,
				"TcpExtTCPMD5NotFound":            0,
				"TcpExtTCPMD5Unexpected":          0,
				"TcpExtTCPMemoryPressures":        0,
				"TcpExtTCPMinTTLDrop":             0,
				"TcpExtTCPOFODrop":                0,
				"TcpExtTCPOFOMerge":               0,
				"TcpExtTCPOFOQueue":               94883,
				"TcpExtTCPOrigDataSent":           183259,
				"TcpExtTCPPartialUndo":            0,
				"TcpExtTCPPrequeueDropped":        0,
				"TcpExtTCPPrequeued":              41815,
				"TcpExtTCPPureAcks":               103466,
				"TcpExtTCPRcvCoalesce":            116830,
				"TcpExtTCPRcvCollapsed":           0,
				"TcpExtTCPRenoFailures":           0,
				"TcpExtTCPRenoRecovery":           0,
				"TcpExtTCPRenoRecoveryFail":       0,
				"TcpExtTCPRenoReorder":            0,
				"TcpExtTCPReqQFullDoCookies":      0,
				"TcpExtTCPReqQFullDrop":           0,
				"TcpExtTCPRetransFail":            0,
				"TcpExtTCPSACKDiscard":            0,
				"TcpExtTCPSACKReneging":           0,
				"TcpExtTCPSACKReorder":            0,
				"TcpExtTCPSYNChallenge":           0,
				"TcpExtTCPSackFailures":           0,
				"TcpExtTCPSackMerged":             35,
				"TcpExtTCPSackRecovery":           2,
				"TcpExtTCPSackRecoveryFail":       0,
				"TcpExtTCPSackShiftFallback":      13,
				"TcpExtTCPSackShifted":            58,
				"TcpExtTCPSchedulerFailed":        0,
				"TcpExtTCPSlowStartRetrans":       0,
				"TcpExtTCPSpuriousRTOs":           0,
				"TcpExtTCPSpuriousRtxHostQueues":  0,
				"TcpExtTCPSynRetrans":             79,
				"TcpExtTCPTSReorder":              0,
				"TcpExtTCPTimeWaitOverflow":       0,
				"TcpExtTCPTimeouts":               61,
				"TcpExtTCPToZeroWindowAdv":        0,
				"TcpExtTCPWantZeroWindowAdv":      0,
				"TcpExtTCPWqueueTooBig":           0,
				"TcpExtTW":                        253,
				"TcpExtTWKilled":                  0,
				"TcpExtTWRecycled":                0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseCompact(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseCompact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_get(t *testing.T) {
	tests := []struct {
		name string
		want map[string]float64
	}{
		{
			name: "Test_get",
			want: map[string]float64{
				"test": 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			get()
		})
	}
}
