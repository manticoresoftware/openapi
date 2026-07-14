# Manticore Search Client
# Copyright (c) 2020-2021, Manticore Software LTD (https://manticoresearch.com)
#
# All rights reserved

import sys
import manticoresearch
import unittest
import importlib
import inspect
from parametrized_test_case import ParametrizedTestCase

configuration = manticoresearch.Configuration(
    host = "http://localhost:9508",
    #username = 'admin',
    #password = '12345678'
    access_token = 'a0818c4ef3fa4b7f575fb9cca722696f3c8fd6fce0c8d9223b5b05bad7d7d0c71'
)

print
suite = unittest.TestSuite()
for module_name in ['test_manual']:
    test = importlib.import_module(module_name)
    for name, obj in inspect.getmembers(test):
        if inspect.isclass(obj) and obj.__name__ != 'ParametrizedTestCase':
            for p_class in inspect.getmro(obj):
                if p_class.__name__ == 'TestCase':
                    suite.addTest(ParametrizedTestCase.parametrize(obj, {'configuration': configuration}))
result = unittest.TextTestRunner(verbosity=2).run(suite)                    
sys.exit(not result.wasSuccessful())
