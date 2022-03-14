export interface Network {
  name: string, // 网络名称
  url: string, // 网络rpc url e.g. http://localhost:7788
  chain_id: string, // chain id
  uid?: number,
}
