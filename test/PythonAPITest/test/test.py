import openapi_client
import unittest
import importlib
import inspect
from parametrized_test_case import ParametrizedTestCase

configuration = openapi_client.Configuration(
    host = "127.0.0.1:9308"
)

print
suite = unittest.TestSuite()
for module_name in ['test_index_api', 'test_search_api']:
    test = importlib.import_module(module_name)
    for name, obj in inspect.getmembers(test):
        if inspect.isclass(obj) and obj.__name__ != 'ParametrizedTestCase':
            for p_class in inspect.getmro(obj):
                if p_class.__name__ == 'TestCase':
                    suite.addTest(ParametrizedTestCase.parametrize(obj, {'configuration': configuration}))
unittest.TextTestRunner(verbosity=2).run(suite)                    
