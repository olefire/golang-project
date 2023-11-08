import requests

user_url = "http://service:8080/user"

paste_url = "http://service:8080/paste"


def create_user(path: str):
    user = {
        "name": "TestName",
        "email": "Test@Email.ru"
    }
    post_user_request = requests.post(path, json=user).json()
    user_id = post_user_request["id"]
    message = f"user ${user_id} successfully created"
    print(message)
    return user_id


def get_user(path: str, user_id):
    get_user_url = path + '/' + user_id
    get_user_request = requests.get(get_user_url).json()
    return get_user_request


def create_paste(path: str, user_id):
    test_paste = {
        "title": "testTitle",
        "paste": "testPasteText",
        "userID": user_id
    }

    post_paste_request = requests.post(paste_url, json=test_paste).json()
    paste_id = post_paste_request["id"]

    message = f"paste ${paste_id} successfully created"
    print(message)
    return paste_id


def get_paste(path: str, paste_id):
    get_paste_url = paste_url + '/' + paste_id

    get_paste_request = requests.get(get_paste_url).json()

    return get_paste_request


def delete_paste(path: str, paste_id):
    delete_paste_url = paste_url + '/' + paste_id
    requests.delete(delete_paste_url)

    message = f"paste ${paste_id} successfully deleted"
    return message


def delete_user(path: str, user_id):
    delete_user_url = user_url + '/' + user_id
    requests.delete(delete_user_url)

    message = f"user ${user_id} successfully deleted"
    return message


user_id = create_user(user_url)
user = get_user(user_url, user_id)

paste_id = create_paste(paste_url, user_id)
paste = get_paste(paste_url, paste_id)

delete_paste_res = delete_paste(paste_url, paste_id)
delete_user_res = delete_user(user_url, user_id)

print(user)
print(paste)

print(delete_user_res)
print(delete_paste_res)
