#To insert the data collected for all years we wrote a script in python so that the data from the csv files will be inserted easily into MySQL database.

import mysql.connector as mysql
from mysql.connector import Error
import pandas as pd
empdata = pd.read_csv('/Users/bhargavbvs/Desktop/Data/esnew/2011.csv', index_col=False, delimiter = ',', keep_default_na=False)
empdata.head()

try:
    conn = mysql.connect(host='prepboost-prod.cppktldlc6tz.ap-south-1.rds.amazonaws.com', database='prepboost', user='root', password='prepboost#123')
    if conn.is_connected():
        cursor = conn.cursor()

        #loop through the data frame
        for i, row in empdata.iterrows():
            #here %S means string values
            print(row)
            sql = "INSERT INTO prepboost.pre_questions (id, exam_id, year, question, que_image_url, difficulty, topic_id, " \
                  "subtopic1_id, subtopic2_id, answer, option1, option2, option3, option4, explanation, learning, source) VALUES " \
                  "(%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s, %s)"
            if row['question_image_url'] == "":
                image = ''

            # val = (row['id'], row['exam_id'], row['year'], row['question'], image, row['difficulty'], row)
            cursor.execute(sql, tuple(row))
            print("Record inserted")
            # the connection is not auto committed by default, so we must commit to save our changes
            conn.commit()
            cursor.reset()
except Error as e:
            print("Error while connecting to MySQL", e)
