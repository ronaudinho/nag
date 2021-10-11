package params

import (
	"fmt"
	"strconv"
	"time"
)

type AheadBehind string

var (
	AheadOrBehind      AheadBehind = "Ahead or Behind"
	BehindOrTied                   = "Behind or Tied"
	AheadOrTied                    = "Ahead or Tied"
	DefaultAheadBehind             = AheadOrBehind
)

type ClutchTime string

var (
	Last5Minutes      ClutchTime = "Last 5 Minutes"
	Last4Minutes                 = "Last 4 Minutes"
	Last3Minutes                 = "Last 3 Minutes"
	Last2Minutes                 = "Last 2 Minutes"
	Last1Minute                  = "Last 1 Minute"
	Last30Seconds                = "Last 30 Seconds"
	Last10Seconds                = "Last 10 Seconds"
	DefaultClutchTime            = Last5Minutes
)

type Conference string

var (
	EastConference    Conference = "East"
	WestConference               = "West"
	NoneConference               = ""
	DefaultConference            = NoneConference
)

// ContextMeasureSimple(_ContextMeasure):
// ContextMeasureDetailed(_ContextMeasure):

type DefenseCategory string

var (
	Overall                DefenseCategory = "Overall"
	Threes                                 = "3 Pointers"
	Twos                                   = "2 Pointers"
	LessThan6ft                            = "Less Than 6Ft"
	LessThan10ft                           = "Less Than 10Ft"
	GreaterThan15ft                        = "Greater Than 15Ft"
	DefaultDefenseCategory                 = Overall
)

type Direction string

var (
	Asc  Direction = "ASC"
	Desc           = "DESC"

	DefaultDirection = Asc
)

type DistanceRange string

var (
	Range5ft             DistanceRange = "5ft Range"
	Range8ft                           = "8ft Range"
	ByZone                             = "By Zone"
	DefaultDistanceRange               = ByZone
)

type DivisionSimple string

var (
	Atlantic              DivisionSimple = "Atlantic"
	Central                              = "Central"
	Northwest                            = "Northwest"
	Pacific                              = "Pacific"
	Southeast                            = "Southeast"
	Southwest                            = "Southwest"
	DefaultDivisionSimple                = Atlantic
)

type Division string

var (
	EastDivision    Division = "East"
	WestDivision             = "West"
	DefaultDivision          = EastDivision
)

type GameScopeSimple string

var (
	Last10                 GameScopeSimple = "Last 10"
	Yesterday                              = "Yesterday"
	DefaultGameScopeSimple                 = Last10
)

// GameScopeDetailed(GameScopeSimple):
type GameScopeDetailed string

var (
	SeasonGameScopeDetailed  GameScopeDetailed = "Season"
	FinalsGameScopeDetailed                    = "Finals"
	DefaultGameScopeDetailed                   = Season
)

type GameSegment string

var (
	FirstHalf          GameSegment = "First Half"
	SecondHalf                     = "Second Half"
	Overtime                       = "Overtime"
	DefaultGameSegment             = FirstHalf
)

var (
	GroupQuantity        = func(i int) string { return strconv.Itoa(i) }
	DefaultGroupQuantity = "5"
)

var (
	LastNGames        = func(i int) string { return strconv.Itoa(i) }
	DefaultLastNGames = "0"
)

type LeagueID string

const (
	NBA             LeagueID = "00"
	ABA                      = "01"
	WNBA                     = "10"
	GLeague                  = "20"
	DefaultLeagueID          = NBA
)

type Location string

var (
	Home            Location = "Home"
	Road                     = "Road"
	DefaultLocation          = Home
)

type MeasureTypeBase string

var (
	Base                   MeasureTypeBase = "Base"
	DefaultMeasureTypeBase                 = Base
)

// MeasureTypeSimple(MeasureTypeBase):
// MeasureTypePlayerGameLogs(MeasureTypeBase):

var (
	NumberOfGames        = func(i int) string { return strconv.Itoa(i) }
	DefaultNumberOfGames = "2147483647"
)

type Outcome string

var (
	Win            Outcome = "W"
	Loss                   = "L"
	DefaultOutcome         = Win
)

// PaceAdjust(_YesNo):
// PaceAdjustNo(_No):
// PlusMinus(_YesNo):
// PlusMinusNo(_No):

type Period string

var (
	AllPeriod      Period = "0"
	FirstPeriod           = "1"
	SecondPeriod          = "2"
	ThirdPeriod           = "3"
	FourthPeriod          = "4"
	QuarterPeriod         = func(i int) string { return strconv.Itoa(i) }
	OvertimePeriod        = func(i int) string { return strconv.Itoa(4 + i) }
	DefaultPeriod         = AllPeriod
)

// StartPeriod(Period):
// EndPeriod(Period):

type PerMode string

const (
	Totals         PerMode = "Totals"
	PerGame                = "PerGame"
	DefaultPerMode         = Totals
)

// PerMode36(PerModeSimple):
// PerMode48(PerModeSimple):
// PerModeTime(PerMode36, PerMode48):
// PerModeDetailed(PerModeTime):

type PlayerExperience string

var (
	Rookie                  PlayerExperience = "Rookie"
	Sophomore                                = "Sophomore"
	Veteran                                  = "Veteran"
	DefaultPlayerExperience                  = Rookie
)

type PlayerOrTeam string

const (
	Player              PlayerOrTeam = "Player"
	Team                             = "Team"
	DefaultPlayerOrTeam              = Team
)

type PlayerOrTeamAbbreviation string

var (
	P                               PlayerOrTeamAbbreviation = "P"
	T                                                        = "T"
	DefaultPlayerOrTeamAbbreviation                          = Team
)

type PlayerPosition string

var (
	Guard                 PlayerPosition = "Guard"
	Forward                              = "Forward"
	Center                               = "Center"
	DefaultPlayerPosition                = Guard
)

type PlayerPositionAbbreviation string

var (
	G                                 PlayerPositionAbbreviation = "G"
	F                                                            = "F"
	C                                                            = "C"
	GF                                                           = "G-F"
	FG                                                           = "F-G"
	FC                                                           = "F-C"
	CF                                                           = "C-F"
	DefaultPlayerPositionAbbreviation                            = G
)

type PlayerScope string

var (
	AllPlayers         PlayerScope = "All Players"
	Rookies                        = "Rookies"
	DefaultPlayerScope             = AllPlayers
)

// TodaysPlayers(_YesNo):
// ActivePlayers(_YesNo):

type PlayType string

var (
	Transition      PlayType = "Transition"
	Isolation                = "Isolation"
	PRBallHandler            = "PRBallHandler"
	PRRollMan                = "PRRollman"
	PostUp                   = "Postup"
	SpotUp                   = "Spotup"
	Handoff                  = "Handoff"
	Cut                      = "Cut"
	OffScreen                = "OffScreen"
	Putbacks                 = "OffRebound"
	Misc                     = "Misc"
	DefaultPlayType          = Transition
)

var (
	PointDiff        = func(i int) string { return strconv.Itoa(i) }
	DefaultPointDiff = "5"
)

type PtMeasureType string

var (
	SpeedDistance        PtMeasureType = "SpeedDistance"
	Rebounding                         = "Rebounding"
	Possessions                        = "Possessions"
	CatchShoot                         = "CatchShoot"
	PullUpShot                         = "PullUpShot"
	Defense                            = "Defense"
	Drives                             = "Drives"
	Passing                            = "Passing"
	ElbowTouch                         = "ElbowTouch"
	PostTouch                          = "PostTouch"
	PaintTouch                         = "PaintTouch"
	Efficiency                         = "Efficiency"
	DefaultPtMeasureType               = SpeedDistance
)

type RangeType string

var DefaultRangeType = "0"

// Rank(_YesNo):
// RankNo(_No):

type RunType string

var DefaultRunType RunType = "each second"

type StartRange string

var DefaultStartRange = "0"

type Scope string

var (
	RSScope      Scope = "RS"
	SScope             = "S"
	RookiesScope       = "Rookies"
	DefaultScope       = SScope
)

// SeasonYear:
var (
	Season = func(t time.Time) string {
		cur := t.Year()
		if t.Month() <= 9 {
			cur = cur - 1
		}
		nxt := strconv.Itoa(cur + 1)[2:]
		return fmt.Sprintf("%d-%s", cur, nxt)
	}
	CurrentSeason = Season(time.Now())
)

// SeasonAll(Season):
// SeasonAll_Time(Season):
// SeasonAllTime(Season):
// SeasonID(SeasonYear):

type SeasonType string

const (
	Regular           SeasonType = "Regular Season"
	PreSeason                    = "Pre Season"
	DefaultSeasonType            = Regular
)

// SeasonTypePlayoffs(SeasonType):
// SeasonTypeAllStar(SeasonTypePlayoffs):

type SeasonSegment string

var (
	PostAllStar          SeasonSegment = "Post All-Star"
	PreAllStar                         = "Pre All-Star"
	DefaultSeasonSegment               = PostAllStar
)

type ShotClockRange string

var (
	Range2224             ShotClockRange = "24-22"
	Range1822                            = "22-18 Very Early"
	Range1518                            = "18-15 Early"
	Range715                             = "15-7 Average"
	Range47                              = "7-4 Late"
	Range04                              = "4-0 Very Late"
	ShotClockOff                         = "ShotClock Off"
	Empty                                = ""
	DefaultShotClockRange                = ShotClockOff
)

/*
func CalculateRange(i int64) ShotClockRange {
	switch {
	case i > 24, i <= 0:
		return Empty
	case 22 < i <= 24:
		return Range2224
	case 18 < i <= 22:
		return Range1822
	case 15 < i <= 18:
		return Range1518
	case 7 < i <= 15:
		return Range715
	case 4 < i <= 7:
		return Range47
	case 0 < i <= 4:
		return Range04
	}
}
*/

type Sorter string

var (
	FGMSorter     Sorter = "FGM"
	FGASorter            = "FGA"
	FG_PCTSorter         = "FG_PCT"
	FG3MSorter           = "FG3M"
	FG3ASorter           = "FG3A"
	FG3_PCTSorter        = "FG3_PCT"
	FTMSorter            = "FTM"
	FTASorter            = "FTA"
	FT_PCTSorter         = "FT_PCT"
	OREBSorter           = "OREB"
	DREBSorter           = "DREB"
	ASTSorter            = "AST"
	STLSorter            = "STL"
	BLKSorter            = "BLK"
	TOVSorter            = "TOV"
	REBSorter            = "REB"
	PTSSorter            = "PTS"
	DateSorter           = "DATE"
	DefaultSorter        = DateSorter
)

type StarterBench string

var (
	Starters            StarterBench = "Starters"
	Bench                            = "Bench"
	DefaultStarterBench              = Starters
)

type Stat string

var (
	PointsStat           Stat = "PTS"
	ReboundsStat              = "REB"
	AssistsStat               = "AST"
	FieldGoalPercentStat      = "FG_PCT"
	FreeThrowPercentStat      = "FT_PCT"
	ThreesPercentStat         = "FG3_PCT"
	StealsStat                = "STL"
	BlocksStat                = "BLK"
	DefaultStat               = PointsStat
)

type StatCategory string

var (
	PointsStatCategory           StatCategory = "Points"
	ReboundsStatCategory                      = "Rebounds"
	AssistsStatCategory                       = "Assists"
	DefenseStatCategory                       = "Defense"
	ClutchStatCategory                        = "Clutch"
	PlaymakingStatCategory                    = "Playmaking"
	EfficiencyStatCategory                    = "Efficiency"
	FastBreakStatCategory                     = "Fast Break"
	ScoringBreakdownStatCategory              = "Scoring Breakdown"
	DefaultStatCategory                       = PointsStatCategory
)

type StatCategoryAbbreviation string

var (
	PTS                             StatCategoryAbbreviation = "PTS"
	FGM                                                      = "FGM"
	FGA                                                      = "FGA"
	FG_PCT                                                   = "FG_PCT"
	FG3M                                                     = "FG3M"
	FG3A                                                     = "FG3A"
	FG3_PCT                                                  = "FG3_PCT"
	FTM                                                      = "FTM"
	FTA                                                      = "FTA"
	OREB                                                     = "OREB"
	DREB                                                     = "DREB"
	AST                                                      = "AST"
	STL                                                      = "STL"
	BLK                                                      = "BLK"
	TOV                                                      = "TOV"
	REB                                                      = "REB"
	DefaultStatCategoryAbbreviation                          = PTS
)

type StatType string

var (
	Traditional     StatType = "Traditional"
	Advanced                 = "Advanced"
	Tracking                 = "Tracking"
	DefaultStatType          = Traditional
)

type TypeGrouping string

var (
	Offensive           TypeGrouping = "offensive"
	Defensive                        = "defensive"
	DefaultTypeGrouping              = Offensive
)
