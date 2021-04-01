# This is a sample Python script.
import requests
import csv
import sys
from datetime import date, timedelta

# Press ⌃R to execute it or replace it with your code.
# Press Double ⇧ to search everywhere for classes, files, tool windows, actions, and settings.

# Press the green button in the gutter to run the script.

insert_query = """
INSERT INTO "public"."historical_eod_data" ("trade_code", "date", "open", "high", "low", "close", "volume") VALUES
"""

value_query = """ ('{trade_code}', '{date}', {open}, {high}, {low}, {close}, {volume})"""


def query_builder(stock):
    trade_code = str(stock[0])
    date = str(stock[1])
    open = float(stock[2])
    high = float(stock[3])
    low = float(stock[4])
    close = float(stock[5])
    volume = stock[6]
    formatted_query = value_query.format(trade_code=trade_code, date=date, open=open, high=high, low=low, close=close,
                                         volume=volume)
    return formatted_query


def data_downloader(date):
    headers = {
        'authority': 'stockbangladesh.com',
        'cache-control': 'max-age=0',
        'sec-ch-ua': '"Google Chrome";v="89", "Chromium";v="89", ";Not A Brand";v="99"',
        'sec-ch-ua-mobile': '?0',
        'upgrade-insecure-requests': '1',
        'origin': 'https://stockbangladesh.com',
        'content-type': 'application/x-www-form-urlencoded',
        'user-agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36',
        'accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9',
        'sec-fetch-site': 'same-origin',
        'sec-fetch-mode': 'navigate',
        'sec-fetch-user': '?1',
        'sec-fetch-dest': 'document',
        'referer': 'https://stockbangladesh.com/download',
        'accept-language': 'en-GB,en-US;q=0.9,en;q=0.8',
        'cookie': '_ga=GA1.2.309986034.1616223702; _gid=GA1.2.776334282.1617089220; sc_is_visitor_unique=rx3183403.1617089253.061065D93CA74F7366BAF04D16699E09.4.4.4.3.3.3.3.3.2; XSRF-TOKEN=eyJpdiI6IlI0ajdDVGdMZEdFR0tSMGZHdEFPcFE9PSIsInZhbHVlIjoiQzBBSDRBd0dmZEZka3VodktZcEJ2NWVWSHJ2dEwwV1hEYXNIK1IwTk1qUFpcL3hlWHF0WTBqend5UUd1dEVDZEUiLCJtYWMiOiIxOWQ4ZjY0NzZlNDhmYWU4NWFmOWQ5ZmViYTJlMzAyMTUxYWYwODQzMTNiYmRhMjRjY2YzYjM2MDFmNDMzMzIyIn0%3D; laravel_session=eyJpdiI6IjhRXC9EdlJSNkR2UFArNGNYTzBCT1d3PT0iLCJ2YWx1ZSI6IkFySUVEU1lUYStjc204RTVhaU1XOWpRSjUzaHlhSVVsV2dwdlcxQ1hOaDhlWHBwZmprQXI4OHZSWjZBSEJQdlAiLCJtYWMiOiIxNDE1NzA3MWRhMzg1NzgyNWZhNzgxYmQ5NjdjODMxZmVlYjVmNDgzNDBkOTk2YmMzOGZlNmRiZWVlNThkZWFiIn0%3D',
    }

    data = {
        '_token': 'bola jabe nah :)',
        'allTogetherByDate': {date},
        'adjusted': ''
    }

    response = requests.post('https://stockbangladesh.com/download', headers=headers, data=data)
    decoded_content = response.content.decode("utf-8")
    cr = csv.reader(decoded_content.splitlines(), delimiter=',')
    my_list = list(cr)
    if len(my_list) > 0:
        #print(len(my_list))
        query = insert_query
        for index, row in enumerate(my_list):
            if len(row) == 7:
                query += query_builder(row)
                if index == len(my_list) - 1:
                    query += ";"
                else:
                    query += ","
        print(query)


if __name__ == '__main__':
    with open("insert.sql", "w") as sys.stdout:
        print("BEGIN;")
        start_date = date(2020, 12, 14)
        end_date = date(2020, 12, 14)
        delta = timedelta(days=1)
        while start_date <= end_date:
            date = start_date.strftime("%Y-%m-%d")
            data_downloader(date)
            start_date += delta
        print("COMMIT;")

# See PyCharm help at https://www.jetbrains.com/help/pycharm/
