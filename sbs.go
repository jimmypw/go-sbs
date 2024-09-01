package sbs

import (
	"errors"
	"strings"
)

// Field Positions, order matters
const (
	F_MESSAGETYPE = iota
	F_TRANSMISSIONTYPE
	F_SESSIONID
	F_AIRCRAFTID
	F_HEXIDENT
	F_FLIGHTID
	F_DATEGEN
	F_TIMEGEN
	F_DATELOG
	F_TIMELOG
	F_CALLSIGN
	F_ALTITUDE
	F_GROUNDSPEED
	F_TRACK
	F_LATITUDE
	F_LONGITUDE
	F_VERTICALRATE
	F_SQUAWK
	F_ALERT
	F_EMERGENCY
	F_SPI
	F_ISONGROUND
)

type SBSMessage interface{}

type Preamble struct {
	SessionID  string
	AircraftID string
	HexID      string
	FlightID   string
	GenDate    string
	GenTime    string
	LogDate    string
	LogTime    string
	Callsign   string
}

// MSG,1 (Identification and Category)
type IDMessage struct {
	Preamble
	Callsign string
}

// MSG,2 ( Surface Position Message)
type SurfacePosition struct {
	Preamble
	Altitude    string
	GroundSpeed string
	Track       string
	Lat         string
	Long        string
	OnGround    string
}

// MSG,3 ( Airborne Position Message)
type AirbornePosition struct {
	Preamble
	Altitude  string
	Lat       string
	Long      string
	Alert     string
	Emergency string
	SPI       string
	OnGround  string
}

// MSG,4 ( Airborne Velocity Message)
type AirborneVelocity struct {
	Preamble
	GroundSpeed  string
	Track        string
	VerticalRate string
}

// MSG,5 ( Surveillance Alt Message)
type SurveillanceAlt struct {
	Preamble
	Altitude string
	Alert    string
	SPI      string
	OnGround string
}

// MSG,6 ( Surveillance ID Message)
type SurveillanceID struct {
	Preamble
	Altitude  string
	Squawk    string
	Alert     string
	Emergency string
	SPI       string
	OnGround  string
}

// MSG,7 ( Air To Air Message)
type AirToAir struct {
	Preamble
	Altitude string
	OnGround string
}

// MSG,8 ( All Call Reply)
type AllCall struct {
	Preamble
	OnGround string
}

func ParseMSG(msg []string) (SBSMessage, error) {
	switch msg[F_TRANSMISSIONTYPE] {
	case "1": // ES Identification and Category
		return ParseIDMessage(msg)
	case "2": // ES Surface Position Message
		return ParseSurfacePosition(msg)
	case "3": // ES Airborne Position Message
		return ParseAirbornePosition(msg)
	case "4": // ES Airborne Velocity Message
		return ParseAirborneVelocity(msg)
	case "5": // Surveillance Alt Message
		return ParseSurveillanceAlt(msg)
	case "6": // Surveillance ID Message
		return ParseSurveillanceID(msg)
	case "7": // Air To Air Message
		return ParseAirToAir(msg)
	case "8": // All Call Reply
		return ParseAllCall(msg)
	default:
		return nil, errors.New("unsupported transmission type")
	}
}

func Parse(message string) (SBSMessage, error) {
	splitstring := strings.Split(message, ",")
	if len(splitstring) != 22 {
		return nil, errors.New("invalid number of fields")
	}

	switch splitstring[F_MESSAGETYPE] {
	case "MSG":
		return ParseMSG(splitstring)
	default:
		return nil, errors.New("unsupported message type")
	}
}

func ParsePreamble(msg []string) Preamble {
	return Preamble{
		SessionID:  msg[F_SESSIONID],
		AircraftID: msg[F_AIRCRAFTID],
		HexID:      msg[F_HEXIDENT],
		FlightID:   msg[F_FLIGHTID],
		GenDate:    msg[F_DATEGEN],
		GenTime:    msg[F_TIMEGEN],
		LogDate:    msg[F_DATELOG],
		LogTime:    msg[F_TIMELOG],
	}
}

func ParseIDMessage(msg []string) (IDMessage, error) {
	return IDMessage{
		Preamble: ParsePreamble(msg),
		Callsign: msg[10],
	}, nil
}

func (m IDMessage) ToString() string {
	return "TODO"
}

func ParseSurfacePosition(msg []string) (SurfacePosition, error) {
	return SurfacePosition{
		Preamble:    ParsePreamble(msg),
		Altitude:    msg[F_ALTITUDE],
		GroundSpeed: msg[F_GROUNDSPEED],
		Track:       msg[F_TRACK],
		Lat:         msg[F_LATITUDE],
		Long:        msg[F_LONGITUDE],
		OnGround:    msg[F_ISONGROUND],
	}, nil
}

func (m SurfacePosition) ToString() string {
	return "TODO"
}

func ParseAirbornePosition(msg []string) (AirbornePosition, error) {
	return AirbornePosition{
		Preamble:  ParsePreamble(msg),
		Altitude:  msg[F_ALTITUDE],
		Lat:       msg[F_LATITUDE],
		Long:      msg[F_LONGITUDE],
		Alert:     msg[F_ALERT],
		Emergency: msg[F_EMERGENCY],
		SPI:       msg[F_SPI],
		OnGround:  msg[F_ISONGROUND],
	}, nil
}

func (m AirbornePosition) ToString() string {
	return "TODO"
}

func ParseAirborneVelocity(msg []string) (AirborneVelocity, error) {
	return AirborneVelocity{
		Preamble:     ParsePreamble(msg),
		GroundSpeed:  msg[F_GROUNDSPEED],
		Track:        msg[F_TRACK],
		VerticalRate: msg[F_VERTICALRATE],
	}, nil
}

func (m AirborneVelocity) ToString() string {
	return "TODO"
}

func ParseSurveillanceAlt(msg []string) (SurveillanceAlt, error) {
	return SurveillanceAlt{
		Preamble: ParsePreamble(msg),
		Altitude: msg[F_ALTITUDE],
		Alert:    msg[F_ALERT],
		SPI:      msg[F_SPI],
		OnGround: msg[F_ISONGROUND],
	}, nil
}

func (m SurveillanceAlt) ToString() string {
	return "TODO"
}

func ParseSurveillanceID(msg []string) (SurveillanceID, error) {
	return SurveillanceID{
		Preamble:  ParsePreamble(msg),
		Altitude:  msg[F_ALTITUDE],
		Squawk:    msg[F_SQUAWK],
		Alert:     msg[F_ALERT],
		Emergency: msg[F_EMERGENCY],
		SPI:       msg[F_SPI],
		OnGround:  msg[F_ISONGROUND],
	}, nil
}

func (m SurveillanceID) ToString() string {
	return "TODO"
}

func ParseAirToAir(msg []string) (AirToAir, error) {
	return AirToAir{
		Preamble: ParsePreamble(msg),
		Altitude: msg[F_ALTITUDE],
		OnGround: msg[F_ISONGROUND],
	}, nil
}

func (m AirToAir) ToString() string {
	return "TODO"
}

func ParseAllCall(msg []string) (AllCall, error) {
	return AllCall{
		Preamble: ParsePreamble(msg),
		OnGround: msg[F_ISONGROUND],
	}, nil
}

func (m AllCall) ToString() string {
	return "TODO"
}
