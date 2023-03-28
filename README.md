**n**ba **a**pi **g**o (nag) is an unofficial NBA Stats API in Go. it is very much a Go port of [nba_api](https://github.com/swar/nba_api).

### status

as this project has been on hiatus for several years, the status of the previously implemented endpoints are detailed below. I will slowly pick up and update the codebase, along with perhaps some refactoring. who knows what the hell I was thinking when making code structures decisions a few years ago.

### endpoints

| endpoint | implemented | up_to_date | test_ok |
|---|---|---|---|
| alltimeleadersgrids | x | ? | x |
| assistleaders | x | ? | x |
| assisttracker |   |   |   |
| boxscoreadvancedv2 | x | ? | x |
| boxscoredefensive | x |   |   |
| boxscorefourfactorsv2 | x | ? | x |
| boxscorematchups |   |   |   |
| boxscoremiscv2 | x | ? | x |
| boxscoreplayertrackv2 | x | ? | x |
| boxscorescoringv2 | x | ? | x |
| boxscoresimilarityscore |   |   |   |
| boxscoresummaryv2 | x | ? | x |
| boxscoretraditionalv2 | x | ? | x |
| boxscoreusagev2 | x | ? | x |
| commonallplayers | x | ? | x |
| commonplayerinfo | x | ? | x |
| commonplayoffseries | x | ? | x |
| commonteamroster | x | ? | x |
| commonteamyears |   |   |   |
| cumestatsplayergames |   |   |   |
| cumestatsplayer |   |   |   |
| cumestatsteamgames |   |   |   |
| cumestatsteam |   |   |   |
| defensehub |   |   |   |
| draftboard |   |   |   |
| draftcombinedrillresults |   |   |   |
| draftcombinenonstationaryshooting |   |   |   |
| draftcombineplayeranthro |   |   |   |
| draftcombinespotshooting |   |   |   |
| draftcombinestats |   |   |   |
| drafthistory |   |   |   |
| fantasywidget |   |   |   |
| franchisehistory |   |   |   |
| franchiseleaders |   |   |   |
| franchiseplayers |   |   |   |
| gamerotation | x | ? | x |
| glalumboxscoresimilarityscore |   |   |   |
| homepageleaders |   |   |   |
| homepagev2 |   |   |   |
| hustlestatsboxscore |   |   |   |
| infographicfanduelplayer |   |   |   |
| leaderstiles |   |   |   |
| leaguedashlineups |   |   |   |
| leaguedashoppptshot |   |   |   |
| leaguedashplayerbiostats |   |   |   |
| leaguedashplayerclutch |   |   |   |
| leaguedashplayerptshot |   |   |   |
| leaguedashplayershotlocations |   |   |   |
| leaguedashplayerstats |   |   |   |
| leaguedashptdefend |   |   |   |
| leaguedashptstats |   |   |   |
| leaguedashptteamdefend |   |   |   |
| leaguedashteamclutch |   |   |   |
| leaguedashteamptshot |   |   |   |
| leaguedashteamshotlocations |   |   |   |
| leaguedashteamstats |   |   |   |
| leaguegamefinder |   |   |   |
| leaguegamelog |   |   |   |
| leaguehustlestatsplayerleaders |   |   |   |
| leaguehustlestatsplayer |   |   |   |
| leaguehustlestatsteamleaders |   |   |   |
| leaguehustlestatsteam |   |   |   |
| leagueleaders |   |   |   |
| leaguelineupviz |   |   |   |
| leagueplayerondetails |   |   |   |
| leagueseasonmatchups |   |   |   |
| leaguestandings |   |   |   |
| leaguestandingsv3 |   |   |   |
| matchupsrollup |   |   |   |
| playbyplay |   |   |   |
| playbyplayv2 | x | ? | x |
| playerawards |   |   |   |
| playercareerbycollege |   |   |   |
| playercareerbycollegerollup |   |   |   |
| playercareerstats |   |   |   |
| playercompare |   |   |   |
| playerdashboardbyclutch |   |   |   |
| playerdashboardbygamesplits |   |   |   |
| playerdashboardbygeneralsplits |   |   |   |
| playerdashboardbylastngames |   |   |   |
| playerdashboardbyopponent |   |   |   |
| playerdashboardbyshootingsplits |   |   |   |
| playerdashboardbyteamperformance |   |   |   |
| playerdashboardbyyearoveryear |   |   |   |
| playerdashptpass |   |   |   |
| playerdashptreb |   |   |   |
| playerdashptshotdefend |   |   |   |
| playerdashptshots |   |   |   |
| playerestimatedmetrics |   |   |   |
| playerfantasyprofilebargraph |   |   |   |
| playerfantasyprofile |   |   |   |
| playergamelog |   |   |   |
| playergamelogs |   |   |   |
| playergamestreakfinder |   |   |   |
| playernextngames |   |   |   |
| playerprofilev2 |   |   |   |
| playervsplayer |   |   |   |
| playoffpicture |   |   |   |
| scoreboard |   |   |   |
| scoreboardv2 | x | ? | x |
| shotchartdetail |   |   |   |
| shotchartleaguewide |   |   |   |
| shotchartlineupdetail |   |   |   |
| synergyplaytypes |   |   |   |
| teamandplayersvsplayers |   |   |   |
| teamdashboardbyclutch |   |   |   |
| teamdashboardbygamesplits |   |   |   |
| teamdashboardbygeneralsplits |   |   |   |
| teamdashboardbylastngames |   |   |   |
| teamdashboardbyopponent |   |   |   |
| teamdashboardbyshootingsplits |   |   |   |
| teamdashboardbyteamperformance |   |   |   |
| teamdashboardbyyearoveryear |   |   |   |
| teamdashlineups |   |   |   |
| teamdashptpass |   |   |   |
| teamdashptreb |   |   |   |
| teamdashptshots |   |   |   |
| teamdetails |   |   |   |
| teamestimatedmetrics |   |   |   |
| teamgamelog | x | ? | x |
| teamgamelogs | x | ? | x |
| teamgamestreakfinder |   |   |   |
| teamhistoricalleaders |   |   |   |
| teaminfocommon |   |   |   |
| teamplayerdashboard |   |   |   |
| teamplayeronoffdetails |   |   |   |
| teamplayeronoffsummary |   |   |   |
| teamvsplayer |   |   |   |
| teamyearbyyearstats |   |   |   |
| videodetails |   |   |   |
| videoevents |   |   |   |
| videostatus |   |   |   |
| winprobabilitypbp |   |   |   |

### LICENSE

while I usually use [WTFPL](http://www.wtfpl.net/faq/), I guess I should use [MIT](LICENSE) here?
