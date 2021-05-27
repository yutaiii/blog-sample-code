import sys
import json
import pymysql

# rds settings
# お試しなので直書き
# secrets managerなどから値を取得する方がベター
# インスタンス直
# host = "<インスタンス直のエンドポイント>"
# proxy経由
host = "<proxyのエンドポイント>"
user = "<ユーザー名>"
passwd = "<パスワード>"
db_name = "<DB名>"

def lambda_handler(event, context):
    print("start!!")
    
    try:
        conn = pymysql.connect(host=host, user=user, passwd=passwd, db=db_name, connect_timeout=5)
    except Exception as e:
        print("error")
        print(e)
    
    with conn.cursor() as cur:
        cur.execute("SELECT * FROM fruits")
        for row in cur:
            print(row)
    
    # TODO implement
    return {
        'statusCode': 200,
        'body': json.dumps('Hello from Lambda!')
    }