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
