// 服务器返回格式
export interface Resp {
  code: Code,
  msg: string,
  data: any,
}

// 服务器code
export enum Code {
  ok = 0,
  err = -1,
}

export interface Base {
  ID: number,
  CreatedAt: string,
  DeletedAt: string,
}
