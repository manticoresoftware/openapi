#!/bin.bash
pip3 install -U setuptools
pip3 install -r requirements.txt
python3 setup.py install
searchd
mysql -P9306 -h0 -e"CREATE TABLE test(content text, name text attribute, cat integer);"
mysql -P9306 -h0 -e"CREATE TABLE test_pq(content text, name text attribute, cat integer) type='pq';"

