import sys, os, json

if __name__ == "__main__":
    ## Deploy
    # data = {"UserName":"sz", "ProcessName": "fib"}
    # POSTDATA = "'" + json.dumps(data) + "'"
    # CMD = "curl -X POST -d " + POSTDATA + " http://localhost:8000/asyncFunc/deploy"
    # print(CMD)
    # result = os.system(CMD)

    ## Invoke
    data = {"UserName":"sz", "ProcessName": "fib", "WasmName": "fib.wasm", "Period": 50000, "Quota": 10000,"Ns": 1}
    POSTDATA = "'" + json.dumps(data) + "'"
    CMD = "curl -X POST -d " + POSTDATA + " http://localhost:8000/asyncFunc/invoke"
    # print(CMD)
    result = os.system(CMD)