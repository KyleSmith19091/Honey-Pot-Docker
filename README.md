# Honeypot + Docker + Golang
A go web server that serves example html files and hosts an api that talks with a postgresql database. 
<div align="center">
<img src="https://user-images.githubusercontent.com/29174023/201514724-5c669ad7-4c1f-4c50-b9c4-d59cb9a88af0.png" width="300px" height="300px" />
<p>Architecture</p>
</div>

If malicious activity is detected then we reroute the request to talk with our "fake database". We detect malicious activity by having a form with 4 input fields, where two of those fields are visible to the user, the other two fields are named in such a way that a bot looking for input fields marked as email and password would find theme easily. 

![Screenshot 2022-10-13 at 19 57 01](https://user-images.githubusercontent.com/29174023/201514906-b02e4f0a-e89e-4071-8ff1-e08e84d55454.png)

So when the form is submitted it will have the email and password fields set in the request. For users interacting with the site using the site normally they will interact with these input fields, which have obfuscated names.

![Screenshot 2022-10-13 at 19 56 54](https://user-images.githubusercontent.com/29174023/201514882-1159c39a-c54c-4b6a-b0f5-a4162a822d58.png)



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

#### Bot - Used to simulate an attack
1. `python bot.py`
