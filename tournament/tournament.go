package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type team struct {
	name                          string
	game, win, loss, draw, points int
}

type scoreBoard map[string]*team

// Tally takes a string and buffer,
// does its calculation and returns some stuff
func Tally(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	board := make(scoreBoard)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if err := board.addGame(line); err != nil {
			return err
		}
	}
	teams := board.getTeams()
	sort.SliceStable(teams, func(i, j int) bool {
		if teams[i].points != teams[j].points {
			return teams[i].points > teams[j].points
		} else if teams[i].win != teams[j].win {
			return teams[i].win > teams[j].win
		}
		return teams[i].name < teams[j].name
	})

	header := "Team                           | MP |  W |  D |  L |  P\n"
	io.WriteString(w, header)
	for _, team := range teams {
		io.WriteString(w, team.String())
	}
	return nil
}

func (b scoreBoard) getTeams() []team {
	teams := make([]team, len(b))
	i := 0
	for _, team := range b {
		teams[i] = *team
		i++
	}
	return teams
}

func (b scoreBoard) addGame(game string) error {
	fields := strings.Split(game, ";")
	if len(fields) != 3 {
		return fmt.Errorf("Game not properly formatted: %s", game)
	}
	home, homeOk := b[fields[0]]
	away, awayOk := b[fields[1]]
	if !homeOk {
		home = &team{name: fields[0]}
		b[fields[0]] = home
	}
	if !awayOk {
		away = &team{name: fields[1]}
		b[fields[1]] = away
	}
	switch fields[2] {
	case "win":
		home.addWin()
		away.addLoss()
	case "loss":
		home.addLoss()
		away.addWin()
	case "draw":
		home.addDraw()
		away.addDraw()
	default:
		return fmt.Errorf("Unknown game condition: %s", fields[2])
	}
	return nil
}

func (t *team) addWin() {
	t.game++
	t.win++
	t.points += 3
}

func (t *team) addLoss() {
	t.game++
	t.loss++
}

func (t *team) addDraw() {
	t.game++
	t.draw++
	t.points++
}

func (t team) String() string {
	return fmt.Sprintf("%-30s | %2d | %2d | %2d | %2d | %2d\n", t.name, t.game, t.win, t.draw, t.loss, t.points)
}
