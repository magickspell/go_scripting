from pb.my.my_pb2 import *

# импортируем все из протобаффа и добавляем селлеров
report = SellerParams()

first_seller = report.result.add()
first_seller.seller_id = 1
first_seller.rating = 11
first_seller.params["some_double"].double = 4.2

second_seller = report.result.add()
second_seller.seller_id = 3
second_seller.rating = 7
second_seller.params["name"].string = 'David'

# выводим селлеров в консоль
print('Report: ')
print(report)

# Записываем в файл
with open('./output.bin', 'wb') as f:
    f.write(report.SerializeToString())

