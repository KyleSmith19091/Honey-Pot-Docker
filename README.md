# Honeypot + Docker + Golang

## How to run
### Dependencies
1. Docker
2. Python 3
    - Selenium `pip install seleniumwire`
    - Chrome web driver `brew install chromewebdriver` or see <a href="https://jonathansoma.com/lede/foundations-2018/classes/selenium/selenium-windows-install/">Windows Install</a>
3. Go

### How to Build and Run
#### Web server and Dbs
1. `docker compose up --build`

2. Access on localhost:8080

#### Bot
1. `python bot.py`
