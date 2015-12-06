package main

import (
  "fmt"
  "github.com/ethanaubuchon/pickems_stats/games"
)

func main() {
  games.ScrapeGameData()

  g := games.GetGames()

  for _, game := range g {
    fmt.Printf("%s %s %s (%d) VS %s (%d)\n", game.Date, game.Time, game.HomeTeam.Short, game.HomeTeam.Score, game.AwayTeam.Short, game.AwayTeam.Score)
  }
}