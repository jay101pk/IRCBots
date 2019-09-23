package bot_test

import (
	"reflect"
	"testing"

	"github.com/AFTERWAKE/IRCBots/dad/bot"
)

func TestCommandMatch(t *testing.T) {
	command := &bot.Command{
		[]*bot.Regex{
			&bot.Regex{"^testpattern$"},
			&bot.Regex{"[^a]([b-z]+)"},
		},
		[]*bot.Response{nil},
		[]string{},
		nil,
	}
	tables := []struct {
		param  string
		expect []string
	}{
		{"test pattern", []string{"test", "est"}},
		{"testpattern", []string{"testpattern"}},
		{"blahpattern", []string{"bl", "l"}},
		{"ttttttpattern", []string{"ttttttp", "tttttp"}},
	}
	for i, table := range tables {
		match := command.Match(table.param)
		if !reflect.DeepEqual(match, table.expect) {
			t.Errorf(
				"Failed [%d]: expected match: %v, got %v",
				i, table.expect, match,
			)
		}

	}
}
