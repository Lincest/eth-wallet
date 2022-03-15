# eth-wallet

一个以太坊钱包。

## tech stack 技术栈

### 前端

- `angular`: [Angular](https://angular.cn/) `version: @angular/cli@13.2.5`
- `primeng`: [PrimeNG (primefaces.org)](https://www.primefaces.org/primeng/#/theming) `version: 13.2.0`

### 后端

- `gin`: [gin-gonic/gin](https://github.com/gin-gonic/gin)
- `gorm`: [GORM - The fantastic ORM library for Golang, aims to be developer friendly.](https://gorm.io/)
- `go-ethereum`: [ethereum/go-ethereum](https://github.com/ethereum/go-ethereum)
- `go-ethereum-hdwallet`: [miguelmota/go-ethereum-hdwallet: Ethereum HD Wallet derivations in Go (golang) (github.com)](https://github.com/miguelmota/go-ethereum-hdwallet)

### 调试

- `ganache`: [Ganache | Overview - Truffle Suite](https://trufflesuite.com/docs/ganache/)
- `geth`: [Go Ethereum](https://geth.ethereum.org/)

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
- [ ] keystore 导入
- [ ] keystore 导出
- [ ] privatekey 导入
- [x] 网络切换





