<?php
require_once(__DIR__ . '/vendor/autoload.php');
$config = new OpenAPI\Client\Configuration();
$config->setHost('http://127.0.0.1:6368');
$apiInstance = new OpenAPI\Client\Api\DefaultApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client() ,$config
);
$body = new \OpenAPI\Client\Model\Insert(); // \OpenAPI\Client\Model\Insert | Inventory item to add
$body->setIndex('testrt2');
$body->setId(102);
$body->setDoc(['title'=>'sample test','gid'=>100]);
try {
    $result = $apiInstance->insert($body);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DefaultApi->insert: ', $e->getMessage(), PHP_EOL;
}
