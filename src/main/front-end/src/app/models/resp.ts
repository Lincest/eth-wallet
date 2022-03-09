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
