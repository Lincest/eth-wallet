# eth-wallet

一个以太坊钱包。

## tech stack 技术栈

### 前端

- `angular`: [Angular](https://angular.cn/) `version: @angular/cli@13.2.5`
- `primeng`: [PrimeNG (primefaces.org)](https://www.primefaces.org/primeng/#/theming) `version: 13.2.0`
- `primeflex`: https://www.primefaces.org/primeflex
- `web3.js`: [web3.utils — web3.js 1.0.0 documentation (web3js.readthedocs.io)](https://web3js.readthedocs.io/en/v1.7.1/web3-utils.html#tobn)

### 后端

- `gin`: [gin-gonic/gin](https://github.com/gin-gonic/gin)
- `gorm`: [GORM - The fantastic ORM library for Golang, aims to be developer friendly.](https://gorm.io/)
- `go-ethereum`: [ethereum/go-ethereum](https://github.com/ethereum/go-ethereum)
- `go-ethereum-hdwallet`: [miguelmota/go-ethereum-hdwallet: Ethereum HD Wallet derivations in Go (golang) (github.com)](https://github.com/miguelmota/go-ethereum-hdwallet)

### 调试

- `ganache`: [Ganache | Overview - Truffle Suite](https://trufflesuite.com/docs/ganache/)
- `geth`: [Go Ethereum](https://geth.ethereum.org/)

### 外部API

- https://min-api.cryptocompare.com/
- 

## file structure 文件结构

`./doc/*`: 开发过程的记录文档

`./src/main/front-end`: 项目前端

`./src/main/back-end`: 项目后端

 ## TODO | Feature

### Base

- [x] 用户注册, 登录

### Wallet

- [x] 用户助记词管理
- [ ] 账户管理
- [ ] 交易管理
- [x] keystore 导入
- [x] keystore 导出
- [x] privatekey 导入
- [x] 网络切换

## API

- `BASE_URL`: `/api/v1` 
- `AUTH_URL`: `/api/v1/auth` 鉴权后url

|                             url                              |                 method                 |                   description                    |
| :----------------------------------------------------------: | :------------------------------------: | :----------------------------------------------: |
|                   `{{BASE_URL}}/register`                    |                 `POST`                 |                     用户注册                     |
|                     `{{BASE_URL}}/login`                     |        `POST {name, password}`         |                     用户登录                     |
|                    `{{AUTH_URL}}/logout`                     |                 `POST`                 |                     用户登出                     |
|                   `{{AUTH_URL}}/mnemonic`                    |           `POST {mnemonic}`            |                    更新助记词                    |
|                    `{{AUTH_URL}}/network`                    |                 `GET`                  |            查询当前用户保存的网络节点            |
|                    `{{AUTH_URL}}/network`                    |    `POST {chain_id, url, name, ID}`    |               新增当前用户网络节点               |
|                    `{{AUTH_URL}}/network`                    |                `DELETE`                |               删除当前用户网络节点               |
|                    `{{AUTH_URL}}/network`                    |    `PUT {chain_id, url, name, ID}`     |               更新当前用户网络节点               |
|      `{{BASE_URL}}/test-rpc?url=url&chain-id=chain-id`       |                 `GET`                  |                  测试网络连通性                  |
|                `{{AUTH_URL}}/current-network`                |                 `GET`                  |               获取当前用户使用网络               |
|                `{{AUTH_URL}}/current-network`                |              `POST {ID}`               |               更新当前用户使用网络               |
|                    `{{AUTH_URL}}/account`                    |                 `POST`                 |         当前用户基于衍生路径新增account          |
|                    `{{AUTH_URL}}/account`                    |                 `GET`                  |             获取当前用户所有account              |
|                   `{{AUTH_URL}}/keystore`                    | `POST form{passphrase keystore(file)}` |       基于keystore和passphrase新增account        |
|                   `{{AUTH_URL}}/keystore`                    |                 `GET`                  | 导出该用户所有的account为zip文件, 密码是用户密码 |
|                    `{{AUTH_URL}}/account`                    |        `POST {private_key_hex}`        |            基于private key新增account            |
|                  `{{AUTH_URL}}/transaction`                  |       `POST {from_address ...}`        |                    创建新交易                    |
| `{{{{AUTH_URL}}/transaction/{transaction-hash=transaction-hash}` |                 `GET`                  |                查询并更新交易状态                |
|                                                              |                                        |                                                  |


