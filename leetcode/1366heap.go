package leetcode

import (
	"container/heap"
	"strings"
)

type Team struct {
	letter string
	score  []int
	index  int
}

type TeamRank []*Team

const maxNumberOfTeams = 26

func (teamRank TeamRank) Len() int {
	return len(teamRank)
}

func (teamRank TeamRank) Less(i, j int) bool {
	for n := 0; n < len(teamRank[i].score); n++ {
		if teamRank[i].score[n] > teamRank[j].score[n] {
			return true
		}

		if teamRank[i].score[n] < teamRank[j].score[n] {
			return false
		}
	}

	return teamRank[i].letter < teamRank[j].letter
}

func (teamRank TeamRank) Swap(i, j int) {
	teamRank[i], teamRank[j] = teamRank[j], teamRank[i]
	teamRank[i].index = i
	teamRank[j].index = j
}

func (teamRank *TeamRank) Push(x any) {
	n := len(*teamRank)
	team := x.(*Team)
	team.index = n
	*teamRank = append(*teamRank, team)
}

func (teamRank *TeamRank) Pop() any {
	old := *teamRank
	n := len(old)
	team := old[n-1]
	old[n-1] = nil
	team.index = -1
	*teamRank = old[0 : n-1]
	return team
}

func (teamRank TeamRank) Result() string {
	results := []string{}

	for teamRank.Len() > 0 {
		topTeam := heap.Pop(&teamRank).(*Team)
		results = append(results, topTeam.letter)
	}

	return strings.Join(results, "")
}

func rankTeams(votes []string) string {
	if len(votes) == 1 {
		return votes[0]
	}

	results := make(map[string][]int)

	for _, vote := range votes {
		for j, letter := range vote {
			teamLetter := string(letter)

			if _, ok := results[teamLetter]; !ok {
				results[teamLetter] = make([]int, maxNumberOfTeams)
			}

			results[teamLetter][j]++
		}
	}

	teamRank := make(TeamRank, len(results))
	i := 0

	for letter, score := range results {
		teamRank[i] = &Team{
			letter: letter,
			score:  score,
			index:  i,
		}
		i++
	}

	heap.Init(&teamRank)

	return teamRank.Result()
}
