# Manticore TypeScript client

Сlient for Manticore Search.


❗ WARNING: this is a development version of the client. The latest release's readme is https://github.com/manticoresoftware/manticoresearch-typescript/tree/5.0.0

## Requirements

Minimum Manticore Search version is 4.2.1 with HTTP protocol enabled.

| Manticore Search | manticoresearch-typescript | Node      |
| ---------------- | -------------------------- | --------- |
| >= 6.3.6         | 5.0.x                      | >= 18.0.0 |
| >= 6.2.0         | 4.0.x                      | >= 18.0.0 |
| >= 6.2.0         | 3.3.1                      | >= 18.0.0 |
| >= 4.2.1         | >= 1.0.x                   | >= 18.0.0 |

### Building

To build and compile the typescript sources to javascript use:
```
npm install
npm run build
```

### Publishing

First build the package then run ```npm publish```

### Consuming

Navigate to the folder of your consuming project and run one of the following commands.

_published:_

```
npm install manticoresearch-ts@5.0.0 --save
```

_unPublished (not recommended):_

```
npm install PATH_TO_GENERATED_PACKAGE --save
```

### Usage

Below code snippet shows exemplary usage of the API. 

```
import {Configuration, IndexApi, SearchApi, UtilsApi} from "manticoresearch-ts";
const config = new Configuration({
	basePath: 'http://localhost:9308',
})
const indexApi = new IndexApi(config);
const searchApi = new SearchApi(config);
await indexApi.insert({index : 'products', id : 1, doc : {title : 'Crossbody Bag with Tassel'}});
await indexApi.insert({index : 'products', id : 2, doc : {title : 'Pet Hair Remover Glove'}});
const res = await searchApi.search({
	index: 'products',
	query: { query_string: {'bag'} },
});
 
```
