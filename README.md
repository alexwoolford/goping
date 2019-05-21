# goping

My eldest son, who's 13, just told me that the ping speed is slow from his PS4.

I wrote this little Go script to ping google.com every minute and publish the statistics to a Kafka topic. I'll write this data to InfluxDB to get a feel for what's normal.

TODO:
1. parameterize the properties
2. figure out how to deploy as a service
3. add logging