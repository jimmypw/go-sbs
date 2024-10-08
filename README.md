# go-sbs

A library for decoding SBS message format

## Spec

copied from http://woodair.net/SBS/Article/Barebones42_Socket_Data.htm


Overview

Users can look at the raw data being sent by the SBS unit by using a Telnet application to listen to port 30003.

The datastream looks like this:

STA,,5,179,400AE7,10103,2008/11/28,14:58:51.153,2008/11/28,14:58:51.153,RM
MSG,4,5,211,4CA2D6,10057,2008/11/28,14:53:49.986,2008/11/28,14:58:51.153,,,408.3,146.4,,,64,,,,,
MSG,8,5,211,4CA2D6,10057,2008/11/28,14:53:50.391,2008/11/28,14:58:51.153,,,,,,,,,,,,0
MSG,4,5,211,4CA2D6,10057,2008/11/28,14:53:50.391,2008/11/28,14:58:51.153,,,408.3,146.4,,,64,,,,,
MSG,3,5,211,4CA2D6,10057,2008/11/28,14:53:50.594,2008/11/28,14:58:51.153,,37000,,,51.45735,-1.02826,,,0,0,0,0
MSG,8,5,812,ABBEE3,10095,2008/11/28,14:53:50.594,2008/11/28,14:58:51.153,,,,,,,,,,,,0
MSG,3,5,276,4010E9,10088,2008/11/28,14:53:49.986,2008/11/28,14:58:51.153,,28000,,,53.02551,-2.91389,,,0,0,0,0
MSG,4,5,276,4010E9,10088,2008/11/28,14:53:50.188,2008/11/28,14:58:51.153,,,459.4,20.2,,,64,,,,,
MSG,8,5,276,4010E9,10088,2008/11/28,14:53:50.594,2008/11/28,14:58:51.153,,,,,,,,,,,,0
MSG,3,5,276,4010E9,10088,2008/11/28,14:53:50.594,2008/11/28,14:58:51.153,,28000,,,53.02677,-2.91310,,,0,0,0,0
MSG,4,5,769,4CA2CB,10061,2008/11/28,14:53:50.188,2008/11/28,14:58:51.153,,,367.7,138.6,,,-2432,,,,,
MSG,8,5,769,4CA2CB,10061,2008/11/28,14:53:50.391,2008/11/28,14:58:51.153,,,,,,,,,,,,0

Decoding this stream isn't difficult.

Message types

There are six message types - MSG, SEL, ID, AIR, STA, CLK. Most data from aircraft is contained in the MSG lines whilst the other types are triggered by user input or system settings. The MSG data was inhibited with a five minute delay in BaseStation versions prior to 1.2.3.145 but from this version onwards is in real time.
ID
	
Type
	
Description
SEL
	 SELECTION CHANGE MESSAGE 	
 Generated when the user changes the selected aircraft in BaseStation.
ID
	 NEW ID MESSAGE 	
 Generated when an aircraft being tracked sets or changes its callsign.
AIR
	 NEW AIRCRAFT MESSAGE 	
 Generated when the SBS picks up a signal for an aircraft that it isn't
 currently tracking.
STA
	 STATUS CHANGE MESSAGE 	
 Generated when an aircraft's status changes according to the time-out
 values in the Data Settings menu.
CLK
	 CLICK MESSAGE 	
 Generated when the user double-clicks (or presses return) on an aircraft
 (i.e. to bring up the aircraft details window).
MSG
	 TRANSMISSION MESSAGE 	
 Generated by the aircraft. There are eight different MSG types.

Transmission messages (MSG) from aircraft may be one of eight types:
ID
	
Type
	  	
Description
MSG,1
	 ES Identification and Category
	
DF17 BDS 0,8
	  
MSG,2
	 ES Surface Position Message 	
DF17 BDS 0,6
	 Triggered by nose gear squat switch.
MSG,3
	 ES Airborne Position Message 	
DF17 BDS 0,5
	  
MSG,4
	 ES Airborne Velocity Message 	
DF17 BDS 0,9
	  
MSG,5
	 Surveillance Alt Message 	
DF4, DF20
	 Triggered by ground radar. Not CRC secured.
  MSG,5 will only be output if  the aircraft has previously sent a
  MSG,1, 2, 3, 4 or 8 signal.
MSG,6
	 Surveillance ID Message 	
DF5, DF21
	 Triggered by ground radar. Not CRC secured.
  MSG,6 will only be output if  the aircraft has previously sent a
  MSG,1, 2, 3, 4 or 8 signal.
MSG,7
	 Air To Air Message 	
DF16
	 Triggered from TCAS.
  MSG,7 is now included in the SBS socket output.
MSG,8
	 All Call Reply 	
DF11
	  Broadcast but also triggered by ground radar

 

Field Data

Each of the above message types may contain up to 22 data fields separated by commas. These fields are:
Field 1:
	 Message type 	 (MSG, STA, ID, AIR, SEL or CLK)
Field 2:
	 Transmission Type 	 MSG sub types 1 to 8. Not used by other message types.
Field 3:
	 Session ID 	 Database Session record number
Field 4:
	 AircraftID 	 Database Aircraft record number
Field 5:
	 HexIdent 	 Aircraft Mode S hexadecimal code
Field 6:
	 FlightID 	 Database Flight record number
Field 7:
	 Date message generated 	  As it says
Field 8:
	 Time message generated 	  As it says
Field 9:
	 Date message logged 	  As it says
Field 10:
	 Time message logged 	  As it says

 The above basic data fields are standard for all messages (Field 2 used only for MSG).

The fields below contain specific aircraft information.
Field 11:
	 Callsign 	 An eight digit flight ID - can be flight number or registration (or even nothing).
Field 12:
	 Altitude 	 Mode C altitude. Height relative to 1013.2mb (Flight Level). Not height AMSL..
Field 13:
	 GroundSpeed 	 Speed over ground (not indicated airspeed)
Field 14:
	 Track 	 Track of aircraft (not heading). Derived from the velocity E/W and velocity N/S
Field 15:
	 Latitude 	 North and East positive. South and West negative.
Field 16:
	 Longitude 	 North and East positive. South and West negative.
Field 17:
	 VerticalRate 	 64ft resolution
Field 18:
	 Squawk 	 Assigned Mode A squawk code.
Field 19:
	 Alert (Squawk change) 	 Flag to indicate squawk has changed.
Field 20:
	 Emergency 	 Flag to indicate emergency code has been set
Field 21:
	 SPI (Ident) 	 Flag to indicate transponder Ident has been activated.
Field 22:
	 IsOnGround 	 Flag to indicate ground squat switch is active

Notes (Courtesy of Edgy):

The socket data outputs a -1 for true, and a 0 for false. Neither means it is not used.

Field 11 (Callsign) is an 8 character (6 bit ASCII subset) field. In BaseStation a NULL is shown as a '@' which is ASCII for NULL. In the cockpit it is just a space on the transponder window, but is sent as a NULL. Therefore, if a crew enter eight spaces in the cockpit this will show in BaseStation as @@@@@@@@.

Field 12 (Altitude) can be 25ft or 100 foot resolution. Mode-C is 100 ft, but many aircraft today send out 25 ft resolution to be able to fly in Europe IFR (RVSM) space. BaseStation only displays Barometer altitude but in the data are HAE (height above ellipsoid), which is sent as the difference between GPS altitude and barometric altitude.

 

Message Content

Each message type contains different field content. In the table below green represents the fields that are sent and grey shows fields for which null data is transmitted. MSG signals contain up to 22 fields and other message types contain up to 10 fields.
  	
1
	
2
	
3
	
4
	
5
	
6
	
7
	
8
	
9
	
10
	  	
11
	
12
	
13
	
14
	
15
	
16
	
17
	
18
	
19
	
20
	
21
	
22
MSG 1
	
MT
	
TT
	
SID
	
AID
	
Hex
	
FID
	
DMG
	
TMG
	
DML
	
TML
	  	
CS
	
	
	
	
	
	
	
	
	
	
	
MSG 2
	  	  	  	  	  	  	  	  	  	  	  	
	
Alt
	
GS
	
Trk
	
Lat
	
Lng
	
	
	
	
	
	
Gnd
MSG 3
	  	  	  	  	  	  	  	  	  	  	  	
	
Alt
	
	
	
Lat
	
LNG
	
	
	
Alrt
	
Emer
	
SPI
	
Gnd
MSG 4
	  	  	  	  	  	  	  	  	  	  	  	
	
	
GS
	
Trk
	
	
	
VR
	
	
	
	
	
MSG 5
	  	  	  	  	  	  	  	  	  	  	  	
	
Alt
	
	
	
	
	
	
	
Alrt
	
	
SPI
	
Gnd
MSG 6
	  	  	  	  	  	  	  	  	  	  	  	
	
Alt
	
	
	
	
	
	
Sq
	
Alrt
	
Emer
	
SPI
	
Gnd
MSG 7
	  	  	  	  	  	  	  	  	  	  	  	
	
Alt
	
	
	
	
	
	
	
	
	
	
Gnd
MSG 8
	  	  	  	  	  	  	  	  	  	  	  	
	
	
	
	
	
	
	
	
	
	
	
Gnd
SEL
	  	  	  	  	  	  	  	  	  	  	  	
CS
	 
ID
	  	  	  	  	  	  	  	  	  	  	  	
CS
	 
AIR
	  	  	  	  	  	  	  	  	  	  	  	  	 
STA
	  	  	  	  	  	  	  	  	  	  	  	  	 
CLK
	  	  	  	
-1
	  	
-1
	  	  	  	  	  	  	 

Notes:

1. STA message uses the callsign field to record status flags based on user time-out values. Values are PL (Position Lost), SL (Signal Lost), RM (Remove), AD (Delete) and OK (used to reset time-outs if aircraft returns into cover).

2. CLK message returns a value of -1 in Fields 4 and 6. Field 5 is null.

3. MSG,7 (Air to Air message) has only recently been included in the socket output.

4. Although aircraft now transmit Heading and True Airspeed these values are not available in the socket output.


From the above table you see that MSG,1 messages only send data for the first eleven fields and the remaining 11 fields are empty. The result is a lot of commas in this (and in other MSG formats).

Examples of each message:
SEL,,496,2286,4CA4E5,27215,2010/02/19,18:06:07.710,2010/02/19,18:06:07.710,RYR1427
ID,,496,7162,405637,27928,2010/02/19,18:06:07.115,2010/02/19,18:06:07.115,EZY691A
AIR,,496,5906,400F01,27931,2010/02/19,18:06:07.128,2010/02/19,18:06:07.128
STA,,5,179,400AE7,10103,2008/11/28,14:58:51.153,2008/11/28,14:58:51.153,RM
CLK,,496,-1,,-1,2010/02/19,18:18:19.036,2010/02/19,18:18:19.036
MSG,1,145,256,7404F2,11267,2008/11/28,23:48:18.611,2008/11/28,23:53:19.161,RJA1118,,,,,,,,,,,
MSG,2,496,603,400CB6,13168,2008/10/13,12:24:32.414,2008/10/13,12:28:52.074,,,0,76.4,258.3,54.05735,-4.38826,,,,,,0
MSG,3,496,211,4CA2D6,10057,2008/11/28,14:53:50.594,2008/11/28,14:58:51.153,,37000,,,51.45735,-1.02826,,,0,0,0,0
MSG,4,496,469,4CA767,27854,2010/02/19,17:58:13.039,2010/02/19,17:58:13.368,,,288.6,103.2,,,-832,,,,,
MSG,5,496,329,394A65,27868,2010/02/19,17:58:12.644,2010/02/19,17:58:13.368,,10000,,,,,,,0,,0,0
MSG,6,496,237,4CA215,27864,2010/02/19,17:58:12.846,2010/02/19,17:58:13.368,,33325,,,,,,0271,0,0,0,0
MSG,7,496,742,51106E,27929,2011/03/06,07:57:36.523,2011/03/06,07:57:37.054,,3775,,,,,,,,,,0
MSG,8,496,194,405F4E,27884,2010/02/19,17:58:13.244,2010/02/19,17:58:13.368,,,,,,,,,,,,0


Interpolation

It can be seen that no single MSG type provides all the data we use in BaseStation and that some data fields are unique to one message type. Callsign is only found in MSG,1, VertRate only in MSG,4 and Squawk in MSG,6.

To collect all 11 data fields for one aircraft would require the reception of at least four MSG types (MSG,1, MSG,3, MSG,4 and MSG,6) but note that MSG,6 is only triggered by ground radar interrogation. If the aircraft is outside any ground radar coverage no MSG,6 will be sent. As MSG,6 is the only message that sends out the squawk code it means this will only be displayed for SBS users who are detecting aircraft within Mode S ground radar coverage.

Likewise MSG,5 and MSG,8 are only sent on interrogation but the data in these types is available in other messages.

MSG,5 and MSG,6 are not CRC secured and will only be received should an aircraft have already sent a MSG,1, 2, 3, 4 or 8.

Ground targets

If Field 22 (IsOnGround) is being sent this will trigger a change of values in Field 12 (Altitude) and in Fields 15 and 16 (Latitude and Longitude). Field 12 (Altitude) will reset to zero and whilst the aircraft remains on the ground no altitude data will be sent.

Positional Accuracy

The Compact Position Report in ADS-B sends Lat/Long data in 17 bits and when airborne this gives accuracy to 5.1 metres. 17 bits equates to four decimal places for Lat/Long values - e.g N54.1234, W145.1234. For ground operations greater positional accuracy is required and so Lat/Long values are extended to five decimal places - e.g. N54.12345, W145.12345 - which gives an accuracy of 1.25 metres.

To accommodate this accuracy into a 17 bit string some data is dropped - the full Lat/Long position data is no longer sent. In BaseStation the missing data needs to be added by the user and this is why Lat/Long values are only interpreted correctly if a location is set in the BaseStation Location Manager. For most users adding a home location in the Location Manager is sufficient but mobile users need to add further locations for the airfields they may be visiting abroad. By abroad I mean intercontinental as BaseStation now plots ground traffic correctly to within approx 2500nm of the set location (it formerly was only 90nm).

The socket data always shows Lat/Long data to five decimal places and can provide full 1.25 metre accuracy where this is sent.

Credits

My thanks go to Andy (Three Miles), Dave Reid and Steve (Edgy) who have posted much information on the Kinetic forum about socket data format. Without their research this page wouldn't exist.

 
BST files


Overview

Basestation has the an option to record data. This is not raw socket data as described above but processed data for the Basestation display.

Unlike raw data, each string in the BST files shows 17 data field values. All are populated, using last known values for each string until a socket MSG updates any values.

The recorded datastream looks like this:


"2018/07/05","02:44:34.126","9004131","896463","ETD44A","United Arab Emirates","0","39000","39000","52.05327","-3.81704","-64","-64","484.6","102.0","8726","2216"
"2018/07/05","02:44:34.142","4736069","484445","KLM656","Netherlands","0","41000","41000","55.11269","-3.75159","0","0","480.8","122.2","25347","6303"
"2018/07/05","02:44:34.153","10672439","A2D937","AAL716","United States","0","40000","40000","51.65419","-3.77826","-64","-64","474.5","95.9","8194","2002"
"2018/07/05","02:44:34.153","10895798","A641B6","DAL132","United States","0","41000","41000","53.92718","-2.84215","-64","-64","494.4","107.3","25364","6314"
"2018/07/05","02:44:34.178","9004131","896463","ETD44A","United Arab Emirates","0","39000","39000","52.05309","-3.81561","0","0","484.6","102.0","8726","2216"
"2018/07/05","02:44:35.122","4736069","484445","KLM656","Netherlands","0","41000","41000","55.11223","-3.75030","0","0","480.8","122.2","25347","6303"
"2018/07/05","02:44:35.124","10672439","A2D937","AAL716","United States","0","40000","40000","51.65408","-3.77647","-64","-64","474.5","95.9","8194","2002"
"2018/07/05","02:44:35.151","9004131","896463","ETD44A","United Arab Emirates","0","39000","39000","52.05286","-3.81393","0","0","484.6","102.0","8726","2216"
"2018/07/05","02:44:35.155","10895798","A641B6","DAL132","United States","0","41000","41000","53.92667","-2.83924","-64","-64","494.4","107.3","25364","6314"
"2018/07/05","02:44:35.454","4736069","484445","KLM656","Netherlands","0","41000","41000","55.11147","-3.74817","0","0","480.8","122.2","25347","6303"
"2018/07/05","02:44:35.460","10672439","A2D937","AAL716","United States","0","40000","40000","51.65396","-3.77462","-64","-64","474.5","95.9","8194","2002"

 

Field Data

The BST file contains 17 data fields separated by commas. These fields are:
Field 1:
	 Date message generated 	 Self evident
Field 2:
	 Time message generated 	 Self evident
Field 3:
	 Mode S Code (Decimal) 	 Aircraft Mode S decimal code
Field 4:
	 Mode S Code (Hex) 	 Aircraft Mode S hexadecimal code
Field 5:
	 Callsign 	 An eight digit flight ID - can be flight number or registration (or even nothing).
Field 6:
	 Country 	 Interpolated from Mode S code using the Countries.dat file.
Field 7:
	 IsOnGround 	 Flag to indicate ground squat switch is active
Field 8:
	 Altitude 	 Mode C altitude. Height relative to 1013.2mb (Flight Level). Not height AMSL..
Field 9:
	 Altitude 	 A placeholder for future development. Same data as above.
Field 10:
	 Latitude 	 North and East positive. South and West negative.
Field 11:
	 Longitude 	 North and East positive. South and West negative.
Field 12:
	 VerticalRate 	  64ft resolution
Field 13:
	 VerticalRate 	 Adjusted data for Basestation screen presentation.
Field 14:
	 GroundSpeed 	 Speed over ground (not indicated airspeed)
Field 15:
	 Track 	 Track of aircraft (not heading). Derived from the velocity E/W and velocity N/S
Field 16:
	 Squawk (Decimal) 	 Assigned Mode A squawk code, decimal.
Field 17:
	 Squawk (Octal) 	 Assigned Mode A squawk code, octal. Cockpit setting.