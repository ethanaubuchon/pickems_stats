package games

import (
  "strings"
  "regexp"
  "strconv"
  "github.com/PuerkitoBio/goquery"
)

const SCRAPE_PATH string = "http://www.nhl.com/ice/schedulebyseason.htm"

// Data Selectors
const TABLE_ROW_SELECTOR string = "#fullPage .contentBlock table.schedTbl tbody tr"
const TABLE_COL_SELECTOR string = "td"
const DATE_SELECTOR string = ".skedStartDateSite"
const TEAM_SELECTOR string = ".teamName a"
const TEAM_SHORT_ATTR string = "rel"
const TIME_SELECTOR string = ".skedStartTimeEST"
const SCORE_REGEX = `FINAL:\ *[A-Z][A-Z][A-Z]\ *\(([1-9]*[0-9]+)\)\ *-\ *[A-Z][A-Z][A-Z]\ *\(([1-9]*[0-9]+)\)`
const NUMBER_REGEX = `[1-9]*[0-9]+`


var games map[string]Game

func AddGame(game Game) {
  if games == nil {
    games = make(map[string]Game)
  }
  games[game.Date + game.Time + game.HomeTeam.Short + game.AwayTeam.Short] = game
}

// func GetGame(date string, time string, homeShort string, awayShort string) (Game, boolean) {
//   if games == nil {
//     return nil, false
//   }
//   return games[string]Game
// }

func GetGames() (map[string]Game) {
  return games
}

func ScrapeGameData() (err error) {
  doc, err := goquery.NewDocument(SCRAPE_PATH)
  if err != nil {
    return
  }

  doc.Find(TABLE_ROW_SELECTOR).Each(func(i int, s *goquery.Selection) {
    game := Game{}
    game.HomeTeam = Team{}
    game.AwayTeam = Team{}

    s.Find(TABLE_COL_SELECTOR).Each(func(j int, row *goquery.Selection) {
      switch j {
      case 0:
        game.Date = row.Find(DATE_SELECTOR).Text()
      case 1:
        game.AwayTeam.Short, _ = row.Find(TEAM_SELECTOR).First().Attr(TEAM_SHORT_ATTR)
      case 2:
        game.HomeTeam.Short, _ = row.Find(TEAM_SELECTOR).First().Attr(TEAM_SHORT_ATTR)
      case 3:
        game.Time = row.Find(TIME_SELECTOR).Text()
      case 4:
        res := strings.Replace(row.Text(), "\n", "", -1)

        if regexp.MustCompile(SCORE_REGEX).MatchString(res) == true {
          scores := regexp.MustCompile(NUMBER_REGEX).FindAllString(res, 2)
          game.AwayTeam.Score, _ = strconv.Atoi(scores[0])
          game.HomeTeam.Score, _ = strconv.Atoi(scores[1])
          AddGame(game)
        }
      }
    })
  })
  return
}