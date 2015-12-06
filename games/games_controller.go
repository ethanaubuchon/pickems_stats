package 'game'

var games map[string]games.Game

func AddGame(game Game) {
  if games == nil {
    games = make(map[string]games.Game)
  }
  games[game.Date + game.Time + game.HomeTeam.short + game.AwayTime.short] = game
}

func GetGame(date string, time string, homeShort string, awayShort string) (Game, boolean) {
  if games == nil {
    return nil, false
  }
  return games[string]games.Game
}

