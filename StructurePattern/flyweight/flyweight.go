package flyweight

import "time"

// 定義constant常數
const (
	// iota 為定義一個不需要值的常數
	// 會隨著contant常數增加，例如TEAM_A 為0 TEAM_A則是1
	TEAM_A = iota
	TEAM_B
)

type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}

type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}

type Team struct {
	ID             uint64
	Name           string
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

// 儲存兩隊比賽資訊
type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}

func getTeamFactory(team int) Team {
	switch team {
	case TEAM_B:
		return Team{
			ID:   2,
			Name: "TEAM_B",
		}
	default:
		return Team{
			ID:   1,
			Name: "TEAM_A",
		}
	}
}

func NewTeamFactory() teamFlyweightFactory {
	// 透過回傳 teamFlyweightFacotry type 並做teamFlyweightFactory初始化的動作
	// 注意：這裡若沒有 0 做初始化 會產生compile error

	return teamFlyweightFactory{
		createdTeams: make(map[int]*Team, 0),
	}
}

// Factory回儲存指向Team當作值的指標
// 以及如果知道 team的名稱 直接透過map 加快搜尋速度
type teamFlyweightFactory struct {
	createdTeams map[int]*Team
}

func (t *teamFlyweightFactory) GetTeam(teamName int) *Team {
	// 如果team有存在的話
	if t.createdTeams[teamName] != nil {
		return t.createdTeams[teamName]
	}
	// team的資訊不在那就從factory要 team的指標
	team := getTeamFactory(teamName)
	t.createdTeams[teamName] = &team

	return t.createdTeams[teamName]
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}
