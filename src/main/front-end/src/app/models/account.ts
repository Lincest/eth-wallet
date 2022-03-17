export interface Account {
  ID: number,
  uid: number,
  derivation_path: string,
  address: string,
  private_key_hex: string
  balance: string
}

export const defaultAccount: Account = {
  ID: 0,
  uid: 0,
  derivation_path: "",
  address: "",
  private_key_hex: "",
  balance: ""
}
