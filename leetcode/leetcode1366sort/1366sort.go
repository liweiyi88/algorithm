package leetcode1366sort

import (
	"sort"
)

type Team struct {
	letter rune
	score  []int
}

type TeamRank []*Team

const maxNumberOfTeams = 26

func (teamRank TeamRank) Len() int {
	return len(teamRank)
}

func (teamRank TeamRank) Less(i, j int) bool {
	for n := 0; n < len(teamRank[i].score); n++ {
		if teamRank[i].score[n] == teamRank[j].score[n] {
			continue
		}

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
}

func (teamRank TeamRank) Result() string {
	results := []rune{}
	sort.Sort(teamRank)

	for i := 0; i < teamRank.Len(); i++ {
		results = append(results, teamRank[i].letter)
	}

	return string(results)
}

func rankTeams(votes []string) string {
	if len(votes) == 1 {
		return votes[0]
	}

	results := make(map[rune][]int)

	for _, vote := range votes {
		for j, letter := range vote {
			teamLetter := letter

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
		}
		i++
	}

	return teamRank.Result()
}
