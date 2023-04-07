package leetcode

import "testing"

func TestRank(t *testing.T) {
	expected := "ACB"
	actual := rankTeams([]string{"ABC", "ACB", "ABC", "ACB", "ACB"})

	if expected != actual {
		t.Errorf("expected: %s but got: %s", expected, actual)
	}

	expected = "XWYZ"
	actual = rankTeams([]string{"WXYZ", "XYZW"})

	if expected != actual {
		t.Errorf("expected: %s but got: %s", expected, actual)
	}

	expected = "ZMNAGUEDSJYLBOPHRQICWFXTVK"
	actual = rankTeams([]string{"ZMNAGUEDSJYLBOPHRQICWFXTVK"})

	if expected != actual {
		t.Errorf("expected: %s but got: %s", expected, actual)
	}

	expected = "VWFHSJARNPEMOXLTUKICZGYQ"
	actual = rankTeams([]string{"FVSHJIEMNGYPTQOURLWCZKAX", "AITFQORCEHPVJMXGKSLNZWUY", "OTERVXFZUMHNIYSCQAWGPKJL", "VMSERIJYLZNWCPQTOKFUHAXG", "VNHOZWKQCEFYPSGLAMXJIUTR", "ANPHQIJMXCWOSKTYGULFVERZ", "RFYUXJEWCKQOMGATHZVILNSP", "SCPYUMQJTVEXKRNLIOWGHAFZ", "VIKTSJCEYQGLOMPZWAHFXURN", "SVJICLXKHQZTFWNPYRGMEUAO", "JRCTHYKIGSXPOZLUQAVNEWFM", "NGMSWJITREHFZVQCUKXYAPOL", "WUXJOQKGNSYLHEZAFIPMRCVT", "PKYQIOLXFCRGHZNAMJVUTWES", "FERSGNMJVZXWAYLIKCPUQHTO", "HPLRIUQMTSGYJVAXWNOCZEKF", "JUVWPTEGCOFYSKXNRMHQALIZ", "MWPIAZCNSLEYRTHFKQXUOVGJ", "EZXLUNFVCMORSIWKTYHJAQPG", "HRQNLTKJFIEGMCSXAZPYOVUW", "LOHXVYGWRIJMCPSQENUAKTZF", "XKUTWPRGHOAQFLVYMJSNEIZC", "WTCRQMVKPHOSLGAXZUEFYNJI"})

	if expected != actual {
		t.Errorf("expected: %s but got: %s", expected, actual)
	}
}
