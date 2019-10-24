# DailyNewsText
This application sends a daily text message of the top headlines in French. 

The app.yaml file is to deploy this appliction on Google Cloud's app engine. 

You need to run Go get github.com/jasonlvhit/gocron from this directory. Then you can do gcloud app deploy.

Replace:
apiKey (for newsapi.org)
accountSid (for Twilio)
authToken
To & From phone numbers

Thanks to, https://newsapi.org/ for the news sources. 

[Sending texts via Twilio with GO](https://www.twilio.com/blog/2017/09/send-text-messages-golang.html)

This app would run better in standard app engine because you could have manual scaling and the app wouldn't be running except when the chron job occured. Could also be integrated with [cloud scheduler.](https://cloud.google.com/scheduler/)
