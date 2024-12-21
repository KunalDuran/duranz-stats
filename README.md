"# duranz-stats" 
 
 Application written in Go programming Language for storing/extracting stats to perform analysis from the ball by ball data available on cricsheet.
 
  - Dataset: Cricsheet [https://cricsheet.org/]
> Currently works with odi, t20, test, ipl data


Stats Provided:

1. Player Stats
    - Batting Stats
        Runs Scored
        Balls Played
        Average
        Strike Rate
        Dot balls
        Number of 6s, 4s, 3s, 2s, 1
        No. of times dismissal (by type, e.g, bowled 10 times)
        Not out
        Ducks
        Highest Score
    - Bowling Stats
        Runs Conceded
        Balls Bowled
        Dots
        Average
        Economy
        6s, 4s, 3s, 2s, 1 Conceded
        Extras conceded
        Wickets taken 
        Best bowling figure
        Maiden
    - Fielding Stats
        Catches
        Run Outs
        Stumpings

    Filters: year, format, vsteam


2. Team Stats
    - Total Matches
    - Win
    - Bat first win
    - Chase win
    - Average score
    - Highest score
    - Lowest score
    - Toss win

    Filters: year, format, vsteam, venue

3. Match Stats (TBD)
    - Scoresheet, Fall of Wickets, Partnerships
    - Batsman vs Bowler
    - run rate graph
    - over by over analysis