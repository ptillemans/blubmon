# Simple WIFI monitor

This app serves to monitor the WiFi connection on my BeagleBoneBlack which has the nasty tendency to loose connection frequently.

As it is intended to run as a cronjob, all output goes to syslog.

to see the output run

    $ grep blubmon /var/log/syslog
