# NGINX Configuration Testing suite
This is a quick setup designed to test changes to nginx locally before making those changes elsewhere.

## Features
1. Sample NGINX config to test changes and features on
2. Simple Golang Webapp to accept requests
3. Locust to script out testing the nginx config

## Usage
1. Make a change to `nginx/my-site.conf`
2. Edit `locust/locustfile.py` to test against your change
3. `docker-compose build`
4. `docker-compose up`
5. Open http://localhost:8119 in your browser.
6. Start testing