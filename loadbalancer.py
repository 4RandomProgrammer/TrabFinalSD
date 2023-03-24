from flask import Flask, jsonify, request
import requests

app = Flask(__name__)
global balance
balance = 0

@app.route('/hello/', methods=['GET'])
def welcome():
    global balance
    value = None
    port = '8090'

    if balance == 0:
        port = '8090'
    else:
        port = '8090'

    value = requests.get('http://localhost:'+ port + '/getData/')
    balance = 1 - balance

    return value.json()

@app.route('/teste/', methods=['POST'])
def teste():
    x = request.get_json()
    
    inverse = {
        "x":x["y"],
        "y":x["x"]
    }

    responseA = requests.post('http://localhost:8090/insert/',json=x)
    # responseB = requests.post('http://localhost:8070/insert/',data=inverse)

    return responseA.json()

if __name__ == '__main__':
    balance = 0
    app.run(host='localhost', port=8080)