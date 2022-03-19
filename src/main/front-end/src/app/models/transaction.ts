export interface TransactionReq {
  from_address: string,
  from_private_key_hex: string,
  to_address: string,
  value: string, // wei
  gas_price: string, // wei
  gas_limit: string
}

export interface TransactionResp {
  ID: number;
  CreatedAt: Date;
  UpdatedAt: Date;
  DeletedAt?: any;
  uid: number;
  hash: string;
  value: string;
  gas_price: string;
  gas_limit: string;
  nonce: string;
  from_address: string;
  to_address: string;
  gas_used: string;
  cost: string;
  status: boolean;
  is_pending: boolean;
  block_number: string;
  block_hash: string;
  network: string;
}
