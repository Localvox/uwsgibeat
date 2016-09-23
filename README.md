[![Build Status](https://travis-ci.org/mrkschan/uwsgibeat.svg?branch=master)](https://travis-ci.org/mrkschan/uwsgibeat)

# uwsgibeat

uwsgibeat is the Beat used for uWSGI monitoring. It is a lightweight agent that reads stats from uWSGI server periodically. uWSGI server must expose its stats via stats server  (http://uwsgi-docs.readthedocs.org/en/latest/StatsServer.html).

## Command Line Options

 - `-path.config` - specifies the location of the configuration file `uwsgibeat.yml` (see etc for a sample version)
