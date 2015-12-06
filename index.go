package main

import (
  "fmt"
  "log"
  "strings"
  "regexp"
  "strconv"
  "github.com/PuerkitoBio/goquery"
  "github.com/ethanaubuchon/pickems_stats/games"
)

func ExampleScrape() (map[string]games.Game) {
  g := make(map[string]games.Game)
  doc, err := goquery.NewDocument("http://www.nhl.com/ice/schedulebyseason.htm")
  if err != nil {
    log.Fatal(err)
  }

  doc.Find("#fullPage .contentBlock table.schedTbl tbody tr").Each(func(i int, s *goquery.Selection) {
    game := games.Game{}
    game.HomeTeam = games.Team{}
    game.AwayTeam = games.Team{}

    s.Find("td").Each(func(j int, row *goquery.Selection) {
      switch j {
      case 0:
        game.Date = row.Find(".skedStartDateSite").Text()
      case 1:
        game.AwayTeam.Short, _ = row.Find(".teamName a").First().Attr("rel")
      case 2:
        game.HomeTeam.Short, _ = row.Find(".teamName a").First().Attr("rel")
      case 3:
        game.Time = row.Find(".skedStartTimeEST").Text()
      case 4:
        r := `FINAL:\ *[A-Z][A-Z][A-Z]\ *\(([1-9]*[0-9]+)\)\ *-\ *[A-Z][A-Z][A-Z]\ *\(([1-9]*[0-9]+)\)`

        res := strings.Replace(row.Text(), "\n", "", -1)

        if regexp.MustCompile(r).MatchString(res) == true {
          scores := regexp.MustCompile(`[1-9]*[0-9]+`).FindAllString(res, 2)
          game.AwayTeam.Score, _ = strconv.Atoi(scores[0])
          game.HomeTeam.Score, _ = strconv.Atoi(scores[1])
          g[game.Date + game.Time + game.AwayTeam.Short + game.HomeTeam.Short] = game
        }
      }
    })
  })
  return g
}

func main() {
  games := ExampleScrape()

  for _, game := range games {
    fmt.Printf("%s %s %s (%d) VS %s (%d)\n", game.Date, game.Time, game.HomeTeam.Short, game.HomeTeam.Score, game.AwayTeam.Short, game.AwayTeam.Score)
  }
}