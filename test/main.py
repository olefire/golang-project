import requests

user_url = "http://localhost:8080/user"
user = {
    "name": "TestName",
    "email": "Test@Email.ru"
}

paste_url = "http://localhost:8080/paste"

post_user_request = requests.post(user_url, json=user).json()

print("post user request result: ", post_user_request)

get_user_url = user_url + '/' + post_user_request["id"]

get_user_request = requests.get(get_user_url).json()

print("get user request result: ", get_user_request)

paste = {
    "title": "testTitle",
    "paste": "testPasteText",
    "userID": post_user_request["id"]
}

post_paste_request = requests.post(paste_url, json=paste).json()

print("post paste request result: ", post_paste_request)

get_paste_url = paste_url + '/' + post_paste_request["id"]

get_paste_request = requests.get(get_paste_url).json()

print("get paste request result: ", get_paste_request)

paste_data = get_paste_request["data"]

delete_paste_url = paste_url + '/' + paste_data["id"]

delete_paste_request = requests.delete(delete_paste_url).json()

print("delete paste request result: ", delete_paste_request)

delete_user_url = user_url + '/' + paste_data["userID"]

delete_user_request = requests.delete(delete_user_url).json()

print("delete user request result: ", delete_user_request)
