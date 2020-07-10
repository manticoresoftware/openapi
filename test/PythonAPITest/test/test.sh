#!/usr/bin/env bash
searchd --config test.conf
mysql -P9306 -h0 <<EOF
INSERT INTO test_pq(id, query) VALUES (1, '@content sample content');
EOF
python test.py
