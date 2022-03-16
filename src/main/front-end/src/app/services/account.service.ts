import { Injectable } from '@angular/core';
import {Code, Resp} from "../models/resp";
import {Account} from "../models/account";
import {map, Observable} from "rxjs";
import {AUTH_URL} from "../models/global";
import {HttpClient} from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class AccountService {

  constructor(
    private http: HttpClient
  ) { }

  // 获取所有账户信息
  getAllAccounts(): Observable<Resp> {
    return this.http.get<Resp>(AUTH_URL + "/account")
  }

  // 上传keystore文件
  uploadKeyStoreFile(keystoreFile: any, passphrase: string): Observable<Resp> {
    const formData = new FormData()
    formData.set("keystore", keystoreFile)
    formData.set("passphrase", passphrase)
    return this.http.post<Resp>(AUTH_URL + "/keystore", formData)
  }

  // 导出keystore文件
  exportKeyStoreFile(): Observable<any> {
    return this.http.get(AUTH_URL + "/keystore", {responseType: 'blob' as 'json'})
  }

  // 上传私钥
  uploadPrivateKey(privateKey: string): Observable<Resp> {
    return this.http.post<Resp>(AUTH_URL + "/account", {private_key_hex: privateKey})
  }

  // 基于衍生路径新增账户
  generateNewAccount(): Observable<Resp> {
    return this.http.post<Resp>(AUTH_URL + "/account", null)
  }
}
