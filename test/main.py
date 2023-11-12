import requests
from bson import ObjectId

user_url = "http://service:8080/user"

paste_url = "http://service:8080/paste"


def create_user(path: str):
    user = {
        "name": "TestName",
        "email": "Test@Email.ru"
    }

    post_user_request = requests.post(path, json=user)
    response = post_user_request.json()
    uid = response["data"]

    assert post_user_request.status_code == 200
    assert response["msg"] == "success"
    assert ObjectId.is_valid(uid)
    return uid


def get_user(path: str, uid):
    get_user_url = path + '/' + uid
    get_user_request = requests.get(get_user_url)
    response = get_user_request.json()

    assert get_user_request.status_code == 200
    assert response["msg"] == "success"
    assert ObjectId.is_valid(response["data"]["id"])


def create_paste(path: str, uid):
    test_paste = {
        "title": "testTitle",
        "paste": "testPasteText",
        "userID": uid
    }

    post_paste_request = requests.post(path, json=test_paste)
    response = post_paste_request.json()
    pid = response["data"]

    assert post_paste_request.status_code == 200
    assert response["msg"] == "success"
    assert ObjectId.is_valid(pid)

    return pid


def get_paste(path: str, pid):
    get_paste_url = path + '/' + pid
    get_paste_request = requests.get(get_paste_url)
    response = get_paste_request.json()

    assert get_paste_request.status_code == 200
    assert response["msg"] == "success"
    assert ObjectId.is_valid(response["data"]["id"])


def delete_paste(path: str, pid):
    delete_paste_url = path + '/' + pid
    delete_paste_request = requests.delete(delete_paste_url)
    response = delete_paste_request.json()

    assert delete_paste_request.status_code == 200
    assert response["msg"] == "success"
    assert ObjectId.is_valid(response["data"])


def delete_user(path: str, uid):
    delete_user_url = path + '/' + uid
    delete_user_request = requests.delete(delete_user_url)
    response = delete_user_request.json()
    assert delete_user_request.status_code == 200
    assert response["msg"] == "success"
    assert ObjectId.is_valid(response["data"])


def run_test():
    user_id = create_user(user_url)
    get_user(user_url, user_id)
    paste_id = create_paste(paste_url, user_id)
    get_paste(paste_url, paste_id)
    delete_paste(paste_url, paste_id)
    delete_user(user_url, user_id)


run_test()
