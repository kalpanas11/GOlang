
from flask import Flask, render_template, request, redirect, url_for
import requests

app = Flask(__name__)

# Go service URL
GO_SERVICE_URL = "http://localhost:8080"

@app.route("/")
def index():
    # Get the list of users from the Go service
    response = requests.get(f"{GO_SERVICE_URL}/users")
    users = response.json()

    return render_template("index.html", users=users)


@app.route("/create", methods=["GET", "POST"])
def create_user():
    if request.method == "POST":
        username = request.form["username"]
        password = request.form["password"]

        # Create a new user in the Go service
        data = {"username": username, "password": password}
        response = requests.post(f"{GO_SERVICE_URL}/users", json=data)

        if response.status_code == 201:
            return redirect(url_for("index"))
        elif response.status_code == 409:
            error = response.json().get("error")
            return render_template("create.html", error=error, users=get_users())
        else:
            error = response.json().get("error")
            return render_template("create.html", error=error, users=get_users())

    return render_template("create.html", users=get_users())

def get_users():
    # Get the list of users from the Go service
    response = requests.get(f"{GO_SERVICE_URL}/users")
    return response.json()




# @app.route("/create", methods=["GET", "POST"])
# def create_user():
#     if request.method == "POST":
#         username = request.form["username"]
#         password = request.form["password"]

#         # Create a new user in the Go service
#         data = {"username": username, "password": password}
#         response = requests.post(f"{GO_SERVICE_URL}/users", json=data)

#         if response.status_code == 201:
#             return redirect(url_for("index"))
#         else:
#             error = response.json().get("error")
#             return render_template("create.html", error=error)

#     # Get the list of users from the Go service
#     response = requests.get(f"{GO_SERVICE_URL}/users")
#     users = response.json()

#     return render_template("create.html", users=users)


if __name__ == "__main__":
    app.run(debug=True, port=5000)























# from flask import Flask, render_template, request, redirect, url_for
# import requests

# app = Flask(__name__)

# API_URL = "http://go-service:8080"

# @app.route('/')
# def index():
#     users = requests.get(f"{API_URL}/users").json
#     return render_template('index.html', users=users)

# @app.route('/create', methods=['GET', 'POST'])
# def create():
#     if request.method == 'POST':
#         name = request.form['name']
#         email = request.form['email']
#         data = {'name': name, 'email': email}
#         requests.post(f"{API_URL}/users", json=data)
#         return redirect(url_for('index'))
#     return render_template('create.html')

# # Other CRUD routes (update, delete) omitted for brevity

# if __name__ == '__main__':
#     app.run(host='0.0.0.0', port=5000, debug=True)