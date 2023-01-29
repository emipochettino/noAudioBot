# [noAudioBot](https://t.me/AudioMurdererBot)

I personally don't like to hear audios so, I just developed this bot to get the transcription of those messages.

for the moment the bot is opened only to me but, the idea is to handle multiple api key
from [assemblyai.com](https://www.assemblyai.com/) so in that way each user can have their own api key and handle their
rate limit

It is developed following some patterns from hexagonal architecture.

### environment variables

```
export ALLOWED_TELEGRAM_USERS={userId1,userId2}
export ASSEMBLY_AI_API_KEY={apiKey}
export TELEGRAM_TOKEN={token}
```
