# i2p-stats

This tool gathers statistics from a running I2P router and uses them to generate various ways of organizing and viewing those statistics.
At this time, the primary way of viewing them is as a static web site, so that it can easily be viewed on an eepsite or shared via github pages.
It also generates a json file for every stat, which can be fetched by referencing the time and date it was gathered.
It is intended to be run about every ten minutes, as a cron job, and for the output to be shared on an eepsite.

