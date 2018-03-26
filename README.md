# Usage

<a href="https://slack.com/oauth/authorize?client_id=26143804727.320430302100&scope=incoming-webhook,commands"><img alt="Add to Slack" height="40" width="139" src="https://platform.slack-edge.com/img/add_to_slack.png" srcset="https://platform.slack-edge.com/img/add_to_slack.png 1x, https://platform.slack-edge.com/img/add_to_slack@2x.png 2x" /></a>

# Development

1. Fire up a PostgreSQL DB

2. Set the following environment variables
  - `SLACK_VERIFICATION_TOKEN`
  - `SQL_CONNECTION_STRING`
  - `SENTRY_DSN`

3. Create a Slack App

4. Configure Slash commands (with something like ngrok)
  - Add match command should point to `/slack/match/add`
  - Undo match command should point to `/slack/match/undo`
  - Get leaderboard command should point to `/slack/league/leadeboard`
