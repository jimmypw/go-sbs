package sbs

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		args    args
		want    SBSMessage
		wantErr bool
	}{
		{
			name: "invalid message",
			args: args{
				message: "asdjisogijdgo",
			},
			wantErr: true,
		},
		{
			name: "valid message, invalid number of parameters",
			args: args{
				message: "MSG,1,asd,asd,ads",
			},
			wantErr: true,
		},
		{
			name: "invalid subtype",
			args: args{
				message: "MSG,0,dfuijsgiud",
			},
			wantErr: true,
		},
		{
			name: "valid msg1",
			args: args{
				message: "MSG,1,145,256,7404F2,11267,2008/11/28,23:48:18.611,2008/11/28,23:53:19.161,RJA1118,,,,,,,,,,,",
			},
			want: IDMessage{
				Preamble: Preamble{
					SessionID:  "145",
					AircraftID: "256",
					HexID:      "7404F2",
					FlightID:   "11267",
					GenDate:    "2008/11/28",
					GenTime:    "23:48:18.611",
					LogDate:    "2008/11/28",
					LogTime:    "23:53:19.161",
				},
				Callsign: "RJA1118",
			},
		},
		//TODO MSG,2
		{
			name: "valid msg3",
			args: args{
				message: "MSG,3,,,7272AB,,,,,,,12000,,,59.52123,-20.53678,,,0,0,0,0",
			},
			want: AirbornePosition{
				Preamble: Preamble{
					SessionID:  "",
					AircraftID: "",
					HexID:      "7272AB",
					FlightID:   "",
					GenDate:    "",
					GenTime:    "",
					LogDate:    "",
					LogTime:    "",
				},
				Altitude:  "12000",
				Lat:       "59.52123",
				Long:      "-20.53678",
				Alert:     "0",
				Emergency: "0",
				SPI:       "0",
				OnGround:  "0",
			},
		},
		{
			name: "valid msg4",
			args: args{
				message: "MSG,4,,,7272AB,,,,,,,,256,143,,,5,,0,0,0,0",
			},
			want: AirborneVelocity{
				Preamble: Preamble{
					SessionID:  "",
					AircraftID: "",
					HexID:      "7272AB",
					FlightID:   "",
					GenDate:    "",
					GenTime:    "",
					LogDate:    "",
					LogTime:    "",
				},
				GroundSpeed:  "256",
				Track:        "143",
				VerticalRate: "5",
			},
		},
		{
			name: "valid msg5",
			args: args{
				message: "MSG,5,,,7272AB,,,,,,,24775,,,,,,,0,0,0,0",
			},
			want: SurveillanceAlt{
				Preamble: Preamble{
					SessionID:  "",
					AircraftID: "",
					HexID:      "7272AB",
					FlightID:   "",
					GenDate:    "",
					GenTime:    "",
					LogDate:    "",
					LogTime:    "",
				},
				Altitude: "24775",
				Alert:    "0",
				SPI:      "0",
				OnGround: "0",
			},
		},
		{
			name: "valid msg6",
			args: args{
				message: "MSG,6,,,7272AB,,,,,,,,,,,,,1153,0,0,0,0",
			},
			want: SurveillanceID{
				Preamble: Preamble{
					SessionID:  "",
					AircraftID: "",
					HexID:      "7272AB",
					FlightID:   "",
					GenDate:    "",
					GenTime:    "",
					LogDate:    "",
					LogTime:    "",
				},
				Altitude:  "",
				Squawk:    "1153",
				Alert:     "0",
				Emergency: "0",
				SPI:       "0",
				OnGround:  "0",
			},
		},
		{
			name: "valid msg7",
			args: args{
				message: "MSG,7,,,7272AB,,,,,,,,,,,,,,,,,",
			},
			want: AirToAir{
				Preamble: Preamble{
					SessionID:  "",
					AircraftID: "",
					HexID:      "7272AB",
					FlightID:   "",
					GenDate:    "",
					GenTime:    "",
					LogDate:    "",
					LogTime:    "",
				},
				Altitude: "",
				OnGround: "",
			},
		},
		{
			name: "valid msg8",
			args: args{
				message: "MSG,8,,,7272AB,,,,,,,,,,,,,,,,,",
			},
			want: AllCall{
				Preamble: Preamble{
					SessionID:  "",
					AircraftID: "",
					HexID:      "7272AB",
					FlightID:   "",
					GenDate:    "",
					GenTime:    "",
					LogDate:    "",
					LogTime:    "",
				},
				OnGround: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
