export interface TransactionReq {
  from_address: string,
  from_private_key_hex: string,
  to_address: string,
  value: string, // wei
  gas_price: string, // wei
  gas_limit: string
}
